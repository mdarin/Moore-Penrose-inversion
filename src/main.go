/*
 * Алгоритм I.10
 * Разложение по сингулярным числам и проблема наименьших квадратов
 * Преобразование к двухдиагональной форме методом Хаусхолдера
 * Разложение двухдиагоняльной матрицы по сингулярным числам с помощью QR-алгоритма
 * Проверка сходимости алгоритма по условию заданной точности
 * Круг задач, для которых может быть использовано разложение по сингулярным числам, весьма шпрок.
 *
 *
 * Псевдообращение матриц(процедура svd)
 *
 * Пусть A - действительная матрица размера m x n. 
 * Матрица X размера n x m представляет собой псевдообратную к A,
 * если она удовлетворяет четырём следующим условиям:
 *  1. AXA = A;
 *  2. XAX = X;
 *  3. (AX)T = AX;
 *  4. (XA)T = XA.
 *
 * Единственное решение, удовлетворяющее этим четырём условиям, обозначают A# или A+.
 *
 * Если матрицу A представить в виде A = UΣVT, то(здесь T транспонирование)
 *    A# = VΣUT,
 * где Σ# = diag(σ#[i])
 *      и 
 *            / 1/σ[i] для σ[i] > 0;
 *   Σ#[i] = < 
 *            \ 0      для σ[i] = 0.
 *
 * Таким образом, псевдообращение можно достаточно просто получить с помощью процедуры svd.
 *
 * Уилкинсон Райнш. "Справочник алгоритмов на языке АЛГОЛ. Линейная алгебра"
 *
 */

package main

import(
	// common purpose
	"fmt"
	_ "os"
	_ "log"
	_ "time"

	// mally platfom
	ma "github.com/mdarin/mally"

	// utils
	_ "sort"
	_ "sync"
	_ "math"
	_ "errors" // errors.New()
)

//
// mian driver
//
func main() {

	fmt.Println()
	fmt.Println()
	fmt.Println(" The Moore–Penrose inverse or pseudoinverse procedure")
	fmt.Println()
	fmt.Println(" A# = V*W*UT,")
	fmt.Println(" где W# = diag(w#[i])")
	fmt.Println("    и")
	fmt.Println("          / 1/w[i] для w[i] > 0;")
	fmt.Println(" w#[i] = <")
	fmt.Println("          \\ 0      для w[i] = 0.")
	fmt.Println()
	fmt.Println()


	/* Model data
	float a[2][2] = { // input matrix A
		{5.0,1.0},
		{2.0,3.0},	
	};
	float a_hash[2][2] = { // pseudoinverse A+ or A#
		{0.230769, −0.076923},
		{−0.153846,0.38461},
	}
	*/


	// input 
	A,_ := ma.New(2,2)
	Apinv,_ := ma.New(2,2)

	// initialize source matrix A and dest pseudoinverse matrix A+ or A#
	ma.AssignArray2D(A, A.Rows(), A.Cols(), &([][]float64{
		{5.0,1.0},
		{2.0,3.0},
	}))

		ma.Mprintf("A", "%f ", A)

	//
	// the Moore–Penrose inverse or pseudoinverse procedure
	//
	// A# = V*W*UT,
	// где W# = diag(w#[i])
	//   и
	//          / 1/w[i] для w[i] > 0;
	// w#[i] = < 
	//          \ 0      для w[i] = 0.
	//

	// prepare internal aux matrices
	W,_ := ma.New(2,1)
	U,_ := ma.New(2,2)
	V,_ := ma.New(2,2)	
	Temp,_ := ma.New(2,2)
	UT,_ := ma.New(2,2)
	I,_ := ma.New(2,2)
	Whash,_ := ma.New(2,2)

	ma.Assign(U,A)
	
	// W are the vector of singular values of a
	// U,V are left and right orthogonal transformation matrices
	ma.SVD(U,W,V)

	// get U transposed
	ma.Transpose(UT,U)
	//ma.Mprintf("U", "%f ", U)
	//ma.Mprintf("UT", "%f ", UT)

	// create identity matrix for W diagonalization
	ma.AssignIdentity(I)
	//ma.Mprintf("I","%f ",I)

	//inverse singular values and diagonalilze W to Whash
	for i := 0; i < Whash.Rows(); i++ {
		value_w,_ := W.Getij(i,0)
		Whash.Setij(i,i,1.0/value_w)
	}
	//ma.Mprintf("W#","%f ",Whash)

	// make the Moore–Penrose inversion
	ma.Mult(Temp,V,Whash)
	ma.Mult(Apinv,Temp,UT)

	// show the rusult
	ma.Mprintf("A+","%f ",Apinv)
}

