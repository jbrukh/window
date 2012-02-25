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
    $ go install -v github.com/jbrukh/window

To run the benchmarks:

    $ go test -test.bench=".*"

### Documentation

A moving window is a FIFO data structure with a fixed maximum size and the property that if the size has been reached and an element is pushed into the window, then the head of the window is pushed out to accomodate it.

Typically, two operations are of importance whilst working with moving windows in a high-performance context. First, pushes are expected to happen often as lots of data comes in. Secondly, the window is typically converted to a simple array for interfacing with secondary processes, such as graphing.

The MovingWindow object implemented in this package optimizes both of these operations by trading off with space complexity. Given two parameters -- size `S` and space parameter `M` -- the window will allocate a single backing array of length `SM` and complexity of background copy operations will be proportional to `1/M`.

As demonstrated in the benchmarks below, a MovingWindow can operate approximately 20x faster than a list-backed implementation and 3-4x faster than a ring-backed implementation. However, it is 5000x more efficient at generating array output because it can take slices of the backing array. It is also worth noting that most of the performance gain is reached quickly with relatively small `M`: for instance, setting `M := S/100` is a good starting point.

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
