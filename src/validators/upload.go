package validators

type Image struct {
	Data string `param:"data" validate:"required,datauri"` //必須パラメータ
}
