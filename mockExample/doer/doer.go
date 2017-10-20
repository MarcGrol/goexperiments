package doer

//go:generate mockgen -destination=../mock_doer/mock_doer.go -package=mock_doer github.com/Duxxie/mockExample/doer Doer

type Doer interface {
	DoSomething(int, string) error
}
