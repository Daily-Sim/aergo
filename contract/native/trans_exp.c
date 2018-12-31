/**
 * @file    trans_exp.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "ast_id.h"
#include "ast_blk.h"
#include "ast_stmt.h"
#include "ir_bb.h"
#include "ir_fn.h"
#include "ir_sgmt.h"

#include "trans_exp.h"

static void
exp_trans_id_ref(trans_t *trans, ast_exp_t *exp)
{
    ast_id_t *id = exp->id;

    ASSERT(id != NULL);

    if (!is_var_id(id))
        return;

    if (is_global_id(id) || is_stack_id(id)) {
        /* The global variable always refers to the stack */
        ASSERT(id->addr >= 0);

        exp->kind = EXP_STACK_REF;
        exp->u_st.is_lval = trans->is_lval;
        exp->u_st.addr = id->addr;
        exp->u_st.offset = id->offset;
    }
    else {
        ASSERT(id->idx >= 0);

        exp->kind = EXP_LOCAL_REF;
        exp->u_lo.is_lval = trans->is_lval;
        exp->u_lo.index = id->idx;
    }
}

static void
exp_trans_lit(trans_t *trans, ast_exp_t *exp)
{
    value_t *val = &exp->u_lit.val;
    ir_sgmt_t *sgmt = trans->ir->sgmt;

    switch (val->type) {
    case TYPE_BOOL:
    case TYPE_UINT64:
    case TYPE_DOUBLE:
        break;

    case TYPE_STRING:
        value_set_i64(val, sgmt_add_raw(sgmt, val_ptr(val), val_size(val) + 1));
        break;

    case TYPE_OBJECT:
        if (is_null_val(val))
            value_set_i64(val, 0);
        else
            value_set_i64(val, sgmt_add_raw(sgmt, val_ptr(val), val_size(val) + 1));
        break;

    default:
        ASSERT1(!"invalid value", val->type);
    }
}

static void
exp_trans_array(trans_t *trans, ast_exp_t *exp)
{
    ast_id_t *id = exp->id;

    exp_trans(trans, exp->u_arr.id_exp);
    exp_trans(trans, exp->u_arr.idx_exp);

    if (is_array_type(&id->meta)) {
        uint32_t offset = 0;
        ast_exp_t *id_exp = exp->u_arr.id_exp;
        ast_exp_t *idx_exp = exp->u_arr.idx_exp;

        if (!is_lit_exp(idx_exp))
            return;

        if (is_stack_ref_exp(id_exp))
            offset += id_exp->u_st.offset;

        /* The following arr_size is stripped arr_size */
        offset += val_i64(&idx_exp->u_lit.val) * exp->meta.arr_size;

        exp->kind = EXP_STACK_REF;
        exp->u_st.addr = id->addr;
        exp->u_st.offset = offset;
    }
    else {
        /* TODO
         * int addr = fn_add_stack_var(trans->fn);
         * ast_exp_t *call_exp = exp_new_call("$map_get", &exp->pos);
         *
         * bb_add_stmt(trans->bb, stmt_new_exp(call_exp, &exp->pos));
         *
         * return <return address of call>; */
    }
}

static void
exp_trans_cast(trans_t *trans, ast_exp_t *exp)
{
    exp_trans(trans, exp->u_cast.val_exp);

    if (!is_primitive_type(&exp->meta) || !is_primitive_type(&exp->u_cast.to_meta)) {
        /* TODO
         * int addr = fn_add_stack_var(trans->fn);
         * ast_exp_t *call_exp = exp_new_call("$concat", &exp->pos);
         *
         * bb_add_stmt(trans->bb, stmt_new_exp(call_exp, &exp->pos));
         *
         * return <return address of call>; */
    }
}

static void
exp_trans_unary(trans_t *trans, ast_exp_t *exp)
{
    bool is_lval = trans->is_lval;
    ast_exp_t *rval_exp = exp->u_un.val_exp;
    ast_exp_t *lval_exp, *bi_exp, *lit_exp;

    switch (exp->u_un.kind) {
    case OP_INC:
    case OP_DEC:
        lval_exp = exp_clone(rval_exp);

        exp_trans_to_lval(trans, lval_exp);
        exp_trans_to_rval(trans, rval_exp);

        trans->is_lval = is_lval;

        lit_exp = exp_new_lit(&exp->pos);
        value_set_i64(&lit_exp->u_lit.val, 1);

        bi_exp = exp_new_binary(exp->u_un.kind == OP_INC ? OP_ADD : OP_SUB, rval_exp,
                                lit_exp, &exp->pos);

        meta_copy(&lit_exp->meta, &rval_exp->meta);
        meta_copy(&bi_exp->meta, &rval_exp->meta);

        bb_set_piggyback(trans->bb, stmt_new_assign(lval_exp, bi_exp, &exp->pos));

        if (exp->u_un.is_prefix)
            *exp = *bi_exp;
        else
            *exp = *rval_exp;
        break;

    case OP_NEG:
    case OP_NOT:
        exp_trans(trans, rval_exp);
        break;

    default:
        ASSERT1(!"invalid operator", exp->u_un.kind);
    }
}

