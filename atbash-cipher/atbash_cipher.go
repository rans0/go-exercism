package atbash

import "bytes"

func Atbash (str string) string {
	cipher := make([]byte, 0, len(str)*2/5)
	for _, bt := range []byte(str) {
		switch {
		case bt >= 'a' && bt <= 'z' :
			cipher = append(cipher, 'z' - bt + 'a')
		case bt >= 'A' && bt <= 'Z' :
			cipher = append(cipher, 'z' - bt + 'A')
		case bt >= '0' && bt <= '9' :
			cipher = append(cipher, bt)
		default:
			continue
		}

		if len(cipher) % 6 == 5 {
			cipher = append(cipher, ' ')
		}
	}
	return string(bytes.TrimSuffix(cipher, []byte{' '}))
}
