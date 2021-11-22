package handlers

import (
	"app-handler/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello application handlers")
}

func Product(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	if r.Method == http.MethodPost {
		var productRequest models.ProductRequest
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&productRequest)
		if err != nil {
			fmt.Fprintf(w, "Unmarshal request error")
			return
		}

		fmt.Println(productRequest)

		productResponse := models.CustomerResponse{
			Id:         1,
			Name:       "Apple",
			Status:     true,
			CreateDate: time.Now(),
		}

		jsonBytes, err := json.Marshal(productResponse)
		if err != nil {
			fmt.Fprintf(w, "Marshal response error")
			return
		}

		io.WriteString(w, string(jsonBytes))
	}

}

func Customer(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	if r.Method == http.MethodPost {

		var customerRequest models.CustomerRequest
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&customerRequest)
		if err != nil {
			fmt.Fprintf(w, "Unmarshal request error")
			return
		}

		fmt.Println(customerRequest)

		customerResponses := []models.CustomerResponse{
			{
				Id:         1,
				Name:       "Tir",
				Status:     true,
				CreateDate: time.Now(),
			},
			{
				Id:         2,
				Name:       "Pra",
				Status:     true,
				CreateDate: time.Now(),
			},
		}

		jsonBytes, err := json.Marshal(customerResponses)
		if err != nil {
			fmt.Fprintf(w, "Marshal response error")
			return
		}

		io.WriteString(w, string(jsonBytes))
	}

}
