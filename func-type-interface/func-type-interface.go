package main

import "net/http"

type IGetter interface {
	Get(string) ([]byte, error)
}

type GetterFunc func(string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// 参数是一个接口
func GetFromSource(getter IGetter, key string) []byte {
	buf, err := getter.Get(key)
	if err == nil {
		return buf
	}
	return nil
}

func test(key string) ([]byte, error) {
	return []byte(key), nil
}

type DB struct{ url string }

func (d DB) Get(key string) ([]byte, error) {
	return []byte("db"), nil
}

/*标注库*/

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main() {
	// 通过 getterFunc 把一个匿名函数转换为一个实现了IGetter接口类型的函数
	//GetterFunc 类型的函数作为参数
	GetFromSource(GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	}), "test")

	GetFromSource(GetterFunc(test), "key")

	db := &DB{"www.baidu.com"}

	GetFromSource(db, "key")
}
