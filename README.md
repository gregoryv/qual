[![Build Status](https://travis-ci.org/gregoryv/qual.svg?branch=master)](https://travis-ci.org/gregoryv/qual)
[![codecov](https://codecov.io/gh/gregoryv/qual/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/qual)
[![Maintainability](https://api.codeclimate.com/v1/badges/83083a5e52d4ffad3288/maintainability)](https://codeclimate.com/github/gregoryv/qual/maintainability)


[qual](https://godoc.org/github.com/gregoryv/qual) - Go package for quality assessment at source code level

## Quick start

    go get github.com/gregoryv/qual

Add a unit test to your project

    func Test_CodeQuality(t *testing.T) {
	    qual.Standard(t)
	}

or if you are really brave, do

    func Test_CodeQuality(t *testing.T) {
	    qual.High(t)
	}

The predefined tests measure

- code complexity
- line width

where the `qual.Standard` test only checks your package and `qual.High`
also includes vendored code.

## Scope

This package helps to improve code readability. Note, it's not
measuring the quality of features in your solution. Those are better
measured with benchmarks and user experience. The quality of the
source code helps developers pass on their intent of their solution to
other developers. We do this in various ways

- documenting
- tests and examples and most importantly
- writing readable code

If we can write readable code, documentation can be minimized and that
is always good, since developers tend to prefer writing code over
documentation.

## Assert

This package also provides an assert func for writing clear and
concise unittests. Test code in itself if written properly is readable
and very contained. It makes a good description in itself if the test
should ever fail. However most packages try to solve this by providing
short named func for operators within the language. E.g.

    assert.Equals(t, a, b, "a and b be should be equal")

This reads fairly well but if you consider adding a couple more of these
statements, readability starts to suffer

    a, b, c, d := SomeFancyFunc()
    assert.Equals(t, a, b, "a and b be should be equal")
    assert.NotEquals(t, a, c, "a and c should be different")
    assert.NotNil(t, d, "d should not be nil")

and get repetitive.
Using the Assert func you can achieve the same thing with

    a, b, c, d := SomeFancyFunc()
    Assert(t, Vars{a,b,c,d},
	  a == b,
	  a != c,
	  d != nil,
	)

As you can see there are no textual descriptions of the different
assertions. We don't have to describe `a == b` with `a and b should be
equal` that is how we read the operator `==`.  Let's assume the above
test fails on the assertion `a != c` then the error would be printed
as

    > a, b, c, d := SomeFancyFunc()
	  failed assert: a != c
        a = 1
        b = 1
        c = 1
        d = nil

What happens is the assert method parses out the context it's called in
and prints the nearest line above it, so we know which statement we are
actually testing. After that the failed assertions are listed and finally
all the vars that are of interest are printed.
More examples are provided in [godoc Assert](http://godoc.org/github.com/gregoryv/qual#Assert)
