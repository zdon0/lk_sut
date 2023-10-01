package decoder

import "golang.org/x/text/encoding/charmap"

var decoder = charmap.Windows1251.NewDecoder()

func Decode(s string) (string, error) {
	return decoder.String(s)
}
