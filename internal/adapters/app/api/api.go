package api 

import (
	"github.com/karankumarshreds/GoHexArch/internal/ports"
)

type Adapter struct {
	// in order to access the underneath layer 
	// we need to have its port 
	arith ports.ArithmeticPort
	db ports.DBPort
}

/* think of the api/application layer as controllers */

func NewAdapter(arith ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error){
	result, err := apiA.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(result, "addition")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error){
	result, err := apiA.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(result, "Subtraction")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error){
	result, err := apiA.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(result, "multiplication")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error){
	result, err := apiA.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apiA.db.AddToHistory(result, "division")
	if err != nil {
		return 0, err
	}
	return result, nil
}

