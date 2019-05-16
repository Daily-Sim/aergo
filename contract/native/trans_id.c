/**
 * @file    trans_id.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "vector.h"
#include "ir_md.h"
#include "ir_fn.h"
#include "ir_bb.h"
#include "trans_blk.h"
#include "trans_stmt.h"
#include "trans_exp.h"
#include "syslib.h"

#include "trans_id.h"

static void
id_trans_param(trans_t *trans, ast_id_t *id)
{
    ast_id_t *param_id;

    if (is_ctor_id(id))
        return;

    /* All functions that are not constructors must be added the contract address as the first
     * argument, and must also be added to the param_ids to reflect the abi */

    param_id = id_new_tmp_var("cont$addr", TYPE_OBJECT, &id->pos);
    param_id->u_var.is_param = true;
    param_id->up = id;
    meta_set_object(&param_id->meta, id->up);

    if (id->u_fn.param_ids == NULL)
        id->u_fn.param_ids = vector_new();

    vector_add_first(id->u_fn.param_ids, param_id);
}

static void
id_trans_ctor(trans_t *trans, ast_id_t *id)
{
    int i, j;
    uint8_t align = 0;
    ir_fn_t *fn = trans->fn;
    vector_t *stmts = vector_new();

    ASSERT(id->up != NULL);
    ASSERT1(is_cont_id(id->up), id->up->kind);
    ASSERT1(fn->heap_usage == 0, fn->heap_usage);

    fn->cont_idx = fn->heap_idx;

    trans->bb = fn->entry_bb;

    vector_foreach(&id->up->u_cont.blk->ids, i) {
        ast_id_t *glob_id = vector_get_id(&id->up->u_cont.blk->ids, i);

        if (is_var_id(glob_id)) {
            fn_add_global(fn, &glob_id->meta);

            if (align == 0)
                align = meta_align(&glob_id->meta);

            stmt_add(stmts, stmt_new_id(glob_id, &glob_id->pos));
        }
        else if (is_tuple_id(glob_id)) {
            vector_foreach(glob_id->u_tup.elem_ids, j) {
                ast_id_t *elem_id = vector_get_id(glob_id->u_tup.elem_ids, j);

                ASSERT1(is_var_id(elem_id), elem_id->kind);

                fn_add_global(fn, &elem_id->meta);

                if (align == 0)
                    align = meta_align(&elem_id->meta);

                stmt_add(stmts, stmt_new_id(elem_id, &elem_id->pos));
            }
        }
    }

    /* A contiguous memory space is required because the contract address must be passed to the
     * member function. */
    if (fn->heap_usage > 0)
        stmt_trans(trans, stmt_make_malloc(fn->heap_idx, fn->heap_usage, align, &id->pos));

    vector_foreach(stmts, i) {
        stmt_trans(trans, vector_get_stmt(stmts, i));
    }

    trans->bb = NULL;
}

static void
set_stack_addr(trans_t *trans, ast_id_t *id)
{
    ast_exp_t *reg_exp;
    ast_exp_t *size_exp;
    ast_exp_t *align_exp;
    ast_exp_t *call_exp;
    ast_exp_t *stk_exp;
    vector_t *arg_exps;
    ir_fn_t *fn = trans->fn;

    if (fn->stack_usage == 0)
        return;

    /* At the beginning of "entry_bb", set the current stack offset to the register */
    reg_exp = exp_new_reg(fn->stack_idx);
    meta_set_int32(&reg_exp->meta);

    arg_exps = vector_new();

    size_exp = exp_new_lit_int(fn->stack_usage, &id->pos);
    meta_set_int32(&size_exp->meta);

    exp_add(arg_exps, size_exp);

    align_exp = exp_new_lit_int(8, &id->pos);
    meta_set_int32(&align_exp->meta);

    exp_add(arg_exps, align_exp);

    call_exp = exp_new_call(FN_ALLOCA, NULL, arg_exps, &id->pos);
    meta_set_int32(&call_exp->meta);

    exp_trans(trans, call_exp);

    vector_add_first(&fn->entry_bb->stmts, stmt_new_assign(reg_exp, call_exp, &id->pos));

    stk_exp = exp_new_global("__STACK_TOP");
    meta_set_int32(&stk_exp->meta);

    /* If there is any stack variable in the function, it has to be restored to the original value
     * at the end of "exit_bb" because "__STACK_TOP" has been changed */
    vector_add_last(&fn->exit_bb->stmts, stmt_new_assign(stk_exp, reg_exp, &id->pos));
}

static void
check_return_stmt(ir_fn_t *fn, src_pos_t *pos)
{
    int i, j;

    vector_foreach(&fn->bbs, i) {
        ir_bb_t *bb = vector_get_bb(&fn->bbs, i);

        /* except unreachable block */
        if (bb->ref_cnt > 0) {
            vector_foreach(&bb->brs, j) {
                ast_stmt_t *stmt;
                ir_br_t *br = vector_get_br(&bb->brs, j);

                if (br->bb != fn->exit_bb)
                    continue;

                stmt = vector_get_last(&bb->stmts, ast_stmt_t);
                if (stmt == NULL || !is_return_stmt(stmt)) {
                    ERROR(ERROR_MISSING_RETURN, pos);
                    return;
                }
            }
        }
    }
}

