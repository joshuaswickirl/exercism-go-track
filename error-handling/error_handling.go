package erratum

import "fmt"

// Use invokes a ResourceOpener and calls the resource's Frob method
// passing the input.
// Opening the resource is retried if there is a TransientError and
// returns any other error. Panics are recovered and if the Panic is a
// FrobError, the resource is Defrob'ed with the defrobTag from the error.
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			default:
				err = fmt.Errorf("recovered value %v of type %T", e, e)
			}
		}
	}()
	resource.Frob(input)
	return err
}
