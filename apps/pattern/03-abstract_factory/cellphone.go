package abstract_factory

type CellPhone interface {
	Call(number string) string
	Send(msg string) string
	CallWithVideo(msg string) string
}
