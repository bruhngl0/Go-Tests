package books

import (
	"fmt"
)

func Info(title, edition string, copies int) string {
	a := 1
	return fmt.Sprintf("Title: %s, Copies: $%d Edition: %s, %d", title, copies, edition, a)
}
