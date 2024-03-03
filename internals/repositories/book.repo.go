package repositories

import (
	"bookstoreGo/internals/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	*sqlx.DB
}

func InitBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db}
}
func (b *BookRepo) FindAll() ([]models.BookModel, error) {
	query := "select * from books"
	result := &[]models.BookModel{}
	if err := b.Select(result, query); err != nil {
		return nil, err
	}
	return *result, nil
}
func (b *BookRepo) FindById(id int) ([]models.BookModel, error) {
	query := "select * from books where id =?"
	result := []models.BookModel{}
	if err := b.Select(&result, query, id); err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BookRepo) DeleteById(id int) error {
	query := "delete from books where id =?"
	if _, err := b.Exec(query, id); err != nil {
		return err
	}
	return nil

}

func (b *BookRepo) SaveBook(body models.BookModel) error {
	query := "insert into books(title,description,author) values (?,?,?)"
	if _, err := b.Exec(query, body.Title, body.Description, body.Author); err != nil {
		return err
	}
	return nil
}

func (b *BookRepo) UpdateById(id int, body models.BookModel) error {
	query := "UPDATE books SET"

	// interface{} dapat menyimpan nilai dari tipe data apa pun
	var args []interface{}

	if body.Title != "" {
		query += " title = ?,"
		args = append(args, body.Title)
	}
	if body.Description != nil && *body.Description != "" {
		query += " description = ?,"
		args = append(args, *body.Description)
	}
	if body.Author != "" {
		query += " author = ?,"
		args = append(args, body.Author)
	}

	// Hapus koma terakhir
	query = strings.TrimSuffix(query, ",")

	query += " WHERE id = ?"
	args = append(args, id)

	// Eksekusi parameterized query dengan Exec, .exec --> tidak mereturn rows (data dari database)
	if _, err := b.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
