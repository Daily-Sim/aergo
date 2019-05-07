/**
 * @file    trans_exp.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "ast_id.h"
#include "ast_blk.h"
#include "ast_stmt.h"
#include "ir_md.h"
#include "ir_abi.h"
#include "ir_fn.h"
#include "ir_bb.h"
#include "ir_sgmt.h"
#include "trans_stmt.h"
#include "syslib.h"

#include "trans_exp.h"

static uint32_t make_array_alloc(trans_t *trans, meta_t *meta);

static void
exp_trans_lit(trans_t *trans, ast_exp_t *exp)
{
    int addr;
    value_t *val = &exp->u_lit.val;
    meta_t *meta = &exp->meta;
    ir_md_t *md = trans->md;

    ASSERT(md != NULL);

    switch (val->type) {
    case TYPE_BOOL:
    case TYPE_BYTE:
    case TYPE_INT128:
    case TYPE_DOUBLE:
        break;

    case TYPE_STRING:
        /* Since val_set_int() is a macro, DO NOT combine the next two lines. */
        addr = sgmt_add_str(&md->sgmt, val_ptr(val));
        value_set_int(val, addr);
        meta_set_int32(meta);
        break;

    case TYPE_OBJECT:
        if (is_null_val(val)) {
            value_set_int(val, 0);
        }
        else {
            /* Same as above */
            addr = sgmt_add_raw(&md->sgmt, val_ptr(val), val_size(val));
            value_set_int(val, addr);
        }
        meta_set_int32(meta);
        break;

    default:
        ASSERT1(!"invalid value", val->type);
    }
}

static void
exp_trans_id(trans_t *trans, ast_exp_t *exp)
{
    ast_id_t *id = exp->id;
    ir_fn_t *fn = trans->fn;

    ASSERT(id != NULL);

    if (is_var_id(id)) {
        if (is_global_id(id)) {
            ASSERT1(id->meta.rel_offset == 0, id->meta.rel_offset);

            /* The global variable always refers to the memory. */
            exp_set_mem(exp, id->meta.base_idx, id->meta.rel_addr, 0);
        }
        else {
            exp_set_reg(exp, id->idx);
        }
    }
    else if (is_cont_id(id)) {
        /* If "id" is a contract identifier, it means the "this" identifier */
        ASSERT1(is_object_meta(&exp->meta), exp->meta.type);

        exp_set_reg(exp, fn->cont_idx);
    }
}

static void
exp_trans_array(trans_t *trans, ast_exp_t *exp)
{
    exp_trans(trans, exp->u_arr.id_exp);
    exp_trans(trans, exp->u_arr.idx_exp);
}

static void
exp_trans_cast(trans_t *trans, ast_exp_t *exp)
{
    exp_trans(trans, exp->u_cast.val_exp);
}

