package main

// 模拟try catch
func try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	try(func(){
		panic("foo")
	}, func(e interface{}) {
		println("catch")
		print(e)
	})
}