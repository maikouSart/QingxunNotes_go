package main

import (
	//"github.com/stretchr/testfi/assert"
	"testing"
)

// 1.测试文件以_test.go结尾
// 2.func TestXxx(*testing.T)
// 3.初始化逻辑放到TestMain中
func HelloTom() string {
	return "jerry"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	if output != expectOutput {
		t.Errorf("expect %s", output)
	}
	//assert.Equal(t, expectOutput, output)
}

func TestManyGoWait(t *testing.T) {
	//测试前的配置、初始化操作
	//代码测试
	//测试后的资源释放操作
}
