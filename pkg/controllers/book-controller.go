package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rohanfizz/go-bookstore/pkg/models"
	"github.com/rohanfizz/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	utils.SendOkResponse(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookid")
	}
	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	utils.SendOkResponse(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	utils.SendOkResponse(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookid")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	utils.SendOkResponse(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err!= nil {
        fmt.Println("Error while parsing bookid")
    }

	bookDetails,db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author!= "" {
        bookDetails.Author = updateBook.Author
    }
	if updateBook.Publication!= "" {
        bookDetails.Publication = updateBook.Publication
    }

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	utils.SendOkResponse(w, res)
}