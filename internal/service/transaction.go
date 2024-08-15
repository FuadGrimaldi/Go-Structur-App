package service

import (
	"errors"
	"go-app/internal/entity"
	"go-app/internal/repository"

	"golang.org/x/net/context"
)

type TransactionService interface {
	FindTransactionByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error)
}

type transactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{repository}
}

var ErrNoTransactionsFound = errors.New("no transactions found for the given user ID")

func (s *transactionService) FindTransactionByUserID(ctx context.Context, userID int64) ([]entity.Transaction, error) {
	transaction, err := s.repository.FindTransactionByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Check if the result is empty
	if len(transaction) == 0 {
		return nil, ErrNoTransactionsFound
	}

	return transaction, nil
}