#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <stdbool.h>
#include <string.h>
#include <assert.h>

#define STONE_T int64_t

typedef struct 
{
    STONE_T *stone_keys;
    int64_t* stone_counts;
    int64_t buffer_length;
    int64_t buffer_capacity;
} StoneWriter;


// default buffer size in sizeof(STONE_T)'s
#define BUFFER_MIN 10 * 1024 * 1024  

static void sw_maybe_grow(StoneWriter* sw) {
    if (sw->buffer_capacity == 0) {
        void* counts = realloc(sw->stone_counts, BUFFER_MIN * sizeof(STONE_T));
        void* keys = realloc(sw->stone_keys, BUFFER_MIN * sizeof(STONE_T));

        if ((counts == NULL) || (keys == NULL)) {
            perror("memory failure\n"); exit(1);
        }

        sw->stone_keys = keys;
        sw->stone_counts = counts;
        sw->buffer_capacity = BUFFER_MIN;
    }
    else if (sw->buffer_length +1 > sw->buffer_capacity) {
        int new_cap = sw->buffer_capacity * 2;
        void* counts = realloc(sw->stone_counts, new_cap * sizeof(STONE_T));
        void* keys = realloc(sw->stone_keys, new_cap * sizeof(STONE_T));
        if ((counts == NULL) || (keys == NULL)) {
            perror("memory failure\n"); exit(1);
        }
        sw->buffer_capacity = new_cap;
        sw->stone_keys = keys;
        sw->stone_counts = counts;
    }
 
}

// maybe broken
static int64_t StoneWriter_count_all(StoneWriter *sw) {
    int64_t tally = 0;
    for (int i = 0; i < sw->buffer_length; i++) {
        tally += (int64_t) sw->stone_counts[i];
    }    

    return tally;
}
static void StoneWriter_free(StoneWriter** sw_ptr_ptr) {
    if (sw_ptr_ptr == NULL) {
        return;
    }
    
    StoneWriter* sw_ptr = *sw_ptr_ptr;
    void * mem1 = (void*) sw_ptr->stone_counts;
    free(mem1);
    void * mem2 = (void*) sw_ptr->stone_keys;
    free(mem2);
    
    memset(sw_ptr, 0, sizeof(StoneWriter));
    sw_ptr = NULL;
}

static StoneWriter StoneWriter_new(void) {
    StoneWriter sw;
    memset(&sw, 0,sizeof(StoneWriter));
    sw.buffer_capacity = 0;
    sw.buffer_length = 0;
    sw.stone_counts = NULL;
    sw.stone_keys = NULL;

    return sw;
}

// static size_t sw_offset_at(StoneWriter *sw, int64_t position) {
//     size_t offset = 0;
//     for (size_t i = 0; i < sw->buffer_length; i++) {
//         offset += i;

//         if (offset)
//     }
// }

// broken
// a crude analogue for iteration
// static STONE_T StoneWriter_get_at(StoneWriter *sw, int64_t position) {
//     // size_t offset = sw_offset_at(sw, position);

//     // if (position == 0) {
//     //     assert(StoneWriter_count(sw) > 0);
//     //     return sw->stone_keys[0];
//     // }

//     position +=1;

//     size_t offset = 0;
//     for (size_t i = 0; i < sw->buffer_length; i++) {
//         offset += i;
//         if (offset > position) {
//             break;
//         }
//     }

//     printf("offset: %lli. count %lli\n", offset, StoneWriter_count(sw));
//     assert(offset < sw->buffer_length);

//     return (sw->stone_keys[offset]);
// }

typedef struct  {
    STONE_T stone;
    int64_t count;
} StoneCountInfo;

static int64_t StoneWriter_get_unique_stones_count(StoneWriter *sw) {
    return sw->buffer_length;
} 

static StoneCountInfo StoneWriter_get_stone_count(StoneWriter *sw, int64_t position) {
    if (position >= sw->buffer_length) {
        perror("out of range\n");
        exit(1);
    }

    STONE_T value = sw->stone_keys[position];
    int64_t count = sw->stone_counts[position];

    StoneCountInfo tuple = {
        .stone = value,
        .count = count
    };

    return tuple;
}


static void StoneWriter_put_stone(StoneWriter *sw, STONE_T value) {
    sw_maybe_grow(sw);
    for (int i = 0; i < sw->buffer_length; i++) {
        if (sw->stone_keys[i] == value) {
            sw->stone_counts[i] +=1;
            return;
        }    
    }

    sw->stone_keys[sw->buffer_length] = value;
    sw->stone_counts[sw->buffer_length] = 1;
    sw->buffer_length+=1;
}


