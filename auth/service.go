package auth 

type AuthorizationCore interface {
	SignUp() error 
	SignIn() error
	RefreshFlow() error 
	BlockFlow() error
}







 



 



