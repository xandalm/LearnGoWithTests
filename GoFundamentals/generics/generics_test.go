package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(Stack[int])

		// check stack is empty
		AssertTrue(t, stackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		// add another thing, pop it back again
		stackOfInts.Push(456)
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, stackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})

	// t.Run("interface stack dx is horrid", func(t *testing.T) {
	// 	myStackOfInts := new(StackOfInts)

	// 	myStackOfInts.Push(1)
	// 	myStackOfInts.Push(2)
	// 	firstNum, _ := myStackOfInts.Pop()
	// 	secondNum, _ := myStackOfInts.Pop()

	// 	// get our ints from out interface{}
	// 	reallyFirstNum, ok := firstNum.(int)
	// 	AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

	// 	reallySecondNum, ok := secondNum.(int)
	// 	AssertTrue(t, ok) // and again!

	// 	AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	// })
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
