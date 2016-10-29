package table

type row struct {
	firstKeyCode int
	secondKeyCode int
	flyTime int
	firstDwellTime int
	secondDwellTime int
}

func (r *row) FirstKeyCode() int {
	return r.firstKeyCode
}

func (r *row) SecondKeyCode() int {
	return r.secondKeyCode
}

func (r *row) FlyTime() int {
	return r.flyTime
}

func (r *row) FirstDwellTime() int {
	return r.firstDwellTime
}

func (r *row) SecondDwellTime() int {
	return r.secondDwellTime
}
