package interactions_test

import (
	"testing"

	"github.com/xandalm/go-specs-greet/domain/interactions"
	"github.com/xandalm/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		got := interactions.Greet("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
