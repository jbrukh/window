## Moving Window Data Structure

Copyright (c) 2012. Jake Brukhman. (jbrukh@gmail.com).
All rights reserved.  See the LICENSE file for BSD-style
license.

------------

### Installation

You need at least the following weekly version of Go:

    go version weekly.2012-02-07 +52ba9506bd99

You can then use the 'go' command to obtain the package:

    $ go get github.com/jbrukh/window
    $ go install -v github.com/jbrukh/widnow

### Documentation

See the comments in the codebase.

### Benchmarks

Here are some benchmarks:

    BenchmarkListS1000         1   42649824000 ns/op
    BenchmarkS1000M1           1    2707254000 ns/op
    BenchmarkS1000M10        100      10864690 ns/op
    BenchmarkS1000M100       100      10774190 ns/op
    BenchmarkS1000M500       100      11953970 ns/op

The first test is running 1,000,000 data points through a 1000-point moving window implemented with a list. On each insert, a slice is generated. The subsequent tests run the same number of data points through a moving window with different size and M combinations (as above). On each insert, a slice is generated as well. See `window_test.go` for details.
