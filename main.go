package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	runManualTests()
}

type manualCase struct {
	name    string
	passed  bool
	details string
}

func runManualTests() {
	cases := []manualCase{
		checkInt("ToIntE 可以解析数字字符串", "123", 123, false),
		checkInt("ToIntE 遇到字母字符串应返回错误", "abc", 0, true),
		checkBool("ToBoolE 可以解析 true", "true", true, false),
		checkBool("ToBoolE 可以解析 false", "false", false, false),
	}

	passed := 0
	for _, tc := range cases {
		if tc.passed {
			fmt.Println("[通过]", tc.name)
			passed++
			continue
		}
		fmt.Printf("[失败] %s: %s\n", tc.name, tc.details)
	}

	fmt.Printf("\n手工测试汇总：通过 %d 项，共 %d 项\n", passed, len(cases))
}

func checkInt(name string, input any, want int, wantErr bool) manualCase {
	got, err := cast.ToIntE(input)

	if wantErr {
		if err != nil {
			return manualCase{name: name, passed: true}
		}
		return manualCase{
			name:    name,
			passed:  false,
			details: fmt.Sprintf("期望返回错误，但得到了数值 %d", got),
		}
	}

	if err != nil {
		return manualCase{
			name:    name,
			passed:  false,
			details: fmt.Sprintf("出现了意外错误：%v", err),
		}
	}
	if got != want {
		return manualCase{
			name:    name,
			passed:  false,
			details: fmt.Sprintf("得到 %d，期望 %d", got, want),
		}
	}

	return manualCase{name: name, passed: true}
}

func checkBool(name string, input any, want bool, wantErr bool) manualCase {
	got, err := cast.ToBoolE(input)

	if wantErr {
		if err != nil {
			return manualCase{name: name, passed: true}
		}
		return manualCase{
			name:    name,
			passed:  false,
			details: fmt.Sprintf("期望返回错误，但得到了数值 %v", got),
		}
	}

	if err != nil {
		return manualCase{
			name:    name,
			passed:  false,
			details: fmt.Sprintf("出现了意外错误：%v", err),
		}
	}
	if got != want {
		return manualCase{
			name:    name,
			passed:  false,
			details: fmt.Sprintf("得到 %v，期望 %v", got, want),
		}
	}

	return manualCase{name: name, passed: true}
}
