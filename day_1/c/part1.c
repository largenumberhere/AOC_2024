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
    FILE* f = fopen(file_path, "r");
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

        // line[4] = '\0';
        // char* right = line + 5;
        char* left = line;

        char* right = strchr(left, ' ');
        * right = '\0';
        right ++;
        lefts[leftsPos++] = atoi(left);
        rights[rightsPos++] = atoi(right);
    }

    free(line);

    qsort(lefts, lines_count, sizeof(int), intCompare);
    qsort(rights, lines_count, sizeof(int), intCompare);

    size_t diff_sum = 0;
    for (int i = 0; i < lines_count; i++) {
        int diff = abs(lefts[i] - rights[i]);
        printf("%i | %i || %i \n", lefts[i], rights[i], diff);
        diff_sum += diff;
    }

    printf("%li\n", diff_sum);
    free(rights);
    free(lefts);

    


    fclose(f);

    

    

}
