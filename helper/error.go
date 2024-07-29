package helper

import "uaspw2/exception"

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfNotFound(err error, errMessage string) {
	if err != nil {
		panic(exception.NewNotFoundError(errMessage))
	}
}
