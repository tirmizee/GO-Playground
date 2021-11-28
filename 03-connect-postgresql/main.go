package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {

		con, err := sql.Open("pgx", "host=0.0.0.0 port=5432 dbname=test user=usr password=pass")
		if err != nil {
			panic(err)
		}

		if err = con.Ping(); err != nil {
			panic(err)
		}

		defer con.Close()

		random := strconv.Itoa(rand.Intn(100))
		username, email := "tirmizee"+random, "tirmizee"+random+"@gmail"
		query := `
			INSERT INTO public.accounts(username, "password", email, created_on) 
			VALUES($1, $2, $3, $4)`
		result, err := con.Exec(query, username, "pass", email, time.Now())
		if err != nil {
			fmt.Fprintf(w, "insert database err")
			return
		}

		count, err := result.RowsAffected()
		if err != nil {
			fmt.Fprintf(w, "rows affected err")
			return
		}

		fmt.Fprintf(w, "insert database "+strconv.Itoa(int(count)))
	})

	http.HandleFunc("/select", func(w http.ResponseWriter, r *http.Request) {

		con, err := sql.Open("pgx", "host=0.0.0.0 port=5432 dbname=test user=usr password=pass")

		if err != nil {
			panic(err)
		}

		if err = con.Ping(); err != nil {
			panic(err)
		}

		defer con.Close()

		rows, err := con.Query(`
			select 
				user_id, 
				username, 
				password, 
				email, 
				created_on 
			from accounts`)

		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		var userId int
		var createdOn time.Time
		var username, password, email string

		for rows.Next() {
			err := rows.Scan(&userId, &username, &password, &email, &createdOn)
			if err != nil {
				fmt.Fprintf(w, err.Error())
				break
			}

			fmt.Fprintf(w, strconv.Itoa(userId)+username+password+email+createdOn.String()+"\n")

		}

		defer rows.Close()

	})

	http.ListenAndServe(":8080", nil)
}
