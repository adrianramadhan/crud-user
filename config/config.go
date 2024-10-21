package config

import (
    "fmt"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    host := os.Getenv("PG_HOST")
    if host == "" {
        host = "localhost"
    }

    port := os.Getenv("PG_PORT")
    if port == "" {
        port = "5432"
    }

    user := os.Getenv("PG_USER")
    if user == "" {
        user = "postgres"
    }

    password := os.Getenv("PG_PASSWORD")
    if password == "" {
        password = "password"
    }

    dbname := os.Getenv("PG_DB")
    if dbname == "" {
        dbname = "golangdb"
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host,
        user,
        password,
        dbname,
        port,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
