package main

import (
	"encoding/base64"
	"machine"
	"time"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyfont/shnm"
)

type Mode int

const (
	modeProfile Mode = iota
	modeTwitter
	modeGithub
	modeWebSite
	modeTinyGoKeebook
	modeKLabTechBook
	modeKLabTechBookBN
	modeLast

	modeDefault = modeProfile
	modeA       = modeProfile
	modeB       = modeTinyGoKeebook
	modeC       = modeTwitter
)

var pages = map[Mode]func() error{
	modeProfile: BadgeProfile(
		"KLab Inc.",
		"MakKi",
		"Daisuke Makiuchi",
		"Twitter: @makki_d",
		"GitHub: makiuchi-d",
		profileImg,
	),
	modeTwitter: BadgeQR(
		"https://twitter.com/makki_d",
		Title{{&freesans.Bold9pt7b, "Twitter"}},
		Body{
			{22, 80, &freesans.Bold12pt7b, "@makki_d"},
		}),
	modeGithub: BadgeQR(
		"https://github.com/makiuchi-d",
		Title{{&freesans.Bold9pt7b, "GitHub"}},
		Body{
			{19, 80, &freesans.Bold12pt7b, "makiuchi-d"},
		}),
	modeWebSite: BadgeQR(
		"https://makiuchi-d.github.io/",
		Title{{&freesans.Bold9pt7b, "Web site"}},
		Body{
			{4, 80, &freesans.Regular9pt7b, "makiuchi-d.github.io"},
		}),
	modeTinyGoKeebook: BadgeQR(
		"https://techbookfest.org/organization/9htjwrJXfaWWnw8Cbg8JsW",
		Title{
			{&shnm.Shnmk12, "あ17"},
			{&freesans.Bold9pt7b, "TinyGo Keeb"},
		},
		Body{
			{8, 50 + 20, &freesans.Bold9pt7b, "TinyGo Keebook"},
			{30, 80 + 20, &freesans.Bold12pt7b, "vol. 1 & 2"},
			//{12, 126, &gophers.Regular58pt, "G    V"},
		}),
	modeKLabTechBook: BadgeQR(
		"https://techbookfest.org/product/aVbpmUVUwehW3Mym5rYrzN",
		Title{
			{&shnm.Shnmk12, "ケ35"},
			{&freesans.Bold9pt7b, "KLab"},
			{&shnm.Shnmk12, "株式会社"},
		},
		Body{
			{4, 58, &freesans.Bold9pt7b, "KLabTechBook 16"},
			{6, 80, &shnm.Shnmk12, "Pythonの抽象基底クラス"},
			{6, 97, &shnm.Shnmk12, "SQLによるDBマイグレ-ション"},
			{6, 114, &shnm.Shnmk12, "現代的GAS開発環境"},
		}),
	modeKLabTechBookBN: BadgeQR(
		"https://www.klab.com/jp/blog/tech/2025/tbf19.html",
		Title{{&freesans.Bold9pt7b, "KLabTechBook"}},
		Body{
			{25, 58, &shnm.Shnmk12, "新刊紹介 & Free PDF"},
			{8, 80, &proggy.TinySZ8pt7b, "https://www.klab.com/jp/"},
			{8, 92, &proggy.TinySZ8pt7b, " blog/tech/2025/tbf19.html"},
		}),
}

