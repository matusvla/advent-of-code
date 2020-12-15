package crt

import "math/big"

func nextIt(a1, a2, q2 *big.Int) {
	helper := big.NewInt(0).Set(a2)

	//a2 = a1 - q2*a2
	a2.Sub(a1, a2.Mul(a2, q2))

	a1.Set(helper)
}

// ExtendedEucleides returns the gcd and integer coeficients m, n such that ma + nb = 1 (Bezout's theorem)
func ExtendedEucleides(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	r_i := a
	r_ip1 := b
	s_i := big.NewInt(1)
	s_ip1 := big.NewInt(0)
	t_i := big.NewInt(0)
	t_ip1 := big.NewInt(1)
	for r_ip1.Cmp(big.NewInt(0)) != 0 {
		q_ip1 := big.NewInt(0).Div(r_i, r_ip1)
		nextIt(r_i, r_ip1, q_ip1)
		nextIt(s_i, s_ip1, s_ip1)
		nextIt(t_i, t_ip1, t_ip1)
		//r_ip1, r_i = r_i-q_ip1*r_ip1, r_ip1
		//s_ip1, s_i = s_i-q_ip1*s_ip1, s_ip1
		//t_ip1, t_i = t_i-q_ip1*t_ip1, t_ip1
	}
	return r_i, s_i, t_i
}
