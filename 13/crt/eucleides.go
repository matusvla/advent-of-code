package crt

// ExtendedEucleides returns the gcd and integer coeficients m, n such that ma + nb = 1 (Bezout's theorem)
func ExtendedEucleides(a, b int64) (int64, int64, int64) {
	r_i := a
	r_ip1 := b
	s_i := int64(1)
	s_ip1 := int64(0)
	t_i := int64(0)
	t_ip1 := int64(1)
	for r_ip1 != 0 {
		q_ip1 := r_i / r_ip1
		r_ip1, r_i = r_i-q_ip1*r_ip1, r_ip1
		s_ip1, s_i = s_i-q_ip1*s_ip1, s_ip1
		t_ip1, t_i = t_i-q_ip1*t_ip1, t_ip1
	}
	return r_i, s_i, t_i
}
