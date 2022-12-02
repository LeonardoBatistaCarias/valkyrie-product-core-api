package product

type Brand int32

const (
	APPLE    Brand = 1
	SAMSUNG  Brand = 2
	XIAOMI   Brand = 3
	MOTOROLA Brand = 4
)

func GetBrandFrom(code int32) Brand {
	return Brand(code)
}
