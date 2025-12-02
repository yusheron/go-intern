# go-intern

这是我的 Go 学习练习仓库，用来一步一步熟悉 Go 语言、命令行工具、测试和后端开发。

目前包含的内容：

- 字符串反转命令行工具（`cmd/reverse`）
- 词频统计命令行工具（`cmd/wordfreq`）
- 文本处理工具包（`pkg/textutil`）
- 基本的单元测试示例
- 简单的 HTTP `/ping` 接口服务（`cmd/server`）

---

## 环境准备

- Go 版本：1.20+（本地安装好 Go）
- 操作系统：Windows（项目路径示例：`C:\Users\18734\go-intern`）

克隆仓库（如果在别的机器上）：

```bash
git clone https://github.com/yusheron/go-intern.git
cd go-intern
```

---

## 项目结构

当前主要目录结构：

```text
go-intern/
  cmd/
    reverse/        # 字符串反转命令行工具
      main.go
    wordfreq/       # 词频统计命令行工具
      main.go
    server/         # 简单的 HTTP 服务（包含 /ping 接口）
      main.go
  pkg/
    textutil/       # 文本工具函数
      reverse.go
      reverse_test.go
      wordcount.go
      wordcount_test.go
      repeat.go
      repeat_test.go
  README.md
  go.mod
```

---

## 功能说明与使用方法

### 1. 字符串反转（cmd/reverse）

从标准输入读取一行字符串，输出反转后的结果。支持中文。

#### 运行方式

在项目根目录下执行：

```bash
go run ./cmd/reverse
```

或者在 GoLand 中运行 `cmd/reverse/main.go` 的 `main` 函数。

#### 示例

输入：

```text
hello
```

输出：

```text
olleh
```

输入：

```text
你好Go
```

输出：

```text
oG好你
```

---

### 2. 词频统计（cmd/wordfreq）

从标准输入读取一行英文文本，按空白字符拆分成单词，统计每个单词出现的次数。

#### 运行方式

在项目根目录下执行：

```bash
go run ./cmd/wordfreq
```

#### 示例

输入：

```text
hello world hello
```

可能的输出（顺序不固定）：

```text
hello 2
world 1
```

---

### 3. 简单 HTTP `/ping` 服务（cmd/server）

启动一个本地 HTTP 服务，提供 `/ping` 接口，用来练习 Go 的 `net/http` 基础。

#### 运行方式

在项目根目录下执行：

```bash
go run ./cmd/server
```

终端会输出类似：

```text
Starting server on :8080
```

表示服务已经在本地 `8080` 端口启动。

#### 测试 `/ping` 接口

在浏览器访问：

```text
http://localhost:8080/ping
```

或者在 PowerShell / 终端中使用：

```bash
curl http://localhost:8080/ping
```

预期返回：

```text
pong
```

---

## 文本工具包（pkg/textutil）

### ReverseString

```go
func ReverseString(s string) string
```

- 功能：反转字符串，正确处理中文等多字节字符。

### WordCount

```go
func WordCount(line string) map[string]int
```

- 功能：统计一行文本中每个单词出现的次数，按空白字符拆分。

### RepeatString

```go
func RepeatString(s string, n int) string
```

- 功能：将字符串 `s` 重复 `n` 次，例如 `RepeatString("go", 3)` 返回 `"gogogo"`。

---

## 运行测试

在项目根目录下执行：

```bash
go test ./...
```

会运行 `pkg/textutil` 下的所有测试，包括：

- `reverse_test.go`
- `wordcount_test.go`
- `repeat_test.go`

---

## 学习记录（说明）

这个仓库是我从零开始学习 Go 时使用的练习项目，目前已经完成：

- 基本命令行工具的实现；
- 使用 `testing` 包编写简单单元测试；
- 使用 Git 和 GitHub 管理每日学习进度；
- 使用 Go 的 `net/http` 标准库实现了一个简单的 HTTP `/ping` 接口。

后续计划 / 当前进度：

- [x] 基本命令行工具（字符串反转、词频统计）
- [x] 使用 `testing` 包编写简单单元测试
- [x] 添加 `RepeatString` 工具函数及测试
- [x] 添加 HTTP 服务基础框架，并实现 `/ping` 接口
- [ ] 实现一个简单的笔记/Todo 后端服务
- [ ] 继续补充更多测试和练习代码