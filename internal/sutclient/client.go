package sutclient

import (
	"context"
	"net/http/cookiejar"
	"sync"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/publicsuffix"

	"lk_sut/internal/config"
)

type SutClient struct {
	client *resty.Client
	mutex  sync.Mutex
}

func NewClient(cfg *config.Config) *SutClient {
	restyClient := resty.New().
		SetBaseURL(cfg.LkSutService.URL).
		SetTimeout(cfg.LkSutService.Timeout).
		OnAfterResponse(afterResponse)

	return &SutClient{
		client: restyClient,
	}
}

func (sc *SutClient) resetCookie() *SutClient {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	sc.client.GetClient().Jar = cookieJar

	return sc
}

func (sc *SutClient) r(ctx context.Context) *resty.Request {
	return sc.client.R().SetContext(ctx)
}
