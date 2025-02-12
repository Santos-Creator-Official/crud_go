package services

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

var Pool *pgxpool.Pool

//postgres
//func InitDatabase() {
//	// Load .env file
//	errLoad := godotenv.Load()
//	if errLoad != nil {
//		log.Fatalf("Error loading .env file: %v", errLoad)
//	}
//
//	dbUser := os.Getenv("DB_USER")
//	dbPass := os.Getenv("DB_PASS")
//	dbHost := os.Getenv("DB_HOST")
//	dbName := os.Getenv("DB_NAME")
//	dbPort := os.Getenv("DB_PORT")
//
//	var err error
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
//
//	Pool, err = pgxpool.New(ctx, connStr)
//	if err != nil {
//		log.Fatalf("Unable to connect to database: %v\n", err)
//	}
//
//	log.Println("Database connection established")
//}

var DB *mongo.Database

// mongoDb
func InitDatabase() {

	// Load .env file
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatalf("Error loading .env file: %v", errLoad)
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)

	if err != nil {
		panic(err)
	}

	DB = client.Database(os.Getenv("MONGO_DATABASE"))
	log.Println("Database connection established")
}

//func CloseDB() {
//	if Pool != nil {
//		Pool.Close()
//		log.Println("Database connection closed")
//	}
//}
