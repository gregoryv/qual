[![Build Status](https://travis-ci.org/gregoryv/asserter.svg?branch=master)](https://travis-ci.org/gregoryv/asserter)
[![codecov](https://codecov.io/gh/gregoryv/asserter/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/asserter)

[asserter](https://godoc.org/github.com/gregoryv/asserter) - Go package oneline assertions

## Quick start

    go get github.com/gregoryv/asserter

In your tests

    func Test_something(t *testing.T) {
       assert := asserter.New(t)
       got, err := something()
       assert(err == nil).Fatal(err)
       assert(got == exp).Errorf("%v, expected %v", got, exp)
	   // or the shorter version when checking got vs. expected
	   assert().Equals(got, exp)
	   // and with optional case message
	   assert().Equals(got, exp, "with no arguments")
    }
