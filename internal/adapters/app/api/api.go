package api 

import (
	"github.com/karankumarshreds/GoHexArch/internal/ports"
)

type Adapter struct {
	// in order to access the underneath layer 
	// we need to have its port 
	arith ports.ArithmeticPort
}

/* think of the api/application layer as controllers */

func (apiA Adapter) GetAddition(a, b int32) (int32, error){
	result, err := apiA.arith.Addition(a, b)
	return result, err
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error){
	result, err := apiA.arith.Subtraction(a, b)
	return result, err
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error){
	result, err := apiA.arith.Multiplication(a, b)
	return result, err
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error){
	result, err := apiA.arith.Division(a, b)
	return result, err
}

