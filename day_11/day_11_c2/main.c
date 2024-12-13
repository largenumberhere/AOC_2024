#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <stdbool.h>
#include <string.h>

#define STONE_T int64_t

typedef struct 
{
    STONE_T *stones;
    int len;
    int cap;
} StoneWriter;


// default buffer size in sizeof(STONE_T)'s
#define BUFFER_MIN 10 * 1024 * 1024  

static void sw_maybe_grow(StoneWriter* sw) {
    if (sw->stones == NULL) {
        // first allocation
        int new_cap = BUFFER_MIN;
        void * m = malloc(sizeof(STONE_T) * (size_t) new_cap);
        if (m == NULL) {
            perror("Failed first allocation\n");
            exit(1);
        }

        sw->stones = (STONE_T*) m;
        sw->cap = new_cap;
        sw->len = 0;
    }
    else if (sw->len + 1 > sw->cap) {
        int new_cap = sw->cap * 2;
        void *m = realloc(sw->stones, sizeof(STONE_T) * (size_t) new_cap);
        if (m == NULL) {
            perror("Failed re-allocation\n");
            exit(1);
        }

        sw->cap = new_cap;
        sw->stones = (STONE_T*)m;
    }

    
}

void StoneWriter_free(StoneWriter** sw_ptr_ptr) {
    StoneWriter* sw_ptr = *sw_ptr_ptr;

    free(sw_ptr->stones);
    sw_ptr->cap = 0;
    sw_ptr->len = 0;

    *sw_ptr_ptr = NULL;
}

static StoneWriter StoneWriter_new(void) {
    StoneWriter writer;
    writer.cap = 0;
    writer.len = 0;
    writer.stones = NULL;
    return writer;
}

static STONE_T StoneWriter_get_at(StoneWriter *sw, int position) {
    return sw->stones[position];
}


static int StoneWriter_len(StoneWriter *sw) {
    return sw->len;
}

static void StoneWriter_put(StoneWriter *sw, STONE_T value) {
    sw_maybe_grow(sw);
    sw->stones[sw->len++] = value;
}

static void StoneWriter_put_many(StoneWriter *sw, STONE_T* values, int values_len) {
    for (int i = 0; i < values_len; i++) {
        StoneWriter_put(sw, values[i]);
    }
}

static void StoneWriter_copy(StoneWriter *destination, StoneWriter* source) {
    for (int i = 0; i < StoneWriter_len(source); i++) {
        StoneWriter_put(destination, StoneWriter_get_at(source, i));
    }
}

static void StoneWriter_replace(StoneWriter *sw, STONE_T value, int pos) {
    if (pos >= sw->len) {
        perror("replacement out of bounds\n");
    }

    sw->stones[pos] = value;
}

static void StoneWriter_print(StoneWriter *sw) {
    STONE_T *stones = sw->stones;
    printf("Stones {\n");
    for (int i = 0; i < sw->len; i++) {
        printf("    %lli\n",stones[i]);
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
        to = to + (digit * pow(10, mul));
        mul ++;

    }

    return to;
}

typedef struct {
    STONE_T one;
    STONE_T two;
    bool has_two;   // if false, two was not used 
} StoneResult;

static StoneResult handle_stone(STONE_T stone) {
    StoneResult output;
    memset(&output, 0, sizeof(StoneResult));
    
    if (stone == 0) {
        output.has_two = false;
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
        output.has_two = false;
    }

    return output;
}

static void split_stones(StoneWriter *sw) {
    int count = StoneWriter_len(sw);

    // make arrray for the results
    StoneResult* results;
    results = calloc(sizeof(*results), count);
    if (results == NULL) {
        perror("failed to allocate\n");
        exit(1);
    }

    // calculate results
    for (int i = 0; i < count; i++) {
        results[i] = handle_stone(StoneWriter_get_at(sw, i));
    }

    // update stones with results
    for (int i = 0; i < count; i++) {
        StoneWriter_replace(sw, results[i].one, i);
    }
    for (int i = 0; i < count; i++) {
        if (results[i].has_two) {
            StoneWriter_put(sw, results[i].two);
        }
    }

    // discard temporary array
    free(results);
}

static int64_t recurse_stone(StoneWriter *sw, int depth, int max_depth, STONE_T starting_stone) {
    printf("Stone %lli Depth %i. Stone count %lli\n",starting_stone,depth, StoneWriter_len(sw));
    if (depth >= max_depth) {
        return StoneWriter_len(sw);
    } else {
       split_stones(sw);
       int64_t new_len  = (int64_t) StoneWriter_len(sw);
       for (int i = 0; i < new_len; i++) {
            // create new array for next level down
            StoneWriter next =  StoneWriter_new();
            StoneWriter* next_ptr = &next;
            StoneWriter_copy(next_ptr, sw);

            int64_t len2 = recurse_stone(next_ptr, depth+1, max_depth, starting_stone);
            
            StoneWriter_free(&next_ptr);

            return len2;
       }
    }

    perror("unrerachable\n");
}
static int64_t iterate_stone(STONE_T stone, int max_depth) {
    StoneWriter sw = StoneWriter_new();
    
    StoneWriter* sw_ptr = &sw;
    StoneWriter_put(sw_ptr, stone);
    int64_t value = recurse_stone(sw_ptr, 0, max_depth, stone);
    
    StoneWriter_free(&sw_ptr);

    return value;
} 

static int64_t iterate_stones(StoneWriter *sw, int max_depth) {
    int64_t tally = 0;
    for (int i = 0; i < StoneWriter_len(sw); i++) {
        STONE_T stone = StoneWriter_get_at(sw, i);

        tally += iterate_stone(stone , max_depth);
    }

    return tally;
}

int main(int argc, char *argv[]) {
    StoneWriter sw = StoneWriter_new();
    StoneWriter* sw_ptr = &sw;
   
    STONE_T stones[] = {5910927,0,1,47,261223,94788,545,7771};
    const int stones_len = sizeof(stones) / 8;

    StoneWriter_put_many(sw_ptr, stones, stones_len);
    // StoneWriter_print(&sw);

    int max_depth = 75;
    int64_t count = iterate_stones(sw_ptr, max_depth);
    printf("count = %i", count);
    

    StoneWriter_free(&sw_ptr);
}

