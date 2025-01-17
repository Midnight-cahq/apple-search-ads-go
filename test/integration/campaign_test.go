//go:build integration
// +build integration

/**
Copyright (C) 2021 Mehmet Gungoren.
This file is part of apple-search-ads-go, a package for working with Apple's
Search Ads API.
apple-search-ads-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
apple-search-ads-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with apple-search-ads-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListCampaigns(t *testing.T) {
	NewClient()
	campaignListResponse, _, err := client.Campaigns.GetAllCampaigns(context.Background(), nil)
	assert.NoError(t, err, "ListCampaigns responded with an error")
	assert.NotEmpty(t, campaignListResponse.Campaigns, "ListCampaigns returned no campaignListResponse")

	firstCampaign := campaignListResponse.Campaigns[0]
	campaignResponse, _, err := client.Campaigns.GetCampaign(context.Background(), firstCampaign.ID)
	assert.NoError(t, err, "GetCampaign responded with an error")
	assert.NotNil(t, campaignResponse.Campaign.Name)
}
