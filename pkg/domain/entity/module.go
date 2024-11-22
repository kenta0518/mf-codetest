package entity

func Entity() []any {
	return concatSlices([]any{
		// User
		&User{},
		&Transaction{},
	})
}

func concatSlices(slices ...[]any) []any {
	var result []any
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
