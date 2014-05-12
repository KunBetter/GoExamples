// ConcurrentByCommunication
// Go语言的并发模型--通过通信来共享内存
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Account interface {
	Withdraw(uint)
	Deposit(uint)
	Balance() int
}

type Bank struct {
	account Account
}

func NewBank(account Account) *Bank {
	return &Bank{account: account}
}

func (bank *Bank) Withdraw(amount uint, actor_name string) {
	fmt.Println("[-]", amount, actor_name)
	bank.account.Withdraw(amount)
}

func (bank *Bank) Deposit(amount uint, actor_name string) {
	fmt.Println("[+]", amount, actor_name)
	bank.account.Deposit(amount)
}

func (bank *Bank) Balance() int {
	return bank.account.Balance()
}

type SimpleAccount struct {
	balance int
}

func NewSimpleAccount(balance int) *SimpleAccount {
	return &SimpleAccount{balance: balance}
}

func (acc *SimpleAccount) Deposit(amount uint) {
	acc.setBalance(acc.balance + int(amount))
}

func (acc *SimpleAccount) Withdraw(amount uint) {
	if acc.balance >= int(amount) {
		acc.setBalance(acc.balance - int(amount))
	} else {
		panic("杰克穷死")
	}
}

func (acc *SimpleAccount) Balance() int {
	return acc.balance
}

func (acc *SimpleAccount) setBalance(balance int) {
	acc.add_some_latency() //增加一个延时函数，方便演示
	acc.balance = balance
}

func (acc *SimpleAccount) add_some_latency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
}

type ConcurrentAccount struct {
	account   *SimpleAccount
	deposits  chan uint
	withdraws chan uint
	balances  chan chan int
}

func NewConcurrentAccount(amount int) *ConcurrentAccount {
	acc := &ConcurrentAccount{
		account:   &SimpleAccount{balance: amount},
		deposits:  make(chan uint),
		withdraws: make(chan uint),
		balances:  make(chan chan int),
	}
	acc.listen()

	return acc
}

/*
	共享资源被封装在一个控制流程中。
	结果就是资源成为了非共享状态。
	没有处理程序能够直接访问或者修改资源。
	你可以看到访问和修改资源的方法实际上并没有执行任何改变。
*/

func (acc *ConcurrentAccount) Balance() int {
	ch := make(chan int)
	acc.balances <- ch
	return <-ch
}

func (acc *ConcurrentAccount) Deposit(amount uint) {
	acc.deposits <- amount
}

func (acc *ConcurrentAccount) Withdraw(amount uint) {
	acc.withdraws <- amount
}

/*
	访问和修改是通过消息和控制流程通信。
	在控制流程中任何访问和修改的动作都是相继发生的。
	当控制流程接收到访问或者修改的请求后会立即执行相关动作。

	select 不断地从各个通道中取出消息，每个通道都跟它们所要执行的操作相一致。
	在 select 声明内部的一切都是相继执行的(在同一个处理程序中排队执行)。
	一次只有一个事件(在通道中接受或者发送)发生，这样就保证了同步访问共享资源。
*/

func (acc *ConcurrentAccount) listen() {
	go func() {
		for {
			select {
			case amnt := <-acc.deposits:
				acc.account.Deposit(amnt)
			case amnt := <-acc.withdraws:
				acc.account.Withdraw(amnt)
			case ch := <-acc.balances:
				ch <- acc.account.Balance()
			}
		}
	}()
}

func main() {
	balance := 80
	b := NewBank(NewConcurrentAccount(balance))

	fmt.Println("初始化余额", b.Balance())

	done := make(chan bool)

	go func() { b.Withdraw(30, "马伊琍"); done <- true }()
	go func() { b.Withdraw(10, "姚笛"); done <- true }()

	//等待 goroutine 执行完成
	<-done
	<-done

	fmt.Println("-----------------")
	fmt.Println("剩余余额", b.Balance())
}
