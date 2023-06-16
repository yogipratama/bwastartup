package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransactionFormatter := CampaignTransactionFormatter{}
	campaignTransactionFormatter.ID = transaction.ID
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt

	return campaignTransactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	campaignTransactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		campaignTransactionFormatter := FormatCampaignTransaction(transaction)
		campaignTransactionsFormatter = append(campaignTransactionsFormatter, campaignTransactionFormatter)
	}

	return campaignTransactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	var userTransactionFormatter UserTransactionFormatter

	userTransactionFormatter.ID = transaction.ID
	userTransactionFormatter.Amount = transaction.Amount
	userTransactionFormatter.Status = transaction.Status
	userTransactionFormatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	userTransactionFormatter.Campaign = campaignFormatter

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	userTransactionsFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		userTransactionFormatter := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransactionFormatter)
	}

	return userTransactionsFormatter
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	transactionFormatter := TransactionFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.CampaignID = transaction.CampaignID
	transactionFormatter.UserID = transaction.UserID
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.Status = transaction.Status
	transactionFormatter.Code = transaction.Code
	transactionFormatter.PaymentURL = transaction.PaymentURL

	return transactionFormatter
}
