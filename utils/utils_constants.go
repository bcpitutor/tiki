package utils

// The key will be generating in runtime.
// sample key format (32bit) "01234567012345670123456701234567"
type EncData struct {
	Key string
}

var Enckey EncData

func (encData *EncData) SetEncKey(key string) {
	encData.Key = key
}

func (encData *EncData) GetEncKey() string {
	return encData.Key
}
