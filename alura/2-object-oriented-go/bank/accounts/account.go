package accounts

type Account interface {
	Withdraw(value float64) string
}
