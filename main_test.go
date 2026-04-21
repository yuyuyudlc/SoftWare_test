package main

import (
	"testing"

	"github.com/spf13/cast"
)

func TestCastSuite(t *testing.T) {
	t.Run("ToIntE", func(t *testing.T) {
		t.Run("可以解析合法数字字符串", func(t *testing.T) {
			got, err := cast.ToIntE("123")
			if err != nil {
				t.Fatalf("出现了意外错误：%v", err)
			}
			if got != 123 {
				t.Fatalf("得到 %d，期望 %d", got, 123)
			}
		})

		t.Run("遇到非法字符串应返回错误", func(t *testing.T) {
			_, err := cast.ToIntE("abc")
			if err == nil {
				t.Fatal("字母字符串输入应返回错误，但实际没有返回错误")
			}
		})
	})

	t.Run("ToBoolE", func(t *testing.T) {
		t.Run("可以解析 true 字符串", func(t *testing.T) {
			got, err := cast.ToBoolE("true")
			if err != nil {
				t.Fatalf("出现了意外错误：%v", err)
			}
			if !got {
				t.Fatal("得到 false，期望 true")
			}
		})

		t.Run("可以解析 false 字符串", func(t *testing.T) {
			got, err := cast.ToBoolE("false")
			if err != nil {
				t.Fatalf("出现了意外错误：%v", err)
			}
			if got {
				t.Fatal("得到 true，期望 false")
			}
		})
	})
}
