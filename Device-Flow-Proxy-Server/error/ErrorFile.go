package error

type Error struct {
	Error            string `json:"error"  validate:"required"`
	ErrorDescription string `json:"error_description" validate:"required"`
}

