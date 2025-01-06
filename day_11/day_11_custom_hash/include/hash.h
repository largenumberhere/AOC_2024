#ifndef HASH_H
#define HASH_H

#include <stdint.h>
#define STONE_T int64_t


int64_t hash_stone(STONE_T stone);
STONE_T unhash_stone(int64_t hash);

#endif