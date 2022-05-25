package weather

type Conditions struct {
	Summary string
}

func ParseResponse(data []byte) (*Conditions, error) {
	return &Conditions{Summary: ""}, nil
}
