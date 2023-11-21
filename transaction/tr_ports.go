// Each movement of money is modeled through a Transaction object, which can be of different
// types (e.g. credit card, debit card, boleto, wire transfer, etc.). Each Transaction type
// has a Finite State Machine (FSM) that defines its current status. The Transaction Event
// object represents transitions in the Transaction's FSM.
//
// A Payment Provider can create a Transaction and update its status through Transaction Events
// as it changes over time. Since an order can be related to multiple payment methods, a
// different Transaction must be created for each of them using the same order ID.
//
// Note: This endpoint is for the exclusive use of payment apps.
package transaction

import (
	"encoding/json"
)

func defaultTransaction() TransactionOpts {
	return TransactionOpts{}
}

// [Required] ID of the Payment Provider that processed this Transaction.
func SetPaymentProviderID(paymentProviderID string) OptTra {
	return func(t *TransactionOpts) {
		t.PaymentProviderID = paymentProviderID
	}
}

// [Required] Object containing the payment method used in this Transaction. See Payment Method.
func SetPaymentMethod(p PaymentMethod) OptTra {
	return func(t *TransactionOpts) {
		t.PaymentMethod = p
	}
}

// [Required] First transaction event that generated this Transaction. See Transaction Events.
func SetFirstEvent(f FirstEvent) OptTra {
	return func(t *TransactionOpts) {
		t.FirstEvent = f
	}
}

// [Optional] Object containing specific info related to this Transaction. See Transaction Info.
func SetInfo(i Info) OptTra {
	return func(t *TransactionOpts) {
		t.Info = i
	}
}

// Initialize a transaction instance
func New(opts ...OptTra) *Transaction {
	dt := defaultTransaction()
	for _, fn := range opts {
		fn(&dt)
	}
	return &Transaction{
		TransactionOpts: dt,
	}
}

// Create a Transaction for a given order.
func (t Transaction) Create(orderID string) (r Transaction, err error) {
	j, err := json.Marshal(t)
	if err != nil {
		return
	}
	br, err := create(j, orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Create a Transaction Event for a given Transaction.
func (t Transaction) CreateEvent(orderID, transactionID string) (r Transaction, err error) {
	j, err := json.Marshal(t)
	if err != nil {
		return
	}
	br, err := createEvent(j, orderID, transactionID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Get all Transactions for a given order.
func GetAll(orderID string) (r []Transaction, err error) {
	b, err := getAll(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Get a specific Transaction for a given order.
func GetById(orderID, transactionID string) (r Transaction, err error) {
	b, err := getByID(orderID, transactionID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}
