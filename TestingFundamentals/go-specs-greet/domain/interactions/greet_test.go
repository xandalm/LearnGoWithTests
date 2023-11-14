package interactions_test

import (
	"testing"

	"github.com/xandalm/go-specs-greet/domain/interactions"
	"github.com/xandalm/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
