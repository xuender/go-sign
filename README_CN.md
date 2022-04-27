# gosign

签名后可自校验的 Go 语言类库。

[历史](http://github.com/xuender/gosign/blob/master/History.md) | [English](http://github.com/xuender/gosign/blob/master/README.md)

## 安装命令行

```shell
go install github.com/xuender/gosign/cmd/gosign@latest
```

## 例子

### 完整性校验

检查执行文件完整性。

```go
package main

import (
	"fmt"

	"github.com/xuender/gosign"
)

func main() {
	if err := gosign.Check("secret_key"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("This file integrity.")
}
```

```shell
go build -o helloword main.go
gosign -s=secret_key helloword
```

### 许可证校验

检查许可证。

```go
package main

import (
	"fmt"
	"os"

	"github.com/xuender/gosign"
)

func main() {
	if len(os.Args) < 2 {
		panic("Miss licence.")
	}

	if err := gosign.Check(os.Args[1]); err != nil {
		panic("Licence FAILED.")
	}

	fmt.Println("Hello Word.")
	fmt.Println("Licence OK.")
}
```

```shell
go build -o helloword main.go
gosign -s=licence_str helloword
# 使用许可证运行
./helloword licence_str
```

### 环境变量校验

检查环境变量。

```go
package main

import (
	"fmt"

	"github.com/xuender/gosign"
)

func main() {
	if err := gosign.CheckEnv("SECRET_KEY"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on safe environment.")
}
```

```shell
go build -o helloword main.go
SECRET_KEY=secret_key gosign -e=SECRET_KEY helloword
# 设置环境变量并运行
SECRET_KEY=secret_key ./helloword
```

### 设备校验

只能运行在签名的设备上。

```go
package main

import (
	"fmt"

	"github.com/xuender/gosign"
)

func main() {
	if err := gosign.CheckMachine(); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine.")
}
```

```shell
go build -o helloword main.go
# 在运行的设备上签名
gosign -m helloword
```

### 多重校验

只能运行在签名的设备上，并校验环境变量。

```go
package main

import (
	"fmt"
	"os"

	"github.com/xuender/gosign"
)

func main() {
	mid, _ := gosign.GetMachineSecret(os.Getenv("SECRET_KEY"))
	if err := gosign.Check(mid); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine and has env.")
}
```

```shell
go build -o helloword main.go
# 在运行的设备上根据环境变量签名
SECRET_KEY=secret_key gosign -m -e=SECRET_KEY helloword
# 设置环境变量并运行
SECRET_KEY=secret_key ./helloword
```

## 说明

使用 gosign 的 Check/CheckEnv/CheckMachine 方法，编译后必须签名。

## License

© xuender, 2022~time.Now

[MIT License](https://github.com/xuender/gosign/blob/master/License)
