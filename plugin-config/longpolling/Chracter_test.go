package longpolling

import (
	"fmt"
	"github.com/goosmesh/goos/plugin-config/longpolling/constants"
	"net/url"
	"testing"
)

func TestCr(t *testing.T)  {
	s := string(rune(1))
	params := "dataId" + constants.WORD_SEPARATOR + "groupId" + constants.WORD_SEPARATOR + "namespaceId" + constants.WORD_SEPARATOR + "md5" + constants.LINE_SEPARATOR + "dataId2" + constants.WORD_SEPARATOR + "groupId2" + constants.WORD_SEPARATOR + "namespaceId2" + constants.WORD_SEPARATOR + "md52"
	fmt.Println(s)
	fmt.Println(params)
	fmt.Println(url.QueryEscape(params))
	fmt.Println(url.QueryUnescape(url.QueryEscape(params)))
}