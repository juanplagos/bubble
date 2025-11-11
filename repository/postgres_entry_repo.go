package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/juanplagos/bubble/model"
)

type EntryRepo interface {
	GetAllEntries() ([]model.Entry, error)
	GetEntryById(id int) (model.Entry, error)
	GetEntryBySlug(slug string) (model.Entry, error)
	CreateEntry(entry *model.Entry) error
	UpdateEntry(id int, entry *model.Entry) error
	DeleteEntry(id int) error
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
	var e model.Entry
	err := repo.pool.QueryRow(
		context.Background(),
		"SELECT id, title, slug, body, author, created_at FROM entries WHERE id = $1",
		id,
	).Scan(&e.ID, &e.Title, &e.Slug, &e.Body, &e.Author, &e.CreatedAt)

	if err != nil {
		return model.Entry{}, err
	}

	return e, nil
}

func (repo *PostgresEntryRepo) GetEntryBySlug(slug string) (model.Entry, error) {
	var e model.Entry
	err := repo.pool.QueryRow(
		context.Background(),
		"SELECT id, title, slug, body, author, created_at FROM entries WHERE slug = $1",
		slug,
	).Scan(&e.ID, &e.Title, &e.Slug, &e.Body, &e.Author, &e.CreatedAt)

	if err != nil {
		return model.Entry{}, err
	}

	return e, nil
}

func (repo *PostgresEntryRepo) CreateEntry(entry *model.Entry) error {
	err := repo.pool.QueryRow(
		context.Background(),
		"INSERT INTO entries (title, slug, body, author, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id",
		entry.Title, entry.Slug, entry.Body, entry.Author,
	).Scan(&entry.ID)
	return err
}

func (repo *PostgresEntryRepo) UpdateEntry(id int, entry *model.Entry) error {
	_, err := repo.pool.Exec(
		context.Background(),
		"UPDATE entries SET title = $1, slug = $2, body = $3, author = $4 WHERE id = $5",
		entry.Title, entry.Slug, entry.Body, entry.Author, id,
	)
	return err
}

func (repo *PostgresEntryRepo) DeleteEntry(id int) error {
	_, err := repo.pool.Exec(
		context.Background(),
		"DELETE FROM entries WHERE id = $1",
		id,
	)
	return err
}
