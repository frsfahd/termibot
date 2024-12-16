package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/dotenv-org/godotenvvault/autoload"
	"github.com/frsfahd/termiBot/internal/chat"
)

type Request struct {
	Messages []chat.Message `json:"messages"`
}

type Response struct {
	Result struct {
		Response string `json:"response"`
	} `json:"result"`
	Success  bool     `json:"success"`
	Errors   []string `json:"errors"`
	Messages []string `json:"messages"`
}

type msgByte []byte

type errMsg struct{ err error }

var (
	baseUrl = os.Getenv("BASE_URL")
	token   = os.Getenv("CF_TOKEN")
	accId   = os.Getenv("CF_ACC_ID")
)

func sendMsg(MsgHistory []chat.Message, endpoint string) tea.Cmd {
	return func() tea.Msg {
		reqUrl, _ := url.Parse(baseUrl)
		reqUrl = reqUrl.JoinPath(accId).JoinPath("/ai/run/").JoinPath(endpoint)

		payload := Request{
			Messages: MsgHistory,
		}
		var buf bytes.Buffer

		json.NewEncoder(&buf).Encode(payload)

		req, _ := http.NewRequest("POST", reqUrl.String(), &buf)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			return errMsg{err: err}
		}

		data, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return errMsg{err: err}
		}
		return msgByte(data)
	}

}
