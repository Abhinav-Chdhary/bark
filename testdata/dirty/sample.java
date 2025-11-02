public class Sample {
    // BARK: Remove this test code before release
    // BARK
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        
        /* BARK: This is a temporary database connection
           Replace with proper connection pooling */
        connectToDatabase();
    }
    
    // Regular comment
    private static void connectToDatabase() {
        // BARK: Hardcoded credentials need to be removed
        String password = "admin123";
        // BARK TODO remove
        String username = "admin";
    }
}

