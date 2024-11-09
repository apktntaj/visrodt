package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Struct untuk data yang akan diterima dari request POST
type HsCodeData struct {
	HsCode string `json:"hs_code"`
}

// Handler function untuk menangani request POST
func PostHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Pastikan metode request adalah POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse body dari request POST ke dalam struct
	var data []HsCodeData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	for _, hsCode := range data {
		fmt.Printf("Received hs_code: %s\n", hsCode.HsCode)
		byteValue := INTR(hsCode.HsCode)

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
				for i, pph := range mfn.Pph {
					ket := "(API)"
					if i == 1 {
						ket = "(NON-API)"
					}
					fmt.Println("PPH:", pph.Value+" "+ket)
				}
			}
		}

		fmt.Printf("\n\nDurasi menarik data INSW: %f detik\n", time.Since(start).Seconds())

	}

	// Kirimkan response kembali ke client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success"}
	json.NewEncoder(w).Encode(response)
}
