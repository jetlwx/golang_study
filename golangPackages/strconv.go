package golangPackages

import (
	"fmt"
	"strconv"
)

func Study() {
	/*
		    常见错误类型
		    // ErrRange 表示值超出范围
		var ErrRange = errors.New("value out of range")

		// ErrSyntax 表示语法不正确
		var ErrSyntax = errors.New("invalid syntax")

		// NumError 记录转换失败
		type NumError struct {
		Func string // 失败的函数名(ParseBool, ParseInt, ParseUint, ParseFloat)
		Num string // 输入的值
		Err error // 失败的原因(ErrRange, ErrSyntax)

	*/

	/*
		　　　　常见参数含义

			　f：要转换的浮点数
			 fmt：格式标记（b、e、E、f、g、G）
			 prec：精度（数字部分的长度，不包括指数部分）
			 bitSize：指定浮点类型（32:float32、64:float64）

			　格式标记：
			'b' (-ddddp±ddd，二进制指数)
			 'e' (-d.dddde±dd，十进制指数)
			 'E' (-d.ddddE±dd，十进制指数)
			 'f' (-ddd.dddd，没有指数)
			 'g' ('e':大指数，'f':其它情况)
			'G' ('E':大指数，'f':其它情况)

			如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
			 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）

	*/
	// int to string
	var s string
	var i int
	i = 5
	s = strconv.Itoa(i)
	fmt.Println("s = ", s) //s =  5
	fmt.Println("-----------------------------------------------------------")

	//string to int
	var s1 string
	s1 = "5"
	i1, err := strconv.Atoi(s1)
	fmt.Println("i1= ", i1, "err", err) // i1=  5 err <nil>

	s2 := "Hello word"
	i2, err2 := strconv.Atoi(s2)
	fmt.Println("i2= ", i2, "err2", err2) //i2=  0 err2 strconv.ParseInt: parsing "Hello word": invalid syntax
	fmt.Println("-----------------------------------------------------------")

	//在字节数组后追加一个bool类型的值
	// AppendBool(dst []byte, b bool) []byte
	// eg
	b := []byte("bool:")
	fmt.Println(string(b)) // bool:
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b)) //bool:true
	fmt.Println("-----------------------------------------------------------")

	//AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
	//　追加成科学记数法
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'e', -1, 32)
	fmt.Println(string(b32)) //float32:3.1415927e+00

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'e', -1, 64)
	fmt.Println(string(b64)) // float64:3.1415926535e+00
	fmt.Println("-----------------------------------------------------------")

	//AppendInt(dst []byte, i int64, base int) []byte
	//十进制和１６进制
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) //int (base 10):-42

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) //int (base 16):-2a
	fmt.Println("-----------------------------------------------------------")

	//AppendQuote(dst []byte, s string) []byte
	//追加双引号
	b1 := []byte("quote:")
	b1 = strconv.AppendQuote(b1, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b1)) //quote:"\"Fran & Freddie's Diner\""
	fmt.Println("-----------------------------------------------------------")

	//AppendQuoteRune(dst []byte, r rune) []byte
	//追加字元类型
	b3 := []byte("rune:")
	b3 = strconv.AppendQuoteRune(b3, '☺')
	fmt.Println(string(b3)) //rune:'☺'
	fmt.Println("-----------------------------------------------------------")
}
