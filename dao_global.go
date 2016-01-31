package main


import (
	"./models"
	"fmt"
	"net/http"
)


func main() {
	models.InitDB("postgres://postgres:postgres@localhost/bookstore?sslmode=disable")
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":8080",nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
