package main

import (
	"errors"
	"fmt"
)

func main() {

	x := []float64{(660.0 + 57.0) / 2, (581.0 + 65.0) / 2, (559.0 + 73.0) / 2, (512.0 + 84.0) / 2, (427.0 + 23.0) / 2, (808.0 + 118.0) / 2, (782.0 + 206.0) / 2, (1057.0 + 388.0) / 2, (1080.0 + 441.0) / 2, (1021.0 + 183.0) / 2, (967.0 + 205.0) / 2, (1027.0 + 195.0) / 2, (1011.0 + 357.0) / 2}
	y := []float64{1.85 / (660 - 57), 1.75 / (581 - 65), 1.75 / (559 - 73), 1.65 / (512 - 84), 1.75 / (427 - 23),
		1.85 / (808 - 118), 1.65 / (782 - 206), 1.7 / (1057 - 388), 1.7 / (1080 - 441), 1.87 / (1021 - 183), 1.87 / (967 - 205), 1.87 / (1027 - 195), 1.7 / (1011 - 357)}
	fmt.Println(UnaryFit(x, y, 2))

}

// 多项式拟合，n为拟合多项式的次数
func UnaryFit(x, y []float64, n int) (Unary, error) {
	i, j, l := len(x), len(y), 0
	if i < j {
		l = i
	} else {
		l = j
	}
	if n++; n <= 0 {
		return nil, errors.New("Illegal input n")
	}
	if n > l {
		return nil, errors.New("Data-set too small")
	}
	group := make([]Linear, n)
	for i := 0; i < n; i++ {
		group[i] = make(Linear, n+1)
	}
	for t := 0; t < l; t++ {
		X, Y := x[t], y[t]
		for i, p := 0, 1.0; i < n; i, p = i+1, p*X {
			for j, q := 0, p; j < n; j, q = j+1, q*X {
				group[i][j] += q
			}
			group[i][n] -= p * Y
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(group[i])
	}
	ans, err := SolveLinearGroup(group...)
	if err != nil {
		return nil, err
	}
	return Unary(ans), nil
}
