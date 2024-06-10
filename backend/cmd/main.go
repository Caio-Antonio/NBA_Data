package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	calc "github.com/Caio-Antonio/NBA_Data"
)

var players calc.Resposta

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}


func getPlayersStats(){
	res, _ := http.Get("https://technicalowl.herokuapp.com/playerhustlestats?seasontype=regular&season=2023")
	var auxPlayers calc.Resposta
	err := json.NewDecoder(res.Body).Decode(&auxPlayers)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
	players = auxPlayers
}

func Hello(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var coef calc.Coeficients
	err := json.NewDecoder(r.Body).Decode(&coef)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	calc.AlterarStats(players.Data[0], coef)
	fmt.Println("hello world")
}

func main (){
	getPlayersStats()
	//calc.AlterarStats(players.Data[0], 3)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}