package erratum

func Use(o ResourceOpener, input string) (err error) {
  res, err := o()
  for err != nil {
    if _, ok := err.(TransientError); !ok {
      return err
    }
    res, err = o()
  }
  defer res.Close()

  defer func() {
    rec := recover();
    if rec != nil {
      frobErr, ok := rec.(FrobError)
      if ok {
        res.Defrob(frobErr.defrobTag)
      }
      err = rec.(error)
    }
  }()

  res.Frob(input)
  return
}

