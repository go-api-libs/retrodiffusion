package retrodiffusion

func (e Error) Error() string {
	return *e.Detail
}
