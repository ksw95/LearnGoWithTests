package dictionary

const (
	ErrKeyNotFound = DictionaryErr("key does not exist")
	ErrKeyInUse    = DictionaryErr("key already in use")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
	case nil:
		return ErrKeyInUse
	default:
		return err
	}

	return nil
}
