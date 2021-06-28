package chain

type department interface {
	execute(*patient)
	setNext(department)
}
