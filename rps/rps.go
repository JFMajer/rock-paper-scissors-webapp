package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper (paper + 1) % 3 = 2
	DRAW         = 0
	PLAYERWINS   = 1
	COMPUTERWINS = 2
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0
	switch computerValue {
	case 1:
		computerChoice = "paper"
	case 2:
		computerChoice = "scissors"
	case 0:
		computerChoice = "rock"
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins"
		winner = COMPUTERWINS
	}

	return Round{Winner: winner, ComputerChoice: computerChoice, RoundResult: roundResult}
}
