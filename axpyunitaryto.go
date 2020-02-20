// +build !amd64 noasm

package vector

// This function is from the gonum repository:
// https://github.com/gonum/gonum/blob/c3867503e73e5c3fee7ab93e3c2c562eb2be8178/internal/asm/f64/axpy.go#L23
func axpyUnitaryTo(dst []float64, alpha float64, x, y []float64) {
	for i, v := range x {
		dst[i] = alpha*v + y[i]
	}
}
