package user

type Serializer interface {
	Decode(input []byte) (*User, error)
	Encode(input *User) ([]byte, error)
}
