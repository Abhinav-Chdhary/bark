public class Sample {
    // Regular class comment
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        
        /* Standard multi-line comment
           explaining the logic */
        connectToDatabase();
    }
    
    // Method documentation
    private static void connectToDatabase() {
        // Implementation details
        String password = getPasswordFromEnv();
    }
    
    private static String getPasswordFromEnv() {
        return System.getenv("DB_PASSWORD");
    }
}

