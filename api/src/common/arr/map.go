package arr

func Map[Tin any, Tout any](arr []Tin, fn func(Tin) Tout) []Tout {
	var out []Tout = make([]Tout, len(arr))
	for _, e := range arr {
		out = append(out, fn(e))
	}
	return out
}
