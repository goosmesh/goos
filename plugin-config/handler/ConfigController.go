package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/goosmesh/goos/core/utils"
	"github.com/goosmesh/goos/plugin-config/entity"
	"github.com/goosmesh/goos/plugin-config/entity/vo"
	"github.com/goosmesh/goos/plugin-config/longpolling"
	utils2 "github.com/goosmesh/goos/plugin-config/longpolling/utils"
	"github.com/goosmesh/goos/plugin-config/service"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 配置文件处理

// 发布配置文件
func Config(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	//fmt.Println(r.Header.Get("Content-Type"))
	if err != nil {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	}
	//else {
	//	fmt.Println(bytes.NewBuffer(body).String())
	//}

	var config entity.Config
	if err := json.Unmarshal([]byte(bytes.NewBuffer(body).String()), &config); err != nil {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	} else {
		if eff, err := service.Config(config); err == nil {
			resp := utils.Succeed(eff)
			if err := json.NewEncoder(w).Encode(resp); err != nil{
				panic(err)
			}
		} else {
			resp := utils.Failed(err.Error())
			if err := json.NewEncoder(w).Encode(resp); err != nil{
				panic(err)
			}
		}
	}



	//fId := r.PostForm["id"]
	//dataId, err := utils.GetParameter("dataId", w, r)
	//if err != nil {
	//	return
	//}
	//groupId, err := utils.GetParameter("groupId", w, r)
	//if err != nil {
	//	return
	//}
	//content, err := utils.GetParameter("content", w, r)
	//if err != nil {
	//	return
	//}
	//fFType, err := utils.GetParameter("type", w, r)
	//if err != nil {
	//	return
	//}
	//var id int64 = 0
	//if len(fId) > 0 {
	//	id, _ = strconv.ParseInt(fId[0], 10, 64)
	//}


}

// 获取配置文件
func GetConfigList(w http.ResponseWriter, r *http.Request)  {


	currentPage, err := utils.GetInt64Parameter("currentPage", true, 1, w, r)
	if err != nil {
		return
	}
	pageSize, err := utils.GetInt64Parameter("pageSize", true, 20, w, r)
	if err != nil {
		return
	}

	queryDataId, err := utils.GetParameter("queryDataId", true, "", w, r)
	if err != nil {
		return
	}

	queryGroup, err := utils.GetParameter("queryGroup", true, "", w, r)
	if err != nil {
		return
	}

	namespace, err := utils.GetInt64Parameter("namespace", false, 0, w, r)
	if err != nil {
		return
	}


	if list, total, err := service.GetConfigList(currentPage, pageSize, namespace, queryDataId, queryGroup); err == nil {
		resp := utils.Succeed(nil).SetList(list).SetPagination(&utils.Pagination{
			Current: currentPage,
			PageSize: pageSize,
			Total: total,
		})
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	} else {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	}
}

// 获取配置文件的配置信息（详细数据）
func GetConfig(w http.ResponseWriter, r *http.Request)  {

	id, err := utils.GetInt64Parameter("id", false, 0, w, r)
	if err != nil {
		return
	}

	if result, err := service.GetConfig(id); err == nil {
		resp := utils.Succeed(nil).PutAll(result)
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	} else {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	}
}

// 删除配置文件
func DeleteConfig(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetInt64Parameter("id", false, 0, w, r)
	if err != nil {
		return
	}

	if eff, err := service.DeleteConfig(id); err == nil {
		resp := utils.Succeed(eff)
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	} else {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	}
}



///////////////////////////////////   sidecar api			/////////////////////////////
///////////////////////////////////   RSA CLIENT API      ///////////////////////////////

