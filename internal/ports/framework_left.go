package ports 

import (
	_"context"
)

type GRPCPort interface {
	Run()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}