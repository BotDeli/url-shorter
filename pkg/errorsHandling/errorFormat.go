package errorsHandling

import "fmt"

func ErrFormat(location, function string, err error) error {
	return fmt.Errorf("%s.%s: %s", location, function, err)
}
