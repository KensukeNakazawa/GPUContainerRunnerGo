package gpu

import "testing"

func TestRound(t *testing.T) {
	t.Run("小数第3位以下が切り捨てられること", func(t *testing.T) {
		target := 1.234
		want := 1.23
		if got := Round(target, 2); got != want {
			t.Errorf("add() = %v, want %v", got, want)
		}
	})
	t.Run("小数第3位以下が切り上げられること", func(t *testing.T) {
		target := 1.235
		want := 1.24
		if got := Round(target, 2); got != want {
			t.Errorf("add() = %v, want %v", got, want)
		}
	})
}
