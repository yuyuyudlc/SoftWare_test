# Go 版本测试管理实施说明

## 1. 开发环境模块

本项目已采用以下模块化测试管理环境：

- Git 代码仓库：使用本地 Git 仓库（可同步到 GitHub/Gitee）
- 持续集成与测试：GitHub Actions（`.github/workflows/ci-go.yml`）自动执行 build/test
- Issue Tracking：使用 GitHub Issues，并通过 `.github/ISSUE_TEMPLATE/bug_report.yml` 规范缺陷提交

## 2. Go 项目配置（替代 Ant/Maven）

由于 Go 项目不生成 Jar，等价目标如下：

- `build`：导出完整可执行产物到 `dist/my-go-project`
- `test`：执行全量单元测试
- `test-smoke`：执行冒烟测试（关键用例子集）

对应配置在 `Makefile` 中：

- `make build`
- `make test`
- `make test-smoke`

## 3. 在持续集成中应用 test 任务进行冒烟测试

GitHub Actions 流程：

1. Checkout 代码
2. 执行 `make build`
3. 执行 `make test-smoke`
4. 执行 `make test`

在 GitHub Actions 日志中可以看到 `go test -v` 输出的测试用例执行日志，满足冒烟测试可追踪要求。

## 4. 缺陷提交流程与格式（Go 项目版）

建议的一般流程：

1. 发现缺陷并复现，确认稳定复现步骤
2. 在 Issue 工具中新建 Bug
3. 按模板填写关键信息：模块、环境、复现步骤、预期/实际、严重级别、日志
4. 指派处理人并设定优先级
5. 修复后关联提交与 PR，并在 CI 通过后关闭 Issue
6. 回归测试并记录结论

建议缺陷字段（最小集合）：

- 标题：`[BUG] 简短描述`
- 影响模块
- 环境信息（OS/Go 版本/分支/提交号）
- 复现步骤
- 预期结果
- 实际结果
- 严重级别
- 日志与截图

以上字段已在 `.github/ISSUE_TEMPLATE/bug_report.yml` 中落地，可直接用于课程实验提交。

## 5. 本次执行结果与缺陷登记

- 已执行：`make build`、`make test-smoke`
- 冒烟测试结果：发现 1 个缺陷（ToBoolE 对 "false" 解析异常）
- 已登记缺陷记录：`defects/BUG-20260421-001-ToBoolE-false.md`

该结果符合测试管理目标：通过 CI 冒烟测试尽早暴露关键缺陷，并进入规范化缺陷跟踪流程。

