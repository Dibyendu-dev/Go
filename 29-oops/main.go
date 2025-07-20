package main

import (
	"fmt"
	"os"
)

// ---------------Banking Sysytem-------------------

// interface =>binding contracts
type BankService interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	ShowBalance()
}

// struct =>class implementation
type Account struct {
	AccountNumber string
	HolderName    string
	Password      string
	Balance       float64
}

type Bank struct {
	Accounts map[string]*Account
}

// constructor fn
func NewBank() *Bank {
	return &Bank{Accounts: make(map[string]*Account)}
}

func NewAccount(accNo, name, password string, initialDeposit float64) *Account {
	return &Account{
		AccountNumber: accNo,
		HolderName:    name,
		Password:      password,
		Balance:       initialDeposit,
	}
}

// method fn for Account struct
func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("❌ Amount must be positive")
		return
	}
	a.Balance += amount
	fmt.Println("✅ Deposited:", amount)

}

func (a *Account) Withdraw(amount float64) {
	if amount > a.Balance {
		fmt.Println("❌ Insufficient Balance.")
		return
	}
	a.Balance -= amount
	fmt.Println("✅ Withdraw:", amount)
}

func (a *Account) ShowBalance() {
	fmt.Printf(" Account: %s | Account Holder Name: %s | Balance:  ₹%.2f\n", a.AccountNumber, a.HolderName, a.Balance)
}

// method fn for Bank struct
func (b *Bank) CreateAccount(accNo, name, password string, deposit float64) {
	if _, exists := b.Accounts[accNo]; exists {
		fmt.Println("❌ Account already exsists")
		return
	}
	b.Accounts[accNo] = NewAccount(accNo, name, password, deposit)
	fmt.Println("✅ Account created for:", name)
}

func (b *Bank) GetAccount(accNo string) *Account {
	return b.Accounts[accNo]
}

func (b *Bank) DeleteAccount(accNo string, password string) {
	acc := b.GetAccount(accNo)
	if acc == nil {
		fmt.Println("❌ Account not found.")
		return
	}
	if acc.Password != password {
		fmt.Println("❌ Incorrect password.")
		return
	}
	delete(b.Accounts, accNo)
	fmt.Println("✅ Account deleted.")
}

func (b *Bank) UpdateAccount(accNo, password, newName, newPassword string) {
	acc := b.GetAccount(accNo)
	if acc == nil {
		fmt.Println("❌ Account not found.")
		return
	}
	if acc.Password != password {
		fmt.Println("❌ Incorrect password.")
		return
	}
	acc.HolderName = newName
	acc.Password = newPassword
	fmt.Println("✅ Account updated.")
}

func (b *Bank) Transfer(fromAcc, toAcc, password string, amount float64) {
	from := b.GetAccount(fromAcc)
	to := b.GetAccount(toAcc)

	if from == nil || to == nil {
		fmt.Println("❌  both accounts not found.")
		return
	}
	if from.Password != password {
		fmt.Println("❌ Incorrect password.")
		return
	}
	if amount > from.Balance {
		fmt.Println("❌ Insufficient funds.")
		return
	}
	from.Balance -= amount
	to.Balance += amount
	fmt.Printf("✅ ₹%.2f transferred from %s to %s\n", amount, fromAcc, toAcc)
}

func main() {
	bank := NewBank()

	for {
		fmt.Println("\n🏦 Bank Menu")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Show Balance")
		fmt.Println("5. Transfer Money")
		fmt.Println("6. Update Account")
		fmt.Println("7. Delete Account")
		fmt.Println("8. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var accNo, name, password string
			var deposit float64
			fmt.Print("Enter Account Number: ")
			fmt.Scan(&accNo)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Set Password: ")
			fmt.Scan(&password)
			fmt.Print("Initial Deposit: ")
			fmt.Scan(&deposit)
			bank.CreateAccount(accNo, name, password, deposit)

		case 2:
			var accNo, password string
			var amount float64
			fmt.Print("Enter Account Number: ")
			fmt.Scan(&accNo)
			fmt.Print("Enter Password: ")
			fmt.Scan(&password)
			acc := bank.GetAccount(accNo)
			if acc != nil && acc.Password == password {
				var service BankService = acc
				fmt.Print("Amount to deposit: ")
				fmt.Scan(&amount)
				service.Deposit(amount)
			} else {
				fmt.Println("❌ Invalid credentials.")
			}
		case 3:
			var accNo, password string
			var amount float64
			fmt.Print("Enter Account Number: ")
			fmt.Scan(&accNo)
			fmt.Print("Enter Password: ")
			fmt.Scan(&password)
			acc := bank.GetAccount(accNo)
			if acc != nil && acc.Password == password {
				var service BankService = acc
				fmt.Print("Amount to withdraw: ")
				fmt.Scan(&amount)
				service.Withdraw(amount)
			} else {
				fmt.Println("❌ Invalid credentials.")
			}
		case 4:
			var accNo, password string
			fmt.Print("Enter Account Number: ")
			fmt.Scan(&accNo)
			fmt.Print("Enter Password: ")
			fmt.Scan(&password)
			acc := bank.GetAccount(accNo)
			if acc != nil && acc.Password == password {
				var service BankService = acc
				service.ShowBalance()
			} else {
				fmt.Println("❌ Invalid credentials.")
			}

		case 5:
			var from, to, password string
			var amount float64
			fmt.Print("From Account Number: ")
			fmt.Scan(&from)
			fmt.Print("To Account Number: ")
			fmt.Scan(&to)
			fmt.Print("Your Password: ")
			fmt.Scan(&password)
			fmt.Print("Amount to transfer: ")
			fmt.Scan(&amount)
			bank.Transfer(from, to, password, amount)

		case 6:
			var accNo, oldPassword, newName, newPassword string
			fmt.Print("Enter Account Number: ")
			fmt.Scan(&accNo)
			fmt.Print("Enter Current Password: ")
			fmt.Scan(&oldPassword)
			fmt.Print("Enter New Name: ")
			fmt.Scan(&newName)
			fmt.Print("Enter New Password: ")
			fmt.Scan(&newPassword)
			bank.UpdateAccount(accNo, oldPassword, newName, newPassword)

		case 7:
			var accNo, password string
			fmt.Print("Enter Account Number: ")
			fmt.Scan(&accNo)
			fmt.Print("Enter Password: ")
			fmt.Scan(&password)
			bank.DeleteAccount(accNo, password)

		case 8:
			fmt.Println("👋 Exiting SimpleBank. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("❌ Invalid option. Try again.")
		
		}
	}
}
