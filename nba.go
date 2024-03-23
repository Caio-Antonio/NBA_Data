package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	PlayerID                         int     `json:"player_id"`
	PlayerName                       string  `json:"player_name"`
	Gp                               int     `json:"gp"`
	SecondsPlayed                    int     `json:"seconds_played"`
	ContestedShots2Pt                int     `json:"contested_shots_2pt"`
	ContestedShots3Pt                int     `json:"contested_shots_3pt"`
	Deflections                      int     `json:"deflections"`
	ChargesDrawn                     int     `json:"charges_drawn"`
	ScreenAstPts                     int     `json:"screen_ast_pts"`
	LooseBallsRecovered              int     `json:"loose_balls_recovered"`
	BoxOuts                          int     `json:"box_outs"`
	DistanceTraveled                 float64 `json:"distance_traveled"`
	AvgSpeed                         float64 `json:"avg_speed"`
	AvgSpeedZscore                   float64 `json:"avg_speed_zscore"`
	ContestedShots2PtPerGameZscore   float64 `json:"contested_shots_2pt_per_game_zscore"`
	ContestedShots3PtPerGameZscore   float64 `json:"contested_shots_3pt_per_game_zscore"`
	DeflectionsPerGameZscore         float64 `json:"deflections_per_game_zscore"`
	ChargesDrawnPerGameZscore        float64 `json:"charges_drawn_per_game_zscore"`
	ScreenAstPtsPerGameZscore        float64 `json:"screen_ast_pts_per_game_zscore"`
	LooseBallsRecoveredPerGameZscore float64 `json:"loose_balls_recovered_per_game_zscore"`
	BoxOutsPerGameZscore             float64 `json:"box_outs_per_game_zscore"`
	DistanceTraveledPerGameZscore    float64 `json:"distance_traveled_per_game_zscore"`
	ContestedShots2PtTotalZscore     float64 `json:"contested_shots_2pt_total_zscore"`
	ContestedShots3PtTotalZscore     float64 `json:"contested_shots_3pt_total_zscore"`
	DeflectionsTotalZscore           float64 `json:"deflections_total_zscore"`
	ChargesDrawnTotalZscore          float64 `json:"charges_drawn_total_zscore"`
	ScreenAstPtsTotalZscore          float64 `json:"screen_ast_pts_total_zscore"`
	LooseBallsRecoveredTotalZscore   float64 `json:"loose_balls_recovered_total_zscore"`
	BoxOutsTotalZscore               float64 `json:"box_outs_total_zscore"`
	DistanceTraveledTotalZscore      float64 `json:"distance_traveled_total_zscore"`
	ContestedShots2PtPer36Zscore     float64 `json:"contested_shots_2pt_per_36_zscore"`
	ContestedShots3PtPer36Zscore     float64 `json:"contested_shots_3pt_per_36_zscore"`
	DeflectionsPer36Zscore           float64 `json:"deflections_per_36_zscore"`
	ChargesDrawnPer36Zscore          float64 `json:"charges_drawn_per_36_zscore"`
	ScreenAstPtsPer36Zscore          float64 `json:"screen_ast_pts_per_36_zscore"`
	LooseBallsRecoveredPer36Zscore   float64 `json:"loose_balls_recovered_per_36_zscore"`
	BoxOutsPer36Zscore               float64 `json:"box_outs_per_36_zscore"`
	DistanceTraveledPer36Zscore      float64 `json:"distance_traveled_per_36_zscore"`
}

type Resposta struct {
	Data []Player `json:"data"`
}

func setPlayers(p *Player, resposta Resposta){
	for i := range resposta.Data{
		
		p = &resposta.Data[i]
		fmt.Println("ID:", p.PlayerID)
		fmt.Println("Nome:", p.PlayerName)
	}

}

func main(){
	res, _ := http.Get("https://technicalowl.herokuapp.com/playerhustlestats?seasontype=regular&season=2023")
	var resposta Resposta
	err := json.NewDecoder(res.Body).Decode(&resposta)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
	if len(resposta.Data) > 0 {
		primeiroItem := resposta.Data[0]
		fmt.Println("ID:", primeiroItem.PlayerID)
		fmt.Println("Nome:", primeiroItem.PlayerName)
	} else {
		fmt.Println("Lista de dados está vazia")
	}
	fmt.Println("Nome:", resposta.Data[1].PlayerName)
	//var p1 Player
	//setPlayers(&p1, resposta)
}
//O vetor com os players ja existe, é o resposta.data