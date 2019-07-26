package outgoing

import "github.com/di-wu/mtga/thread"

const (
	AuthenticateMethod      thread.LogMethod = "Authenticate"
	CrackBoosterMethod      thread.LogMethod = "PlayerInventory.CrackBoostersV3"
	GetProductCatalogMethod thread.LogMethod = "PlayerInventory.GetProductCatalog"
	RedeemWildCardBulk      thread.LogMethod = "PlayerInventory.RedeemWildCardBulk"
	SetPetSelectionMethod   thread.LogMethod = "PlayerInventory.SetPetSelection"
	TrackDetailMethod       thread.LogMethod = "Quest.GetTrackDetail"

	CreateDeckMethod      thread.LogMethod = "Deck.CreateDeckV3"
	DeleteDeckMethod      thread.LogMethod = "Deck.DeleteDeck"
	UpdateDeckMethod      thread.LogMethod = "Deck.UpdateDeckV3"
	DraftStatusMethod     thread.LogMethod = "Draft.DraftStatus"
	MakePickMethod        thread.LogMethod = "Draft.MakePick"
	AIPracticeMethod      thread.LogMethod = "Event.AIPractice"
	ClaimPrizeMethod      thread.LogMethod = "Event.ClaimPrize"
	DeckSubmitMethod      thread.LogMethod = "Event.DeckSubmitV3"
	DraftMethod           thread.LogMethod = "Event.Draft"
	DropMethod            thread.LogMethod = "Event.Drop"
	GetPlayerCourseMethod thread.LogMethod = "Event.GetPlayerCourseV2"
	JoinMethod            thread.LogMethod = "Event.Join"
	JoinQueueMethod       thread.LogMethod = "Event.JoinQueue"
	PayEntryMethod        thread.LogMethod = "Event.PayEntry"

	PurchaseProductMethod thread.LogMethod = "Mercantile.PurchaseProduct"
)
