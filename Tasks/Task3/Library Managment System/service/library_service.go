package service

import (
	model "liabrayManagment/model"
)

type Book = model.Book
type Member = model.Member

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int) bool
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}

type Laibrary struct {
	Books  map[int]Book
	Member map[int]Member
}

func NewLaibrary() *Laibrary {
	return &Laibrary{
		Books:  make(map[int]Book),
		Member: make(map[int]Member),
	}
}

func (l *Laibrary) AddBook(book Book) bool {
	l.Books[book.Id] = book
	return true
}

func (l *Laibrary) RemoveBook(bookId int) bool {
	_, ok := l.Books[bookId]
	if !ok {
		return false
	}
	delete(l.Books, bookId)
	return true

}

func (l *Laibrary) BorrowBook(bookId int, memberId int) bool {

	member, ok := l.Member[memberId]

	if !ok {
		return false
	}

	book, ok := l.Books[bookId]

	if !ok || book.Status == "Borrowed" {
		return false
	}

	book.Status = "Borrowed"
	l.Books[bookId] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Member[memberId] = member

	return true
}

func (l *Laibrary) ReturnBook(bookId int, memberId int) bool {

	member, ok := l.Member[memberId]
	if !ok {
		return false
	}

	book, ok := l.Books[bookId]
	if !ok {
		return false
	}

	borrowedBook := member.BorrowedBooks

	index := -1

	for i, book := range borrowedBook {
		if bookId == book.Id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}
	member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)
	l.Member[memberId] = member

	book.Status = "Available"
	l.Books[bookId] = book

	return true
}

func (l *Laibrary) ListAvailableBooks() []Book {

	var books []Book

	for _, book := range l.Books {
		if book.Status == "Available" {
			books = append(books, book)
		}
	}

	return books
}

func (l *Laibrary) ListBorrowedBooks(memberID int) []Book {

	return l.Member[memberID].BorrowedBooks
}
