package TheatrumOrbis

type Strings struct {
	Lookup []chunk
}

type chunk struct {
	offset uint64
	count uint
	entries []entry
}

type entry struct {
	l uint
	v string
}

func (s *Strings) Add(n string) int {

	return 0
}


func (s *Strings) Get(n int) string {
	return ""
}
