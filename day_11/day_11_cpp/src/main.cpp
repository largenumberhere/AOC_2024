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

typedef struct 
{
    std::unordered_map <STONE_T, int64_t> inner; 
} StoneWriter;


static void StoneWriter_clear(StoneWriter* sw) {
    sw->inner.clear();

}


static int64_t StoneWriter_count_all(StoneWriter *sw) {
    int64_t tally = 0;
    for (auto item : sw->inner) {
        tally+= item.second;
    }
    
    return tally;
}

static void StoneWriter_free(StoneWriter sw) {
    // freeing is implicit
}

static StoneWriter StoneWriter_new(void) {
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
        if (pair.second!=0) {
            count += 1;
        }
    }

    return count;
} 




static void StoneWriter_put_stone(StoneWriter *sw, STONE_T value) {
    auto val = sw->inner.find(value);

    int64_t count = 0;
    if (val != sw->inner.end()) {
        count = (*val).second;
    }

    
    count += 1;
    sw->inner[value] = count;
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
    int64_t new_count = 0 + count;
    if (val != sw->inner.end()) {
        new_count = (*val).second + count;
    }
    
    sw->inner.insert_or_assign(value, new_count);

}

static int64_t StoneWriter_get_unique_stone_count(StoneWriter *sw, STONE_T value) {
    auto pair_ptr = sw->inner.find(value);
    int64_t count = 0;
    if (pair_ptr != sw->inner.end()) {
        count = (*pair_ptr).second;
    }

    return count;
}

// make a deep copy of the data
static void StoneWriter_copy(StoneWriter *destination, StoneWriter* source) {
    for (auto item : source->inner) {
        destination->inner[item.first] = item.second;
    }
}

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
    return iter.pos != iter.end;
}

static StoneCountInfo StoneWriterIter_next(StoneIter *iter) {
    if (!StoneWriterIter_has_more(*iter)) {
        perror("failed to next");
        exit(1);
    }
    auto current = *iter->pos;
    iter->pos++;

    
    StoneCountInfo info;
    info.stone = current.first;
    info.count = current.second;
    
    return info;
}

static void StoneWriter_replace_stone(StoneWriter* sw, STONE_T from ,STONE_T to, int64_t count) {
    if (count == 0) {
        return;
    }

    auto from_item = sw->inner.find(from);
    if (from_item == sw->inner.end()) {
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
    if (to_item == sw->inner.end()) {
        new_to = count;
    }

    sw->inner[to] = new_to;
    sw->inner[from] = new_from;
    
}



static void StoneWriter_print(StoneWriter *sw) {
    printf("StoneWriter {\n");
    for (auto item : sw->inner) {
        if (item.second !=0) {
            printf("    %lli: %lli\n", item.first, item.second);
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
        to = to + (digit * (int64_t) std::pow((double) 10, (double) mul));
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
        to = to + (digit * (int64_t)pow((double)10, (double)mul));
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

static void split_stones(StoneWriter *sw, StoneWriter* tmp) {
    StoneWriter_clear(tmp);
    StoneWriter_copy(tmp, sw);

    StoneIter iter = StoneWriter_iter(tmp);
    while(StoneWriterIter_has_more(iter)) {
        StoneCountInfo stone_info = StoneWriterIter_next(&iter);

        StoneResult stone_result = eval_stone(stone_info.stone);
        StoneWriter_replace_stone(sw, stone_info.stone, stone_result.one, stone_info.count);

        // add seccondary stone
        if (stone_result.has_two) {
            StoneWriter_put_stone_count(sw, stone_result.two, stone_info.count);
        }
    }

}

int main(void) {
    StoneWriter sw = StoneWriter_new();
    
    STONE_T stones[] = {125, 17};
    const int stones_len = sizeof(stones) / 8;

    StoneWriter_put_array(&sw, stones, stones_len);
    StoneWriter_print(&sw);

    StoneWriter temporary = StoneWriter_new();
    int max_depth = 75;
    for (int i = 0; i < max_depth; i++) {
        split_stones(&sw, &temporary);
    }
    int64_t count = StoneWriter_count_all(&sw);
    printf("count = %lli\n", count);
    int64_t unique = StoneWriter_get_unique_stones_count(&sw);
    printf("unique = %lli\n", unique);
    StoneWriter_free(sw);
}

