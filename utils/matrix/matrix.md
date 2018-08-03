 
 ```go

package matrix

import "math/big"

type Matrix struct {
	rows, columns int        // the number of rows and columns.
	data          []*big.Int // the contents of the matrix as one long slice.
}

// Set lets you define the value of a matrix at the given row and
// column.

// Set 允许您在给定的行和列中定义矩阵的值。
func (A *Matrix) Set(r int, c int, val *big.Int) {

	A.data[findIndex(r, c, A)] = val
}

// Get retrieves the contents of the matrix at the row and column.
// Get 检索行和列中矩阵的内容
func (A *Matrix) Get(r, c int) *big.Int {

	return A.data[findIndex(r, c, A)]
}

// Column returns a slice that represents a column from the matrix.
// This works by examining each row, and adding the nth element of
// each to the column slice.

// Column 返回表示矩阵中的列的切片。通过检查每一行, 并添加第 n 个元素的
// 每个列切片。
func (A *Matrix) Column(n int) []*big.Int {
	col := make([]*big.Int, A.rows)

	for i := 1; i <= A.rows; i++ {
		col[i-1] = A.Row(i)[n-1]
	}
	return col
}

// Row returns a slice that represents a row from the matrix.

// Row 返回表示矩阵中的行的切片
func (A *Matrix) Row(n int) []*big.Int {
	return A.data[findIndex(n, 1, A):findIndex(n, A.columns+1, A)]
}

// Multiply multiplies two matrices together and return the resulting matrix.
// For each element of the result matrix, we get the dot product of the
// corresponding row from matrix A and column from matrix B.

// Multiply 将两个矩阵相乘, 并返回结果矩阵。对于结果矩阵的每个元素, 从矩阵 B 的矩阵 A 和列中得到对应行的点乘积。
func Multiply(A, B Matrix) *Matrix {
	C := Zeros(A.rows, B.columns)
	for r := 1; r <= C.rows; r++ {
		A_row := A.Row(r)
		for c := 1; c <= C.columns; c++ {
			B_col := B.Column(c)
			C.Set(r, c, dotProduct(A_row, B_col))
		}
	}
	return &C
}

// Identity creates an identity matrix with n rows and n columns.  When you
// multiply any matrix by its corresponding identity matrix, you get the
// original matrix.  The identity matrix looks like a zero-filled matrix with
// a diagonal line of one's starting at the upper left.

// Identity 创建具有 n 行和 n 列的标识矩阵。 当您将任何矩阵乘以其相应的标识矩阵时, 就会得到原始矩阵。 标识矩阵看起来像一个零填充矩阵, 其对角线从左上角开始。
func Identity(n int) Matrix {
	A := Zeros(n, n)
	for i := 0; i < len(A.data); i++ {
		A.data[i] = big.NewInt(1)
	}
	return A
}

// Zeros creates an r x c sized matrix that's filled with zeros.  The initial
// state of an int is 0, so we don't have to do any initialization.

//Zeros 创建一个 r x c 大小矩阵填充零。 int 的初始状态为 0, 因此我们不必进行任何初始化。
func Zeros(r, c int) Matrix {
	return Matrix{r, c, make([]*big.Int, r*c)}
}

// New creates an r x c sized matrix that is filled with the provided data.
// The matrix data is represented as one long slice.

// New 创建一个以所提供的数据填充的 r x c 大小矩阵。
// 矩阵数据表示为一个长切片

func New(r, c int, data []*big.Int) Matrix {
	if len(data) != r*c {

		panic("[]*big.Int data provided to matrix.New is great than the provided capacity of the matrix!")
	}
	A := Zeros(r, c)
	A.data = data
	return A
}

// findIndex takes a row and column and returns the corresponding index
// from the underlying data slice.

// findIndex 获取行和列, 并从基础数据切片返回相应的索引。
func findIndex(r, c int, A *Matrix) int {
	return (r-1)*A.columns + (c - 1)
}

// dotProduct calculates the algebraic dot product of two slices.  This is just
// the sum  of the products of corresponding elements in the slices.  We use
// this when we multiply matrices together.

// dotProduct 计算两个切片的代数点乘积。 这只是切片中对应元素的产品的总和。 我们用
//当我们将矩阵相乘时。
func dotProduct(a, b []*big.Int) *big.Int {
	total := new(big.Int)
	x := new(big.Int)
	z := new(big.Int)

	alen := len(a)
	for i := 0; i < alen; i++ {

		y := x.Mul(a[i], b[i])
		total = z.Add(total, y)

		//total = total.Add(total, total.Mul(a[i], b[i]))
	}

	return total
}
```