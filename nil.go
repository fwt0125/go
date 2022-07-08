package main
import "fmt"
// 定义一个结构体
type MyImplement struct{}
// 实现fmt.Stringer的String方法
func (m *MyImplement) String() string {
	return "hi"
}
// 在函数中返回fmt.Stringer接口
func GetStringer() fmt.Stringer {
	// 赋nil
	var s *MyImplement = nil
	// 返回变量
	return s
}
func main() {
	// 判断返回值是否为nil
	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}

	if GetStringerNil() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}
}

func GetStringerNil() fmt.Stringer {
	var s *MyImplement = nil
	if s == nil {
		return nil
	}
	return s
}

//第 9 行，实现 fmt.Stringer 的 String() 方法。
//第 21 行，s 变量此时被 fmt.Stringer 接口包装后，实际类型为 *MyImplement，值为 nil 的接口。
//第 27 行，使用 GetStringer() 的返回值与 nil 判断时，虽然接口里的 value 为 nil，但 type 带有 *MyImplement 信息，使用 == 判断相等时，依然不为 nil。