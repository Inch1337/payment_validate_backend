package payments

type Payment interface {
	Pay(amount int) error
}
