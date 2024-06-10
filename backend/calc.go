package calc

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

type Coeficients struct {
	AvgSpeedCoef int16 `json:"AvgSpeedCoef"`
	BoxOutsCoef int16 `json:"BoxOutsCoef"`
	ChargesCoef int16 `json:"ChargesCoef"`
	ContestedShots2PtCoef int16 `json:"ContestedShots2PtCoef"`
	ContestedShots3PtCoef int16 `json:"ContestedShots3PtCoef"`
	DeflectionsCoef int16 `json:"DeflectionsCoef"`
	DistanceTraveledCoef int16 `json:"DistanceTraveledCoef"`
	LooseBallsRecoveredCoef int16 `json:"LooseBallsRecoveredCoef"`
	ScreenAstPtsCoef int16 `json:"ScreenAstPtsCoef"`
}

type Resposta struct {
	Data []Player `json:"data"`
}

func setPlayers(p *Player, players Resposta){
	for i := range players.Data{
		
		p = &players.Data[i]
		fmt.Println("ID:", p.PlayerID)
		fmt.Println("Nome:", p.PlayerName)
	}

}

func AlterarStats(p Player, coef Coeficients){
	p.AvgSpeed = p.AvgSpeed * float64(coef.AvgSpeedCoef)
	p.BoxOuts = p.BoxOuts * int(coef.BoxOutsCoef)
	p.ChargesDrawn = p.ChargesDrawn * int(coef.ChargesCoef)
	p.ContestedShots2Pt = p.ContestedShots2Pt * int(coef.ContestedShots2PtCoef)
	p.ContestedShots3Pt = p.ContestedShots3Pt * int(coef.ContestedShots3PtCoef)
	p.Deflections = p.Deflections * int(coef.DeflectionsCoef)
	p.DistanceTraveled = p.DistanceTraveled * float64(coef.DistanceTraveledCoef)
	p.LooseBallsRecovered = p.LooseBallsRecovered * int(coef.LooseBallsRecoveredCoef)
	p.ScreenAstPts = p.ScreenAstPts * int(coef.ScreenAstPtsCoef)
	fmt.Println(p.PlayerName)
	fmt.Println(p.AvgSpeed)
	fmt.Println(p.BoxOuts)
	fmt.Println(p.ChargesDrawn)
	fmt.Println(p.ContestedShots2Pt)
	fmt.Println(p.ContestedShots3Pt)
	fmt.Println(p.Deflections)
	fmt.Println(p.DistanceTraveled)
	fmt.Println(p.LooseBallsRecovered)
	fmt.Println(p.ScreenAstPts)
}

func main(){
	res, _ := http.Get("https://technicalowl.herokuapp.com/playerhustlestats?seasontype=regular&season=2023")
	var players Resposta
	err := json.NewDecoder(res.Body).Decode(&players)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
	// if len(players.Data) > 0 {
	// 	primeiroItem := players.Data[0]
	// 	fmt.Println("ID:", primeiroItem.PlayerID)
	// 	fmt.Println("Nome:", primeiroItem.PlayerName)
	// } else {
	// 	fmt.Println("Lista de dados está vazia")
	// }
	//fmt.Println("Nome:", players.Data[1].PlayerName)
	//var p1 Player
	//setPlayers(&p1, resposta)
	//AlterarStats(players.Data[0], 3)
}
//O vetor com os players ja existe, é o resposta.data
//todo - logica para conseguir os inputs dos coeficientes via post e multiplicar