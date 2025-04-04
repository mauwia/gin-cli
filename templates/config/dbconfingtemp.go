package configTemplates

// MainTemplate returns the basic Gin server code as a string.
// It replaces the placeholder with the provided module name.
func PostgresDBConfigTemplate(moduleName string) string {
	return `
	
	package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgresDB(config *AppConfig) *gorm.DB {
	dbHost := config.DB_HOST
	dbPort := config.DB_PORT
	dbUser := config.DB_USER
	dbPassword := config.DB_PASSWORD
	dbName := config.DB_NAME

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db

}


`
}
func MongoDBConfigTemplate(moduleName string) string {
	return `
	
	package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB(config *AppConfig) *mongo.Database {
	mongoURI := config.MONGODB_URI
	dbName := config.MONGODB_DATABASE

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database(dbName)
	return db

}


`
}