static void
exp_trans_binary(trans_t *trans, ast_exp_t *exp)
{
    exp_trans(trans, exp->u_bin.l_exp);
    exp_trans(trans, exp->u_bin.r_exp);

    if (exp->u_bin.kind == OP_ADD && is_string_type(&exp->meta)) {
        /* TODO
         * int addr = fn_add_stack();
         * ast_exp_t *call exp = exp_new_call("$concat", &exp->pos);
         * bb_add_stmt(trans->bb, stmt_new_exp(call_exp, &exp->pos));
         * return exp_new_stack(addr, offset, &exp->pos);
        */
    }
}

static void
exp_trans_ternary(trans_t *trans, ast_exp_t *exp)
{
    exp_trans(trans, exp->u_tern.pre_exp);
    exp_trans(trans, exp->u_tern.in_exp);
    exp_trans(trans, exp->u_tern.post_exp);

    if (is_lit_exp(exp->u_tern.pre_exp)) {
        /* Maybe we should do this in optimizer */
        meta_t meta = exp->meta;

        if (val_bool(&exp->u_tern.pre_exp->u_lit.val))
            *exp = *exp->u_tern.in_exp;
        else
            *exp = *exp->u_tern.post_exp;

        meta_copy(&exp->meta, &meta);
    }
}

static void
exp_trans_access(trans_t *trans, ast_exp_t *exp)
{
    ast_id_t *qual_id = exp->u_acc.id_exp->id;
    ast_id_t *fld_id = exp->id;

    exp_trans(trans, exp->u_acc.id_exp);
    exp_trans(trans, exp->u_acc.fld_exp);

    if (is_fn_id(fld_id))
        /* we will transform this to call expression in the generator */
        return;

    ASSERT(qual_id != NULL);

    exp->kind = EXP_STACK_REF;
    exp->u_st.addr = qual_id->addr;
    exp->u_st.offset = fld_id->offset;
}

static void
exp_trans_call(trans_t *trans, ast_exp_t *exp)
{
    ast_id_t *id = exp->id;

    // FIXME bb_add_stmt(trans->bb, stmt_new_exp(exp, &exp->pos));

    if (is_map_type(&exp->meta))
        return;

    if (id->u_fn.ret_id != NULL) {
        ast_id_t *ret_id = id->u_fn.ret_id;

        if (is_tuple_id(ret_id)) {
            int i;
            array_t *exps = array_new();
            array_t *var_ids = &ret_id->u_tup.var_ids;

            ASSERT(trans->fn != NULL);

            for (i = 0; i < array_size(var_ids); i++) {
                ast_id_t *var_id = array_get_id(var_ids, i);
                ast_exp_t *ref_exp;

                ASSERT1(var_id->offset == 0, var_id->offset);

                fn_add_stack(trans->fn, var_id);

                ref_exp = exp_new_stack_ref(var_id->addr, 0, &exp->pos);
                meta_copy(&ref_exp->meta, &var_id->meta);

                array_add_last(exps, ref_exp);
            }

            exp->kind = EXP_TUPLE;
            exp->u_tup.exps = exps;
        }
        else {
            fn_add_stack(trans->fn, ret_id);

            ASSERT1(ret_id->offset == 0, ret_id->offset);

            exp->kind = EXP_STACK_REF;
            exp->u_st.addr = ret_id->addr;
            exp->u_st.offset = 0;
        }
    }
}

static void
exp_trans_sql(trans_t *trans, ast_exp_t *exp)
{
}

static void
exp_trans_tuple(trans_t *trans, ast_exp_t *exp)
{
    int i;
    array_t *exps = exp->u_tup.exps;

    for (i = 0; i < array_size(exps); i++) {
        exp_trans(trans, array_get_exp(exps, i));
    }
}

static void
exp_trans_init(trans_t *trans, ast_exp_t *exp)
{
    int i;
    array_t *exps = exp->u_init.exps;

    for (i = 0; i < array_size(exps); i++) {
        exp_trans(trans, array_get_exp(exps, i));
    }
}

void
exp_trans(trans_t *trans, ast_exp_t *exp)
{
    switch (exp->kind) {
    case EXP_NULL:
        return;

    case EXP_ID_REF:
        exp_trans_id_ref(trans, exp);
        break;

    case EXP_LOCAL_REF:
    case EXP_STACK_REF:
        return;

    case EXP_LIT:
        exp_trans_lit(trans, exp);
        break;

    case EXP_ARRAY:
        exp_trans_array(trans, exp);
        break;

    case EXP_CAST:
        exp_trans_cast(trans, exp);
        break;

    case EXP_UNARY:
        exp_trans_unary(trans, exp);
        break;

    case EXP_BINARY:
        exp_trans_binary(trans, exp);
        break;

    case EXP_TERNARY:
        exp_trans_ternary(trans, exp);
        break;

    case EXP_ACCESS:
        exp_trans_access(trans, exp);
        break;

    case EXP_CALL:
        exp_trans_call(trans, exp);
        break;

    case EXP_SQL:
        exp_trans_sql(trans, exp);
        break;

    case EXP_TUPLE:
        exp_trans_tuple(trans, exp);
        break;

    case EXP_INIT:
        exp_trans_init(trans, exp);
        break;

    default:
        ASSERT1(!"invalid expression", exp->kind);
    }
}

/* end of trans_exp.c */
