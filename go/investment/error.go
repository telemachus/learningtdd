package investment

import (
	"fmt"
	"strings"
)

// ConversionErrors stores a list of failed attempts to convert unknown
// currencies.
type ConversionErrors []string

func (ce ConversionErrors) Error() string {
	s := ""
	if len(ce) > 1 {
		s = "s"
	}

	return fmt.Sprintf(
		"investment: currency conversion failure%s: [%s]",
		s,
		strings.Join(ce, ", "),
	)
}
