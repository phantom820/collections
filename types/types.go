package types

type Key interface {
	comparable
	Ordered()
}
