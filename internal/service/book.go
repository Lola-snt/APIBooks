package service

import "database/sql"

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "Insert into books (title, author, genre) values(?,?,?)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(lastInsertID)
	return nil
}

func (s BookService) GetBooks() ([]Book, error) {
	query := "select id, title, author, genre from books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	query := "select id, title, author, genre from books where id = ? "
	row, _ := s.db.Query(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "update books set title=?, author=?, genre=? where id=?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)

	return err
}
func (s *BookService) DeleteBook(id int) error {
	query := "delete from books where id=?"
	_, err := s.db.Exec(query, id)
	return err
}
