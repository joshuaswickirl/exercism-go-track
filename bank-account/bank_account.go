package account

import "sync"

// Account is a bank account with a balance, open/closed state, and a mutex.
type Account struct {
	sync.RWMutex
	closed  bool
	balance int64
}

// Open constructs a new Account. It returns nil if the inital deposit is
// negative.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{closed: false, balance: initialDeposit}
}

// Close return pays out the balance and closes the account. If the account
// is already closed it returns 0, false.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	payout = a.balance
	a.balance = 0
	a.closed = true
	return payout, true
}

// Balance returns the current balance. If the Account is closed 0, false
// is returned.
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit updates the Account balance to calculated amount. If the deposit
// would bring the balance below zero, the transaction is aborted. In this
// case the exiting balance is returned with ok=false.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	if b := a.balance + amount; b < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}
