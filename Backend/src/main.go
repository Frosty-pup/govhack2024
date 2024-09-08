package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sdomino/scribble"
)

type TrialPayload struct {
	ID    int
	Data1 string
	Data2 string
}

var x, _ = scribble.New("./database", nil)

func respond(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		out := TrialPayload{ID: 123, Data1: "Data1component", Data2: "Data2 cp,[jiodgkl;sdfjgl;]"}
		respio, _ := json.Marshal(out)
		fmt.Fprintln(w, string(respio))

		fmt.Println(r.FormValue("testy"))
		fmt.Println(r.FormValue("tasty"))

	case http.MethodPost:
		ID := 1
		Idstring := fmt.Sprintf("%d", ID)
		fmt.Printf("%T", Idstring)
		testy := r.FormValue("testy")
		tasty := r.FormValue("tasty")
		data := TrialPayload{ID: ID, Data1: testy, Data2: tasty}
		if err := x.Write("transactions", Idstring, data); err != nil {
			fmt.Println(err)
		}

	}

}

func Interaction() {
	run := true
	scanner := bufio.NewScanner(os.Stdin)

	for {
		for scanner.Scan() {
			if !run {
				break
			}

			input := scanner.Text()
			if input == "end" {
				run = false
			}
			fmt.Println("I heard ya!")
		}

	}

}

func main() {

	go Interaction()

	http.HandleFunc("/tom/", respond)

	fmt.Println("Running")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Error")
	}

}
