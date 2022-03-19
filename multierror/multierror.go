package multierror

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/tigorlazuardi/epub-scraper/logger"
	"github.com/tigorlazuardi/epub-scraper/unsafeutils"
)

type MultiError []error

func (multierror MultiError) Display() json.RawMessage {
	w := make(map[int]interface{}, len(multierror))
	for i, err := range multierror {
		index := i + 1
		if e, ok := err.(logger.Display); ok {
			w[index] = e.Display()
			continue
		}
		if b, err := json.Marshal(err); err != nil || unsafeutils.GetString(b) == "{}" {
			w[index] = err.Error()
		} else {
			w[index] = json.RawMessage(b)
		}
	}
	e, _ := json.Marshal(w)
	return e
}

func (multierror MultiError) Error() string {
	s := strings.Builder{}
	for i, err := range multierror {
		index := i + 1
		s.WriteString(strconv.Itoa(index))
		s.WriteString(". ")
		s.WriteString(err.Error())

		if index != len(multierror) {
			s.WriteString("| ")
		}
	}
	return s.String()
}

func NewMultiError(capacity int) MultiError {
	return make(MultiError, 0, capacity)
}
