[qual](https://godoc.org/github.com/gregoryv/qual) - Go package for quality assessment at src level

## Quick start

    go get github.com/gregoryv/qual

Add a unit test to your project

    func Test_CodeQuality(t *testing.T) {
	    //  < 5 good, ..., > 9 is bad
	    CyclomaticComplexity(5, false, t)
		// < 80 good, ..., > 100 is bad
		SourceWidth(80, false, t)
    }
