package transformations

func TransformationsAvailable() map[string]func(params TransformationParams) {
	return map[string]func(params TransformationParams){
		"--blur":     Blur,
		"--negative": Negative,
	}
}
