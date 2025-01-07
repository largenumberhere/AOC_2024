#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <stdbool.h>
#include <string.h>
#include <assert.h>
#include <iostream>
#include <unordered_map>
#define STONE_T int64_t
#define MAX_KEYS 3875
#include "../include/hash.hpp"
typedef struct 
{
    // STONE_T stone_keys_arr[MAX_KEYS];
    // int64_t stone_counts_arr[MAX_KEYS];
    // int64_t length;
    std::unordered_map <STONE_T, int64_t> inner; 
} StoneWriter;


static void StoneWriter_clear(StoneWriter* sw) {
    sw->inner.clear();
    // memset(sw->stone_counts_arr, 0, sizeof(sw->stone_counts_arr));
    // memset(sw->stone_keys_arr, 0, sizeof(sw->stone_keys_arr));
}



static int64_t StoneWriter_count_all(StoneWriter *sw) {
    // for (auto ptr = sw->inner.begin(); ptr != sw->inner.end(); ptr ++) {
    //     int64_t item = *ptr;

    // }

    int64_t tally = 0;
    for (auto item : sw->inner) {
        tally+= item.second;
    }
    
    // for (int i = 0; i < sw->length; i++) {
    //     tally += (int64_t) sw->stone_counts_arr[i];
    // }    

    return tally;
}

static void StoneWriter_free(StoneWriter sw) {
    // freeing is implicit


}

static StoneWriter StoneWriter_new(void) {
    // std::multiset<int64_t> inner;
    std::unordered_map<STONE_T, int64_t> inner = std::unordered_map<STONE_T, int64_t>();
    
    StoneWriter sw {
        .inner = inner,
    };

    return sw;
}

typedef struct  {
    STONE_T stone;
    int64_t count;
} StoneCountInfo;

static int64_t StoneWriter_get_unique_stones_count(StoneWriter *sw) {
    int64_t count = 0;
    for (auto pair : sw->inner) {
        count += 1;
    }
    // int64_t count = 0;
    // for (int i = 0; i < sw->length; i++) {
    //     if (sw->stone_counts_arr[i] > 0) {
    //         count +=1;
    //     }
    // }

    return count;

} 

// static StoneCountInfo StoneWriter_get_stone_info(StoneWriter *sw, int64_t position) {
//     // int64_t i = 0;
//     // bool found = false;
//     // std::pair<int64_t, int64_t> values;
//     // for (auto val : sw->inner) {
//     //     if (i==position) {
//     //         found = true;
            
//     //     }
//     //     i+=1;
//     // }

//     // StoneCountInfo tuple = {
//     //     .stone = value,
//     //     .count = count
//     // };


//     // if (position >= sw->length) {
//     //     perror("out of range\n");
//     //     exit(1);
//     // }

//     // int64_t stone = unhash_stone(position);
//     // int64_t value = (int64_t) stone;
//     // int64_t count = sw->stone_counts_arr[position];

//     // StoneCountInfo tuple = {
//     //     .stone = value,
//     //     .count = count
//     // };

//     // return tuple;


// }


static void StoneWriter_put_stone(StoneWriter *sw, STONE_T value) {
    auto val = sw->inner.find(value);

    auto count = (*val).second;
    if (val == sw->inner.begin()) {
        count = 0;
    }

    
    count += 1;
    sw->inner.insert_or_assign(value, count);


    // int64_t hash = hash_stone(value);

    // if (hash >= sw->length) {
    //     perror("hash too big");
    //     exit(1);
    // }

    // if (sw->stone_keys_arr[hash] != value) {
    //     sw->stone_keys_arr[hash] = value;
    //     sw->stone_counts_arr[hash] = 1;
    // } else {
    //     sw->stone_counts_arr[hash] +=1;
    // }
    // sw_maybe_grow(sw);
    // for (int i = 0; i < sw->buffer_length; i++) {
    //     if (sw->stone_keys[i] == value) {
    //         sw->stone_counts[i] +=1;
    //         return;
    //     }    
    // }



    // sw->stone_keys[sw->buffer_length] = value;
    // sw->stone_counts[sw->buffer_length] = 1;
    // sw->buffer_length+=1;
}