static void StoneWriter_put_array(StoneWriter *sw, STONE_T* values, int values_len) {
    for (int i = 0; i < values_len; i++) {
        StoneWriter_put_stone(sw, values[i]);
    }
}

static void StoneWriter_put_stone_count(StoneWriter *sw, STONE_T value, int64_t count) {
    if (count <= 0) {
        // perror("zero stones");
        // exit(1);
        return;
    }

    // ensure stone is initialized
    StoneWriter_put_stone(sw, value);
    count -=1;


    if ((count) > 0) {
        // get position of stone
        int64_t pos = -1;
        for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
            if (sw->stone_keys[i] == value) {
                pos = i;
                break;
            }
        }

        if (pos == -1) {
            perror("Failed to find stone\n");
            exit(1);
        }

        sw->stone_counts[pos] += (count);
    }
}

static int64_t StoneWriter_get_unique_stone_count(StoneWriter *sw, STONE_T value) {
    // get position of stone
    int64_t pos = -1;
    for (int64_t i = 0; i < StoneWriter_count_all(sw); i++) {
        if (sw->stone_keys[i] == value) {
            pos = i;
            break;
        }
    }

    if (pos == -1) {
        return 0;
    }
    // if (pos == -1) {
    //     perror("Failed to find stone\n");
    //     exit(1);
    // }

    int64_t count = sw->stone_counts[pos];

    return count;
}

// make a deep copy of the data
static void StoneWriter_copy(StoneWriter *destination, StoneWriter* source) {
    // copy over all items
    for (size_t i = 0; i < StoneWriter_get_unique_stones_count(source) ; i++) {
        StoneCountInfo info = StoneWriter_get_stone_count(source, i);
        StoneWriter_put_stone_count(destination, info.stone, info.count);
    }

    return;
}

static void StoneWrier_replace_stone(StoneWriter* sw, STONE_T from ,STONE_T to, int64_t count) {
    // update the from stone
    bool from_updated = false;
    for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
        StoneCountInfo info = StoneWriter_get_stone_count(sw, i);
        
        // update the from stone
        if (info.stone == from) {
            sw->stone_counts[i] -= count;      
            if (sw->stone_counts[i] < 0) {
                perror("negative stones");
                exit(1);
            }
            from_updated = true;
            break;
        } 
    }

    if (!from_updated) {
        perror("from not upated :(");
        exit(1);
    }

    // update the to stone
    StoneWriter_put_stone_count(sw, to, count);
    // bool to_updated = false;
    // for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
    //     StoneCountInfo info = StoneWriter_get_stone_count(sw, i);

    //     if (info.stone == to) {
    //         sw->stone_counts[i] += count;
    //         to_updated = true;
    //         break;
    //     }
    // }

    // if (!to_updated) {
    //     perror("to not upated :(");
    //     exit(1);
    // }
}

// broken
// static void StoneWriter_replace(StoneWriter *sw, STONE_T value, int64_t pos) {
//     printf("replacing at pos %lli in count %lli", pos, StoneWriter_count_all(sw));
    
//     // get offset in the buffers of the item at position `pos`
    
//     pos+=1;
//     int64_t offset = -1;
//     int64_t tally = 0;

//     for (int i = 0; i < StoneWriter_count_all(sw); i++) {
//         tally += sw->stone_counts[i];
//         if (tally > pos) {
//             offset = i;
//             break; 
//         }
//     }


//     assert(offset!=-1);
    
//     STONE_T* key =  &(sw->stone_keys[offset]);
//     int* count = &(sw->stone_counts[offset]);

//     // printf("key %i, count %i\n", *key, *count);

//     if ((*count) >= 1) {
//         // decrease count and put new anywhre it fits
//         *count -=1;
//         StoneWriter_put_stone(sw, value);
//         return;
//     } else if ((*count) == 0) {
//         (*key) = value;
//         *count = 1;

//         return;
//     } 
    
//     printf("%lli\n", value);
//     perror("invalid count");
//     exit(1);
// }

static void StoneWriter_print(StoneWriter *sw) {
    printf("Stones {\n");
    printf("  (len)%i\n", sw->buffer_length);
    printf("  (cap)%i\n", sw->buffer_capacity);
    for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
        StoneCountInfo info = StoneWriter_get_stone_count(sw, i);
        if (info.count !=0) {     
            printf("    %lli: %lli\n", info.stone, info.count);
        }
    }
    printf("}\n");
}



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

// static void split_stones(StoneWriter *sw) {
//     int64_t count = StoneWriter_count_all(sw);

//     // make arrray for the results
//     StoneResult* results = calloc(sizeof(*results), count);
//     if (results == NULL) {
//         perror("failed to allocate\n");
//         exit(1);
//     }

//     // calculate results
//     for (int64_t i = 0; i < count; i++) {
//         results[i] = handle_stone(StoneWriter_get_at(sw, i));
//     }

