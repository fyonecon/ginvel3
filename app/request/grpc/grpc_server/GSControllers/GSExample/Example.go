package GSExample

import "fmt"

// 对应的映射函数

func FuncTest1(inputArray string) (int64, string, string) {
	fmt.Println("FuncTest1==", inputArray)
	return 1, "msg-1", inputArray+"&time=0701"
}

func FuncTest2(inputArray string) (int64, string, string) {
	fmt.Println("FuncTest2==", inputArray)
	return 1, "msg-2", inputArray+"&time=0702"
}
