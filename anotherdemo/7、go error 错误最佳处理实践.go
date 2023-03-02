package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func main() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}

//https://mp.weixin.qq.com/s?__biz=MzI2MzEwNTY3OQ==&mid=2648983868&idx=1&sn=87601fc86d898619bb915b0766848bc6&chksm=f25014dcc5279dca44299e41b26fa447348fd1fb3283fbfc23ed518b0214a0b62601f1e54f30&scene=27
