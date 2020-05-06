package cmd

func ThrowIf(err error) {

	if err != nil {
		panic(err)
	}
}
