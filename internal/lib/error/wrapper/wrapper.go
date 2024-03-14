package wrapper

import "fmt"

func Wrap(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}
