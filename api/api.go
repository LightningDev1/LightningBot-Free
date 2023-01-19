package api

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/LightningDev1/LightningBot-Free/constants"
	"github.com/LightningDev1/LightningBot-Free/http"
	"github.com/LightningDev1/discordgo"
	"github.com/bitly/go-simplejson"
	"github.com/go-ping/ping"
)

type TranslateResult struct {
	Original            string
	Result              string
	SourceLanguage      string
	DestinationLanguage string
	Confidence          float64
	Error               error
}

func TTS(text, language string) ([]byte, error) {
	ttsUrl := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s", url.QueryEscape(text), url.QueryEscape(language))

	httpResponse := http.Get(ttsUrl)
	if httpResponse.Error != nil {
		return nil, httpResponse.Error
	}

	return httpResponse.BodyBytes, nil
}

func Translate(text, language string) *TranslateResult {
	language = strings.ToLower(language)
	language = strings.Split(language, "_")[0]

	if _, ok := constants.LANGUAGES[language]; !ok {
		if _, ok := constants.LANGUAGE_SPECIAL_CASES[language]; ok {
			language = constants.LANGUAGE_SPECIAL_CASES[language]
		} else if _, ok := constants.LANGUAGE_CODES[language]; ok {
			language = constants.LANGUAGE_CODES[language]
		} else {
			return &TranslateResult{
				Error: fmt.Errorf("invalid destination language"),
			}
		}
	}

	translateUrl := fmt.Sprintf("https://translate.google.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s", url.QueryEscape(language), url.QueryEscape(text))

	httpResponse := http.Get(translateUrl)
	if httpResponse.Error != nil {
		return &TranslateResult{
			Error: httpResponse.Error,
		}
	}

	jsonResult, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		return &TranslateResult{
			Error: err,
		}
	}

	translateData := jsonResult.GetIndex(0).GetIndex(0)
	return &TranslateResult{
		Original:            translateData.GetIndex(1).MustString(),
		Result:              translateData.GetIndex(0).MustString(),
		SourceLanguage:      jsonResult.GetIndex(2).MustString(),
		DestinationLanguage: language,
		Confidence:          jsonResult.GetIndex(6).MustFloat64(),
	}
}

func ChangeSettings(session *discordgo.Session, settings map[string]any, settingsPath bool) error {
	settingsUrl := ""
	if settingsPath {
		settingsUrl = "https://discord.com/api/v10/users/@me/settings"
	} else {
		settingsUrl = "https://discord.com/api/v10/users/@me"
	}

	_, err := session.Request("PATCH", settingsUrl, settings)

	return err
}

func Ping(host string) (time.Duration, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return time.Duration(0), err
	}
	pinger.Count = 3
	pinger.SetPrivileged(true)
	err = pinger.Run()
	if err != nil {
		return time.Duration(0), err
	}
	stats := pinger.Statistics()
	return stats.AvgRtt, nil
}

func GetPublicIP() string {
	httpResponse := http.Get("https://ipinfo.io/ip")
	if httpResponse.Error != nil {
		return "Error"
	}
	return httpResponse.Body
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "Error"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func GetMacAddress() string {
	currentIP := GetLocalIP()
	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {
		addrs, err := interf.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if strings.Contains(addr.String(), currentIP) {
				netInterface, err := net.InterfaceByName(interf.Name)
				if err != nil {
					continue
				}
				return strings.ToUpper(netInterface.HardwareAddr.String())
			}
		}
	}
	return "Error"
}

func CreatePaste(text string) (string, error) {
	httpResponse := http.Post("https://paste.lightning-bot.com/documents", text)
	if httpResponse.Error != nil {
		return "", httpResponse.Error
	}

	paste, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		return "", err
	}

	return "https://paste.lightning-bot.com/" + paste.Get("key").MustString(), nil
}
