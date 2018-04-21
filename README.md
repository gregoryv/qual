[![Build Status](https://travis-ci.org/gregoryv/qual.svg?branch=master)](https://travis-ci.org/gregoryv/qual)
[![codecov](https://codecov.io/gh/gregoryv/qual/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/qual)
[![Maintainability](https://api.codeclimate.com/v1/badges/83083a5e52d4ffad3288/maintainability)](https://codeclimate.com/github/gregoryv/qual/maintainability)


[qual](https://godoc.org/github.com/gregoryv/qual) - Go package for quality assessment at src level

## Quick start

    go get github.com/gregoryv/qual

Add a unit test to your project

    func Test_CodeQuality(t *testing.T) {
	    qual.Standard(t)
	}

or use a more custom set of metrics

    func Test_CodeQualtiy(t *testing.T) {
	    //  < 6 good, ..., > 9 is bad
	    CyclomaticComplexity(5, false, t)
        // < 81 good, ..., > 100 is bad
        SourceWidth(80, false, t)
    }
