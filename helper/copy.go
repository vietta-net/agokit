package helper

type Copy interface {
	To(data interface{}) (err error)
	From(data interface{}) (err error)
}
