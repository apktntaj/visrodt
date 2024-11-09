package main

import (
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []Insw `json:"data"`
}

type Insw struct {
	HsCode string `json:"hs_code"`
	Mfn    []Mfn  `json:"mfn"`
}

type Mfn struct {
	Bm  []Bm  `json:"bm"`
	Ppn []Ppn `json:"ppn"`
	Pph []Pph `json:"pph"`
}

type Bm struct {
	Value string `json:"bm"`
}

type Ppn struct {
	Value string `json:"ppn"`
}

type Pph struct {
	Value string `json:"pph"`
}

// TODO: Refactor your smelly code

func main() {

	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/", PostHandler)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func mock() []byte {
// 	file, err := os.Open("hscode.json")

// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}

// 	defer file.Close()
// 	byteValue, err := io.ReadAll(file)

// 	if err != nil {
// 		log.Fatalf("Error reading file: %v", err)
// 	}

// 	return byteValue
// }
