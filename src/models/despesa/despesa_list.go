package despesa

type ListModel struct {
	BalancoTotal float64     `json:"balancototal,omitempty"`
	Elements     interface{} `json:"elements"`
}
