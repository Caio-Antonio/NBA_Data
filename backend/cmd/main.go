package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	calc "github.com/Caio-Antonio/NBA_Data"
)

var players calc.Resposta

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
	body, _ := io.ReadAll(r.Body)
	coef, _ := strconv.Atoi(string(body))
	calc.AlterarStats(players.Data[0], coef)
	fmt.Println("hello world")
}

func main (){
	getPlayersStats()
	//calc.AlterarStats(players.Data[0], 3)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}