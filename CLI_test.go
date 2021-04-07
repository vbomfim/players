package poker_test

import (
	"strings"
	"testing"

	poker "github.com/vbomfim/players"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		RecordPlayerWinFromUserInput(t, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		RecordPlayerWinFromUserInput(t, "Cleo")
	})
}

func RecordPlayerWinFromUserInput(t testing.TB, player string) {
	in := strings.NewReader(player + " wins\n")
	playerStore := &poker.StubPlayerStore{}
	cli := &poker.CLI{playerStore, in}
	cli.PlayPoker()

	poker.AssertPlayerWin(t, playerStore, player)
}
