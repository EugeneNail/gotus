package environment

type Environment int

const (
	Development Environment = iota
	Production
)

var stringValues = map[Environment]string{
	Development: "development",
	Production:  "production",
}

func (environment Environment) ToString() string {
	return stringValues[environment]
}
