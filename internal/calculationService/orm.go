package calculationservice

type Calculation struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculatuionRequest struct {
	Expression string `json:"expression"`
}
