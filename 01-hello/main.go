package main

import (
	"encoding/json"
	"fmt"
	"hello-app/models"
	"io"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello Application")
	})

	http.HandleFunc("/json", func(rw http.ResponseWriter, r *http.Request) {

		customer := models.Customer{
			Id:         1,
			FirstName:  "Pratya",
			LastName:   "Yeekhaday",
			CreateDate: time.Now(),
		}

		data, err := json.Marshal(customer)
		if err != nil {
			fmt.Println(err.Error())
		}

		io.WriteString(rw, string(data))
	})

	http.HandleFunc("/custom", func(rw http.ResponseWriter, r *http.Request) {

		employee := models.Employee{
			EmpId:       1,
			EmpFullname: "Pratya Yeekhaday",
			IsActive:    true,
			CreateDate:  time.Now(),
		}

		data, err := json.Marshal(employee)
		if err != nil {
			fmt.Println(err.Error())
		}

		io.WriteString(rw, string(data))
	})

	fmt.Println("server 8080 started")
	http.ListenAndServe(":8080", nil)

}
