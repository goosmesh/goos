package utils

import (
	"errors"
	"github.com/goosmesh/goos/plugin-config/longpolling/constants"
	"strings"
)

func GenConfigKey(dataId string, groupId string, namespaceId string) string {
	return dataId + constants.WORD_SEPARATOR + groupId + constants.WORD_SEPARATOR + namespaceId
}

func ParseConfigKey(key string) (dataId string, groupId string, namespaceId string, err error) {
	items := strings.Split(key, constants.WORD_SEPARATOR)
	if len(items)!= 3 {
		return "", "", "", errors.New("not a key")
	}
	return items[0], items[1], items[2], nil
}