package main

import "fmt"

type vector = []float64
type matrix []vector

func (m matrix) inverse() matrix {
	le := len(m)
	for _, v := range m {
		if len(v) != le {
			// 首先得是一个方阵
			panic("Not a square matrix")
		}
	}
	aug := make(matrix, le)
	for i := 0; i < le; i++ {
		// 毕竟是初等变换法
		aug[i] = make(vector, 2*le)
		copy(aug[i], m[i])
		// 同时另外一个矩阵的对角线一定是1（单位矩阵）
		aug[i][i+le] = 1
	}
	// 使其转化为简化行阶梯型矩阵(行初等变换)
	aug.toReducedRowEchelonForm()
	inv := make(matrix, le)

	for i := 0; i < le; i++ {
		inv[i] = make(vector, le)
		copy(inv[i], aug[i][le:])
	}
	return inv
}

func (m matrix) toReducedRowEchelonForm() {
	lead := 0
	rowCount, colCount := len(m), len(m[0])
	for r := 0; r < rowCount; r++ {
		if colCount <= lead {
			return
		}
		i := r

		for m[i][lead] == 0 {
			i++
			if rowCount == i {
				i = r
				lead++
				if colCount == lead {
					return
				}
			}
		}

		m[i], m[r] = m[r], m[i]
		if div := m[r][lead]; div != 0 {
			for j := 0; j < colCount; j++ {
				m[r][j] /= div
			}
		}

		for k := 0; k < rowCount; k++ {
			if k != r {
				mult := m[k][lead]
				for j := 0; j < colCount; j++ {
					m[k][j] -= m[r][j] * mult
				}
			}
		}
		lead++
	}
}

func (m matrix) print(title string) {
	fmt.Println(title)
	for _, v := range m {
		fmt.Printf("% f\n", v)
	}
	fmt.Println()
}

func main() {
	a := matrix{{1, 2, 3}, {4, 1, 6}, {7, 8, 9}}
	a.inverse().print("Inverse of A is:\n")

	b := matrix{{2, -1, 0}, {-1, 2, -1}, {0, -1, 2}}
	b.inverse().print("Inverse of B is:\n")
}
