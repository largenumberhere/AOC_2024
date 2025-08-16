#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>
#include <unordered_set>
#include <unordered_map>
#include <vector>
#include <string>
#include <memory>
#include <iostream>
#include <algorithm>

int main() {
    std::unordered_set<std::string> unique_keys;
    
    const char* file_path = "sample_input.txt";
    FILE* f = fopen(file_path, "r");
    if (f == NULL) {
        fprintf(stderr, "failed to open %s\n", file_path);
        exit(1);
    }

    size_t buff_cap = 0;
    char* line = NULL;
    while (1) {
        ssize_t line_len = getline(&line, &buff_cap , f );
        if (line_len < 0) {
            break;
        }

        char* left = line;
        char* right = strchr(left, '-');
        * right = '\0';
        right ++;

        std::string left_str = std::string({left[0], left[1]});
        std::string right_str = std::string({right[0], right[1]});
        unique_keys.insert(left_str);
        unique_keys.insert(right_str);
    }

    std::unordered_map<std::string_view, std::unordered_set<std::string_view>> connections_map;
    if (buff_cap > 0) {
        line[0] = '\0';
    }

    rewind(f);
    while(1) {
        ssize_t line_len = getline(&line, &buff_cap , f );
        if (line_len < 0) {
            break;
        }
        
        char* left = line;
        char* right = strchr(left, '-');
        * right = '\0';
        right ++;

        std::string left_str = std::string({left[0], left[1]});
        std::string right_str = std::string({right[0], right[1]});
        
        std::string_view left_ref = std::string_view(*unique_keys.find(left_str));
        std::string_view right_ref = std::string_view(*unique_keys.find(right_str));

        // left -> right
        if (!connections_map.contains(left_ref)) {
            connections_map.insert({left_ref, std::unordered_set<std::string_view>()});
        }
        connections_map[left_ref].insert(right_ref);

        // right -> left
        if (!connections_map.contains(right_ref)) {
            connections_map.insert({right_ref, std::unordered_set<std::string_view>()});
        }
        connections_map[right_ref].insert(left_ref);

    }
        
    size_t i = 0;
    for (auto pair : connections_map) {
        auto left = pair.first;
        for (auto right : pair.second) {
            i++;
        }
    }
    std::cout << "connections count: " << i << "\n"; 

    // fetch all the triangles
    std::unordered_set<std::string> triangles;
    for (std::string_view key : unique_keys) {
        auto children1 = connections_map[key];
        for (auto child1 : children1) {
            if (child1 == key) {
                continue;
            }
            auto children2 = connections_map[child1];
            for (auto child2 : children2) {
                if (child2 == key || child2 == child1) {
                    continue;
                }
                if (!connections_map[child2].contains(key)) {
                    continue;
                }

                if (!connections_map[child2].contains(child1)) {
                    continue;
                }
                
                std::string_view arr[3];
                arr[0] = key;
                arr[1] = child1;
                arr[2] = child2;

                std::sort(std::begin(arr), std::end(arr));
                std::string s;
                s.append(arr[0]);
                s.append(" ");
                s.append(arr[1]);
                s.append(" ");
                s.append(arr[2]);
        
                if (!triangles.contains(s)) {
                    triangles.insert(s);
                }
            }
        }
    }

    std::vector<std::string> removeable;
    for (auto tptr = triangles.begin(); tptr != triangles.end(); tptr++) {
        if (tptr->at(0) != 't' && tptr->at(3) != 't' && tptr->at(6) != 't') {
            removeable.push_back(*tptr);
        }
    }
    for (auto r : removeable) {
        triangles.extract(r);
    }

    size_t count = 0;
    for (auto s : triangles) {
        count ++;
    }
    std::cout << "final count: " << count << "\n"; 

    free(line);
    fclose(f);
}
