package body

type NoneBody struct {
}

func (n *NoneBody) Data() []byte {
	return nil
}
