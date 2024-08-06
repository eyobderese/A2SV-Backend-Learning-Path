package main

import (
	"fmt"
	controller "liabrayManagment/controller"
)

func main() {

	fmt.Println("WELLCOME TO ABREHOT LIBRARY COMPANION")
	fmt.Println("=================LOGIN===========================")
	login()
	operation()

}

func login() {
	fmt.Print("UserName... ")
	var username int
	fmt.Scanln(&username)

	fmt.Print("Full Name... ")
	var fullName string
	fmt.Scanln(&fullName)
	isLogin := controller.Login(username, fullName)

	for !isLogin {
		fmt.Println("It is Invalid username or name , Please insert the right username and name")

		fmt.Print("UserName... ")
		var username int
		fmt.Scanln(&username)

		fmt.Print("Full Name... ")
		var fullName string
		fmt.Scanln(&fullName)
		isLogin = controller.Login(username, fullName)

	}

	fmt.Println("Well come You are sucess fully logding")

}

func operation() {
	for {
		fmt.Println("Please choose an option:")
		fmt.Println("1. Add book")
		fmt.Println("2. Remove book")
		fmt.Println("3. Borrow book")
		fmt.Println("4. Return book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books")
		fmt.Println("7. Exit")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			// Call the function to add a book
			fmt.Println("Add book selected")
			addBook()
		case 2:
			// Call the function to remove a book
			fmt.Println("Remove book selected")
			removeBook()
		case 3:
			// Call the function to borrow a book
			fmt.Println("Borrow book selected")
			borrowBook()
		case 4:
			// Call the function to return a book
			fmt.Println("Return book selected")
			returnBook()
		case 5:
			// Call the function to list available books
			fmt.Println("List available books selected")
			listAvailableBooks()
		case 6:
			// Call the function to list borrowed books
			fmt.Println("List borrowed books selected")
			listBorrowedBooks()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please enter a number between 1 and 7.")
		}
	}
}

func addBook() {
	var title string
	var id int
	var author string

	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)

	fmt.Print("Enter book id: ")
	fmt.Scanln(&id)

	fmt.Print("Enter book author: ")
	fmt.Scanln(&author)
	isCreated := controller.AddBook(title, id, author)

	if isCreated {
		fmt.Println("Book added successfully!")
	} else {
		fmt.Println("Book added Unsuccessfully!")
	}

}

func removeBook() {
	var bookId int

	fmt.Print("Enter book id to remove: ")
	fmt.Scanln(&bookId)

	isRemoved := controller.RemoveBook(bookId)

	if isRemoved {
		fmt.Println("Book removed successfully!")
	} else {
		fmt.Println("Failed to remove book. Please check the book id and try again.")
	}
}

func borrowBook() {
	var bookId int

	fmt.Print("Enter book id to borrow: ")
	fmt.Scanln(&bookId)

	isBorrowed := controller.BorrowBook(bookId)

	if isBorrowed {
		fmt.Println("Book borrowed successfully!")
	} else {
		fmt.Println("Failed to borrow book. Please check the book id and try again.")
	}
}

func returnBook() {
	var bookId int

	fmt.Print("Enter book id to return: ")
	fmt.Scanln(&bookId)

	isReturned := controller.ReturnBook(bookId)

	if isReturned {
		fmt.Println("Book returned successfully!")
	} else {
		fmt.Println("Failed to return book. Please check the book id and try again.")
	}
}

func listAvailableBooks() {
	books := controller.ListAvailableBooks()

	fmt.Println("Available books:")
	for _, book := range books {
		fmt.Printf("Title: %s\nAuthor: %s\nID: %d\n\n", book.Title, book.Author, book.Id)
	}
}

func listBorrowedBooks() {
	books := controller.ListBorrowedBooks()

	fmt.Println("Borrowed books:")
	for _, book := range books {
		fmt.Printf("Title: %s\nAuthor: %s\nID: %d\n\n", book.Title, book.Author, book.Id)
	}
}
