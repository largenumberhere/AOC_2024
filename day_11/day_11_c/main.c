#include <stdint.h> 
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define stone_t int32_t

static stone_t *stones; 
static int stones_cap;
static int stones_len;

typedef struct Stone {
    int value;
    int count;
};

int count_digits(stone_t value) {
    // int64_t vin = value;
    int size = 0;
    while (value !=0) {
        size++;
        value /=10;
    }

    // printf("%lli has count of digits %lli\n", vin, size);
    return size;
}

stone_t right_half(stone_t value, int length) {
    stone_t to = 0;
    stone_t mul = 0;
    for (int i = 0; (i < (length/2)) && (value!=0) ;i++) {
        // pop off last digit
        stone_t digit = value % 10;
        value /= 10;
        to = to + (digit * pow(10, mul));
        mul ++;

    }

    return to;
}

stone_t left_half(stone_t value, int length) {
   stone_t vin = value; 
    // discard first half
    for (int i = 0; i < (length/2); i++) {
        value /=10;
    }

   
    stone_t to = 0;
    stone_t mul = 0;
    for (int i = 0; (i < (length/2)) && (value!=0) ;i++) {
        // pop off last digit
        stone_t digit = value % 10;
        value /= 10;
        to = to + (digit * pow(10, mul));
        mul ++;

    }

    // printf("%d -> %d\n", vin, to);

    return to;
}

void append_stone(stone_t value) {
    
    
    if (stones_len +1 > stones_cap) {
        printf("reallocating\n");
        int new_cap = stones_cap * 2;

        void * m = realloc(stones, sizeof(stone_t) * new_cap);
        if (m == NULL) {
            perror("failed to reallocate\n");
        }
        
        stones_cap = new_cap;
        stones = (stone_t*) m;
    }

    
    stones[stones_len++] = value; 
    
}

void iterate_stones() {
    int len = stones_len;
    for (int i = 0; i < len; i++) {
        int digits = count_digits(stones[i]);

        // If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
        if (stones[i] == 0) {
            stones[i] = 1;

        // If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
        } else if ((digits % 2) == 0) {
            stone_t right = right_half(stones[i], digits);
            stone_t left = left_half(stones[i], digits);
            
            // printf("splitting %lli\n", stones[i]);
            // printf("left %lli\n", left);
            // printf("right %lli\n", right);
            

            stones[i] = left;
            append_stone(right);
            // printf("appened: %lli", stones[stones_len-1]);
        
        // If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
        } else {
            stones[i] = stones[i] * 2024;
        }
    }
}

void print_stones() {
    printf("Stones: ");
    for (int i = 0; i < stones_len; i++) {
        printf("%lli, ", stones[i]);
    }
    printf("\n");
}

int main() {
    // initialize array
    const int size = 4*1024*1024;

    void *m = malloc(sizeof(stone_t) * size);
    if (m == NULL) {
        perror("allocation failure");
    } 
    memset(m, 0, stones_cap);

    stones = (stone_t*) m;
    stones_cap = size;
    stones_len = 0;


    // set the starting values
    // char* input = "5910927 0 1 47 261223 94788 545 7771";
    // stones[0] = 5910927;
    // stones[1] = 0;
    // stones[2] = 1;
    // stones[3] = 47;
    // stones[4] = 261223;
    // stones[5] = 94788;
    // stones[6] = 545;
    // stones[7] = 7771;
    // stones_len = 8;

    stones[0] = 125;
    stones[1] = 17;
    stones_len = 2;

    print_stones();
    for (int i = 0; i < 75; i++) {
        iterate_stones();
        printf("iteration %i. Length %i\n",i+1, stones_len);
        // print_stones();
    }

    printf("count of stones: %i\n", stones_len);





    return 0;
}