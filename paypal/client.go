package paypal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Adam0120/gopay"
	"github.com/Adam0120/gopay/pkg/util"
	"github.com/Adam0120/gopay/pkg/xhttp"
	"github.com/Adam0120/gopay/pkg/xlog"
)

// Client PayPal支付客
type Client struct {
	Clientid    string
	Secret      string
	Appid       string
	AccessToken string
	ExpiresIn   int
	IsProd      bool
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
}

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool) (client *Client, err error) {
	if clientid == util.NULL || secret == util.NULL {
		return nil, gopay.MissPayPalInitParamErr
	}
	client = &Client{
		Clientid:    clientid,
		Secret:      secret,
		IsProd:      isProd,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) doPayPalGet(ctx context.Context, uri string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + uri
	if !c.IsProd {
		url = baseUrlSandbox + uri
	}
	httpClient := xhttp.NewClient()
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Url: %s", url)
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPost(ctx context.Context, bm gopay.BodyMap, path string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + path
	if !c.IsProd {
		url = baseUrlSandbox + path
	}
	httpClient := xhttp.NewClient()
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPatch(ctx context.Context, patchs []*Patch, path string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + path
	if !c.IsProd {
		url = baseUrlSandbox + path
	}
	httpClient := xhttp.NewClient()
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		jb, _ := json.Marshal(patchs)
		xlog.Debugf("PayPal_RequestBody: %s", string(jb))
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Patch(url).SendStruct(patchs).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}
