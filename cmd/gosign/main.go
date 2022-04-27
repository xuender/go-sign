package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuender/gosign"
)

func main() {
	defer errorHanlder()

	var (
		isCheck, isMachine bool
		envName, secret    string
	)

	flag.Usage = usage
	flag.BoolVar(&isCheck, "c", false, "check sign")
	flag.BoolVar(&isMachine, "m", false, "check machine")
	flag.StringVar(&envName, "e", "", "environment variable")
	flag.StringVar(&secret, "s", "", "secret key")
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()

		return
	}

	if envName != "" {
		secret = os.Getenv(envName)
	}

	if isMachine {
		secret = Panic1(gosign.GetMachineSecret(secret))
	}

	if isCheck {
		check(secret)

		return
	}

	for _, arg := range flag.Args() {
		path := Panic1(abs(arg))
		sum := gosign.NewSign(path, []byte(secret))
		base := filepath.Base(path)

		if !sum.Hasgosign {
			panic(fmt.Sprintf("file %s not use gosign.", base))
		}

		Panic(sum.Sign())
		fmt.Fprintf(flag.CommandLine.Output(), "%s: sign OK.\n", base)
	}
}

func check(secret string) {
	for _, arg := range flag.Args() {
		path := Panic1(abs(arg))
		err := gosign.CheckFile(path, secret)

		if err == nil {
			fmt.Fprintf(flag.CommandLine.Output(), "%s: sign OK.\n", filepath.Base(path))

			continue
		}

		if errors.Is(gosign.ErrSignFailed, err) {
			fmt.Fprintf(flag.CommandLine.Output(), "%s: sign Failed.\n", filepath.Base(path))

			continue
		}

		panic(err)
	}
}

func errorHanlder() {
	if err := recover(); err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "gosign: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "gosign [%s]\n\n", gosign.ModVersion)
	fmt.Fprintf(flag.CommandLine.Output(), "Self verification of golang lib.\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [path ...]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nMod: %s\nGit: %s\n", gosign.ModPath, gosign.ModSum)
}

func abs(path string) (string, error) {
	if path == "" {
		return "", gosign.ErrFileName
	}

	if path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return path, err
		}

		path = home + path
	}

	return filepath.Abs(path)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Panic1[T any](t T, err error) T {
	Panic(err)

	return t
}
