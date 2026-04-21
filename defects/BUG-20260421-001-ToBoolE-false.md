# [BUG] ToBoolE 解析 "false" 返回 true

- 缺陷编号: BUG-20260421-001
- 状态: Open
- 严重级别: High
- 影响模块: 类型转换（bool）
- 发现方式: CI 冒烟测试（`make test-smoke`）

## 环境信息

- OS: macOS
- Go: 以本地 `go version` 为准
- 分支: main/master（按实际）
- 提交号: 待填写

## 复现步骤

1. 在项目根目录执行 `make test-smoke`
2. 查看 `TestCastSuite/ToBoolE/可以解析_false_字符串` 用例

## 预期结果

- `cast.ToBoolE("false")` 返回 `false`

## 实际结果

- 返回值为 `true`，测试失败

## 证据日志

```text
=== RUN   TestCastSuite/ToBoolE/可以解析_false_字符串
main_test.go:46: 得到 true，期望 false
--- FAIL: TestCastSuite/ToBoolE/可以解析_false_字符串
```

## 处理建议

- 排查 `cast` 包中布尔字符串解析逻辑
- 增加针对 "false"、"0"、大小写混合的表驱动测试
