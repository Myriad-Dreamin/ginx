package serial

type methodDescription struct {
	method   MethodType
	name     string
	requests []ObjectDescription
	replies  []ObjectDescription
}

