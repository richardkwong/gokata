package main

/*

Mobbing Exercise

https://www.cyber-dojo.org/creator/choose_problem?type=group

You task is to implement a tennis scoring program.
Summary of tennis scoring:

1.
A game is won by the first player to have won at least four points in total and at least two points more than the opponent.

2.
The running score of each game is described in a manner peculiar to tennis: scores from zero to three points are described as "love", "fifteen", "thirty", and "forty" respectively.

3.
If at least three points have been scored by each player, and the scores are equal, the score is "deuce".

4.
If at least three points have been scored by each side and a player has one more point than his opponent, the score of the game is "advantage" for the player in the lead.

[source http://en.wikipedia.org/wiki/Tennis#Scoring]

*/

type Game struct {
	player1Score int
	player2Score int
}

func main() {
	// game := NewGame()
}

func NewGame() *Game {
	return &Game{player1Score: 0, player2Score: 0}
}

func (g *Game) getScore() string {

	player1Score := evalScore(g.player1Score)
	player2Score := evalScore(g.player2Score)

	if player1Score == player2Score {
		return player1Score + " All"
	}

	return ""
}

func evalScore(score int) string {
	switch score {
	case 0:
		return "Love"
	default:
		return ""
	}
}

func (g *Game) getWinner() int {
	if (g.player1Score >= 4 && (g.player1Score - g.player2Score >= 2)) {
		return 1
	} else if (g.player2Score >= 4 && (g.player2Score - g.player1Score >= 2)) {
		return 2
	}
	
	return 0
}
