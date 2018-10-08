package M09_proxy

type Subject interface {
	Do() string
}

type REalSubject struct {
}

func (REalSubject) Do() string {
	return "real"
}

type Proxy struct {
	real REalSubject
}

func (p Proxy) Do() string {
	var res string
	res += "pre:"

	res += p.real.Do()

	res += ":after"

	return res
}
