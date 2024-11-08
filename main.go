package main

import (
	"encoding/json"
	"fmt"
)

// type Response struct {
// 	Code    string `json:"code"`
// 	Message string `json:"message"`
// 	Data    []Data `json:"data"`
// }

// type Data struct {
// 	HsCode string `json:"hs_code"`
// 	Bab    string `json:"bab"`
// 	// NewMfn []NewMfn `json:"new_mfn"`
// }

type BM struct {
	BM string `json:"bm"`
}

type PPN struct {
	PPN string `json:"ppn"`
}

type PPH struct {
	PPH string `json:"pph"`
}
type NewMFN struct {
	Regulation  string `json:"regulation"`
	IssuedAt    string `json:"issued_at"`
	EffectiveAt string `json:"effective_at"`

	BM  []BM  `json:"bm"`
	PPN []PPN `json:"ppn"`
	PPH []PPH `json:"pph"`
}

type Response struct {
	NewMFN []NewMFN `json:"new_mfn"`
}

func main() {
	// Exctract("01012100")
	// file, err := os.Open("hscode.json")
	// if err != nil {
	// 	log.Fatalf("Error opening file: %v", err)
	// }
	// defer file.Close()

	byteValue := Exctract("01012100")

	var response Response
	json.Unmarshal(byteValue, &response)

	fmt.Println(response)

	for _, newMFN := range response.NewMFN {
		fmt.Println("Regulation:", newMFN.Regulation)
		fmt.Println("Issued At:", newMFN.IssuedAt)
		fmt.Println("Effective At:", newMFN.EffectiveAt)
		for _, bm := range newMFN.BM {
			fmt.Println("BM:", bm.BM)

		}
		for _, ppn := range newMFN.PPN {
			fmt.Println("PPN:", ppn.PPN)

		}
		for _, pph := range newMFN.PPH {
			fmt.Println("PPH:", pph.PPH)
		}
	}
}
