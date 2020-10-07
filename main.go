package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la página principal!")
	fmt.Println("Endpoint homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/prueba", returnPrueba).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

type Form_struct struct {
	DataForm [][]int `json:"dataform"`
}

type Respuesta struct {
	ArrayValido    string
	ArrayOriginal  [][]int
	ArrayProcesado [][]int
}

func returnPrueba(w http.ResponseWriter, r *http.Request) {
	var Form Form_struct
	err := json.NewDecoder(r.Body).Decode(&Form)
	if err != nil {
		fmt.Fprintln(w, "Error al leer %s", err)
	}

	var arrays = Form.DataForm
	arrayValido := true
	_ = arrayValido
	var hLen int = len(arrays)
	if hLen > 0 {
		for _, array := range arrays {
			if len(array) != hLen {
				arrayValido = false
			}
		}
	} else {
		arrayValido = false
	}

	var arrayRotate [][]int

	if arrayValido == true {
		arrayRotate = RotarArray(arrays, hLen)
		_ = arrayRotate
	}

	var strArrayValido string
	if arrayValido == true {
		strArrayValido = "Array válido"
	} else {
		strArrayValido = "Array inválido"
	}

	var respuesta = Respuesta{strArrayValido, arrays, arrayRotate}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}

func RotarArray(array [][]int, n int) [][]int {
	var newArray [][]int
	newArray = make([][]int, n)
	for i := 0; i < n; i++ {
		newArray[i] = make([]int, n)
		for j := 0; j < n; j++ {
			newArray[i][j] = array[n-j-1][i]
		}
	}

	fmt.Println(newArray)

	return newArray
}

func main() {
	handleRequests()
}
