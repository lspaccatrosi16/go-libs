package enum

type Enum interface {
	String() string
	FromI(i int) Enum
	IsValid() bool
}

const InvalidString = "invalid"

func All[T Enum]() []string {
	options := []string{}

	t := *new(T)
	i := 0

	for {
		thisOpt := t.FromI(i)
		if thisOpt.String() != InvalidString {
			options = append(options, thisOpt.String())
			i++
		} else {
			break
		}
	}

	return options
}