//     // update stones with results
//     for (int i = 0; i < count; i++) {
//         StoneWriter_replace(sw, results[i].one, i);
//     }
//     for (int i = 0; i < count; i++) {
//         if (results[i].has_two) {
//             StoneWriter_put_stone(sw, results[i].two);
//         }
//     }

//     // discard temporary array
//     free(results);
// }

// static int64_t recurse_stone(StoneWriter *sw, int depth, int max_depth, STONE_T starting_stone) {
//     printf("Stone %lli Depth %i. Stone count %i\n",starting_stone,depth, StoneWriter_count(sw));
//     if (depth >= max_depth) {
//         return StoneWriter_count(sw);
//     } else {
//        split_stones(sw);
//     //    int64_t new_len  = (int64_t) StoneWriter_count(sw);
//     //    for (int i = 0; i < new_len; i++) {
//             // create new array for next level down
//             StoneWriter next =  StoneWriter_new();
//             StoneWriter* next_ptr = &next;
//             StoneWriter_copy(next_ptr, sw);

//             int64_t len2 = recurse_stone(next_ptr, depth+1, max_depth, starting_stone);
            
//             StoneWriter_free(&next_ptr);

//             return len2;
//     //    }
//     }

//     perror("unrerachable\n");
// }
// static int64_t iterate_stone(STONE_T stone, int max_depth) {
//     StoneWriter sw = StoneWriter_new();
    
//     StoneWriter* sw_ptr = &sw;
//     StoneWriter_put(sw_ptr, stone);
//     int64_t value = recurse_stone(sw_ptr, 0, max_depth, stone);
    
//     StoneWriter_free(&sw_ptr);

//     return value;
// } 



// static int64_t iterate_stone(STONE_T stone, int max_depth) {
//     StoneWriter sw = StoneWriter_new();
//     StoneWriter_put_stone(&sw, stone);

//     for (int i = 0; i < max_depth; i++) {
//         int64_t current_stones = StoneWriter_count_all(&sw);

//         for (int64_t pos = current_stones-1; pos >=0; pos --) {
//             STONE_T stone = StoneWriter_get_at(&sw, pos);
            
//             StoneResult res = handle_stone(stone);
//             StoneWriter_replace(&sw, res.one, pos);
//             if (res.has_two) {
//                 StoneWriter_put_stone(&sw, res.two);
//             }
//         }
//     }

//     int64_t count = StoneWriter_count_all(&sw);
//     StoneWriter* sw_ptr = &sw;
//     StoneWriter_free(&sw_ptr);

//     return count;
// }


static void iterate_stones(StoneWriter *sw, int max_depth) {
    // int64_t tally = 0;
    StoneWriter tmp = StoneWriter_new(); 
    StoneWriter* tmp_ptr = &tmp;
    StoneWriter_copy(tmp_ptr, sw);
    // StoneWriter_print(tmp_ptr);

    
    for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(tmp_ptr); i++) {
        StoneCountInfo stone_info = StoneWriter_get_stone_count(tmp_ptr, i);
        // printf("stone %lli has count: %i\n", stone_info.stone, stone_info.count);

        StoneResult stone_result = eval_stone(stone_info.stone);
        // printf("stone %lli evaluated to %lli and %lli\n", stone_info.stone, stone_result.one, stone_result.two);


        // replace primary stone
        // printf("replacing %i instances of %lli with %i\n", stone_info.count, stone_info.stone, stone_result.one);
        StoneWrier_replace_stone(sw, stone_info.stone, stone_result.one, stone_info.count);

        // add seccondary stone
        if (stone_result.has_two) {
            StoneWriter_put_stone_count(sw, stone_result.two, stone_info.count);
        }
    }

    StoneWriter_free(&tmp_ptr);
    // StoneWriter_print(sw);
    // _sleep(10000);
    

    // int64_t all = StoneWriter_count_all(sw);
    // return all;
}

int main(void) {
    StoneWriter sw = StoneWriter_new();
    StoneWriter* sw_ptr = &sw;
   
    STONE_T stones[] = {5910927,0,1,47,261223,94788,545,7771};
    // STONE_T stones[] = {125,17};
    const int stones_len = sizeof(stones) / 8;

    StoneWriter_put_array(sw_ptr, stones, stones_len);
    StoneWriter_print(sw_ptr);

    int max_depth = 75;
    for (int i = 0; i < max_depth; i++) {
        iterate_stones(sw_ptr, max_depth);
    }
    int64_t count = StoneWriter_count_all(&sw);
 
    StoneWriter_print(&sw);    
    printf("count = %lli\n", count);
    StoneWriter_free(&sw_ptr);
}

