package transaction

import (
	"encoding/json"
	"github.com/netoudi/codepay-process-transaction/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (k *KafkaPresenter) Bind(input interface{}) error {
	k.ID = input.(process_transaction.TransactionDtoOutput).ID
	k.Status = input.(process_transaction.TransactionDtoOutput).Status
	k.ErrorMessage = input.(process_transaction.TransactionDtoOutput).ErrorMessage
	return nil
}

func (k *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	return j, nil
}
