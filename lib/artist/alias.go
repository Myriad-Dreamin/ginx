package artist

func Q(desc ...interface{}) RequestObj {
	return Request(desc...)
}

func T(name string, base interface{}) *transferClass {
	return Transfer(name, base)
}

func QT(name string, base interface{}) RequestObj {
	return Q(T(name, base))
}
