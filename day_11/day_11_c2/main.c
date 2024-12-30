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

static void StoneWriter_clear(StoneWriter* sw) {
    sw->buffer_length = 0;
}

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

static int64_t StoneWriter_count_all(StoneWriter *sw) {
    int64_t tally = 0;
    for (int i = 0; i < sw->buffer_length; i++) {
        tally += (int64_t) sw->stone_counts[i];
    }    

    return tally;
}

static void StoneWriter_free(StoneWriter sw) {
    void *mem1 = (void*) sw.stone_counts;
    free(mem1);
    void* mem2 = (void*) sw.stone_keys;
    free(mem2);
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

typedef struct  {
    STONE_T stone;
    int64_t count;
} StoneCountInfo;

static int64_t StoneWriter_get_unique_stones_count(StoneWriter *sw) {
    return sw->buffer_length;
} 

static StoneCountInfo StoneWriter_get_stone_info(StoneWriter *sw, int64_t position) {
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

    int64_t count = sw->stone_counts[pos];

    return count;
}

// make a deep copy of the data
static void StoneWriter_copy(StoneWriter *destination, StoneWriter* source) {
    // copy over all items
    for (size_t i = 0; i < StoneWriter_get_unique_stones_count(source) ; i++) {
        StoneCountInfo info = StoneWriter_get_stone_info(source, i);
        StoneWriter_put_stone_count(destination, info.stone, info.count);
    }

    return;
}

// duplicate the stone writer
static StoneWriter StoneWriter_dup(StoneWriter* source) {
    StoneWriter sw = StoneWriter_new();
    StoneWriter_copy(&sw, source);
    return sw;
}

static void StoneWrier_replace_stone(StoneWriter* sw, STONE_T from ,STONE_T to, int64_t count) {
    // update the from stone
    bool from_updated = false;
    for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
        StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
        
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
}



static void StoneWriter_print(StoneWriter *sw) {
    printf("Stones {\n");
    printf("  (len)%i\n", sw->buffer_length);
    printf("  (cap)%i\n", sw->buffer_capacity);
    for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(sw); i++) {
        StoneCountInfo info = StoneWriter_get_stone_info(sw, i);
        if (info.count !=0) {     
            printf("    %lli: %lli\n", info.stone, info.count);
        }
    }
    printf("}\n");
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

    for (int64_t i = 0; i < StoneWriter_get_unique_stones_count(tmp); i++) {
        StoneCountInfo stone_info = StoneWriter_get_stone_info(tmp, i);

        StoneResult stone_result = eval_stone(stone_info.stone);
        StoneWrier_replace_stone(sw, stone_info.stone, stone_result.one, stone_info.count);

        // add seccondary stone
        if (stone_result.has_two) {
            StoneWriter_put_stone_count(sw, stone_result.two, stone_info.count);
        }
    }

    // StoneWriter_free(tmp);
}

int main(void) {
    StoneWriter sw = StoneWriter_new();
   
    STONE_T stones[] = {125,17};
    const int stones_len = sizeof(stones) / 8;

    StoneWriter_put_array(&sw, stones, stones_len);
    StoneWriter_print(&sw);

    StoneWriter temporary = StoneWriter_new();
    int max_depth = 75;
    for (int i = 0; i < max_depth; i++) {
        split_stones(&sw, max_depth, &temporary);
    }
    int64_t count = StoneWriter_count_all(&sw);
 
    StoneWriter_print(&sw);    
    printf("count = %lli\n", count);
    StoneWriter_free(sw);
}

