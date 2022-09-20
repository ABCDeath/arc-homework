package vector

type Vector struct {
	x, y int
}

func (v Vector) Add(vector Vector) Vector {
	return New(v.x+vector.x, v.y+vector.y)
}

func New(x, y int) Vector {
	return Vector{x: x, y: y}
}
