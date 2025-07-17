package implementations

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/interfaces"
	"context"
	"time"
)

// PaymentMethodUsecase implements the payment method usecase interface
type PaymentMethodUsecase struct {
	repo *repository.RepositoryManager
}

// NewPaymentMethodUsecase creates a new payment method usecase
func NewPaymentMethodUsecase(repo *repository.RepositoryManager) interfaces.PaymentMethodUsecase {
	return &PaymentMethodUsecase{repo: repo}
}

// CreatePaymentMethod creates a new payment method
func (u *PaymentMethodUsecase) CreatePaymentMethod(ctx context.Context, req interfaces.CreatePaymentMethodRequest) (*models.PaymentMethod, error) {
	paymentMethod := &models.PaymentMethod{
		Name:      req.Name,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
	}

	if paymentMethod.Status == "" {
		paymentMethod.Status = models.StatusAktif
	}

	err := u.repo.PaymentMethod.Create(ctx, paymentMethod)
	if err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

// GetPaymentMethod retrieves a payment method by ID
func (u *PaymentMethodUsecase) GetPaymentMethod(ctx context.Context, id uint) (*models.PaymentMethod, error) {
	return u.repo.PaymentMethod.GetByID(ctx, id)
}

// GetPaymentMethodByName retrieves a payment method by name
func (u *PaymentMethodUsecase) GetPaymentMethodByName(ctx context.Context, name string) (*models.PaymentMethod, error) {
	return u.repo.PaymentMethod.GetByName(ctx, name)
}

// UpdatePaymentMethod updates a payment method
func (u *PaymentMethodUsecase) UpdatePaymentMethod(ctx context.Context, id uint, req interfaces.UpdatePaymentMethodRequest) (*models.PaymentMethod, error) {
	paymentMethod, err := u.repo.PaymentMethod.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		paymentMethod.Name = *req.Name
	}
	if req.Status != nil {
		paymentMethod.Status = *req.Status
	}

	err = u.repo.PaymentMethod.Update(ctx, paymentMethod)
	if err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

// DeletePaymentMethod deletes a payment method
func (u *PaymentMethodUsecase) DeletePaymentMethod(ctx context.Context, id uint) error {
	return u.repo.PaymentMethod.Delete(ctx, id)
}

// ListPaymentMethods lists payment methods with pagination
func (u *PaymentMethodUsecase) ListPaymentMethods(ctx context.Context, limit, offset int) ([]*models.PaymentMethod, error) {
	return u.repo.PaymentMethod.List(ctx, limit, offset)
}

// GetPaymentMethodsByStatus retrieves payment methods by status
func (u *PaymentMethodUsecase) GetPaymentMethodsByStatus(ctx context.Context, status models.StatusUmum) ([]*models.PaymentMethod, error) {
	return u.repo.PaymentMethod.GetByStatus(ctx, status)
}

// PaymentUsecase implements the payment usecase interface
type PaymentUsecase struct {
	repo *repository.RepositoryManager
}

// NewPaymentUsecase creates a new payment usecase
func NewPaymentUsecase(repo *repository.RepositoryManager) interfaces.PaymentUsecase {
	return &PaymentUsecase{repo: repo}
}

// CreatePayment creates a new payment
func (u *PaymentUsecase) CreatePayment(ctx context.Context, req interfaces.CreatePaymentRequest) (*models.Payment, error) {
	payment := &models.Payment{
		TransactionID: req.TransactionID,
		MethodID:      req.MethodID,
		Amount:        req.Amount,
		Status:        req.Status,
		PaymentDate:   req.PaymentDate,
		CreatedBy:     req.CreatedBy,
	}

	if payment.Status == "" {
		payment.Status = models.TransactionStatusSukses
	}

	err := u.repo.Payment.Create(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// GetPayment retrieves a payment by ID
func (u *PaymentUsecase) GetPayment(ctx context.Context, id uint) (*models.Payment, error) {
	return u.repo.Payment.GetByID(ctx, id)
}

// UpdatePayment updates a payment
func (u *PaymentUsecase) UpdatePayment(ctx context.Context, id uint, req interfaces.UpdatePaymentRequest) (*models.Payment, error) {
	payment, err := u.repo.Payment.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.TransactionID != nil {
		payment.TransactionID = *req.TransactionID
	}
	if req.MethodID != nil {
		payment.MethodID = *req.MethodID
	}
	if req.Amount != nil {
		payment.Amount = *req.Amount
	}
	if req.Status != nil {
		payment.Status = *req.Status
	}
	if req.PaymentDate != nil {
		payment.PaymentDate = req.PaymentDate
	}

	err = u.repo.Payment.Update(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// DeletePayment deletes a payment
func (u *PaymentUsecase) DeletePayment(ctx context.Context, id uint) error {
	return u.repo.Payment.Delete(ctx, id)
}

// ListPayments lists payments with pagination
func (u *PaymentUsecase) ListPayments(ctx context.Context, limit, offset int) ([]*models.Payment, error) {
	return u.repo.Payment.List(ctx, limit, offset)
}

// GetPaymentsByTransaction retrieves payments by transaction ID
func (u *PaymentUsecase) GetPaymentsByTransaction(ctx context.Context, transactionID uint) ([]*models.Payment, error) {
	return u.repo.Payment.GetByTransactionID(ctx, transactionID)
}

// GetPaymentsByMethod retrieves payments by payment method ID
func (u *PaymentUsecase) GetPaymentsByMethod(ctx context.Context, methodID uint) ([]*models.Payment, error) {
	return u.repo.Payment.GetByMethodID(ctx, methodID)
}

// GetPaymentsByStatus retrieves payments by status
func (u *PaymentUsecase) GetPaymentsByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Payment, error) {
	return u.repo.Payment.GetByStatus(ctx, status)
}

// GetPaymentsByDateRange retrieves payments by date range
func (u *PaymentUsecase) GetPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Payment, error) {
	return u.repo.Payment.GetByDateRange(ctx, startDate, endDate)
}

// CashFlowUsecase implements the cash flow usecase interface
type CashFlowUsecase struct {
	repo *repository.RepositoryManager
}

// NewCashFlowUsecase creates a new cash flow usecase
func NewCashFlowUsecase(repo *repository.RepositoryManager) interfaces.CashFlowUsecase {
	return &CashFlowUsecase{repo: repo}
}

// CreateCashFlow creates a new cash flow
func (u *CashFlowUsecase) CreateCashFlow(ctx context.Context, req interfaces.CreateCashFlowRequest) (*models.CashFlow, error) {
	cashFlow := &models.CashFlow{
		UserID:    req.UserID,
		Type:      req.FlowType,
		Source:    req.Description,
		Amount:    req.Amount,
		Date:      req.FlowDate,
		CreatedBy: req.CreatedBy,
	}

	err := u.repo.CashFlow.Create(ctx, cashFlow)
	if err != nil {
		return nil, err
	}

	return cashFlow, nil
}

// GetCashFlow retrieves a cash flow by ID
func (u *CashFlowUsecase) GetCashFlow(ctx context.Context, id uint) (*models.CashFlow, error) {
	return u.repo.CashFlow.GetByID(ctx, id)
}

// UpdateCashFlow updates a cash flow
func (u *CashFlowUsecase) UpdateCashFlow(ctx context.Context, id uint, req interfaces.UpdateCashFlowRequest) (*models.CashFlow, error) {
	cashFlow, err := u.repo.CashFlow.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.UserID != nil {
		cashFlow.UserID = *req.UserID
	}
	if req.FlowType != nil {
		cashFlow.Type = *req.FlowType
	}
	if req.Amount != nil {
		cashFlow.Amount = *req.Amount
	}
	if req.Description != nil {
		cashFlow.Source = *req.Description
	}
	if req.FlowDate != nil {
		cashFlow.Date = *req.FlowDate
	}

	err = u.repo.CashFlow.Update(ctx, cashFlow)
	if err != nil {
		return nil, err
	}

	return cashFlow, nil
}

// DeleteCashFlow deletes a cash flow
func (u *CashFlowUsecase) DeleteCashFlow(ctx context.Context, id uint) error {
	return u.repo.CashFlow.Delete(ctx, id)
}

// ListCashFlows lists cash flows with pagination
func (u *CashFlowUsecase) ListCashFlows(ctx context.Context, limit, offset int) ([]*models.CashFlow, error) {
	return u.repo.CashFlow.List(ctx, limit, offset)
}

// GetCashFlowsByUser retrieves cash flows by user ID
func (u *CashFlowUsecase) GetCashFlowsByUser(ctx context.Context, userID uint) ([]*models.CashFlow, error) {
	return u.repo.CashFlow.GetByUserID(ctx, userID)
}

// GetCashFlowsByOutlet retrieves cash flows by outlet ID
func (u *CashFlowUsecase) GetCashFlowsByOutlet(ctx context.Context, outletID uint) ([]*models.CashFlow, error) {
	// Since CashFlow model doesn't have OutletID, we'll return all cash flows
	// This could be enhanced to filter by user's outlet if needed
	return u.repo.CashFlow.List(ctx, 100, 0)
}

// GetCashFlowsByType retrieves cash flows by type
func (u *CashFlowUsecase) GetCashFlowsByType(ctx context.Context, flowType models.CashFlowType) ([]*models.CashFlow, error) {
	return u.repo.CashFlow.GetByType(ctx, flowType)
}

// GetCashFlowsByDateRange retrieves cash flows by date range
func (u *CashFlowUsecase) GetCashFlowsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.CashFlow, error) {
	return u.repo.CashFlow.GetByDateRange(ctx, startDate, endDate)
}

// GetTotalByTypeAndDateRange retrieves total cash flows by type and date range
func (u *CashFlowUsecase) GetTotalByTypeAndDateRange(ctx context.Context, flowType models.CashFlowType, startDate, endDate time.Time) (float64, error) {
	return u.repo.CashFlow.GetTotalByType(ctx, flowType, startDate, endDate)
}

// TransactionUsecase implements the transaction usecase interface
type TransactionUsecase struct {
	repo *repository.RepositoryManager
}

// NewTransactionUsecase creates a new transaction usecase
func NewTransactionUsecase(repo *repository.RepositoryManager) interfaces.TransactionUsecase {
	return &TransactionUsecase{repo: repo}
}

// CreateTransaction creates a new transaction
func (u *TransactionUsecase) CreateTransaction(ctx context.Context, req interfaces.CreateTransactionRequest) (*models.Transaction, error) {
	transaction := &models.Transaction{
		InvoiceNumber:   req.InvoiceNumber,
		TransactionDate: req.TransactionDate,
		UserID:          req.UserID,
		CustomerID:      req.CustomerID,
		OutletID:        req.OutletID,
		TransactionType: req.TransactionType,
		Status:          req.Status,
		CreatedBy:       req.CreatedBy,
	}

	if transaction.Status == "" {
		transaction.Status = models.TransactionStatusSukses
	}

	err := u.repo.Transaction.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// GetTransaction retrieves a transaction by ID
func (u *TransactionUsecase) GetTransaction(ctx context.Context, id uint) (*models.Transaction, error) {
	return u.repo.Transaction.GetByID(ctx, id)
}

// GetTransactionByInvoiceNumber retrieves a transaction by invoice number
func (u *TransactionUsecase) GetTransactionByInvoiceNumber(ctx context.Context, invoiceNumber string) (*models.Transaction, error) {
	return u.repo.Transaction.GetByInvoiceNumber(ctx, invoiceNumber)
}

// UpdateTransaction updates a transaction
func (u *TransactionUsecase) UpdateTransaction(ctx context.Context, id uint, req interfaces.UpdateTransactionRequest) (*models.Transaction, error) {
	transaction, err := u.repo.Transaction.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.InvoiceNumber != nil {
		transaction.InvoiceNumber = *req.InvoiceNumber
	}
	if req.TransactionDate != nil {
		transaction.TransactionDate = *req.TransactionDate
	}
	if req.UserID != nil {
		transaction.UserID = *req.UserID
	}
	if req.CustomerID != nil {
		transaction.CustomerID = req.CustomerID
	}
	if req.OutletID != nil {
		transaction.OutletID = *req.OutletID
	}
	if req.TransactionType != nil {
		transaction.TransactionType = *req.TransactionType
	}
	if req.Status != nil {
		transaction.Status = *req.Status
	}

	err = u.repo.Transaction.Update(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// DeleteTransaction deletes a transaction
func (u *TransactionUsecase) DeleteTransaction(ctx context.Context, id uint) error {
	return u.repo.Transaction.Delete(ctx, id)
}

// ListTransactions lists transactions with pagination
func (u *TransactionUsecase) ListTransactions(ctx context.Context, limit, offset int) ([]*models.Transaction, error) {
	return u.repo.Transaction.List(ctx, limit, offset)
}

// GetTransactionsByCustomer retrieves transactions by customer ID
func (u *TransactionUsecase) GetTransactionsByCustomer(ctx context.Context, customerID uint) ([]*models.Transaction, error) {
	return u.repo.Transaction.GetByCustomerID(ctx, customerID)
}

// GetTransactionsByUser retrieves transactions by user ID
func (u *TransactionUsecase) GetTransactionsByUser(ctx context.Context, userID uint) ([]*models.Transaction, error) {
	return u.repo.Transaction.GetByUserID(ctx, userID)
}

// GetTransactionsByOutlet retrieves transactions by outlet ID
func (u *TransactionUsecase) GetTransactionsByOutlet(ctx context.Context, outletID uint) ([]*models.Transaction, error) {
	return u.repo.Transaction.GetByOutletID(ctx, outletID)
}

// GetTransactionsByStatus retrieves transactions by status
func (u *TransactionUsecase) GetTransactionsByStatus(ctx context.Context, status models.TransactionStatus) ([]*models.Transaction, error) {
	return u.repo.Transaction.GetByStatus(ctx, status)
}

// GetTransactionsByDateRange retrieves transactions by date range
func (u *TransactionUsecase) GetTransactionsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*models.Transaction, error) {
	return u.repo.Transaction.GetByDateRange(ctx, startDate, endDate)
}

// TransactionDetailUsecase implements the transaction detail usecase interface
type TransactionDetailUsecase struct {
	repo *repository.RepositoryManager
}

// NewTransactionDetailUsecase creates a new transaction detail usecase
func NewTransactionDetailUsecase(repo *repository.RepositoryManager) interfaces.TransactionDetailUsecase {
	return &TransactionDetailUsecase{repo: repo}
}

// CreateTransactionDetail creates a new transaction detail
func (u *TransactionDetailUsecase) CreateTransactionDetail(ctx context.Context, req interfaces.CreateTransactionDetailRequest) (*models.TransactionDetail, error) {
	transactionDetail := &models.TransactionDetail{
		TransactionType: req.TransactionType,
		TransactionID:   req.TransactionID,
		ProductID:       req.ProductID,
		SerialNumberID:  req.SerialNumberID,
		Quantity:        req.Quantity,
		UnitPrice:       req.UnitPrice,
		TotalPrice:      req.TotalPrice,
		CreatedBy:       req.CreatedBy,
	}

	err := u.repo.TransactionDetail.Create(ctx, transactionDetail)
	if err != nil {
		return nil, err
	}

	return transactionDetail, nil
}

// GetTransactionDetail retrieves a transaction detail by ID
func (u *TransactionDetailUsecase) GetTransactionDetail(ctx context.Context, id uint) (*models.TransactionDetail, error) {
	return u.repo.TransactionDetail.GetByID(ctx, id)
}

// UpdateTransactionDetail updates a transaction detail
func (u *TransactionDetailUsecase) UpdateTransactionDetail(ctx context.Context, id uint, req interfaces.UpdateTransactionDetailRequest) (*models.TransactionDetail, error) {
	transactionDetail, err := u.repo.TransactionDetail.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.TransactionType != nil {
		transactionDetail.TransactionType = *req.TransactionType
	}
	if req.TransactionID != nil {
		transactionDetail.TransactionID = *req.TransactionID
	}
	if req.ProductID != nil {
		transactionDetail.ProductID = req.ProductID
	}
	if req.SerialNumberID != nil {
		transactionDetail.SerialNumberID = req.SerialNumberID
	}
	if req.Quantity != nil {
		transactionDetail.Quantity = *req.Quantity
	}
	if req.UnitPrice != nil {
		transactionDetail.UnitPrice = *req.UnitPrice
	}
	if req.TotalPrice != nil {
		transactionDetail.TotalPrice = *req.TotalPrice
	}

	err = u.repo.TransactionDetail.Update(ctx, transactionDetail)
	if err != nil {
		return nil, err
	}

	return transactionDetail, nil
}

// DeleteTransactionDetail deletes a transaction detail
func (u *TransactionDetailUsecase) DeleteTransactionDetail(ctx context.Context, id uint) error {
	return u.repo.TransactionDetail.Delete(ctx, id)
}

// ListTransactionDetails lists transaction details with pagination
func (u *TransactionDetailUsecase) ListTransactionDetails(ctx context.Context, limit, offset int) ([]*models.TransactionDetail, error) {
	return u.repo.TransactionDetail.List(ctx, limit, offset)
}

// GetTransactionDetailsByTransaction retrieves transaction details by transaction ID
func (u *TransactionDetailUsecase) GetTransactionDetailsByTransaction(ctx context.Context, transactionID uint) ([]*models.TransactionDetail, error) {
	return u.repo.TransactionDetail.GetByTransactionID(ctx, transactionID)
}

// GetTransactionDetailsByProduct retrieves transaction details by product ID
func (u *TransactionDetailUsecase) GetTransactionDetailsByProduct(ctx context.Context, productID uint) ([]*models.TransactionDetail, error) {
	return u.repo.TransactionDetail.GetByProductID(ctx, productID)
}

// DeleteTransactionDetailsByTransaction deletes transaction details by transaction ID
func (u *TransactionDetailUsecase) DeleteTransactionDetailsByTransaction(ctx context.Context, transactionID uint) error {
	return u.repo.TransactionDetail.DeleteByTransactionID(ctx, transactionID)
}