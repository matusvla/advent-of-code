package crt

import (
	"errors"
	"fmt"
	"math/big"
)

func ChineseRemainderTheorem(a, n []int) (*big.Int, error) {
	if len(a) != len(n) || len(a) == 0 {
		return nil, errors.New("invalidpassed slices")
	}

	var a1, a2, n1, n2, x *big.Int
	var err error
	n1 = big.NewInt(int64(n[0]))
	a1 = big.NewInt(int64(a[0]) % int64(n[0]))
	for i := 1; i < len(a); i++ {
		n2 = big.NewInt(int64(n[i]))
		a2 = big.NewInt(int64(a[i]) % int64(n[i]))
		x, err = ChineseRemainderTheorem2Equations(a1, a2, n1, n2)
		if err != nil {
			return nil, err
		}
		a1.Set(x)
		n1.Mul(n1, n2)
	}
	return x, nil
}

func ChineseRemainderTheorem2Equations(a1, a2, n1, n2 *big.Int) (*big.Int, error) {
	gcd, m1f, m2f := ExtendedEucleides(n1.Int64(), n2.Int64())
	m1 := big.NewInt(m1f)
	m2 := big.NewInt(m2f)
	if gcd > 1 {
		return nil, fmt.Errorf("Numbers %v, %v are not coprime.", n1, n2)
	}
	// x := a1*m2*n2 + a2*m1f*n1
	x := big.NewInt(0)
	helper := big.NewInt(0)
	x.Mul(a1, m2)
	x.Mul(x, n2)
	helper.Mul(a2, m1)
	helper.Mul(helper, n1)
	x.Add(x, helper)
	//for x < 0 {
	//	x += n1 * n2
	//}
	zero := big.NewInt(0)
	mult := big.NewInt(0).Mul(n1, n2)
	corr := big.NewInt(0).Div(x, mult)
	corr.Mul(corr, mult)
	corr.Abs(corr)
	if x.Cmp(zero) == -1 {
		x.Add(x, corr).Add(x, mult)
	} else if x.Cmp(zero) == 1 {
		x.Sub(x, corr)
	}
	return x, nil
}
