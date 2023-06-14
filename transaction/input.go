package transaction

import "bwastartup/user"

type GetTransactionsCampaignInput struct {
	CampaignID int `uri:"id" binding:"required"`
	User       user.User
}
