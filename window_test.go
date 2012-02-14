package window

import "testing"
import "container/list"

func BenchmarkS1000M1(b *testing.B) {
    m(1000,1,b)
}

func BenchmarkS1000M10(b *testing.B) {
    m(1000,10,b)
}

func BenchmarkS1000M100(b *testing.B) {
    m(1000,100,b)
}

func BenchmarkS1000M500(b *testing.B) {
    m(1000,500,b)
}

// m will, given the size and multiple, run
// 1000 times the size worth of data through
// the moving window and generate a slice 
// each time
func m(size, multiple int, b *testing.B) {
    b.StopTimer()
    w := New(size, multiple)
    TIMES := 1000*size
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < TIMES; j++ {
            w.PushBack(float64(i))
            s := w.Slice()
            // we'll throw in extra processing,
            // just in case of optimization
            s[0] = 1.
        }
    }
}

func BenchmarkListS1000(b *testing.B) {
    l(1000,b)
}

func l(size int, b *testing.B) {
    b.StopTimer()
    lst := list.New()
    TIMES := 1000*size
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < TIMES; j++ {
            lst.PushBack(float64(j))
            if (lst.Len() > size) {
                lst.Remove(lst.Front())
            }
            s := make([]float64, 0, lst.Len())
            for e := lst.Front(); e != nil; e = e.Next() {
                s = append(s, e.Value.(float64))
            }
        }
    }
}
