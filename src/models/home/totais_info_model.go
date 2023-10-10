package home

type TotaisInfoModel struct {
	BalancoMes     float64 `json:"balancoMes,omitempty"`
	TotalPago      float64 `json:"totalPago,omitempty"`
	PorcentagemPie uint    `json:"porcentagemPie,omitempty"`
}
