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
