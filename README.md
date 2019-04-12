# Moore-Penrose-inversion
The Moore–Penrose inverse or pseudoinverse procedure

```
SVD of a Square Matrix

If the matrix A is square, N × N say, then U, V, and W are all square matrices
of the same size. Their inverses are also trivial to compute: U and V are orthogonal,
so their inverses are equal to their transposes; W is diagonal, so its inverse is the
diagonal matrix whose elements are the reciprocals of the elements w j. 
It now follows immediately that the inverse of A is

A−1 = V · [diag (1/w j )] · UT

The only thing that can go wrong with this construction is for one of the w j ’s
to be zero, or (numerically) for it to be so small that its value is dominated by
roundoff error and therefore unknowable. If more than one of the w j ’s have this
problem, then the matrix is even more singular. So, first of all, SVD gives you a
clear diagnosis of the situation.

[Numerical Recipes in C]
```
