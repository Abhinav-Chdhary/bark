#include <iostream>

// BARK: Remove debug output before production
// BARK
class Sample {
public:
    void process() {
        std::cout << "Debug: Processing..." << std::endl;
        // BARK: Replace with proper logging framework
    }
    
    // Regular comment
    void normalFunction() {
        /* BARK: Optimize this algorithm later */
        for (int i = 0; i < 1000000; i++) {
            // Do something
            // BARK needs optimization
        }
    }
};

