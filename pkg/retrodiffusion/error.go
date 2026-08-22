package retrodiffusion

import (
	"fmt"
	"strings"
)

func (e APIError) Error() string {
	b := &strings.Builder{}

	for i, d := range e.Detail {
		if i > 0 {
			b.WriteString("; ")
		}

		for j, loc := range d.Loc {
			if j > 0 {
				b.WriteByte('.')
			}

			b.WriteString(loc)
		}
		if len(d.Loc) > 0 {
			b.WriteString(": ")
		}

		b.WriteString(d.Msg)

		if d.Type != "" {
			fmt.Fprintf(b, " (%s)", d.Msg, d.Type)
		}
	}

	return b.String()
}

// func (e Error) Error() string {
// 	return fmt.Sprintf("%s: %s", e.Detail.Code, e.Detail.Message)
// }
