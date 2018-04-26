package main

import "fmt"

type Account struct {
	Email    string
	Password string
	Rank     int
}

func updatePassword(account Account, password string) {
	account.Password = password
	fmt.Printf("account.Password(func)  : %s\n", account.Password)
	fmt.Printf("account.Password(pointer): %p\n", &account)
}

func (account Account) updatePasswordMethod(password string) {
	account.Password = password
	fmt.Printf("account.Password(func) : %s\n", account.Password)
	fmt.Printf("account.Password(pointer): %p\n", &account)
}

func updatePasswordP(account *Account, password string) {
	account.Password = password
	fmt.Printf("accountP.Password(func)  : %s\n", account.Password)
	fmt.Printf("accountP.Password(pointer): %p\n", account)
}

func main() {
	fmt.Println("--function--")
	account := Account{
		Email:    "sample@example.com",
		Password: "password",
		Rank:     1,
	}
	fmt.Printf("account.Password(pointer): %p\n", &account)
	fmt.Printf("account.Password(before): %s\n", account.Password)
	updatePassword(account, "new password")
	fmt.Printf("account.Password(after) : %s\n", account.Password)

	fmt.Println("-----------------")
	fmt.Println("--method--")
	accountM := Account{
		Email:    "sample@example.com",
		Password: "password",
		Rank:     1,
	}
	fmt.Printf("account.Password(pointer): %p\n", &accountM)
	fmt.Printf("account.Password(before): %s\n", accountM.Password)
	accountM.updatePasswordMethod("new password")
	fmt.Printf("account.Password(after) : %s\n", accountM.Password)

	fmt.Println("-----------------")
	fmt.Println("--pointer--")
	accountP := &Account{
		Email:    "sample@example.com",
		Password: "password",
		Rank:     1,
	}
	fmt.Printf("accountP.Password(pointer): %p\n", accountP)
	fmt.Printf("accountP.Password(before): %s\n", accountP.Password)
	updatePasswordP(accountP, "new password")
	fmt.Printf("accountP.Password(after) : %s\n", accountP.Password)
}
