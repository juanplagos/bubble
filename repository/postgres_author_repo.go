package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/juanplagos/bubble/model"
)

type AuthorRepo interface {
	GetAllAuthors() ([]model.Author, error)
	GetAuthorByUsername(username string) (model.Author, error)
	GetAuthorByEmail(email string) (model.Author, error)
	CreateAuthor(author *model.Author) error
	UpdateAuthor(username string, author *model.Author) error
	DeleteAuthor(username string) error
}

type PostgresAuthorRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresAuthorRepo(pool *pgxpool.Pool) *PostgresAuthorRepo {
	return &PostgresAuthorRepo{
		pool: pool,
	}
}

func (repo *PostgresAuthorRepo) GetAllAuthors() ([]model.Author, error) {
	rows, err := repo.pool.Query(
		context.Background(),
		"SELECT username, email, password FROM authors",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []model.Author

	for rows.Next() {
		var a model.Author
		err := rows.Scan(&a.Username, &a.Email, &a.Password)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return authors, nil
}

func (repo *PostgresAuthorRepo) GetAuthorByUsername(username string) (model.Author, error) {
	var a model.Author
	err := repo.pool.QueryRow(
		context.Background(),
		"SELECT username, email, password FROM authors WHERE username = $1",
		username,
	).Scan(&a.Username, &a.Email, &a.Password)

	if err != nil {
		return model.Author{}, err
	}

	return a, nil
}

func (repo *PostgresAuthorRepo) GetAuthorByEmail(email string) (model.Author, error) {
	var a model.Author
	err := repo.pool.QueryRow(
		context.Background(),
		"SELECT username, email, password FROM authors WHERE email = $1",
		email,
	).Scan(&a.Username, &a.Email, &a.Password)

	if err != nil {
		return model.Author{}, err
	}

	return a, nil
}

func (repo *PostgresAuthorRepo) CreateAuthor(author *model.Author) error {
	_, err := repo.pool.Exec(
		context.Background(),
		"INSERT INTO authors (username, email, password) VALUES ($1, $2, $3)",
		author.Username, author.Email, author.Password,
	)
	return err
}

func (repo *PostgresAuthorRepo) UpdateAuthor(username string, author *model.Author) error {
	_, err := repo.pool.Exec(
		context.Background(),
		"UPDATE authors SET email = $1, password = $2 WHERE username = $3",
		author.Email, author.Password, username,
	)
	return err
}

func (repo *PostgresAuthorRepo) DeleteAuthor(username string) error {
	_, err := repo.pool.Exec(
		context.Background(),
		"DELETE FROM authors WHERE username = $1",
		username,
	)
	return err
}
