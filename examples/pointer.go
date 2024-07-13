package examples

type checkStr struct {
	name string
}

func change(v *checkStr) {
	v.name = "changed"
}

func Ccc() {
	v := checkStr{name: "original"}
	change(&v)
	println(v.name)
}
