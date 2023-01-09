package models

type Recorrencia struct {
	Id            uint   `json:"id,omitempty"`
	Meses         uint32 `json:"meses,omitempty"`
	DiaVencimento uint8  `json:"diaVencimento,omitempty"`
}
