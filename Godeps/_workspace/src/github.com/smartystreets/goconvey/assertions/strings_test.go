package assertions

import (
	"testing"
)

func TestShouldStartWith(t *testing.T) {
	fail(t, so("", ShouldStartWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("", ShouldStartWith, "asdf", "asdf"), "This assertion requires exactly 1 comparison values (you provided 2).")

	pass(t, so("", ShouldStartWith, ""))
	pass(t, so("superman", ShouldStartWith, "super"))
	fail(t, so("superman", ShouldStartWith, "bat"), "Expected 'superman' to start with 'bat' (but it didn't)!")
	fail(t, so("superman", ShouldStartWith, "man"), "Expected 'superman' to start with 'man' (but it didn't)!")

	fail(t, so(1, ShouldStartWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func TestShouldNotStartWith(t *testing.T) {
	fail(t, so("", ShouldNotStartWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("", ShouldNotStartWith, "asdf", "asdf"), "This assertion requires exactly 1 comparison values (you provided 2).")

	fail(t, so("", ShouldNotStartWith, ""), "Expected '<empty>' NOT to start with '<empty>' (but it did)!")
	fail(t, so("superman", ShouldNotStartWith, "super"), "Expected 'superman' NOT to start with 'super' (but it did)!")
	pass(t, so("superman", ShouldNotStartWith, "bat"))
	pass(t, so("superman", ShouldNotStartWith, "man"))

	fail(t, so(1, ShouldNotStartWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func TestShouldEndWith(t *testing.T) {
	fail(t, so("", ShouldEndWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("", ShouldEndWith, "", ""), "This assertion requires exactly 1 comparison values (you provided 2).")

	pass(t, so("", ShouldEndWith, ""))
	pass(t, so("superman", ShouldEndWith, "man"))
	fail(t, so("superman", ShouldEndWith, "super"), "Expected 'superman' to end with 'super' (but it didn't)!")
	fail(t, so("superman", ShouldEndWith, "blah"), "Expected 'superman' to end with 'blah' (but it didn't)!")

	fail(t, so(1, ShouldEndWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func TestShouldNotEndWith(t *testing.T) {
	fail(t, so("", ShouldNotEndWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("", ShouldNotEndWith, "", ""), "This assertion requires exactly 1 comparison values (you provided 2).")

	fail(t, so("", ShouldNotEndWith, ""), "Expected '<empty>' NOT to end with '<empty>' (but it did)!")
	fail(t, so("superman", ShouldNotEndWith, "man"), "Expected 'superman' NOT to end with 'man' (but it did)!")
	pass(t, so("superman", ShouldNotEndWith, "super"))

	fail(t, so(1, ShouldNotEndWith, 2), "Both arguments to this assertion must be strings (you provided int and int).")
}

func TestShouldContainSubstring(t *testing.T) {
	fail(t, so("asdf", ShouldContainSubstring), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("asdf", ShouldContainSubstring, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fail(t, so(123, ShouldContainSubstring, 23), "Both arguments to this assertion must be strings (you provided int and int).")

	pass(t, so("asdf", ShouldContainSubstring, "sd"))
	fail(t, so("qwer", ShouldContainSubstring, "sd"), "Expected 'qwer' to contain substring 'sd' (but it didn't)!")
}

func TestShouldNotContainSubstring(t *testing.T) {
	fail(t, so("asdf", ShouldNotContainSubstring), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("asdf", ShouldNotContainSubstring, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fail(t, so(123, ShouldNotContainSubstring, 23), "Both arguments to this assertion must be strings (you provided int and int).")

	pass(t, so("qwer", ShouldNotContainSubstring, "sd"))
	fail(t, so("asdf", ShouldNotContainSubstring, "sd"), "Expected 'asdf' NOT to contain substring 'sd' (but it didn't)!")
}

func TestShouldBeBlank(t *testing.T) {
	fail(t, so("", ShouldBeBlank, "adsf"), "This assertion requires exactly 0 comparison values (you provided 1).")
	fail(t, so(1, ShouldBeBlank), "The argument to this assertion must be a string (you provided int).")

	fail(t, so("asdf", ShouldBeBlank), "Expected 'asdf' to be blank (but it wasn't)!")
	pass(t, so("", ShouldBeBlank))
}

func TestShouldNotBeBlank(t *testing.T) {
	fail(t, so("", ShouldNotBeBlank, "adsf"), "This assertion requires exactly 0 comparison values (you provided 1).")
	fail(t, so(1, ShouldNotBeBlank), "The argument to this assertion must be a string (you provided int).")

	fail(t, so("", ShouldNotBeBlank), "Expected value to NOT be blank (but it was)!")
	pass(t, so("asdf", ShouldNotBeBlank))
}
