package repository

import (
    "context"
    "fmt"
    "log"
    "os"
    "github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgresPool() *pgxpool.Pool {
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    host := os.Getenv("POSTGRES_HOST")
    port := os.Getenv("POSTGRES_PORT")
    dbname := os.Getenv("POSTGRES_DB")
    
	fmt.Println("se conectando a:", user, host, port, dbname)

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
        user, password, host, port, dbname)
    
    pool, err := pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("não foi possível criar a pool: %v\n", err)
    }
    
    err = pool.Ping(context.Background())
    if err != nil {
        log.Fatalf("não foi possível pingar o banco de dados: %v\n", err)
    }
    
    fmt.Println("pool do postgres pronta")
    
    return pool 
}