/**
 * @file    ir_sgmt.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _IR_SGMT_H
#define _IR_SGMT_H

#include "common.h"

#define SGMT_INIT_CAPACITY          10

#ifndef _IR_SGMT_T
#define _IR_SGMT_T
typedef struct ir_sgmt_s ir_sgmt_t;
#endif /* ! _IR_SGMT_T */

struct ir_sgmt_s {
    int cap;
    int size;
    uint32_t offset;

    uint32_t *lens;
    uint32_t *addrs;
    char **datas;
};

int sgmt_add_global(ir_sgmt_t *sgmt, type_t type);
int sgmt_add_raw(ir_sgmt_t *sgmt, void *ptr, uint32_t len);

static inline void
sgmt_init(ir_sgmt_t *sgmt)
{
    sgmt->cap = SGMT_INIT_CAPACITY;
    sgmt->size = 0;
    sgmt->offset = 0;

    sgmt->lens = xmalloc(sizeof(uint32_t) * sgmt->cap);
    sgmt->addrs = xmalloc(sizeof(uint32_t) * sgmt->cap);
    sgmt->datas = xmalloc(sizeof(char *) * sgmt->cap);
}

static inline ir_sgmt_t *
sgmt_new(void)
{
    ir_sgmt_t *sgmt = xmalloc(sizeof(ir_sgmt_t));

    sgmt_init(sgmt);

    return sgmt;
}

#endif /* ! _IR_SGMT_H */
