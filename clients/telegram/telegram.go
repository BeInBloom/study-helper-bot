package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	errs "github.com/BeInBloom/study-helper-bot/lib/errors"
)

type client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) client {
	return client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

//============================================================

func (c *client) GetUpdates(offset int, limit int) ([]Update, error) {
	const errMsg = "cant get updates"

	q := url.Values{}
	q.Add("offset", strconv.Itoa(limit))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, errs.Wrap(errMsg, err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errs.Wrap(errMsg, err)
	}

	if !res.Ok {
		return nil, errs.Wrap(errMsg, err)
	}

	return res.Result, nil
}

func (c *client) SendMessage(chatID int, text string) error {
	const errMsg = "cant send message"

	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		errs.Wrap(errMsg, err)
	}

	return nil
}

//============================================================

func (c *client) doRequest(methor string, query url.Values) (data []byte, err error) {
	const errMsg = "cant do request"

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, methor),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errs.Wrap(errMsg, err)
	}

	req.URL.RawQuery = query.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, errs.Wrap(errMsg, err)
	}

	defer func() { res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errs.Wrap(errMsg, err)
	}

	return body, nil
}

// ============================================================
func newBasePath(token string) string {
	return "bot" + token
}
