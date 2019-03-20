package utils

import (
	"github.com/goosmesh/goos/core/utils/alg"
	"github.com/goosmesh/goos/plugin-config/entity"
	"github.com/goosmesh/goos/plugin-config/longpolling/constants"
	"github.com/goosmesh/goos/plugin-config/service"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

// 和配置文件 md5 操作相关
// logical util

// key : md5
// 解析传输协议(w为字段分隔符，l为每条数据分隔符， dataId，groupId，namespaceId，md5)：D w G w N w MD5 l
func ConfigMD5ToMap(md5Datas string) map[string] string {
	result := map[string]string{}
	lines := strings.Split(md5Datas, constants.LINE_SEPARATOR)
	for _, line := range lines {
		if line != "" {
			items := strings.Split(line, constants.WORD_SEPARATOR)
			if len(items) == 4 {
				result[GenConfigKey(items[0], items[1], items[3])] = items[4]
			}
		}
	}
	return result
}
// 改变的key 转为response
func ChangedConfigToResponse(data []string) string {
	str := ""
	for _, key := range data {
		str += key + constants.LINE_SEPARATOR
	}
	return url.QueryEscape(str)
}

// 过滤已经改变的配置文件，返回key
func FilterChangedConfig(data map[string] string) []string {
	var result []string

	for key, md5 := range data {
		// get parse
		dataId, groupId, namespaceId, err := ParseConfigKey(key)
		if err == nil {
			m, err := getConfigMD5(dataId, groupId, namespaceId)
			if err == nil && m != md5 {
				result = append(result, key)
			}
		}
	}

	return result
}

func getConfigMD5(dataId string, groupId string, namespaceId string) (md5 string, err error) {
	config, err := service.GetConfigByQuery(dataId, groupId, namespaceId)
	if err != nil {
		return "", err
	}
	if config == nil {
		return "", errors.New("can not find config : " + dataId + " " + groupId + " " + namespaceId)
	}
	data, err := alg.RsaDecrypt(config.(entity.Config).Content)
	if err != nil {
		return "", err
	}
	m := alg.MD5(data)
	return m, nil
}