//func longOperation(ctx context.Context, ch chan<- string) {
//	// Simulate long operation.
//	// Change it to more than 10 seconds to get server timeout.
//	select {
//	case <-time.After(time.Second * 2):
//		ch <- "Successful result1."
//		ch <- "Successful result2."
//	case <-ctx.Done():
//		close(ch)
//	}
//}
// 基于md5的 long polling 配置变化监听
func ConfigLongPollListener(w http.ResponseWriter, r *http.Request)  {

	// get header long pulling metas

	body, err := ioutil.ReadAll(r.Body)
	log.Info(bytes.NewBuffer(body).String())

	//fmt.Println(r.Header.Get("Content-Type"))
	if err != nil {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	}
	//else {
	//	fmt.Println(bytes.NewBuffer(body).String())
	//}

	var configMd5Data vo.ConfigMd5Data

	err = json.Unmarshal([]byte(bytes.NewBuffer(body).String()), &configMd5Data)
	if err != nil {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
		return
	}


	timeout, err := utils.GetInt64Header("Long-Pulling-Timeout", false, 30000, nil, r)
	if err != nil {
		_, _ = fmt.Fprint(w, "")
		return
	}

	noHangup, err := utils.GetHeader("Long-Pulling-Timeout-No-Hangup", true, "true", nil, r)
	if err != nil {
		_, _ = fmt.Fprint(w, "")
		return
	}
	nh, err := strconv.ParseBool(noHangup)
	if err != nil {
		nh = true
	}

	//probeModify, err := utils.GetParameter("Listening-Configs", false, "", nil, r)
	//if err != nil {
	//	_, _ = fmt.Fprint(w, "")
	//	return
	//}
	if configMd5Data.ListeningConfigs == "" {
		_, _ = fmt.Fprint(w, "")
		return
	}
	probeModifyDecode, err := url.QueryUnescape(configMd5Data.ListeningConfigs)
	if err != nil {
		_, _ = fmt.Fprint(w, "")
		return
	}

	clientMd5Map := utils2.ConfigMD5ToMap(probeModifyDecode)
	if len(clientMd5Map) == 0 {
		_, _ = fmt.Fprint(w, "")
		return
	}
	fmt.Println(clientMd5Map)


	// first test go long polling
	notifier, ok := w.(http.CloseNotifier)
	if !ok {
		panic("Expected http.ResponseWriter to be an http.CloseNotifier")
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)
	exe, clientClose := longpolling.AddLongPolling(ctx, ch, clientMd5Map, nh)
	go exe()

	select {
	case result := <-ch:
		_, _ = fmt.Fprint(w, result)
		cancel()
		return
	case <-time.After(time.Millisecond * time.Duration(timeout)):
		_, _ = fmt.Fprint(w, "")
	case <-notifier.CloseNotify():
		clientClose()
		fmt.Println("Client has disconnected.")
	}
	cancel()
	<-ch

}

// 传入配置文件MD5，比较文件MD5，不一致则返回RSA加密数据
//func RsaGetConfig(w http.ResponseWriter, r *http.Request) {
//	body, err := ioutil.ReadAll(r.Body)
//	//fmt.Println(r.Header.Get("Content-Type"))
//	if err != nil {
//		resp := utils.Failed(err.Error())
//		if err := json.NewEncoder(w).Encode(resp); err != nil{
//			panic(err)
//		}
//	}
//	//else {
//	//	fmt.Println(bytes.NewBuffer(body).String())
//	//}
//
//	var configQuery vo.ConfigQuery
//	if err := json.Unmarshal([]byte(bytes.NewBuffer(body).String()), &configQuery); err != nil {
//		resp := utils.Failed(err.Error())
//		if err := json.NewEncoder(w).Encode(resp); err != nil{
//			panic(err)
//		}
//	} else {
//		if len(configQuery.ConfigList) == 0 {
//			resp := utils.Succeed(nil)
//			if err := json.NewEncoder(w).Encode(resp); err != nil{
//				panic(err)
//			}
//			return
//		}
//		if result, err := service.RsaGetConfig(configQuery); err == nil {
//			resp := utils.Succeed(result)
//			if err := json.NewEncoder(w).Encode(resp); err != nil{
//				panic(err)
//			}
//		} else {
//			resp := utils.Failed(err.Error())
//			if err := json.NewEncoder(w).Encode(resp); err != nil{
//				panic(err)
//			}
//		}
//	}
//}

// 客户端获取配置文件API（core api）
// 客户端根据 dataId, groupId, namespaceId 获取配置信息
func GetConfigClient(w http.ResponseWriter, r *http.Request)  {
	dataId, err := utils.GetParameter("dataId", false, "", w, r)
	if err != nil {
		return
	}
	groupId, err := utils.GetParameter("groupId", false, "", w, r)
	if err != nil {
		return
	}
	namespaceId, err := utils.GetParameter("namespaceId", true, "", w, r)
	if err != nil {
		return
	}

	if result, err := service.GetConfigByQuery(dataId, groupId, namespaceId); err == nil {
		//resp := utils.Succeed(nil).PutAll(result)
		_, _ = w.Write([] byte(result.(entity.Config).Content))
		//if err := json.NewEncoder(w).Encode(resp); err != nil{
		//	panic(err)
		//}
	} else {
		resp := utils.Failed(err.Error())
		if err := json.NewEncoder(w).Encode(resp); err != nil{
			panic(err)
		}
	}
}