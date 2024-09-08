package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sdomino/scribble"
)

type User struct {
	Personid       string `json:"personid"`
	WeeklyIncome   string `json:"weeklyincome"`
	Employmenttype string `json:"employmenttype"`
}

type Transaction struct {
	TransactionID string `json:"transactionid"`
	Description   string `json:"description"`
	Datetime      string `json:"datetime"`
	Isdeduction   string `json:"Isdeduction"`
}

type counter struct {
	Transactioncounter int
}

var x, _ = scribble.New("./database", nil)

func individualtransactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ID := r.FormValue("id")
		Transactionreq := Transaction{}

		if err := x.Read("transactions", ID, &Transactionreq); err != nil {
			fmt.Println(err)
		}

		respio, _ := json.Marshal(Transactionreq)
		fmt.Fprintln(w, string(respio))

		fmt.Println(r.FormValue("testy"))
		fmt.Println(r.FormValue("tasty"))

	case http.MethodPost:
		Counter := counter{}
		if err := x.Read("config", "counter", &Counter); err != nil {
			fmt.Println(err)
		}

		ID := Counter.Transactioncounter + 1
		Counter.Transactioncounter = ID
		if err := x.Write("config", "counter", Counter); err != nil {
			fmt.Println(err)
		}

		Idstring := fmt.Sprintf("%d", ID)
		fmt.Printf("%T", Idstring)
		fmt.Println(Idstring)
		description := r.FormValue("description")
		datetimey := r.FormValue("datetime")
		isdeduction := r.FormValue("isdeduction")
		data := Transaction{TransactionID: Idstring, Description: description, Datetime: datetimey, Isdeduction: isdeduction}
		if err := x.Write("transactions", Idstring, data); err != nil {
			fmt.Println(err)
		}

	}
}

func returnuser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		user := User{}

		if err := x.Read("config", "user", &user); err != nil {
			fmt.Println(err)
		}

		respio, _ := json.Marshal(user)
		fmt.Fprintln(w, string(respio))

	case http.MethodPost:

		personid := r.FormValue("personid")
		weeklyincome := r.FormValue("weeklyincome")
		employmenttype := r.FormValue("employmenttype")
		data := User{Personid: personid, WeeklyIncome: weeklyincome, Employmenttype: employmenttype}
		if err := x.Write("config", "user", data); err != nil {
			fmt.Println(err)
		}

	}
}

func alltrans(w http.ResponseWriter, r *http.Request) {
	records, err := x.ReadAll("transactions")
	if err != nil {
		fmt.Println("Error", err)
	}

	fishies := []Transaction{}
	for _, f := range records {
		fishFound := Transaction{}
		if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
			fmt.Println("Error", err)
		}
		fishies = append(fishies, fishFound)
	}
	xx, _ := json.Marshal(fishies)
	fmt.Fprintln(w, string(xx))

}

func main() {

	http.HandleFunc("/individualtransactions/", individualtransactions)
	http.HandleFunc("/alltransactions/", alltrans)
	http.HandleFunc("/returnuser/", returnuser)

	fmt.Println("Running")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Error")
	}

}
