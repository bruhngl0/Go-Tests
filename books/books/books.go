package books

import (
	"fmt"
)

func Info(title, edition string, copies int) string {
	return fmt.Sprintf("Title: %s, Copies: $%d Edition: %s", title, copies, edition)
}
