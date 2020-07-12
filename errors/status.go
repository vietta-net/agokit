package errors

type Status interface {
	To(status interface{}) (err error)
}