static void
exp_trans_unary(trans_t *trans, ast_exp_t *exp)
{
    op_kind_t op = exp->u_un.kind;
    ast_exp_t *val_exp = exp->u_un.val_exp;
    ast_exp_t *var_exp, *bi_exp, *lit_exp;

    switch (op) {
    case OP_INC:
    case OP_DEC:
        /* Clone value expression because we have to transform it to "x op 1" */
        var_exp = exp_clone(val_exp);

        exp_trans(trans, var_exp);
        exp_trans(trans, val_exp);

        lit_exp = exp_new_lit_int(1, &exp->pos);
        meta_copy(&lit_exp->meta, &val_exp->meta);

        if (!exp->u_un.is_prefix) {
            ast_exp_t *reg_exp = exp_new_reg(fn_add_register(trans->fn, &val_exp->meta));

            meta_copy(&reg_exp->meta, &val_exp->meta);

            bb_add_stmt(trans->bb, stmt_new_assign(reg_exp, val_exp, &exp->pos));
            val_exp = reg_exp;
        }

        bi_exp = exp_new_binary(op == OP_INC ? OP_ADD : OP_SUB, val_exp, lit_exp, &exp->pos);
        meta_copy(&bi_exp->meta, &val_exp->meta);

        bb_add_stmt(trans->bb, stmt_new_assign(var_exp, bi_exp, &exp->pos));
        *exp = *val_exp;
        break;

    case OP_NEG:
    case OP_NOT:
        exp_trans(trans, val_exp);
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
}

static void
exp_trans_ternary(trans_t *trans, ast_exp_t *exp)
{
    exp_trans(trans, exp->u_tern.pre_exp);
    exp_trans(trans, exp->u_tern.in_exp);
    exp_trans(trans, exp->u_tern.post_exp);

    if (is_lit_exp(exp->u_tern.pre_exp)) {
        /* Maybe we should do this in optimizer (if exists) */
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
    ast_exp_t *qual_exp = exp->u_acc.qual_exp;
    ast_id_t *fld_id = exp->id;

    exp_trans(trans, qual_exp);

    if (is_fn_id(fld_id)) {
        ASSERT1(is_reg_exp(qual_exp), qual_exp->kind);
        return;
    }

    if (is_reg_exp(qual_exp))
        /* Access relative address based on "qual_exp". */
        exp_set_mem(exp, qual_exp->meta.base_idx, fld_id->meta.rel_addr, fld_id->meta.rel_offset);
    else
        /* Only the offset value of the current field is stored. (see exp_gen_access()) */
        exp->meta.rel_offset = fld_id->meta.rel_offset;
}

static void
exp_trans_call(trans_t *trans, ast_exp_t *exp)
{
    int i;
    meta_t *meta = &exp->meta;
    ast_exp_t *id_exp = exp->u_call.id_exp;
    ast_id_t *fn_id = exp->id;
    ir_fn_t *fn = trans->fn;

    if (exp->u_call.kind != FN_UDF && exp->u_call.kind != FN_CTOR) {
        sys_fn_t *sys_fn = SYS_FN(exp->u_call.kind);

        ASSERT(id_exp == NULL);

        exp->u_call.qname = sys_fn->qname;
        md_add_abi(trans->md, syslib_abi(sys_fn));
        return;
    }

    exp_trans(trans, id_exp);

    if (fn_id->up != trans->id)
        md_add_abi(trans->md, abi_new(fn_id));

    if (is_ctor_id(fn_id) || is_lib_id(fn_id->up)) {
        /* The constructor does not change the parameter, it always returns address. */
        vector_foreach(exp->u_call.arg_exps, i) {
            exp_trans(trans, vector_get_exp(exp->u_call.arg_exps, i));
        }
        return;
    }

    /* Since non-constructor functions are added the contract base address as a first parameter,
     * we must also add the address as a call argument here. */
    if (exp->u_call.arg_exps == NULL)
        exp->u_call.arg_exps = vector_new();

    if (is_access_exp(id_exp)) {
        ast_exp_t *qual_exp = id_exp->u_acc.qual_exp;

        ASSERT1(is_reg_exp(qual_exp), qual_exp->kind);
        ASSERT1(is_object_meta(&qual_exp->meta), qual_exp->meta.type);

        /* If the expression is of type "x.y()", pass "x" as the first argument. */
        vector_add_first(exp->u_call.arg_exps, qual_exp);
    }
    else {
        ast_exp_t *param_exp;

        ASSERT1(fn->cont_idx == 0, fn->cont_idx);

        /* If the expression is of type "x()", pass my first parameter as the first argument. */
        param_exp = exp_new_reg(0);
        meta_set_int32(&param_exp->meta);

        vector_add_first(exp->u_call.arg_exps, param_exp);
    }

    vector_foreach(exp->u_call.arg_exps, i) {
        exp_trans(trans, vector_get_exp(exp->u_call.arg_exps, i));
    }

    if (fn_id->u_fn.ret_id != NULL) {
        uint32_t reg_idx = fn_add_register(fn, meta);
        ast_exp_t *reg_exp;

        reg_exp = exp_new_reg(reg_idx);
        meta_copy(&reg_exp->meta, meta);

        /* We have to clone it because the call expression itself is transformed */
        bb_add_stmt(trans->bb, stmt_new_assign(reg_exp, exp_clone(exp), &exp->pos));

        if (is_array_meta(&fn_id->meta) || is_struct_meta(&fn_id->meta)) {
            uint32_t tmp_idx = fn_add_register(trans->fn, meta);
            ast_exp_t *addr_exp, *cpy_exp;

            /* If the return value is an array or struct, we must copy the value because we do
             * share memory space between the caller and the callee */
            stmt_trans_malloc(trans, meta, tmp_idx);

            addr_exp = exp_new_reg(tmp_idx);
            meta_set_int32(&addr_exp->meta);

            cpy_exp = syslib_new_memcpy(trans, addr_exp, reg_exp, meta_memsz(meta), &exp->pos);

            bb_add_stmt(trans->bb, stmt_new_exp(cpy_exp, &exp->pos));

            exp_set_reg(exp, tmp_idx);
        }
        else {
            exp_set_reg(exp, reg_idx);
        }
    }
}

static void
exp_trans_sql(trans_t *trans, ast_exp_t *exp)
{
    /* TODO */
}

static void
exp_trans_tuple(trans_t *trans, ast_exp_t *exp)
{
    int i;
    vector_t *elem_exps = exp->u_tup.elem_exps;

    vector_foreach(elem_exps, i) {
        exp_trans(trans, vector_get_exp(elem_exps, i));
    }
}

static void
make_map_init(trans_t *trans, ast_exp_t *exp)
{
    int i;
    uint32_t reg_idx;
    fn_kind_t kind;
    meta_t *meta = &exp->meta;
    vector_t *elem_exps = exp->u_init.elem_exps;
    ast_exp_t *l_exp, *r_exp;

    //ASSERT(exp->u_init.is_topmost);
    ASSERT1(meta->elem_cnt == 2, meta->elem_cnt);

    reg_idx = fn_add_register(trans->fn, meta);

    l_exp = exp_new_reg(reg_idx);
    meta_set_int32(&l_exp->meta);

    if (is_int64_meta(meta->elems[0]) && is_int64_meta(meta->elems[1]))
        kind = FN_MAP_NEW_I64_I64;
    else if (is_int64_meta(meta->elems[0]))
        kind = FN_MAP_NEW_I64_I32;
    else if (is_int64_meta(meta->elems[1]))
        kind = FN_MAP_NEW_I32_I64;
    else
        kind = FN_MAP_NEW_I32_I32;

    r_exp = exp_new_call(kind, NULL, NULL, meta->pos);
    meta_set_int32(&r_exp->meta);

    stmt_trans(trans, stmt_new_assign(l_exp, r_exp, meta->pos));

    vector_foreach(elem_exps, i) {
        /* key-value pair */
        ast_exp_t *kvp_exp = vector_get_exp(elem_exps, i);
        ast_exp_t *k_exp, *v_exp;
        ast_exp_t *call_exp;
        vector_t *arg_exps = vector_new();

        ASSERT1(is_init_exp(kvp_exp), kvp_exp->kind);
        ASSERT1(vector_size(kvp_exp->u_init.elem_exps) == 2,
                vector_size(kvp_exp->u_init.elem_exps));

        k_exp = vector_get_first(kvp_exp->u_init.elem_exps, ast_exp_t);
        v_exp = vector_get_last(kvp_exp->u_init.elem_exps, ast_exp_t);

        exp_trans(trans, k_exp);
        exp_trans(trans, v_exp);

        exp_add(arg_exps, l_exp);
        exp_add(arg_exps, k_exp);
        exp_add(arg_exps, v_exp);

        if (is_int64_meta(&k_exp->meta) && is_int64_meta(&v_exp->meta))
            kind = FN_MAP_PUT_I64_I64;
        else if (is_int64_meta(&k_exp->meta))
            kind = FN_MAP_PUT_I64_I32;
        else if (is_int64_meta(&v_exp->meta))
            kind = FN_MAP_PUT_I32_I64;
        else
            kind = FN_MAP_PUT_I32_I32;

        call_exp = exp_new_call(kind, NULL, arg_exps, meta->pos);
        meta_set_void(&call_exp->meta);

        stmt_trans(trans, stmt_new_exp(call_exp, meta->pos));
    }

    exp_set_reg(exp, reg_idx);
}

static void
make_static_init(trans_t *trans, ast_exp_t *exp)
{
    int i;
    uint32_t offset = 0;
    uint32_t size;
    char *raw;
    meta_t *meta = &exp->meta;
    vector_t *elem_exps = exp->u_init.elem_exps;

    if (is_map_meta(meta)) {
        exp->u_init.is_static = false;
        return;
    }

    size = meta_memsz(meta);
    raw = xcalloc(size);

    if (is_array_meta(meta)) {
        ASSERT(meta->dim_sizes[0] > 0);
        ASSERT2(meta_align(meta) > 0, meta->type, meta_align(meta));
        ASSERT2((ptrdiff_t)(raw + offset) % meta_align(meta) == 0, raw, offset);

        if (meta_align(meta) == 8)
            *(int64_t *)(raw + offset) = meta->dim_sizes[0];
        else
            *(int *)(raw + offset) = meta->dim_sizes[0];

        offset += meta_align(meta);
    }

    exp->u_init.is_static = true;

    vector_foreach(elem_exps, i) {
        uint32_t write_sz;
        ast_exp_t *elem_exp = vector_get_exp(elem_exps, i);
        meta_t *elem_meta = &elem_exp->meta;
        value_t *elem_val = &elem_exp->u_lit.val;

        exp_trans(trans, elem_exp);

        if ((is_lit_exp(elem_exp) && is_int_val(elem_val) &&
             !mpz_fits_slong_p(val_mpz(elem_val)) && !mpz_fits_ulong_p(val_mpz(elem_val))) ||
            is_map_meta(elem_meta) ||
            (!is_lit_exp(elem_exp) && (!is_init_exp(elem_exp) || !elem_exp->u_init.is_static))) {
            exp->u_init.is_static = false;
            return;
        }

        ASSERT1(is_lit_exp(elem_exp), elem_exp->kind);

        /* Only struct, or array which is a member of struct, is stored in separate memory. */
        if (is_struct_meta(elem_meta) || (is_struct_meta(meta) && is_tuple_meta(elem_meta))) {
            uint32_t addr;
            ir_md_t *md = trans->md;

            ASSERT1(is_ptr_val(elem_val), elem_val->type);
            ASSERT1(val_size(elem_val) > 0, val_size(elem_val));

            addr = sgmt_add_raw(&md->sgmt, val_ptr(elem_val), val_size(elem_val));

            value_set_int(elem_val, addr);
            meta_set_int32(elem_meta);
        }

        offset = ALIGN(offset, meta_align(elem_meta));
        write_sz = value_serialize(elem_val, raw + offset, elem_meta);

        ASSERT2(write_sz <= meta_memsz(elem_meta), write_sz, meta_memsz(elem_meta));
        ASSERT3(offset + write_sz <= size, offset, write_sz, size);

        offset += write_sz;
    }

    ASSERT2(offset <= size, offset, size);

    exp_set_lit(exp, NULL);
    value_set_ptr(&exp->u_lit.val, raw, offset);
}

static void
make_elem_count(trans_t *trans, meta_t *meta, uint32_t reg_idx, uint32_t offset, uint32_t size)
{
    ast_exp_t *l_exp, *r_exp;

    l_exp = exp_new_mem(reg_idx, meta->rel_addr, offset);
    r_exp = exp_new_lit_int(size, meta->pos);

    /* TODO We need more precise measure for meta */
    if (meta_align(meta) == 8) {
        meta_set_int64(&l_exp->meta);
        meta_set_int64(&r_exp->meta);
    }
    else {
        meta_set_int32(&l_exp->meta);
        meta_set_int32(&r_exp->meta);
    }

    bb_add_stmt(trans->bb, stmt_new_assign(l_exp, r_exp, meta->pos));
}

static void
make_dynamic_init(trans_t *trans, ast_exp_t *exp)
{
    int i;
    meta_t *meta = &exp->meta;
    vector_t *elem_exps = exp->u_init.elem_exps;
    uint32_t reg_idx = meta->base_idx;
    uint32_t offset = meta->rel_offset;

    if (exp->u_init.is_topmost) {
        reg_idx = fn_add_register(trans->fn, meta);

        stmt_trans_malloc(trans, meta, reg_idx);
        exp_set_reg(exp, reg_idx);
    }

    if (is_array_meta(meta)) {
        ASSERT1(meta->dim_sizes[0] > 0, meta->dim_sizes[0]);
        ASSERT2(meta_align(meta) > 0, meta->type, meta_align(meta));

        make_elem_count(trans, meta, reg_idx, offset, meta->dim_sizes[0]);
        offset += meta_align(meta);
    }

    vector_foreach(elem_exps, i) {
        ast_exp_t *elem_exp = vector_get_exp(elem_exps, i);
        meta_t *elem_meta = &elem_exp->meta;
        ast_exp_t *l_exp, *r_exp;

        /* Only struct, or array which is a member of struct, is stored in separate memory. */
        if (is_struct_meta(elem_meta) || (is_struct_meta(meta) && is_tuple_meta(elem_meta))) {
            uint32_t tmp_idx = fn_add_register(trans->fn, elem_meta);

            stmt_trans_malloc(trans, elem_meta, tmp_idx);

            offset = ALIGN32(offset);

            elem_meta->base_idx = tmp_idx;
            elem_meta->rel_offset = 0;

            exp_trans(trans, elem_exp);

            l_exp = exp_new_mem(reg_idx, 0, offset);
            meta_set_int32(&l_exp->meta);

            r_exp = exp_new_reg(tmp_idx);
            meta_set_int32(&r_exp->meta);

            bb_add_stmt(trans->bb, stmt_new_assign(l_exp, r_exp, &exp->pos));

            offset += meta_regsz(elem_meta);
        }
        else if (is_map_meta(elem_meta)) {
            offset = ALIGN32(offset);

            make_map_init(trans, elem_exp);
            ASSERT1(is_reg_exp(elem_exp), elem_exp->kind);

            l_exp = exp_new_mem(reg_idx, 0, offset);
            meta_set_int32(&l_exp->meta);

            r_exp = exp_new_reg(elem_exp->meta.base_idx);
            meta_set_int32(&r_exp->meta);

            bb_add_stmt(trans->bb, stmt_new_assign(l_exp, r_exp, &exp->pos));

            offset += meta_regsz(elem_meta);
        }
        else {
            offset = ALIGN(offset, meta_align(elem_meta));

            elem_meta->base_idx = reg_idx;
            elem_meta->rel_offset = offset;

            exp_trans(trans, elem_exp);

            if (!is_init_exp(elem_exp)) {
                l_exp = exp_new_mem(reg_idx, 0, offset);
                meta_copy(&l_exp->meta, elem_meta);

                bb_add_stmt(trans->bb, stmt_new_assign(l_exp, elem_exp, &exp->pos));
            }

            offset += meta_memsz(elem_meta);
        }
    }
}

static void
exp_trans_init(trans_t *trans, ast_exp_t *exp)
{
    meta_t *meta = &exp->meta;

    ASSERT1(is_tuple_meta(meta) || is_struct_meta(meta) || is_map_meta(meta), meta->type);

    make_static_init(trans, exp);

    if (is_map_meta(meta))
        make_map_init(trans, exp);
    else if (!exp->u_init.is_static)
        make_dynamic_init(trans, exp);
}

static uint32_t
make_map_alloc(trans_t *trans, meta_t *meta)
{
    fn_kind_t kind;
    uint32_t reg_idx;
    ast_exp_t *l_exp, *r_exp;

    ASSERT1(is_map_meta(meta), meta->type);
    ASSERT1(meta->elem_cnt == 2, meta->elem_cnt);

    reg_idx = fn_add_register(trans->fn, meta);

    l_exp = exp_new_reg(reg_idx);
    meta_set_int32(&l_exp->meta);

    if (is_int64_meta(meta->elems[0]) && is_int64_meta(meta->elems[1]))
        kind = FN_MAP_NEW_I64_I64;
    else if (is_int64_meta(meta->elems[0]))
        kind = FN_MAP_NEW_I64_I32;
    else if (is_int64_meta(meta->elems[1]))
        kind = FN_MAP_NEW_I32_I64;
    else
        kind = FN_MAP_NEW_I32_I32;

    r_exp = exp_new_call(kind, NULL, NULL, meta->pos);
    meta_set_int32(&r_exp->meta);

    stmt_trans(trans, stmt_new_assign(l_exp, r_exp, meta->pos));

    return reg_idx;
}

static uint32_t
make_struct_alloc(trans_t *trans, meta_t *meta)
{
    int i;
    uint32_t offset = 0;
    uint32_t reg_idx;
    ast_exp_t *l_exp, *r_exp;

    ASSERT1(is_struct_meta(meta), meta->type);
    ASSERT1(meta->elem_cnt > 0, meta->elem_cnt);

    reg_idx = fn_add_register(trans->fn, meta);

    stmt_trans_malloc(trans, meta, reg_idx);

    for (i = 0; i < meta->elem_cnt; i++) {
        meta_t *elem_meta = meta->elems[i];

        if (is_array_meta(elem_meta) || is_struct_meta(elem_meta) || is_map_meta(elem_meta)) {
            uint32_t mem_idx;

            l_exp = exp_new_mem(reg_idx, 0, offset);
            meta_set_int32(&l_exp->meta);

            if (is_array_meta(elem_meta))
                mem_idx = make_array_alloc(trans, elem_meta);
            else if (is_struct_meta(elem_meta))
                mem_idx = make_struct_alloc(trans, elem_meta);
            else
                mem_idx = make_map_alloc(trans, elem_meta);

            r_exp = exp_new_reg(mem_idx);
            meta_set_int32(&r_exp->meta);

            bb_add_stmt(trans->bb, stmt_new_assign(l_exp, r_exp, elem_meta->pos));
        }

        offset += ALIGN(meta_regsz(elem_meta), meta_align(elem_meta));
    }

    return reg_idx;
}

static uint32_t
make_elem_alloc(trans_t *trans, meta_t *meta, uint32_t reg_idx, uint32_t offset, int dim_idx)
{
    int i;

    ASSERT2(dim_idx < meta->max_dim, dim_idx, meta->max_dim);
    ASSERT2(meta->dim_sizes[dim_idx] > 0, dim_idx, meta->dim_sizes[dim_idx]);
    ASSERT2(meta_align(meta) > 0, meta->type, meta_align(meta));

    /* the count of elements */
    make_elem_count(trans, meta, reg_idx, offset, meta->dim_sizes[dim_idx]);
    offset += meta_align(meta);

    if (dim_idx == meta->max_dim - 1) {
        if (is_struct_meta(meta) || is_map_meta(meta)) {
            uint32_t mem_idx;
            ast_exp_t *l_exp, *r_exp;

            for (i = 0; i < meta->dim_sizes[dim_idx]; i++) {
                l_exp = exp_new_mem(reg_idx, 0, offset);
                meta_set_int32(&l_exp->meta);

                if (is_struct_meta(meta))
                    mem_idx = make_struct_alloc(trans, meta);
                else
                    mem_idx = make_map_alloc(trans, meta);

                r_exp = exp_new_reg(mem_idx);
                meta_set_int32(&r_exp->meta);

                bb_add_stmt(trans->bb, stmt_new_assign(l_exp, r_exp, meta->pos));

                offset += ALIGN(meta_regsz(meta), meta_align(meta));
            }
        }
        else {
            offset += ALIGN(meta_regsz(meta), meta_align(meta)) * meta->dim_sizes[dim_idx];
        }
    }
    else {
        for (i = 0; i < meta->dim_sizes[dim_idx]; i++) {
            offset = make_elem_alloc(trans, meta, reg_idx, offset, dim_idx + 1);
        }
    }

    return offset;
}

static uint32_t
make_array_alloc(trans_t *trans, meta_t *meta)
{
    uint32_t reg_idx = fn_add_register(trans->fn, meta);

    ASSERT1(is_array_meta(meta), meta->type);

    stmt_trans_malloc(trans, meta, reg_idx);

    make_elem_alloc(trans, meta, reg_idx, 0, 0);

    return reg_idx;
}

static void
exp_trans_alloc(trans_t *trans, ast_exp_t *exp)
{
    meta_t *meta = &exp->meta;

    if (is_array_meta(meta))
        exp_set_reg(exp, make_array_alloc(trans, meta));
    else if (is_struct_meta(meta))
        exp_set_reg(exp, make_struct_alloc(trans, meta));
    else if (is_map_meta(meta))
        exp_set_reg(exp, make_map_alloc(trans, meta));
    else
        ASSERT1(!"invalid allocation type", meta->type);
}

void
exp_trans(trans_t *trans, ast_exp_t *exp)
{
    ASSERT(exp != NULL);

    switch (exp->kind) {
    case EXP_NULL:
        break;

    case EXP_LIT:
        exp_trans_lit(trans, exp);
        break;

    case EXP_ID:
        exp_trans_id(trans, exp);
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

    case EXP_ALLOC:
        exp_trans_alloc(trans, exp);
        break;

    case EXP_GLOBAL:
    case EXP_REG:
    case EXP_MEM:
        break;

    default:
        ASSERT1(!"invalid expression", exp->kind);
    }
}

/* end of trans_exp.c */
