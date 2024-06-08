package util_test

import (
	"math"

	"github.com/fish-trainer/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing util", func() {
	DescribeTable("function MergeOrdered",
		func(in1 []int, in2 []int, expect []int) {
			var x []int = util.MergeOrdered(in1, in2)
			Expect(x).To(Equal(expect))
		},
		Entry("works with ints", []int{2, 3, 19}, []int{1, 4, 8, 20}, []int{1, 2, 3, 4, 8, 19, 20}),
		Entry("works if in1 empty", []int{}, []int{1, 4, 8, 20}, []int{1, 4, 8, 20}),
		Entry("works if in2 empty", []int{2, 3, 19}, []int{}, []int{2, 3, 19}),
	)
	DescribeTable("function MergeWithFuncSimple",
		func(in1 []int, in2 []int, expect []int) {
			var x []int = util.MergeWithFuncSimple(
				in1,
				in2,
				func(x int, y int) int {
					return x - y
				})
			Expect(x).To(Equal(expect))
		},
		Entry("works with ints", []int{2, 3, 19}, []int{1, 4, 8, 20}, []int{1, 2, 3, 4, 8, 19, 20}),
		Entry("works if in1 empty", []int{}, []int{1, 4, 8, 20}, []int{1, 4, 8, 20}),
		Entry("works if in2 empty", []int{2, 3, 19}, []int{}, []int{2, 3, 19}),
	)
	{
		type point struct {
			x float64
			y float64
		}
		type point_on_x_axis float64
		distance_f := func(p point, x point_on_x_axis) int {
			return int(math.Ceil(p.x - float64(x)))
		}
		conv1 := func(p point) float64 { return p.x }
		conv2 := func(x point_on_x_axis) float64 { return float64(x) }
		DescribeTable("function MergeWithFunc",
			func(in1 []point, in2 []point_on_x_axis, expect []float64) {
				var x []float64 = util.MergeWithFunc(
					in1,
					in2,
					distance_f,
					conv1,
					conv2,
				)
				Expect(x).To(Equal(expect))
			},
			Entry("works with ints", []point{{2.1, 3.5}, {4.1, 2.1}}, []point_on_x_axis{1.0, 4.0, 8.2, 20.0}, []float64{1.0, 2.1, 4.0, 4.1, 8.2, 20.0}),
			//Entry("works if in1 empty", []int{}, []int{1, 4, 8, 20}, []int{1, 4, 8, 20}),
			//Entry("works if in2 empty", []int{2, 3, 19}, []int{}, []int{2, 3, 19}),
		)
	}
})
