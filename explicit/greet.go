package subpackage

import "fmt"

// Greet greets someone.
func Greet(name string) string {
	return fmt.Sprintf(Format, name)
}
