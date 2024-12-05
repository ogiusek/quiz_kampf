// value object
package valid

type Validable interface {
	Valid() error
}

// Value() (driver.Value, error)
