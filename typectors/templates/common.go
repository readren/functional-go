package fung

func catch(errPtr *interface{}) {
	*errPtr = recover()
}
