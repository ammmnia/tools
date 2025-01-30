package errs

import "github.com/pkg/errors"

func Unwrap(err error) error {
	for err != nil {
		unwrap, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			break
		}
		err = unwrap.Unwrap()
	}
	return err
}

func Wrap(err error) error {
	return errors.WithStack(err)
}

func WrapMsg(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	withMessage := errors.WithMessage(err, toString(msg, kv))
	return errors.WithStack(withMessage)
}

func WrapMessage(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	return errors.New(msg)
}
