package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func getParameter(name string, r *http.Request) []string  {
	v := r.PostForm[name]

	if len(v) == 0 {
		v = r.URL.Query()[name]
	}
	return v
}

func getHeader(name string, r *http.Request) []string {
	v := r.Header[name]
	return v
}

func GetParameter(name string, nullAble bool, defaultValue string, w http.ResponseWriter, r *http.Request) (value string, err error) {
	v := getParameter(name, r)

	if len(v) == 0 {
		if !nullAble {
			if w != nil {
				resp := Failed("goos required parameter " + name)
				if err := json.NewEncoder(w).Encode(resp); err != nil{
					return "", err
				} else {
					return "", errors.New("goos required parameter " + name)
				}
			} else {
				return "", errors.New("goos required header " + name)
			}
		} else {
			return defaultValue, nil
		}
	} else {
		return v[0], nil
	}
}

func GetHeader(name string, nullAble bool, defaultValue string, w http.ResponseWriter, r *http.Request) (value string, err error) {
	v := getHeader(name, r)

	if len(v) == 0 {
		if !nullAble {
			if w != nil {
				resp := Failed("goos required header " + name)
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					return "", err
				} else {
					return "", errors.New("goos required header " + name)
				}
			} else {
				return "", errors.New("goos required header " + name)
			}
		} else {
			return defaultValue, nil
		}
	} else {
		return v[0], nil
	}
}

func GetInt64Parameter(name string, nullAble bool, defaultValue int64, w http.ResponseWriter, r *http.Request) (value int64, err error) {
	if v, e := GetParameter(name, nullAble, strconv.FormatInt(defaultValue, 10), w, r); e != nil {
		return 0, e
	} else {
		if v, e := strconv.ParseInt(v, 10, 64); e == nil {
			return v, e
		} else {
			if !nullAble {
				return v, e
			} else {
				return defaultValue, nil
			}
		}

	}
}

func GetInt64Header(name string, nullAble bool, defaultValue int64, w http.ResponseWriter, r *http.Request) (value int64, err error) {
	if v, e := GetHeader(name, nullAble, strconv.FormatInt(defaultValue, 10), w, r); e != nil {
		return 0, e
	} else {
		if v, e := strconv.ParseInt(v, 10, 64); e == nil {
			return v, e
		} else {
			if !nullAble {
				return v, e
			} else {
				return defaultValue, nil
			}
		}

	}
}

func GetIntHeader(name string, nullAble bool, defaultValue int, w http.ResponseWriter, r *http.Request) (value int, err error) {
	if v, e := GetHeader(name, nullAble, strconv.Itoa(defaultValue), w, r); e != nil {
		return 0, e
	} else {
		if v, e := strconv.Atoi(v); e == nil {
			return v, e
		} else {
			if !nullAble {
				return v, e
			} else {
				return defaultValue, nil
			}
		}

	}
}
