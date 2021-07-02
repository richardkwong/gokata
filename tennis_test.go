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

func assert(t testing.TB, condition bool, message string) {
	if !condition {
		t.Error(message)
	}
}
