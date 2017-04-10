package edm

type Binary []byte

// UnmarshalText calls url.Parse
// func (b *Binary) UnmarshalText(p []byte) error {
// 	src := p
// 	dstLen := base64.StdEncoding.DecodedLen(len(src))
// 	dst := make([]byte, dstLen)
// 	_, err := base64.StdEncoding.Decode(dst, src)
// 	if err == nil {
// 		*b = dst
// 	}
// 	return err
// }

// func (b Binary) MarshalJSON() ([]byte, error) {
// 	src := b
// 	dstLen := base64.StdEncoding.EncodedLen(len(src))
// 	dst := make([]byte, dstLen)
// 	base64.StdEncoding.Encode(dst, src)
// 	return dst, nil
// }