var profileImg, _ = base64.StdEncoding.DecodeString(`
///////9ALaqqqAFbabf///////u1ghb/bbQF7rn9bf////tv/6ANqqqpIXfQr/+///9v/t3AKu21WAGqsfbb///7//v3aIdVSqwEvdC7/v/+7/2vfbQKqqqkkVZo3Vf/3//v/dfYAtVVKACroKv9vv//fu99akqqCJUBVKD2r////fu717oCqWVIJFVRW/V//d/v7v1VCRQQIgFVQK1f////vXuu1gKhRRAASVC26r3/+/fe9VVQCBAFSASiq1//3b7d616oAqIH///wCKqqv///933rVUBQBX//lECrbf//7b2uqoI////QAv+gSra773/v97VUL/9/arV/gKqq/737eqjVSf6r+gAAAP4VV///39/qqqH9/bwAAAf4SVr7/3rqsFUH1pbX4CQAD0Snf7XvvdoQHqtra39AoAHxHX7/vdaoKp99ur31AApBUJ17/u77VoB7StdVX/oACD56v9u7rdQKXbdt7qvoJJAseX9/7vZVQHratTX1P4ACR+K3+rtbqQF9t9vWrt+CQAPhf6/t7KSQ6tptf/Vq8BKR4j39tqtSBPdtt+qvtX8ABHi/73u1IIDqtq79tNu8JCB8N3utVUog7dr71tttV4EIHz/u9togCPatf9Vtq6+EIh09u1tVRIDrd8Vuq3V14ICH7+/taQAk7avntd2u1qAoB976tqRSAPbenVq2t3t5AkPLr1VSAJB6vB2t1/uv3gAQz/raqJAA+3xe9rXu/fpSQsLfaoIEiH/1uqtfAAJfAACL9aqoQAF/9V29v9AB/gkkRr7UghJI4NH9VXQKqAcAAAPralBAAGbx3ev8IAIXJKoDfakJBIDBw+8+wIlQgcAAq660oBAjgU/79xIiCijyRCL10gRBAomP+/6IiKECoBA/XlShBBOHL/gQIiIUqDSCv+upCBBDBy+40IiRQgKwID+1VECBAw4+aQoiJBCoLQk/7VIECAsWPngAiIlKEiwgP/tJUECPHv5pOAIiAUSMBL/9UAIICRj9YFLoiKggSkA//VSIQQ4a/OUwUiIAiooSf/9IAAQN+f3gRaCIK1AigD//UkSQT/n55RFKIgKFAok//8ASAA9p9aBFIUv/UKMAP//VCEkvafWlEJQF/woTin//8EAAH2PzgERCo/tBQ4A///EkkqXt/5USkAfTFBGiP9/wkAA35/6ASAqn00KJgL//9CQAP7f+FRKgF88QRcQ///oRAC7v/kBACpfPCiGgv//+iCv9h/wVFUBfjUFJ4j///gUFdqf+QIAqn88oEaA//3+IH9WX/BRVAT+MBUXhP///4Iavl/VCgKo/vKAhoD+//+QXW5f0KCoAP6wVSeJ////hB/0f8QVBVP68gBFgP//98Af/F/BQFAi/PCqF4j////CnqVfCCqKg/vCAUeC//3/wB/+fwUAICP90KolgP////Bbf5wgqoqL/8QAhwj0l//8f/fgugRAD/8BVCeC+3//9H//VCpRKov/SAKGiAAB//z3f+C6CIBP/wVQVwDAAv/8//ektKQqD/8oCgaJEt///3/f0EISgS/+DlFOAABf/f3t/6UooFQP/i4IDiQP////t/2gBEoBb/8Coq4Bt////+W2CKIBVD/5/0g6JC7///+n/UJIqAEv+/6gvAD/////4AQT8gVUv+C+SvhJ/////+FxQ9CoAi/o/SGpAP//7/+gBBf0QqivoAEHeCT/+//f5RFH8RBCP6UBRfiB//////CEHvJKkU8ASh8wEP+v///4URrQgBf1Kil48IT/v///+Qhd0iqXvQEE7uIA/6v/+hhCO0SCPv/UKbPQUv+v//9aKJvBKL//QSnuwgL+of3Q/oJfKARf7+jFV8CD/6r38F4oPYKiu7+gsbtIK/6pP9P/Ao0oEP/9/8LXAg/+rR//+chQAoq///t/fBCf/tJf//HCKqhQ9u/tyq0AH/ttR//o8FwChD/++9vYRF/+qoX/4NI5KCC/+/ndfBBf+qqhf4Q8gEUUvt/q6ugB//3VUf+gNAUQgjv/udtxSP/6tUg8Ag5QRFCf/vj3oAf9+1VEDxAvgpEKBfvEuoQPv/3aqS8En6AEoK/f4n8BD//2qqSDwC/4qBUAdRAsED7/+6EQAqIP6AKASV+FPEQ///oEAlD4T/9IVSAgKPEC///0AAAAaA/9AgCKgICkAv//8CBAJR4L//BVICoXwBf///oACQADw///oASA2gH////0CAARBUP//qAAAj4A////++VVRCDx////6qvAX////+60gBCB0f///VQVaK///3/723akED53////////v///9u2SgQQsf/////+///////+22VRAA5j//////////p/V+2ohEkF3L/////////p//1WqlAABbz/////////6P/f+1UlJFBvP//////6ooF/9t2qkAAAXz///////9fA/7+21UqiggvP/1XdqgANCX/q7WqgCCAf5/+/f/3QDQB//7e1KqCIQPr9AAAAAHAhf9v9WqgKAhA/gAAAAAWhCH/+r+qqoEiAH/QkkkqVwEJ/7/q9qRUAJIHaAAAALhQIP/tv1qSAUoAHdQEpKAqBID//+qAAKgAiID/6AALwAAIP+19UERFUgABvUKVAMEqQD//wAIAIABJEAu4AEgQAAv/+gAABVVVAEAD/UgCBVf/3/4SACCQAEkFIe1CCEAq/3/6QEiKqqkAIAP//+EX/4CP/AAACqREVIiT///gBde+0=`)

var (
	display                          uc8151.Device
	btnA, btnB, btnC, btnUp, btnDown machine.Pin
)

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	btnA = machine.BUTTON_A
	btnB = machine.BUTTON_B
	btnC = machine.BUTTON_C
	btnUp = machine.BUTTON_UP
	btnDown = machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	mode := modeDefault

	for {
		display.ClearBuffer()
		display.Display()
		display.WaitUntilIdle()

		err := pages[mode]()
		if err != nil {
			println(err.Error())
			display.ClearBuffer()
			tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 11, err.Error(), black)
		}

		display.Display()
		display.WaitUntilIdle()

		mode = selectMode(mode)
	}
}

func selectMode(m Mode) Mode {
	for {
		switch {
		case btnA.Get():
			return modeA
		case btnB.Get():
			return modeB
		case btnC.Get():
			return modeC
		case btnDown.Get():
			m++
			if m == modeLast {
				m = 0
			}
			return m
		case btnUp.Get():
			m--
			if m < 0 {
				m = modeLast - 1
			}
			return m
		}

		time.Sleep(200 * time.Millisecond)
	}
}
