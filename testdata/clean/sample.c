#include <stdio.h>

// Clean C file without bark comments
int main() {
    printf("Production mode\n");
    
    /* Regular comment */
    int x = 42;
    
    /* Proper implementation */
    int* ptr = malloc(sizeof(int) * 100);
    if (ptr == NULL) {
        return 1;
    }
    
    free(ptr);
    
    return 0;
}