static void StoneWriter_put_array(StoneWriter *sw, STONE_T* values, int values_len) {
    for (int i = 0; i < values_len; i++) {
        StoneWriter_put_stone(sw, values[i]);
    }
}

static void StoneWriter_put_stone_count(StoneWriter *sw, STONE_T value, int64_t count) {
    if (count <= 0) {
        return;
    }

    auto val = sw->inner.find(value);
    auto new_count = (*val).second + count;
    if (val == sw->inner.begin()) {
        new_count = 0;
    }

    
    sw->inner.insert_or_assign(value, new_count);


    
    // int64_t hash = hash_stone(value);
    // sw->stone_counts_arr[hash]+= count;

    
    // if (count <= 0) {
    //     return;
    // }

    // // ensure stone is initialized
    // StoneWriter_put_stone(sw, value);
    // count -=1;


    // if ((count) > 0) {
    //     // get position of stone
    //     int64_t pos = -1;
    //     for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
    //         if (sw->stone_keys[i] == value) {
    //             pos = i;
    //             break;
    //         }
    //     }

    //     if (pos == -1) {
    //         perror("Failed to find stone\n");
    //         exit(1);
    //     }

    //     sw->stone_counts[pos] += (count);
    // }
}

static int64_t StoneWriter_get_unique_stone_count(StoneWriter *sw, STONE_T value) {
    // for (auto ptr = sw->inner.begin(); ptr!= sw->inner.end(); ptr++) {
    //     auto item = *ptr;
    return (*sw->inner.find(value)).second;

    // }
    
    // int64_t hash = hash_stone(value);
    // return sw->stone_counts_arr[hash];

    // // get position of stone
    // int64_t pos = -1;
    // for (int64_t i = 0; i < StoneWriter_count_all(sw); i++) {
    //     if (sw->stone_keys[i] == value) {
    //         pos = i;
    //         break;
    //     }
    // }

    // if (pos == -1) {
    //     return 0;
    // }

    // int64_t count = sw->stone_counts[pos];

    // return count;
}

// make a deep copy of the data
static void StoneWriter_copy(StoneWriter *destination, StoneWriter* source) {
    // copy over all items
    // for (size_t i = 0; i < StoneWriter_get_unique_stones_count(source) ; i++) {
    //     StoneCountInfo info = StoneWriter_get_stone_info(source, i);
    //     StoneWriter_put_stone_count(destination, info.stone, info.count);
    // }
    
    // memcpy(destination, source, sizeof(StoneWriter));


    destination->inner.clear();
    for (auto item : source->inner) {
        destination->inner.insert_or_assign(item.first, item.second);
    }

    // std::copy(*source, *destination);


    // return;
}

// duplicate the stone writer
// static StoneWriter StoneWriter_dup(StoneWriter* source) {
//     StoneWriter sw = StoneWriter_new();
//     StoneWriter_copy(&sw, source);
//     return sw;
// }

typedef struct {
    StoneWriter *sw;
    std::unordered_map<STONE_T, int64_t>::iterator pos;
    std::unordered_map<STONE_T, int64_t>::iterator end;
} StoneIter;

static StoneIter StoneWriter_iter(StoneWriter* sw) {
    auto pos = sw->inner.begin();
    auto end = sw->inner.end();

    StoneIter iter = {
        .sw=sw,
        .pos = pos,
        .end = end
    };

    return iter;
}

static bool StoneWriterIter_has_more(StoneIter iter) {
    auto next = iter.pos++;

    return next != iter.end;
}

static StoneCountInfo StoneWriterIter_next(StoneIter *iter) {
    
    
    auto current = *iter->pos;
    iter->pos++;

    
    StoneCountInfo info;
    info.stone = current.first;
    info.count = current.second;
    
    return info;
    // StoneCountInfo info = StoneWriter_get_stone_info(iter->sw, iter->pos);
    // iter->pos+=1;
    // return info;
}

