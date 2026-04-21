package main

import (
	"testing"

	"github.com/spf13/cast"
)

func TestToStringE(t *testing.T) {
	t.Run("可以将整数转换为字符串", func(t *testing.T) {
		got, err := cast.ToStringE(123)
		if err != nil {
			t.Fatalf("出现了意外错误：%v", err)
		}
		if got != "123" {
			t.Fatalf("得到 %q，期望 %q", got, "123")
		}
	})

	t.Run("可以将布尔值转换为字符串", func(t *testing.T) {
		got, err := cast.ToStringE(true)
		if err != nil {
			t.Fatalf("出现了意外错误：%v", err)
		}
		if got != "true" {
			t.Fatalf("得到 %q，期望 %q", got, "true")
		}
	})

	t.Run("nil 应转换为空字符串", func(t *testing.T) {
		got, err := cast.ToStringE(nil)
		if err != nil {
			t.Fatalf("出现了意外错误：%v", err)
		}
		if got != "" {
			t.Fatalf("得到 %q，期望空字符串", got)
		}
	})
}
