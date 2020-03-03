package splitwise

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/oauth1"
)

// URLs
const (
	AccessTokenURL        = "https://secure.splitwise.com/oauth/access_token"
	RequestTokenURl       = "https://secure.splitwise.com/oauth/request_token"
	AuthorizationTokenURL = "https://secure.splitwise.com/oauth/authorize"

	GetCurrentUserURL = "https://secure.splitwise.com/api/v3.0/get_current_user"
)

type Opts struct {
	ConsumerKey    string
	ConsumerSecret string
}

// Client is the client for Splitwise API
type Client struct {
	config oauth1.Config
	token  *oauth1.Token
}

// NewClient returns returns a new splitwise client
func NewClient(opts Opts) (*Client, error) {
	config := oauth1.Config{
		ConsumerKey:    opts.ConsumerKey,
		ConsumerSecret: opts.ConsumerSecret,
		Endpoint: oauth1.Endpoint{
			AccessTokenURL:  AccessTokenURL,
			RequestTokenURL: RequestTokenURl,
			AuthorizeURL:    AuthorizationTokenURL,
		},
	}

	client := Client{
		config: config,
	}

	return &client, nil
}

// GetAndSetAccessToken gets access token, access token secret
func (c *Client) GetAccessToken() (string, string, error) {
	requestToken, requestSecret, err := c.config.RequestToken()
	if err != nil {
		return "", "", err
	}
	authorizationURL, err := c.config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", err
	}
	fmt.Println("Please sign in at: ", authorizationURL.String())
	fmt.Printf("Choose whether to grant the application access.\nPaste " +
		"the oauth_verifier parameter from the " +
		"address bar: ")

	var verifier string
	fmt.Scanf("%s", &verifier)
	accessToken, accessSecret, err := c.config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		return "", "", err
	}

	token := oauth1.NewToken(accessToken, accessSecret)
	return token.Token, token.TokenSecret, nil
}

func (c *Client) SetAccessTokens(token, tokenSecret string) {
	c.token = &oauth1.Token{
		Token:       token,
		TokenSecret: tokenSecret,
	}
}

// GetCurrentUser returns current user who authorized Oauth
func (c *Client) GetCurrentUser() (*User, error) {
	type userReponse struct {
		User User `json:"user"`
	}
	ctx := context.Background()
	httpClient := c.config.Client(ctx, c.token)

	resp, err := httpClient.Get(GetCurrentUserURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := userReponse{}
	b, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	return &response.User, nil
}
