package main

import (
	"fmt"
	"testing"
)

var game *Game

func Setup() {
	game = NewGame()
}

func TestNewGame(t *testing.T) {
	Setup()

	assert(t, game.player1Score == 0 && game.player2Score == 0, fmt.Sprintf(`player 1 score or player 2 score is not zero (player 1 score is %d; player 2 score is %d)`, game.player1Score, game.player2Score))
}

func TestGetScore(t *testing.T) {
	Setup()

	score := game.getScore()

	assert(t, score == "Love All", fmt.Sprintf(`score != "Love All"; score == %s`, score))
}

func TestEvalScore(t *testing.T) {
	t.Run("Returns Love", func(t *testing.T) {
		score := 0

		result := evalScore(score)

		assert(t, result == "Love", fmt.Sprintf(`result != "Love"; result == %s`, result))
	})
}

func TestWinnerPlayerOne(t *testing.T) {
	t.Run("Returns Player One Wins", func(t *testing.T) {
		Setup()
		game.player1Score = 4
		game.player2Score = 2
	
		result := game.getWinner()

		assert(t, result == 1, fmt.Sprintf(`result != 1; result == %d`, result))
	})
}

func TestWinnerPlayerTwo(t *testing.T) {
	t.Run("Returns Player Two Wins", func(t *testing.T) {
		Setup()
		game.player1Score = 0
		game.player2Score = 4
	
		result := game.getWinner()

		assert(t, result == 2, fmt.Sprintf(`result != 2; result == %d`, result))
	})
}

func TestNoWinner(t *testing.T) {
	t.Run("Returns No Winner", func(t *testing.T) {
		Setup()
		game.player1Score = 3
		game.player2Score = 4
	
		result := game.getWinner()

		assert(t, result == 0, fmt.Sprintf(`result != 0; result == %d`, result))
	})
}

func assert(t testing.TB, condition bool, message string) {
	if !condition {
		t.Error(message)
	}
}
