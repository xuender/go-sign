# gosign

Self verification after sign of golang lib.

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
# sign on the final running machine
SECRET_KEY=secret_key gosign -m -e=SECRET_KEY helloword
# set env and run
SECRET_KEY=secret_key ./helloword
```

## PS

Use gosign and Check/CheckEnv/CheckMachine must be signed, otherwise it cannot run after build.
