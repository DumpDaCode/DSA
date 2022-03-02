package algorithms

type Matrix [][]int

// Naive Approach
func (A Matrix) NaiveMultiply(B Matrix) Matrix {
	n := len(A)
	var C Matrix = make(Matrix, len(A))
	for i := range C {
		C[i] = make([]int, len(A[i]))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = 0
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// Divide and conquer approach
func (A Matrix) RecursiveMultiply(B Matrix) Matrix {
	n := len(A)
	var C Matrix = make(Matrix, len(A))
	for i := range C {
		C[i] = make([]int, len(A[i]))
	}
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	} else {
		m := len(A) / 2

		// Just to avoid any issues further
		// Update: Turns out slices are very funny
		var A11 Matrix = make(Matrix, len(A))
		copy(A11, A)
		A11 = A11[0:m]
		for i, v := range A11 {
			A11[i] = v[0:m]
		}

		var B11 Matrix = make(Matrix, len(B))
		copy(B11, B)
		B11 = B11[0:m]
		for i, v := range B11 {
			B11[i] = v[0:m]
		}

		var A12 Matrix = make(Matrix, len(A))
		copy(A12, A)
		A12 = A12[0:m]
		for i, v := range A12 {
			A12[i] = v[m:]
		}

		var B12 Matrix = make(Matrix, len(B))
		copy(B12, B)
		B12 = B12[0:m]
		for i, v := range B12 {
			B12[i] = v[m:]
		}

		var A21 Matrix = make(Matrix, len(A))
		copy(A21, A)
		A21 = A21[m:]
		for i, v := range A21 {
			A21[i] = v[0:m]
		}

		var B21 Matrix = make(Matrix, len(B))
		copy(B21, B)
		B21 = B21[m:]
		for i, v := range B21 {
			B21[i] = v[0:m]
		}

		var A22 Matrix = make(Matrix, len(A))
		copy(A22, A)
		A22 = A22[m:]
		for i, v := range A22 {
			A22[i] = v[m:]
		}

		var B22 Matrix = make(Matrix, len(B))
		copy(B22, B)
		B22 = B22[m:]
		for i, v := range B22 {
			B22[i] = v[m:]
		}

		/*
			Normal Way of doing it
		*/
		C11 := A11.RecursiveMultiply(B11).Add(A12.RecursiveMultiply(B21))
		C12 := A11.RecursiveMultiply(B12).Add(A12.RecursiveMultiply(B22))
		C21 := A21.RecursiveMultiply(B11).Add(A22.RecursiveMultiply(B21))
		C22 := A21.RecursiveMultiply(B12).Add(A22.RecursiveMultiply(B22))

		/*
			Uncomment below lines for Strassen's algorithm
		*/
		// S1 := B12.Sub(B22)
		// S2 := A11.Add(A12)
		// S3 := A21.Add(A22)
		// S4 := B21.Sub(B11)
		// S5 := A11.Add(A22)
		// S6 := B11.Add(B22)
		// S7 := A12.Sub(A22)
		// S8 := B21.Add(B22)
		// S9 := A11.Sub(A21)
		// S10 := B11.Add(B12)

		// P1 := A11.RecursiveMultiply(S1)
		// P2 := S2.RecursiveMultiply(B22)
		// P3 := S3.RecursiveMultiply(B11)
		// P4 := A22.RecursiveMultiply(S4)
		// P5 := S5.RecursiveMultiply(S6)
		// P6 := S7.RecursiveMultiply(S8)
		// P7 := S9.RecursiveMultiply(S10)

		// C11 := P5.Add(P4).Sub(P2).Add(P6)
		// C12 := P1.Add(P2)
		// C21 := P3.Add(P4)
		// C22 := P5.Add(P1).Sub(P3).Sub(P7)

		// Copying Data back to matrix
		for i := range A {
			for j := range A[0] {
				if i < m && j < m {
					C[i][j] = C11[i][j]
				} else if i < m && j >= m {
					C[i][j] = C12[i][j-m]
				} else if i >= m && j < m {
					C[i][j] = C21[i-m][j]
				} else {
					C[i][j] = C22[i-m][j-m]
				}
			}
		}
		return C
	}
}

func (A Matrix) Add(B Matrix) Matrix {
	var C Matrix = make(Matrix, len(A))
	for i := range C {
		C[i] = make([]int, len(A[i]))
	}
	for i := range A {
		for j := range A[i] {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func (A Matrix) Sub(B Matrix) Matrix {
	var C Matrix = make(Matrix, len(A))
	for i := range C {
		C[i] = make([]int, len(A[i]))
	}
	for i := range A {
		for j := range A[i] {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}
