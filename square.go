package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}


func (in Square) End() Point {
	
	return Point {x:in.start.x+int(in.a),y:in.start.y+int(in.a)}
}

func (in Square) Area() uint {
	return in.a*in.a
}

func (in Square) Perimeter() uint {
	return in.a*4
}