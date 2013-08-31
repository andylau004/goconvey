package convey

import "fmt"

type expectation func(actual interface{}, expected []interface{}) string

func ShouldEqual(actual interface{}, expected []interface{}) string {
	switch {
	case len(expected) == 0:
		return fmt.Sprintf("This expectation requires a second value (only one provided: '%v').", actual)
	case len(expected) > 1:
		return fmt.Sprintf("This expectation only accepts 2 values to be compared (and %v were provided).", len(expected)+1)
	case actual != expected[0]:
		return fmt.Sprintf("'%v' should equal '%v' (but it doesn't)!", actual, expected[0])
	}
	return ""
}

func ShouldBeNil(actual interface{}, expected []interface{}) string {
	if actual != nil {
		return fmt.Sprintf("'%v' should have been nil!", actual)
	}
	return ""
}

/*

	// Equality
X	So(thing, ShouldEqual, thing2)
	So(thing, ShouldNotEqual, thing2)
	So(thing, ShouldMarshalLike, thing2)

	// Sameness
	So(thing, ShouldBe, thing2)
	So(thing, ShouldNotBe, thing2)
	So(thing, ShouldPointTo, thing2)
	So(thing, ShouldNotPointTo, thing2)
X	So(thing, ShouldBeNil, thing2)
	So(thing, ShouldNotBeNil, thing2)
	So(thing, ShouldBeTrue)
	SO(thing, ShouldBeFalse)

	// Interfaces
	So(1, ShouldImplement, Interface)
	So(1, ShouldNotImplement, OtherInterface)

	// Type checking
	So(1, ShouldBeAn, int)
	So(1, ShouldNotBeAn, int)
	So("1", ShouldBeA, string)
	So("1", ShouldNotBeA, string)

	// Quantity comparison
	So(1, ShouldBeGreaterThan, 0)
	So(1, ShouldBeGreaterThanOrEqualTo, 0)
	So(1, ShouldBeLessThan, 2)
	So(1, ShouldBeLessThanOrEqualTo, 2)

	// Tolerences
	So(1.1, ShouldBeWithin(.1).Of, 1)
	So(1.1, ShouldNotBeWithin(.1).Of, 2)

	// Collections
	So([]int{}, ShouldBeEmpty)
	So([]int{1}, ShouldNotBeEmpty)
	So([]int{1, 2, 3}, ShouldContain, 1)
	So([]int{1, 2, 3}, ShouldNotContain, 4)
	So(1, ShouldBeIn, []int{1, 2, 3})
	So(4, ShouldNotBeIn, []int{1, 2, 3})

	// Strings
	So("asdf", ShouldStartWith, "as")
	So("asdf", ShouldNotStartWith, "df")
	So("asdf", ShouldEndWith, "df")
	So("asdf", ShouldNotEndWith, "df")
	So("(asdf)", ShouldBeSurroundedWith, "()")
	So("(asdf)", ShouldNotBeSurroundedWith, "[]")

*/
