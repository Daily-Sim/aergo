/**
 * @file    gen_bb.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "gen_id.h"
#include "gen_stmt.h"
#include "gen_exp.h"
#include "gen_util.h"

#include "gen_bb.h"

void
bb_gen(gen_t *gen, ir_bb_t *bb)
{
    int i;
    int instr_cnt = gen->instr_cnt;
    BinaryenExpressionRef *instrs = gen->instrs;
    BinaryenExpressionRef body;

    gen->instr_cnt = 0;
    gen->instrs = NULL;

    for (i = 0; i < array_size(&bb->ids); i++) {
        gen_add_instr(gen, id_gen(gen, array_get_id(&bb->ids, i)));
    }

    for (i = 0; i < array_size(&bb->stmts); i++) {
        gen_add_instr(gen, stmt_gen(gen, array_get_stmt(&bb->stmts, i)));
    }

    body = BinaryenBlock(gen->module, NULL, gen->instrs, gen->instr_cnt,
                         BinaryenTypeNone());

    bb->rb = RelooperAddBlock(gen->relooper, body);

    gen->instr_cnt = instr_cnt;
    gen->instrs = instrs;
}

void
br_gen(gen_t *gen, ir_bb_t *bb)
{
    int i;

    ASSERT(bb->rb != NULL);

    for (i = 0; i < array_size(&bb->brs); i++) {
        ir_br_t *br = array_get_br(&bb->brs, i);
        BinaryenExpressionRef cond = NULL;

        ASSERT(br->bb->rb != NULL);

        if (br->cond_exp != NULL)
            cond = exp_gen(gen, br->cond_exp, &br->cond_exp->meta, false);

        RelooperAddBranch(bb->rb, br->bb->rb, cond, NULL);
    }
}

/* end of gen_bb.c */