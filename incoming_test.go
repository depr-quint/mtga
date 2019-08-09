package mtga

import (
	"github.com/di-wu/mtga/thread/incoming/deck"
	"github.com/di-wu/mtga/thread/incoming/quest"
	"testing"

	"github.com/di-wu/mtga/thread/incoming/front_door"
)

func TestIncomingConnectionDetails(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]1/01/2000 0:00:00 AM`,
			`<== FrontDoor.ConnectionDetails(0)`,
			`{`,
			`	"sessionId": "00000000-0000-0000-0000-000000000000",`,
			`	"isQueued": "False"`,
			`}`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.Incoming.OnConnectionDetails(func(details front_door.ConnectionDetails) {
		callback = true
		if details.SessionId != "00000000-0000-0000-0000-000000000000" || details.IsQueued != "False" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestIncomingGetPreconDecks(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]1/01/2000 0:00:00 AM`,
			`<== Deck.GetPreconDecks(0)`,
			`[`,
			`	{`,
			`		"id": "00000000-0000-0000-0000-000000000000",`,
			`		"name": "?=?Loc/Decks/Precon/Precon_Dimir_Manipulation",`,
			`		"description": "UB Control",`,
			`		"format": "precon",`,
			`		"resourceId": "00000000-0000-0000-0000-000000000000",`,
			`		"deckTileId": null,`, // TODO: what if not null?
			`		"mainDeck": [`,
			`			{`,
			`				"id": "00000",`,
			`				"quantity": 4`,
			`			},`,
			`			{`,
			`				"id": "00001",`,
			`				"quantity": 1`,
			`			}`,
			`		],`,
			`		"sideboard": [`,
			`			{`,
			`				"id": "00001",`,
			`				"quantity": 3`,
			`			}`,
			`		],`,
			`		"lastUpdated": "2000-01-01T00:00:00.0000000Z",`,
			`		"lockedForUse": true,`,
			`		"lockedForEdit": true,`,
			`		"cardBack": null,`,
			`		"isValid": true`,
			`	}`,
			`]`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.Incoming.OnGetPreconDecks(func(decks []deck.PreconDeck) {
		callback = true
		if len(decks) != 1 {
			t.Error()
		}
		d := decks[0]
		if d.Description != "UB Control" || len(d.MainDeck) != 2 || len(d.Sideboard) != 1 || d.IsValid == false {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestIncomingGetPlayerQuests(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]1/01/2000 0:00:00 AM`,
			`<== Quest.GetPlayerQuests(0)`,
			`[`,
			`	{`,
			`		"questId": "00000000-0000-0000-0000-000000000000",`,
			`		"goal": 20,`,
			`		"locKey": "Quests/Quest_Simic_Manipulator",`,
			`		"tileResourceId": "00000000-0000-0000-0000-000000000000",`,
			`		"treasureResourceId": "00000000-0000-0000-0000-000000000000",`,
			`		"questTrack": "Default",`,
			`		"isNewQuest": true,`,
			`		"endingProgress": 0,`,
			`		"startingProgress": 0,`,
			`		"canSwap": true,`,
			`		"inventoryUpdate": null,`, // TODO: what if not null?
			`		"chestDescription": {`,
			`			"image1": "ObjectiveReward_XPCoinLarge",`,
			`			"image2": null,`,
			`			"image3": null,`,
			`			"prefab": "RewardPopup3DIcon_XPCoin",`,
			`			"referenceId": null,`, // TODO: what if not null?
			`			"headerLocKey": "MainNav/EventRewards/Gold_And_XP_Reward",`,
			`			"descriptionLocKey": null,`, // TODO: what if not null?
			`			"quantity": "500",`,
			`			"locParams": {`,
			`				"number1": 500,`,
			`				"number2": 500`,
			`			},`,
			`			"availableDate": "0001-01-01T00:00:00"`,
			`		},`,
			`		"hoursWaitAfterComplete": 0`,
			`	}`,
			`]`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.OnGetPlayerQuests(func(quests []quest.PlayerQuest) {
		callback = true
		if len(quests) != 1 {
			t.Error()
		}
		q := quests[0]
		if q.Goal != 20 || q.CanSwap != true || q.ChestDescription.Image2 != "" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
