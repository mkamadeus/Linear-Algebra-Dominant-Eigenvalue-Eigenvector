package main

import (
	"fmt"
	"errors"
)

func multiply(x, y [][]float32) ([][]float32, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication.")
	}

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return out, nil
}

func main() {
	iter := 1000

	A := [][]float32{{0,1,0},{0,0,1},{4,-17,8}}
	x := [][]float32{{1},{1},{1}}

	n:=3
	val:=float32(0.0)
	for i:=0; i<n;i++ {	
		fmt.Println(A[i])
	}

	for i:=0; i < iter; i++ {
		x1, _ := multiply(A,x)
		
		mx := x1[0][0]
		for j:=1; j < n; j++ {
			if(x1[j][0]>mx){
				mx=x1[j][0]
			}
		}
		
		for j:=0; j < n ; j++ {
			x1[j][0] /= mx
		}
		fmt.Println("Iteration",i+1, x1)

		val = mx
		x = x1
	}

	fmt.Println("Dominant eigenvalue =", val)
	
	fmt.Print("Dominant eigenvector =", "<")
	for i:=0;i<n;i++ {
		if(i!=n-1){
			fmt.Print(x[i][0],",")
		}else{
			fmt.Print(x[i][0])
		}
	}
	fmt.Println(">")
}