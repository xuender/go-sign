# gosign

[![CircleCI](https://circleci.com/gh/xuender/gosign.svg?style=shield)](https://circleci.com/gh/xuender/gosign)
[![GoDoc](https://godoc.org/github.com/xuender/gosign?status.svg)](https://pkg.go.dev/github.com/xuender/gosign)

Self verification after sign of golang lib.

[Changelog](http://github.com/xuender/gosign/blob/master/History.md) | [中文](http://github.com/xuender/gosign/blob/master/README_CN.md)

## Install cmd

```shell
go install github.com/xuender/gosign/cmd/gosign@latest
```

## Examples

### base

Check the integrity of the execution file to prevent tampering or virus intrusion.

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

### licence

Check license string.

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
# run
./helloword licence_str
```

### env

Check environment variables.

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
# set env and run
SECRET_KEY=secret_key ./helloword
```

### machine

Only run on the sign machine.

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
# sign on the final running machine
gosign -m helloword
```

### complex

Only run on the sign machine and has env.

```go
package main

import (
	"fmt"
	"os"

	"github.com/xuender/gosign"
)

func main() {
	mid := gosign.GetMachineSecret(os.Getenv("SECRET_KEY"))
	if err := gosign.Check(mid); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine and has env.")
}
```

```shell
go build -o helloword main.go
# sign on the final running machine
SECRET_KEY=secret_key gosign -m -e=SECRET_KEY helloword
# set env and run
SECRET_KEY=secret_key ./helloword
```

## PS

Use gosign and Check/CheckEnv/CheckMachine must be signed, otherwise it cannot run after build.

## License

© xuender, 2022~time.Now

[MIT License](https://github.com/xuender/gosign/blob/master/License)
