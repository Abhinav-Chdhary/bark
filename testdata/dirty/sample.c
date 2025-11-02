#include <stdio.h>

// BARK: Remove debug printf statements
// BARK
int main() {
    printf("Debug mode\n");  // BARK: Delete this line
    
    /* Regular comment */
    int x = 42;
    
    /* BARK: This malloc needs proper error handling */
    int* ptr = malloc(sizeof(int) * 100);
    
    // BARK plain test
    free(ptr);

    return 0;
}