static void StoneWrier_replace_stone(StoneWriter* sw, STONE_T from ,STONE_T to, int64_t count) {
    auto from_item = sw->inner.find(from);
    if (from_item == sw->inner.begin()) {
        perror("attempt to replace unitialized item");
        exit(1);
    }
    
    int64_t new_from = (*from_item).second - count;
    if (new_from <0) {
        perror("negative from!");
        exit(1);
    }

    auto to_item = sw->inner.find(to);
    int64_t new_to = (*to_item).second + count;
    if (to_item == sw->inner.begin()) {
        new_to = 0;
    }



    

    
    sw->inner.insert_or_assign(to, new_to);
    sw->inner.insert_or_assign(from, new_from);
    


    // if (count == 0) {
    //     return;
    // }
    
    // int64_t to_hash = hash_stone(to);
    // int64_t from_hash = hash_stone(from);

    // if ((sw->stone_counts_arr[from_hash] - count) < 0) {

    //     printf("replacing %lli with %lli, count %lli\n", from, to, count);
    //     printf("count of %lli: %lli\n", from, sw->stone_counts_arr[from_hash]);
    //     printf("negative stones!\n");
    //     exit(0);
    // }
    // sw->stone_counts_arr[from_hash] -=count;
    // sw->stone_counts_arr[to_hash] += count;
    // // if ((sw->stone_counts_arr[to_hash] + count) < 0) {
        
    // // }
    

    // // update the from stone
    // bool from_updated = false;
    // for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
    //     StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
        
    //     // update the from stone
    //     if (info.stone == from) {
    //         sw->stone_counts[i] -= count;      
    //         if (sw->stone_counts[i] < 0) {
    //             perror("negative stones");
    //             exit(1);
    //         }
    //         from_updated = true;
    //         break;
    //     } 
    // }

    // if (!from_updated) {
    //     perror("from not upated :(");
    //     exit(1);
    // }

    // // update the to stone
    // StoneWriter_put_stone_count(sw, to, count);
}



static void StoneWriter_print(StoneWriter *sw) {
    // for (auto item : sw->inner) {
    //     std::cout << "  " << item.first << ":" << item.second << "\n";
    // }
    // std::cout << "\n";

    // printf("Stones {\n");
    // printf("  (len)%i\n", sw->buffer_length);
    // printf("  (cap)%i\n", sw->buffer_capacity);
    // for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
    //     StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
    //     if (info.count !=0) {     
    //         printf("    %lli: %lli\n", info.stone, info.count);
    //     }
    // }
    // printf("}\n");

    // printf("Stones {\n");
    // for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
    //     StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
    //     if (info.count !=0) {
    //         printf("    %lli: %lli\n", info.stone, info.count);
    //     }
    // }

    StoneIter iter = StoneWriter_iter(sw);
    while (StoneWriterIter_has_more(iter)) {
        StoneCountInfo info = StoneWriterIter_next(&iter);
        if (info.count !=0) {
            printf("    %lli: %lli\n", info.stone, info.count);
        }
    }
    

    printf("}\n");

}

static void Stonewriter_print_keys(StoneWriter *sw) {
    printf("Stone keys: { \n");
    StoneIter iter = StoneWriter_iter(sw);
    int i = 0;
    while (StoneWriterIter_has_more(iter)) {
        // StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
        StoneCountInfo info = StoneWriterIter_next(&iter);
        printf("case %i: stone=%lli; break;\n", i, info.stone);
        i++;
    }
 
    // for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
    //     StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
    //     printf("case %i: stone=%lli; break;\n", i, info.stone);
    // }
    printf("\n");
}

// int64_t pow(int64_t a, int64_t b) {
//     for (int64_t i = 0; i < b; i++) {
//         a *= b;
//     }

//     return a;
// }

