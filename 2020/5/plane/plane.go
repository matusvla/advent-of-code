package plane

type plane map[int]struct{}

func NewPlane(firstSeat, lastSeat int) plane {
	p := make(map[int]struct{})
	for i := firstSeat; i <= lastSeat; i++ {
		p[i] = struct{}{}
	}
	return p
}

func (p plane) MarkOccupied(i int) {
	delete(p, i)
}

func (p plane) EmptySeats() (result []int) {
	for key := range p {
		result = append(result, key)
	}
	return
}
