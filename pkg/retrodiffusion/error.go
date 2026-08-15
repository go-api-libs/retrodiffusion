package retrodiffusion

import "fmt"

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Detail.Code, e.Detail.Message)
}
