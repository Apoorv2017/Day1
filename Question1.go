package main

import "fmt"

type Matrix struct {
	rows    int
	columns int
	matrix  [][]int
}

func (m Matrix) get_rows() int {
	return m.rows

}
func (m Matrix) get_columns() int {
	return m.columns

}
func (m *Matrix) set_entry(val, i, j int) {
	m.matrix[i][j] = val
}
func (m1 Matrix) add(m2 Matrix) [][]int {
	temp := make([][]int, m1.rows)
	for i := range temp {
		temp[i] = make([]int, m1.columns)
	}
	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m1.columns; j++ {
			temp[i][j] = m1.matrix[i][j] + m2.matrix[i][j]
		}
	}
	return temp

}
func (m Matrix) print() {
	for i := 0; i < m.rows; i++ {
		fmt.Println(m.matrix[i])
	}

}

func main() {
	m1 := Matrix{1, 3, [][]int{{1, 2, 3}}}
	m2 := Matrix{1, 3, [][]int{{1, 2, 3}}}
	m1.print()
	m2.print()
	fmt.Println(m1.get_rows())
	fmt.Println(m1.get_columns())
	fmt.Println(m1.add(m2))
	m1.set_entry(4, 0, 2)
	m1.print()

}
