package error_handling

type PaymentServiceError struct {
	Err error
}

func (err PaymentServiceError) Error() string {
	return err.Err.Error()
}
