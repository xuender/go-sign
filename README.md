# gosign

Self verification of golang lib.

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

## PS

Use gosign and Check/CheckEnv/CheckMachine must be signed, otherwise it cannot run after build.
