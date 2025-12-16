package service

import (
	"go-ecommerce-api/internal/domain"
	"go-ecommerce-api/internal/dto"
	"log"
)


type UserService struct{

}



func (s UserService) SignUp(input dto.UserSignup)(string,error){
	log.Println(input)

	return "this-my-token-for-now",nil
}

func (s UserService) FindUserByEmail(email string) (*domain.User,error){
	
	return nil,nil
}

func (s UserService) Login (input any) (string , error){
	return "",nil

}

func (s UserService) GetVerificationCode(e domain.User) (int, error){
	return 0,nil
}

func (s UserService) VerifyCode(id uint,code int) error{
	return nil
}

func (s UserService) CreateProfile(id uint ,input any) error{
return nil
}

func (s UserService) GetProfile(id uint)(*domain.User,error){
	return nil,nil
}


func (s UserService) UpdateProfile(id uint,input any) error{
	return nil
}

func (s UserService) BecomeSeller(id uint,input any)(string,error){
	return "",nil 
}

func (s UserService) FindCart(id uint) ([]interface{} ,error){
	return nil,nil
}