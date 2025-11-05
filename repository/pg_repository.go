package repository

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/juanplagos/bubble/model"
    "github.com/jackc/pgx/v5/pgxpool"
)

type EntryRepo interface {
    GetAllEntries() ([]model.Entry, error)
    GetEntryById(id int) (model.Entry, error)
    GetEntryBySlug(slug string) (model.Entry, error) 
}

type PostgresEntryRepo struct {
    pool *pgxpool.Pool
}

func (repo *PostgresEntryRepo) GetAllEntries() ([]model.Entry, error) {
    rows, err := repo.pool.Query(context.Background(), "SELECT id, title, slug, body, author, created_at FROM entries")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var entries []model.Entry

    for rows.Next() {
        var e model.Entry

        err := rows.Scan(&e.ID, &e.Title, &e.Slug, &e.Body, &e.Author, &e.CreatedAt)
        if err != nil {
            return nil, err
        }
        entries = append(entries, e)
    }

    if rows.Err() != nil {
        return nil, rows.Err()
    }

    return entries, nil 
}

func (repo *PostgresEntryRepo) GetEntryById(id int) (model.Entry, error) {
    return model.Entry{}, nil
}

func (repo *PostgresEntryRepo) GetEntryBySlug(slug string) (model.Entry, error) {
    return model.Entry{}, nil
}

var Pool *pgxpool.Pool

func InitDB() {
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    host := os.Getenv("POSTGRES_HOST")
    port := os.Getenv("POSTGRES_PORT")
    dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

    var err error
    Pool, err = pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("não foi possível criar a pool: %v\n", err)
    }

    err = Pool.Ping(context.Background())
    if err != nil {
        log.Fatalf("não foi possível pingar o banco de dados: %v\n", err)
    }
    fmt.Println("pool do postgres pronta")
}

func CloseDB() {
    if Pool != nil {
        Pool.Close()
    }
}