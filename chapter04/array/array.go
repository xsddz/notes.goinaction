package array

import (
	"fmt"
	"reflect"
)

// Run 数组测试入口
func Run() {
	fmt.Printf("================================================================================\n")

	/*
		1. 数组是一个长度固定的数据类型，用于存储一段具有相同的类型的元素的连续块。一旦声明，
		数组里存储的数据类型和数组长度就都不能改变了。如果需要存储更多的元素，就需要先创建一个
		更长的数组，再把原来数组里的值复制到新数组里。
		2. 在 Go 语言中声明变量时，总会使用对应类型的零值来对变量进行初始化。
		3. 数组变量的类型包括数组长度和每个元素的类型。只有这两部分都相同的数组，才是类型相同的
		数组，才能互相赋值。
		4. 复制数组指针，只会复制指针的值，而不会复制指针所指向的值。
		5. 在函数之间传递变量时，总是以值的方式传递的。如果这个变量是一个数组，意味着整个数组，
		不管有多长，都会完整复制，并传递给函数。
	*/

	// 声明一个包含 5 个元素的整型数组，并设置为零值
	var arr01 [5]int
	fmt.Printf("arr01:\t%T\t%v\n", arr01, arr01)

	// 声明一个包含 5 个元素的整型数组，用具体值初始化每个元素
	arr02 := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("arr02:\t%T\t%v\n", arr02, arr02)

	// 声明一个整型数组，用具体值初始化每个元素，容量由初始化值的数量决定
	arr03 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("arr03:\t%T\t%v\n", arr03, arr03)

	// 声明一个有 5 个元素的数组，用具体值初始化索引为 1 和 2 的元素，其余元素保持零值
	arr04 := [5]int{1: 10, 2: 20}
	fmt.Printf("arr04:\t%T\t%v\n", arr04, arr04)

	// 声明包含 5 个元素的指向整数的数组，用整型指针初始化索引为 0 和 2 的数组元素
	arr05 := [5]*int{0: new(int), 2: new(int)}
	fmt.Printf("arr05:\t%T\t%v\t%v\n", arr05, arr05, makePointerArrVal(arr05))

	// 把 arr02 的值复制到 arr01，复制之后，两个数组的值完全一样
	arr01 = arr02
	fmt.Printf("\ncopy arr02 to arr01:\n")
	fmt.Printf("arr01:\t%T\t%v\n", arr01, arr01)
	fmt.Printf("arr02:\t%T\t%v\n", arr02, arr02)

	// 修改索引为 2 的元素的值
	arr01[2] = 14101 // 修改 arr01 的值不会影响 arr02
	arr02[2] = 14102
	arr03[2] = 14103
	arr04[2] = 14104
	*arr05[2] = 14105
	fmt.Printf("\nmodify index 2 value of arr[0-5]:\n")
	fmt.Printf("arr01:\t%T\t%v\n", arr01, arr01)
	fmt.Printf("arr02:\t%T\t%v\n", arr02, arr02)
	fmt.Printf("arr03:\t%T\t%v\n", arr03, arr03)
	fmt.Printf("arr04:\t%T\t%v\n", arr04, arr04)
	fmt.Printf("arr05:\t%T\t%v\t%v\n", arr05, arr05, makePointerArrVal(arr05))

	fmt.Printf("================================================================================\n")

	// 声明第一个包含 3 个元素的指向字符串的指针数组
	var arr06 [3]*string
	fmt.Printf("arr06:\t%T\t%v\t%v\n", arr06, arr06, makePointerArrVal(arr06))

	// 声明第二个包含 3 个元素的指向字符串的指针数组，使用字符串指针初始化这个数组
	arr07 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*arr07[0] = "Red"
	*arr07[1] = "Blue"
	*arr07[2] = "Green"
	fmt.Printf("arr07:\t%T\t%v\t%v\n", arr07, arr07, makePointerArrVal(arr07))

	// 将 arr07 复制给 arr06，复制之后，两个数组指向同一组字符串
	arr06 = arr07
	fmt.Printf("\ncopy arr07 to arr06:\n")
	fmt.Printf("arr06:\t%T\t%v\t%v\n", arr06, arr06, makePointerArrVal(arr06))
	fmt.Printf("arr07:\t%T\t%v\t%v\n", arr07, arr07, makePointerArrVal(arr07))

	// 修改索引为 1 的指针元素的值
	*arr06[1] = "Blueeeeeeeeee" // 修改 arr06 的值会影响 arr07
	fmt.Printf("\nmodify index 1 pointer point value of arr6:\n")
	fmt.Printf("arr06:\t%T\t%v\t%v\n", arr06, arr06, makePointerArrVal(arr06))
	fmt.Printf("arr07:\t%T\t%v\t%v\n", arr07, arr07, makePointerArrVal(arr07))

	// 修改索引为 1 的指针值
	string := "sixsixsix-Blue"
	arr06[1] = &string
	fmt.Printf("\nmodify index 1 pointer value of arr6:\n")
	fmt.Printf("arr06:\t%T\t%v\t%v\n", arr06, arr06, makePointerArrVal(arr06))
	fmt.Printf("arr07:\t%T\t%v\t%v\n", arr07, arr07, makePointerArrVal(arr07))

	fmt.Printf("================================================================================\n")

	// 声明一个二维整型数组，两个维度分别存储 4 个元素和 2 个元素
	var darr01 [4][2]int
	// 为每个元素赋值
	darr01[0][0] = 10
	darr01[0][1] = 20
	darr01[1][0] = 30
	darr01[1][1] = 40
	fmt.Printf("darr01:\t%T\t%v\n", darr01, darr01)

	// 使用数组字面量来声明并初始化一个二维整型数组
	darr02 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Printf("darr02:\t%T\t%v\n", darr02, darr02)

	// 声明并初始化外层数组中索引为 1 个和 3 的元素
	darr03 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Printf("darr03:\t%T\t%v\n", darr03, darr03)

	// 声明并初始化外层数组和内层数组的单个元素
	darr04 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Printf("darr04:\t%T\t%v\n", darr04, darr04)

	fmt.Printf("================================================================================\n")
}

// makePointerArrVal 生成指针数组值的字符串格式
func makePointerArrVal(arr interface{}) string {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return ""
	}

	result := fmt.Sprint("[")
	for i := 0; i < v.Len(); i++ {
		if i != 0 {
			result += fmt.Sprint(" ")
		}

		if v.Index(i).IsNil() {
			result += fmt.Sprint(nil)
		} else {
			result += fmt.Sprint(v.Index(i).Elem())
		}

	}
	result += fmt.Sprint("]")

	return result
}

// makeIntPointerArrVal 生成int指针数组值的字符串格式
func makeIntPointerArrVal(arr []*int) string {
	result := fmt.Sprint("[")
	for pindex, pval := range arr {
		if pindex != 0 {
			result += fmt.Sprint(" ")
		}

		if pval != nil {
			result += fmt.Sprint(*pval)
		} else {
			result += fmt.Sprint(nil)
		}
	}
	result += fmt.Sprint("]")

	return result
}

// makeStringPointerArrVal 生成string指针数组值的字符串格式
func makeStringPointerArrVal(arr []*string) string {
	result := fmt.Sprint("[")
	for pindex, pval := range arr {
		if pindex != 0 {
			result += fmt.Sprint(" ")
		}

		if pval != nil {
			result += fmt.Sprint(*pval)
		} else {
			result += fmt.Sprint(nil)
		}
	}
	result += fmt.Sprint("]")

	return result
}
