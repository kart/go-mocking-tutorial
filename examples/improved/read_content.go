package improved

import (
	"errors"
	"github.com/kart/go-mocking-tutorial/base"
)

// ReadContent reads the content from a provided file while satisfying the
// following contract:
//	a) If provided name is `nil` or empty, nil content and "no name"
//		error is returned
//	b) If the stat on the file fails, nil content and error is returned.
//	c) If size of the file is less than 10 bytes, nil content and
//		"too small" error is returned.
//	d) If we fail to open the file, nil content and error is returned.
//	e) If we fail to read the file, nil content and error is returned.
//	f) If read returns less than size of the content, nil content and
//		"partial read" error is returned.
//	g) Finally, if none of the above happens, the content of the provided
//		file is returned with nil error.
func ReadContent(name string) ([]byte, error) {
	if len(name) == 0 {
		return nil, errors.New("no name")
	}
	info, err := base.AppOs.Stat(name)
	if err != nil {
		return nil, err
	}
	if info.Size() < 10 {
		return nil, errors.New("too small")
	}
	file, err := base.AppOs.Open(name)
	if err != nil {
		return nil, err
	}
	content := make([]byte, info.Size())
	n, err := file.Read(content)
	if err != nil {
		return nil, err
	}
	if n < int(info.Size()) {
		return nil, errors.New("partial read")
	}
	return content, nil
}
