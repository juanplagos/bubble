package repository

import (
    "context"

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

func NewPostgresEntryRepo(pool *pgxpool.Pool) *PostgresEntryRepo {
    return &PostgresEntryRepo{
        pool: pool,
    }
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