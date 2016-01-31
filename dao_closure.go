package main
import (
	"./models"
	"database/sql"
	"log"
	"net/http"
	"fmt"
)

type Env struct {
	db *sql.DB
}

func main(){
	db, err := models.NewDB("postgres://postgres:postgres@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	http.Handle("/books", booksIndex(env))
	http.ListenAndServe(":8080", nil)
}

func booksIndex(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := models.AllBooksDep(env.db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	})
}
