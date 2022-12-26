package body

type RawBody struct {
	data    string
	rawType RawType
}

func (r *RawBody) Data() []byte {
	return []byte(r.data)
}
