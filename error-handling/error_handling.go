package erratum

// Use opens a resource and recovers if it panics.
func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	for {
		resource, err = o()
		if _, ok := err.(TransientError); ok {
			continue
		} else if err != nil {
			return
		}
		break
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				resource.Defrob(e.defrobTag)
				err = e
			}
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	resource.Frob(input)
	return
}
