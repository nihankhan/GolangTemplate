package handlers

import (
	"log"
	"net/http"
	"html/template"

	"github.com/nihankhan/GolangTemplate/db"
)

type Nihan struct {
	ID          int
	CashPickUp  string
	CashDropOff string
	Amount      int
}

func Query(resp http.ResponseWriter, req *http.Request) {

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM Nihan.Garda")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var data []Nihan

	for rows.Next() {
		ebu := Nihan{}

		err := rows.Scan(&ebu.ID, &ebu.CashPickUp, &ebu.CashDropOff, &ebu.Amount)

		if err != nil {
			log.Println(err)

			return
		}

		data = append(data, ebu)

		// fmt.Fprintf(resp, "%v\n", data)
	}

		tmpl, err := template.ParseGlob("/home/nihan/Documents/GolangTemplate/templates/*.html")

		if err != nil {
			panic(err)
		}

		tmpl.ExecuteTemplate(resp, "index.html", data)

		// fmt.Fprintf(resp, "%v\n", data)
}