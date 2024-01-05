package product

type Product interface {
	Price() float64
}

type small struct {
	price float64
}

func (s small) Price() float64 {
	return s.price
}

type medium struct {
	price float64
}

func (m medium) Price() float64 {
	return m.price + ((m.price * 3) / 100)
}

type large struct {
	price float64
}

func (l large) Price() float64 {
	return l.price + ((l.price * 6) / 100) + 2500
}

func Factory(typeProduct string, price float64) (product Product, err string) {
	switch typeProduct {
	case "SMALL":
		product = small{price: price}
	case "MEDIUM":
		product = medium{price: price}
	case "LARGE":
		product = large{price: price}
	default:
		err = "This type product not exist"
	}
	return

}
