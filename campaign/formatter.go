package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	ShortDescription string                         `json:"short_description"`
	Description      string                         `json:"description"`
	ImageURL         string                         `json:"image_url"`
	GoalAmount       int                            `json:"goal_amount"`
	CurrentAmount    int                            `json:"current_amount"`
	UserID           int                            `json:"user_id"`
	Slug             string                         `json:"slug"`
	Perks            []string                       `json:"perks"`
	BackerCount      int                            `json:"backer_count"`
	User             CampaignDetailUserFormatter    `json:"user"`
	Images           []CampaignDetailImageFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignDetailImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageURL = ""
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.BackerCount = campaign.BackerCount
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	splitPerk := strings.Split(campaign.Perks, ",")

	for _, perk := range splitPerk {
		trimSpacePerk := strings.TrimSpace(perk)
		perks = append(perks, trimSpacePerk)
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User

	campaignDetailUserFormatter := CampaignDetailUserFormatter{}
	campaignDetailUserFormatter.Name = user.Name
	campaignDetailUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = campaignDetailUserFormatter

	images := []CampaignDetailImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignDetailImageFormatter := CampaignDetailImageFormatter{}
		campaignDetailImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignDetailImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignDetailImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
