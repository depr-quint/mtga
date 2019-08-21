package mtga

import (
	"github.com/di-wu/mtga/thread/minus_one"
	"testing"
)

func TestMinusOneEventMatchCreated(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]8/11/2019 9:28:18 AM`,
		`(-1) Incoming Event.MatchCreated {`,
		`	"controllerFabricUri": "http://10.50.4.105:55885/match/v1/3fb1fc0e-a01e-4e44-9b6c-9026edb6a719",`,
		`	"matchEndpointHost": "client.arenamatch-a.east.magic-the-gathering-arena.com",`,
		`	"matchEndpointPort": 9505,`,
		`	"opponentScreenName": "Diwu#93074",`,
		`	"opponentIsWotc": false,`,
		`	"matchId": "3fb1fc0e-a01e-4e44-9b6c-9026edb6a719",`,
		`	"opponentRankingClass": "Gold",`,
		`	"opponentRankingTier": 4,`,
		`	"opponentMythicPercentile": 0.0,`,
		`	"opponentMythicLeaderboardPlace": 0,`,
		`	"eventId": "Play",`,
		`	"opponentAvatarSelection": "Avatar_Basic_Vannifar",`,
		`	"opponentCardBackSelection": "CardBack_LocketSimic",`,
		`	"opponentPetSelection": "M20_BattlePass",`,
		`	"opponentPetModSelections": [ "Level.2" ],`,
		`	"avatarSelection": "Avatar_Basic_Lazav",`,
		`	"cardbackSelection": "CardBack_M20_Chandra",`,
		`	"petModSelections": [],`,
		`	"battlefield": "GRN"`,
		`}`,
	}
	var callback bool
	parser := Parser{}
	parser.OnEventMatchCreated(func(match minus_one.MatchCreated) {
		callback = true
		if match.OpponentRankingClass != "Gold" || len(match.OpponentPetModSelections) != 1 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestMinusOneTrackProgressUpdated(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]8/11/2019 9:33:55 AM`,
		`(-1) Incoming TrackProgress.Updated [`,
		`	{`,
		`		"trackName": "EarlyPlayerProgression",`,
		`		"trackTier": 0`,
		`	}`,
		`]`,
	}
	var callback bool
	parser := Parser{}
	parser.OnTrackProgressUpdated(func(update []minus_one.TrackProgress) {
		callback = true
		if len(update) != 1 || update[0].TrackName != "EarlyPlayerProgression" || update[0].TrackTier != 0 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestMinusOneInventoryUpdated(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]8/11/2019 9:33:55 AM`,
		`(-1) Incoming Inventory.Updated {`,
		`	"delta": {`,
		`		"gemsDelta": 0,`,
		`		"boosterDelta": [],`,
		`		"cardsAdded": [],`,
		`		"decksAdded": [],`,
		`		"vanityItemsAdded": [],`,
		`		"vanityItemsRemoved": [],`,
		`		"draftTokensDelta": 0,`,
		`		"goldDelta": 0,`,
		`		"sealedTokensDelta": 0,`,
		`		"vaultProgressDelta": 0.3,`,
		`		"wcCommonDelta": 0,`,
		`		"wcMythicDelta": 0,`,
		`		"wcRareDelta": 0,`,
		`		"wcUncommonDelta": 0,`,
		`		"artSkinsAdded": [],`,
		`		"voucherItemsDelta": []`,
		`	},`,
		`	"aetherizedCards": [`,
		`		{`,
		`			"grpId": 69208,`,
		`			"goldAwarded": 0,`,
		`			"gemsAwarded": 0,`,
		`			"set": ""`,
		`		}`,
		`	],`,
		`	"context": "PlayerReward.OnMatchCompletedDaily",`,
		`	"xpGained": 0`,
		`}`,
	}
	var callback bool
	parser := Parser{}
	parser.OnInventoryUpdated(func(update minus_one.InventoryUpdate) {
		callback = true
		if update.Delta.VaultProgressDelta != 0.3 || update.AetherizedCards[0].GrpId != 69208 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestMinusOneRankUpdated(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]8/10/2019 1:06:17 PM`,
		`(-1) Incoming Rank.Updated {`,
		`	"playerId": "00000000000000000000000000",`,
		`	"seasonOrdinal": 8,`,
		`	"newClass": "Bronze",`,
		`	"oldClass": "Bronze",`,
		`	"newLevel": 1,`,
		`	"oldLevel": 2,`,
		`	"oldStep": 2,`,
		`	"newStep": 0,`,
		`	"wasLossProtected": false,`,
		`	"rankUpdateType": "Limited"`,
		`}`,
	}
	var callback bool
	parser := Parser{}
	parser.OnRankUpdated(func(update minus_one.RankUpdate) {
		callback = true
		if update.SeasonOrdinal != 8 || update.RankUpdateType != "Limited" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
