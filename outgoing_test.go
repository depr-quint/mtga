package mtga

import (
	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/outgoing/log"
	"github.com/di-wu/mtga/thread/outgoing/log/client"
	"testing"

	"github.com/di-wu/mtga/thread/outgoing"
)

func TestOutgoingAuth(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]1/01/2000 0:00:00 AM`,
		`==> Authenticate(0):`,
		`{`,
		`	"jsonrpc": "2.0",`,
		`	"method": "Authenticate",`,
		`	"params": {`,
		`		"ticket": "rAnDOmChAraCt3rS",`,
		`		"clientVersion": "1595.718832"`,
		`	},`,
		`	"id": "0"`,
		`}`,
	}
	var callback1, callback2 bool
	parser := Parser{}
	parser.OnTreadLog(func(log thread.Log) {
		callback1 = true
	})
	parser.Outgoing.OnAuthenticate(func(auth outgoing.Authenticate) {
		callback2 = true
		if auth.Ticket != "rAnDOmChAraCt3rS" || auth.ClientVersion != "1595.718832" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback1 || !callback2 {
		t.Error()
	}
}

func TestOutgoingInfoUserDeviceSpecs(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]1/01/2000 0:00:00 AM`,
		`==> Log.Info(0):`,
		`{`,
		`	"jsonrpc": "2.0",`,
		`	"method": "Log.Info",`,
		`	"params": {`,
		`		"messageName": "Client.UserDeviceSpecs",`,
		`		"humanContext": "User Device Specs",`,
		`		"payloadObject": {`,
		`			"graphicsDeviceName": "MGTA Ti",`,
		`			"graphicsDeviceType": "Direct3D11",`,
		`			"graphicsDeviceVendor": "MGTA",`,
		`			"graphicsDeviceVersion": "Direct3D 11.0 [level 11.1]",`,
		`			"graphicsMemorySize": 3072,`,
		`			"graphicsMultiThreaded": true,`,
		`			"graphicsShaderLevel": 50,`,
		`			"deviceUniqueIdentifier": "0000000000000000000000000000000000000000",`,
		`			"deviceModel": "MGTA-0",`,
		`			"deviceType": "Desktop",`,
		`			"operatingSystem": "Windows 10 (10.0.0) 64bit",`,
		`			"operatingSystemFamily": "Windows",`,
		`			"processorCount": 16,`,
		`			"processorFrequency": 3700,`,
		`			"processorType": "MGTA Tab",`,
		`			"systemMemorySize": 16336,`,
		`			"maxTextureSize": 16384,`,
		`			"isWindowed": false,`,
		`			"gameResolution": {`,
		`				"width": 1920,`,
		`				"height": 1080`,
		`			},`,
		`			"monitorResolution": {`,
		`				"width": 1920,`,
		`				"height": 1080`,
		`			},`,
		`			"monitorSupportedResolutions": [`,
		`				{`,
		`					"width": 1280,`,
		`					"height": 720,`,
		`					"validForWindow": true,`,
		`					"validForFullscreen": true`,
		`				},`,
		`				{`,
		`					"width": 1920,`,
		`					"height": 1080,`,
		`					"validForWindow": false,`,
		`					"validForFullscreen": true`,
		`				}`,
		`			],`,
		`			"playerId": "00000000000000000000000000"`,
		`		},`,
		`		"transactionId": "00000000-0000-0000-0000-000000000000"`,
		`	},`,
		`	"id": "0"`,
		`}`,
	}
	var callback1, callback2 bool
	parser := Parser{}
	parser.Outgoing.OnLogInfo(func(info log.Info) {
		callback1 = true
		if info.MessageName != "Client.UserDeviceSpecs" ||
			info.HumanContext != "User Device Specs" {
			t.Error()
		}
	})
	parser.Outgoing.OnUserDeviceSpecs(func(specs client.UserDeviceSpecs) {
		callback2 = true
		if specs.GraphicsDeviceName != "MGTA Ti" || specs.GraphicsMemorySize != 3072 || specs.DeviceType != "Desktop" ||
			specs.GameResolution.Width != specs.MonitorResolution.Width || len(specs.MonitorSupportedResolutions) != 2 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback1 || !callback2 {
		t.Error()
	}
}

func TestOutgoingConnected(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]1/01/2000 0:00:00 AM`,
		`==> Log.Info(0):`,
		`{`,
		`	"jsonrpc": "2.0",`,
		`	"method": "Log.Info",`,
		`	"params": {`,
		`		"messageName": "Client.Connected",`,
		`		"humanContext": "Client connected to FrontDoor at: ProdB - client.arenagame-b.east.magic-the-gathering-arena.com:9405",`,
		`		"payloadObject": {`,
		`			"playerId": "00000000000000000000000000",`,
		`			"screenName": "Diwu#93074",`,
		`			"clientVersion": "1595.718832",`,
		`			"settings": {`,
		`				"gameplay": {`,
		`					"disableEmotes": false,`,
		`					"evergreenKeywordReminders": true,`,
		`					"autoTap": true,`,
		`					"autoOrderTriggeredAbilities": true,`,
		`					"autoChooseReplacementEffects": true,`,
		`					"showPhaseLadder": true,`,
		`					"allPlayModesToggle": true`,
		`				},`,
		`				"audio": {`,
		`					"master": 10.0,`,
		`					"music": 100.0,`,
		`					"effects": 100.0,`,
		`					"voice": 100.0,`,
		`					"ambience": 100.0,`,
		`					"playInBackground": true`,
		`				},`,
		`				"language": {`,
		`					"language": "English"`,
		`				}`,
		`			}`,
		`		},`,
		`		"transactionId": "00000000-0000-0000-0000-000000000000"`,
		`	},`,
		`	"id": "0"`,
		`}`,
	}
	var callback bool
	parser := Parser{}
	parser.OnConnected(func(conn client.Connected) {
		callback = true
		if conn.ScreenName != "Diwu#93074" || conn.Settings.Audio.Master != 10.0 ||
			conn.Settings.Gameplay.AutoTap == false || conn.Settings.Language.Language != "English" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
