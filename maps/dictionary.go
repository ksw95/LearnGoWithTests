package dictionary

const (
	ErrKeyNotFound   = DictionaryErr("value cannot be searched as key is not found")
	ErrKeyInUse      = DictionaryErr("cannot add new key-value, as key is already in use")
	ErrNoKeyToUpdate = DictionaryErr("value cannot be updated as key not found")
	ErrNoKeyToDelete = DictionaryErr("key cannot be updated as key is not found")
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

func (d Dictionary) Update(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		return ErrNoKeyToUpdate
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}

}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
