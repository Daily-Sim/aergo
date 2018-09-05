/**
 * @file    strbuf.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "util.h"

#include "strbuf.h"

void
strbuf_init(strbuf_t *sb)
{
    sb->size = STRBUF_INIT_SIZE;
    sb->offset = 0;
    sb->buf = xmalloc(sb->size + 1);
    sb->buf[0] = '\0';
}

void
strbuf_reset(strbuf_t *sb)
{
    sb->offset = 0;
    sb->buf[0] = '\0';
}

void
strbuf_append_str(strbuf_t *sb, char *str, int str_len)
{
    if (sb->offset + str_len > sb->size) {
        sb->size += max(sb->size, str_len);
        sb->buf = realloc(sb->buf, sb->size + 1);
    }

    memcpy(sb->buf + sb->offset, str, str_len);

    sb->offset += str_len;
    sb->buf[sb->offset] = '\0';
}

void
strbuf_append_char(strbuf_t *sb, char c)
{
    if (sb->offset + 1 > sb->size) {
        sb->size *= 2;
        sb->buf = realloc(sb->buf, sb->size + 1);
    }

    sb->buf[sb->offset++] = c;
    sb->buf[sb->offset] = '\0';
}


void
strbuf_copy(strbuf_t *src, strbuf_t *dest)
{
    strbuf_reset(dest);
    strbuf_append_str(dest, src->buf, src->offset);
}

/* end of strbuf.c */
