package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
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

func main() {

	start := time.Now()

	// byteValue := Exctract("01012100")
	byteValue := mock()

	var response Response
	json.Unmarshal(byteValue, &response)

	fmt.Println(response.Code)
	fmt.Println(response.Message)

	for _, insw := range response.Data {
		fmt.Println(insw.HsCode)

		for _, mfn := range insw.Mfn {
			for _, bm := range mfn.Bm {
				fmt.Println("BM:", bm.Value)
			}
			for _, ppn := range mfn.Ppn {
				fmt.Println("PPN:", ppn.Value)
			}
			for _, pph := range mfn.Pph {
				fmt.Println("PPH:", pph.Value)
			}
		}
	}

	fmt.Printf("\n\nDurasi menarik data INSW: %f detik\n", time.Since(start).Seconds())
}

func mock() []byte {
	file, err := os.Open("hscode.json")

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()
	byteValue, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return byteValue
}
