package shape

type Rectangle struct {
	Panjang float32
	Lebar   float32
}

func (r *Rectangle) Area() float32 {
	luas := r.Lebar * r.Panjang
	return luas
}