static void
id_trans_fn(trans_t *trans, ast_id_t *id)
{
    ast_id_t *ret_id = id->u_fn.ret_id;
    meta_t addr_meta;
    ir_md_t *md = trans->md;
    ir_fn_t *fn;

    ASSERT(id->up != NULL);
    ASSERT1(is_cont_id(id->up), id->up->kind);
    ASSERT(md != NULL);

    id_trans_param(trans, id);

    fn = fn_new(id);

    meta_set_int32(&addr_meta);

    /* All heap variables access memory by adding relative offset to this register */
    fn->heap_idx = fn_add_register(fn, &addr_meta);

    /* All stack variables access memory by adding relative offset to this register */
    fn->stack_idx = fn_add_register(fn, &addr_meta);

    if (ret_id != NULL)
        /* All return values are stored in this register */
        fn->ret_idx = fn_add_register(fn, &ret_id->meta);

    /* It is used internally for binaryen, not for us (see fn_gen()) */
    fn->reloop_idx = fn_add_register(fn, &addr_meta);

    trans->fn = fn;

    if (is_ctor_id(id))
        id_trans_ctor(trans, id);
    else
        /* The "cont_idx" is always 0 because it is prepended to parameters */
        fn->cont_idx = 0;

    trans->bb = fn->entry_bb;

    if (id->u_fn.blk != NULL)
        blk_trans(trans, id->u_fn.blk);

    set_stack_addr(trans, id);

    if (trans->bb != NULL) {
        bb_add_branch(trans->bb, NULL, fn->exit_bb);
        fn_add_basic_blk(fn, trans->bb);
    }

    fn_add_basic_blk(fn, fn->exit_bb);

    if (ret_id != NULL) {
        ast_exp_t *arg_exp;
        ast_stmt_t *ret_stmt;

        if (is_ctor_id(id)) {
            /* The contract address is returned at the end of "exit_bb" */
            arg_exp = exp_new_reg(fn->cont_idx);
            meta_set_int32(&arg_exp->meta);

            ret_stmt = stmt_new_return(arg_exp, &id->pos);
            ret_stmt->u_ret.ret_id = ret_id;

            vector_add_last(&fn->exit_bb->stmts, ret_stmt);
        }
        else {
            check_return_stmt(fn, &id->pos);
        }
    }

    trans->fn = NULL;
    trans->bb = NULL;

    if (is_public_id(id) || id->is_used)
        md_add_fn(md, fn);
}

static void
id_trans_contract(trans_t *trans, ast_id_t *id)
{
    int i;
    ast_blk_t *up = trans->blk;
    ast_blk_t *blk = id->u_cont.blk;

    ASSERT(blk != NULL);

    if (id->is_imported) {
        vector_foreach(&blk->ids, i) {
            ast_id_t *elem_id = vector_get_id(&blk->ids, i);

            if (is_fn_id(elem_id))
                id_trans_param(trans, elem_id);
        }
        return;
    }

    trans->id = id;
    trans->md = md_new(id->name);
    trans->blk = blk;

    /* "blk->stmts" is no longer used. */
    vector_foreach(&blk->ids, i) {
        id_trans(trans, vector_get_id(&blk->ids, i));
    }

    ir_add_md(trans->ir, trans->md);

    trans->blk = up;
    trans->md = NULL;
    trans->id = NULL;
}

static void
id_trans_interface(trans_t *trans, ast_id_t *id)
{
    int i;
    ast_blk_t *blk = id->u_itf.blk;

    ASSERT(blk != NULL);

    vector_foreach(&blk->ids, i) {
        ast_id_t *elem_id = vector_get_id(&blk->ids, i);

        ASSERT1(is_fn_id(elem_id), elem_id->kind);
        ASSERT(!is_ctor_id(elem_id));

        /* If the interface type is used as a parameter, we can invoke it with the interface
         * function, so transform the parameter here */

        id_trans_param(trans, elem_id);
    }
}

static void
id_trans_label(trans_t *trans, ast_id_t *id)
{
    id->u_lab.stmt->label_bb = bb_new();
}

void
id_trans(trans_t *trans, ast_id_t *id)
{
    ASSERT(id != NULL);

    switch (id->kind) {
    case ID_FN:
        id_trans_fn(trans, id);
        break;

    case ID_CONT:
        id_trans_contract(trans, id);
        break;

    case ID_ITF:
        id_trans_interface(trans, id);
        break;

    case ID_LABEL:
        id_trans_label(trans, id);
        break;

    case ID_VAR:
    case ID_STRUCT:
    case ID_ENUM:
    case ID_TUPLE:
    case ID_LIB:
        break;

    default:
        ASSERT1(!"invalid identifier", id->kind);
    }
}

/* end of trans_id.c */
