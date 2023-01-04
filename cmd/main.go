package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gildemberg-santos/retrospectiva/pkg"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	http.HandleFunc("/", RestrospectivaHttp)
	http.ListenAndServe(":8080", nil)
}

type DataResult struct {
	UserId    string `json:"user_id"`
	CompanyId string `json:"company_id"`
	Link      string `json:"link"`
}

func RestrospectivaHttp(w http.ResponseWriter, r *http.Request) {
	var d = DataResult{}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			http.Error(w, http.StatusText(http.StatusBadRequest)+": Nenhuma informação recebida", http.StatusBadRequest)
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest)+": json.NewDecoder", http.StatusBadRequest)
			return
		}
	}

	if d.UserId == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+": Nenhuma informação encontrada", http.StatusBadRequest)
		return
	}

	var data bson.Raw = pkg.FindByIdClient(d.UserId).(bson.Raw)
	d.CompanyId = data.Lookup("id-empresa").StringValue()

	json.NewEncoder(w).Encode(d)
}
