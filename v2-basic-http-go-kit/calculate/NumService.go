package calculate

import "context"

type Service interface {
	Calculate(ctx context.Context, in CalculateIn) ResultAck
}

type CalculateIn struct {
	A    int    `json:"a"`
	B    int    `json:"b"`
	Type string `json:"type"`
}

type ResultAck struct {
	Res int `json:"res"`
}

type BaseServer struct{}

func (b BaseServer) Calculate(ctx context.Context, in CalculateIn) ResultAck {
	//TODO implement me
	if in.Type == "plus" {
		return ResultAck{Res: in.A + in.B}
	} else if in.Type == "minus" {
		return ResultAck{Res: in.A - in.B}
	}
	return ResultAck{Res: 0}
}

func NewService() Service {
	return &BaseServer{}
}
