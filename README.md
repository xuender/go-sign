# go-sign

[![GoCI](https://github.com/xuender/go-sign/workflows/Go/badge.svg)](https://github.com/xuender/go-sign/actions)
[![codecov](https://codecov.io/gh/xuender/go-sign/branch/main/graph/badge.svg?token=QL31K7FRZ6)](https://codecov.io/gh/xuender/go-sign)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-sign)](https://goreportcard.com/report/github.com/xuender/go-sign)
[![GoDoc](https://godoc.org/github.com/xuender/go-sign?status.svg)](https://pkg.go.dev/github.com/xuender/go-sign)
[![Gitter](https://badges.gitter.im/xuender-go-sign/community.svg)](https://gitter.im/xuender-go-sign/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![GitHub license](https://img.shields.io/github/license/xuender/go-sign)](https://github.com/xuender/go-sign/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xuender/go-sign)](https://github.com/xuender/go-sign/issues)
[![GitHub stars](https://img.shields.io/github/stars/xuender/go-sign)](https://github.com/xuender/go-sign/stargazers)

Self verification after sign of golang lib.

[Changelog](http://github.com/xuender/go-sign/blob/master/History.md) | [中文](http://github.com/xuender/go-sign/blob/master/README_CN.md)

## Install cmd

```shell
go install github.com/xuender/go-sign/cmd/sign@latest
```

## Examples

### base

Check the integrity of the execution file to prevent tampering or virus intrusion.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-sign"
)

func main() {
  if err := sign.Check("secret_key"); err != nil {
    panic(err)
  }

  fmt.Println("Hello Word.")
  fmt.Println("This file integrity.")
}
```

```shell
go build -o helloword main.go
sign -s=secret_key helloword
```

### licence

Check license string.

```go
package main

import (
  "fmt"
  "os"

  "github.com/xuender/go-sign"
)

func main() {
  if len(os.Args) < 2 {
    panic("Miss licence.")
  }

  if err := sign.Check(os.Args[1]); err != nil {
    panic("Licence FAILED.")
  }

  fmt.Println("Hello Word.")
  fmt.Println("Licence OK.")
}
```

```shell
go build -o helloword main.go
sign -s=licence_str helloword
# run
./helloword licence_str
```

### env

Check environment variables.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-sign"
)

func main() {
  if err := sign.CheckEnv("SECRET_KEY"); err != nil {
    panic(err)
  }

  fmt.Println("Hello Word.")
  fmt.Println("Run on safe environment.")
}
```

```shell
go build -o helloword main.go
SECRET_KEY=secret_key sign -e=SECRET_KEY helloword
# set env and run
SECRET_KEY=secret_key ./helloword
```

### machine

Only run on the sign machine.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-sign"
)

func main() {
  if err := sign.CheckMachine(); err != nil {
    panic(err)
  }

  fmt.Println("Hello Word.")
  fmt.Println("Run on sign machine.")
}
```

```shell
go build -o helloword main.go
# sign on the final running machine
sign -m helloword
```

### complex

Only run on the sign machine and has env.

```go
package main

import (
  "fmt"
  "os"

  "github.com/xuender/go-sign"
)

func main() {
  mid := sign.GetMachineSecret(os.Getenv("SECRET_KEY"))
  if err := sign.Check(mid); err != nil {
    panic(err)
  }

  fmt.Println("Hello Word.")
  fmt.Println("Run on sign machine and has env.")
}
```

```shell
go build -o helloword main.go
# sign on the final running machine
SECRET_KEY=secret_key sign -m -e=SECRET_KEY helloword
# set env and run
SECRET_KEY=secret_key ./helloword
```

## PS

Use  sign and Check/CheckEnv/CheckMachine must be signed, otherwise it cannot run after build.

## License

© xuender, 2022~time.Now

[MIT License](https://github.com/xuender/go-sign/blob/master/License)