static STONE_T right_half(STONE_T value, int length) {
    STONE_T to = 0;
    STONE_T mul = 0;
    for (int i = 0; (i < (length/2)) && (value!=0) ;i++) {
        // pop off last digit
        STONE_T digit = value % 10;
        value /= 10;
        to = to + (digit * pow(10, mul));
        mul ++;

    }

    return to;
}

static int count_digits(STONE_T value) {
    int size = 0;
    while (value !=0) {
        size++;
        value /=10;
    }

    return size;
}


static STONE_T left_half(STONE_T value, int length) {
   STONE_T vin = value; 
    // discard first half
    for (int i = 0; i < (length/2); i++) {
        value /=10;
    }

   
    STONE_T to = 0;
    STONE_T mul = 0;
    for (int i = 0; (i < (length/2)) && (value!=0) ;i++) {
        // pop off last digit
        STONE_T digit = value % 10;
        value /= 10;
        to = to + (digit * (size_t)pow(10, mul));
        mul ++;

    }

    return to;
}

typedef struct {
    STONE_T one;
    STONE_T two;
    bool has_two;   // if false, two was not used 
} StoneResult;

static StoneResult eval_stone(STONE_T stone) {
    StoneResult output;
    memset(&output, 0, sizeof(StoneResult));
    
    if (stone == 0) {
        output.has_two = false;
        output.two = -1;
        output.one = 1;
    } else if ((count_digits(stone)%2) == 0) {
        int digits = count_digits(stone);

        STONE_T right = right_half(stone, digits);
        STONE_T left = left_half(stone, digits);

        output.one = left;
        output.two = right;
        output.has_two = true;
    } else {
        output.one = stone * 2024;
        output.two = -1;
        output.has_two = false;
    }

    return output;
}

static void split_stones(StoneWriter *sw, int max_depth, StoneWriter* tmp) {
    // StoneWriter tmp = StoneWriter_dup(sw);
    StoneWriter_clear(tmp);
    StoneWriter_copy(tmp, sw);

    // for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(tmp); i++) {
    StoneIter iter = StoneWriter_iter(tmp);
    while(StoneWriterIter_has_more(iter)) {
        // StoneCountInfo stone_info = StoneWriter_get_stone_info(tmp, i);
        StoneCountInfo stone_info = StoneWriterIter_next(&iter);

        StoneResult stone_result = eval_stone(stone_info.stone);
        StoneWrier_replace_stone(sw, stone_info.stone, stone_result.one, stone_info.count);

        // add seccondary stone
        if (stone_result.has_two) {
            StoneWriter_put_stone_count(sw, stone_result.two, stone_info.count);
        }
    }

    // StoneWriter_free(tmp);
}

// #define i_implement
// #define i_type Stones
// #define i_key int
// #define i_val int
// #include "stc/cmap.h"
// Stones idnames = {0};


int main(void) {
    


    // exit(0);
    StoneWriter sw = StoneWriter_new();
    
    STONE_T stones[] = {5910927,0,1,47,261223,94788,545,7771};
    // STONE_T stones[] = {125, 17};
    const int stones_len = sizeof(stones) / 8;

    StoneWriter_put_array(&sw, stones, stones_len);
    // StoneWriter_print(&sw);
    // StoneWriter_put_stone(&sw, 17);
    // StoneWrier_replace_stone(&sw, 17, 2, 2);
    
    StoneWriter_print(&sw);



    StoneWriter temporary = StoneWriter_new();
    int max_depth = 75;
    for (int i = 0; i < max_depth; i++) {
        split_stones(&sw, max_depth, &temporary);
    }
    int64_t count = StoneWriter_count_all(&sw);
 
    // StoneWriter_print(&sw);    
    // Stonewriter_print_keys(&sw);

    printf("count = %lli\n", count);
    int64_t unique = StoneWriter_get_unique_stones_count(&sw);
    printf("unique = %lli\n", unique);
    StoneWriter_free(sw);
}

