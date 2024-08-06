package controller

import (
	model "liabrayManagment/model"
	service "liabrayManagment/service"
)

type Library = service.Laibrary
type Member = model.Member
type Book = model.Book

var library = service.NewLaibrary()

var MEMBERID int

func Login(memberID int, name string) bool {
	var newMember Member
	newMember.Id = memberID
	newMember.Name = name
	library.Member[memberID] = newMember
	MEMBERID = memberID
	return true

}

func AddBook(title string, id int, author string) bool {
	var newBook = Book{
		Title:  title,
		Id:     id,
		Author: author,
		Status: "Available",
	}

	library.AddBook(newBook)

	return true

}

func RemoveBook(bookId int) bool {
	isRemoved := library.RemoveBook(bookId)
	return isRemoved
}

//BorrowBook(bookID int, memberID int)

func BorrowBook(bookID int) bool {
	isSucess := library.BorrowBook(bookID, MEMBERID)
	if isSucess == true {
		return true
	}
	return false
}

func ReturnBook(bookID int) bool {
	isSucess := library.ReturnBook(bookID, MEMBERID)
	if isSucess == true {
		return true
	}
	return false
}

func ListAvailableBooks() []Book {
	return library.ListAvailableBooks()
}

func ListBorrowedBooks() []Book {
	return library.ListBorrowedBooks(MEMBERID)
}

/*
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book


*/
