[![Build Status](https://travis-ci.org/gregoryv/qual.svg?branch=master)](https://travis-ci.org/gregoryv/qual)
[![codecov](https://codecov.io/gh/gregoryv/qual/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/qual)
[![Maintainability](https://api.codeclimate.com/v1/badges/83083a5e52d4ffad3288/maintainability)](https://codeclimate.com/github/gregoryv/qual/maintainability)


[qual](https://godoc.org/github.com/gregoryv/qual) - Quality assessment at source code level

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

