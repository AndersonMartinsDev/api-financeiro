package despesa

const (
	FIXA      string = "FIXA"
	PARCELADA string = "PARCELADA"
	UNICA     string = "UNICA"
)

func ValueOf(s string) interface{} {
	switch s {
	case "FIXA":
		return FIXA
	case "PARCELADA":
		return PARCELADA
	case "UNICA":
		return UNICA
	}
	return ""
}
