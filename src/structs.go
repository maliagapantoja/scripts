package src

type EgrhStruct struct {
	CreatedTime int64 `json:"createdTime"`
}
type Response struct {
	Egrh EgrhStruct `json:"egrh"`
}
type OutPutMessage struct {
	Message string
}