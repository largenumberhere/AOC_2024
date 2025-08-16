#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <math.h>
#include <string.h>

int intCompare(const void* a, const void* b) {
    int a_val = * (int*)a;
    int b_val = * (int*) b;

    if (a_val < b_val) {
        return -1;
    } else if (a_val > b_val) {
        return 1;
    }

    return 0;
}

int main() {
    const char* file_path = "my_input.txt";
    FILE* f = fopen("my_input.txt", "r");
    if (f == NULL) {
        fprintf(stderr, "Failed to open file %s\n", file_path);
        exit(1);
    }


    // read file count of lines
    size_t lines_count = 0;
    ssize_t buff_cap = 0;
    char* line = NULL;
    while (getline(&line, &buff_cap, f) > 0) {
        if (strlen(line) == 0 || line[0] == '\n' || line[0] == '\r') {
            continue;
        }
        lines_count++;
    }
    
    int* lefts = calloc(lines_count, sizeof(int));
    int* rights = calloc(lines_count, sizeof(int));
    int leftsPos = 0;
    int rightsPos = 0;

    rewind(f);
    while (1) {
        ssize_t line_len = getline(&line, &buff_cap , f );

        if (line_len < 0) {
            break;
        }

   
        char* left = line;
        char* right = strchr(left, ' ');
        * right = '\0';
        right ++;
        lefts[leftsPos++] = atoi(left);
        rights[rightsPos++] = atoi(right);
    }

    free(line);

    int* map = calloc(999999, sizeof(int));
    for (int i = 0; i < lines_count; i++) {
        int r = rights[i];
        map[rights[i]] ++;
    }


    ssize_t tally = 0;
    for (int i = 0; i < lines_count; i++) {
        int count = map[lefts[i]];
        ssize_t product = lefts[i] * count; 
        printf("%i | %i || %li (%i) \n", lefts[i], rights[i], product, count);
        tally += product;
    }

    printf("similarity : %li\n", tally);
    
    free(map);
    free(rights);
    free(lefts);
    fclose(f);
}
