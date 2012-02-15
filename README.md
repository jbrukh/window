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

To run the benchmarks:

    $ go test -test.bench=".*"

### Documentation

See the comments in the codebase.

### Benchmarks

Here are some benchmarks:

     BenchmarkListS1000             10         188497800 ns/op
     BenchmarkRingS1000             50          35760520 ns/op
     BenchmarkS1000M1                1        2817453000 ns/op
     BenchmarkS1000M10             200           9012795 ns/op
     BenchmarkS1000M100            200           8807370 ns/op
     BenchmarkS1000M500            200           8471190 ns/op
     BenchmarkSlicifyList        50000             41656 ns/op
     BenchmarkSlicifyRing        50000             48598 ns/op
     BenchmarkSlicifyWindow  200000000              9.30 ns/op

The first set of benchmarks measures the time it takes for 1,000,000 points to go through a size 1000 list-based moving window, ring-based moving window, and a MovingWindow under different values of its space parameter M. The later test measures how long it takes to convert the moving windows into a slice, a common operation.
