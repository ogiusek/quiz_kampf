package valid

func ValidMany(objects []Validable) error {
	for _, v := range objects {
		if v == nil {
			continue
		}
		if err := v.Valid(); err != nil {
			return err
		}
	}
	return nil
}
