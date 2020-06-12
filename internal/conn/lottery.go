package conn

import (
	"context"
	"fmt"

	"github.com/dghubble/sling"

	"github.com/donech/tool/xlog"
)

const twoToneSphereHost = "http://www.cwl.gov.cn/cwl_admin/kjxx/findKjxx/forIssue?name=ssq&code=%s"

type SSQResponse struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Result  []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		Date        string `json:"date"`
		Red         string `json:"red"`
		Blue        string `json:"blue"`
		Blue2       string `json:"blue2"`
		Sales       string `json:"sales"`
		Poolmoney   string `json:"poolmoney"`
		Content     string `json:"content"`
		Addmoney    string `json:"addmoney"`
		Addmoney2   string `json:"addmoney2"`
		DetailsLink string `json:"detailsLink"`
		VideoLink   string `json:"videoLink"`
		Msg         string `json:"msg"`
		Z2Add       string `json:"z2add"`
		M2Add       string `json:"m2add"`
		Prizegrades []struct {
			Type      int    `json:"type"`
			Typenum   string `json:"typenum"`
			Typemoney string `json:"typemoney"`
		} `json:"prizegrades"`
	} `json:"result"`
}

type LotteryClient struct {
}

func NewLotteryClient() *LotteryClient {
	return &LotteryClient{}
}

func (c *LotteryClient) GetTwoToneSphere(ctx context.Context, period string) SSQResponse {
	url := fmt.Sprintf(twoToneSphereHost, period)
	xlog.S(ctx).Info("发起请求：", url)
	ssq := SSQResponse{}
	slingC := sling.New()
	slingC.Get(url).Add("referer", "go tes")
	slingC.Add("Cookie", "_Jo0OQK=36E2440B5D0808590F2B8E8E16280C78448693143B5B5E35C67FC6F8FBB046E6269F101AF90D1CAE075F58B53BEC6C5B9449B54F045CF741DCB2E11BDCE3495366DCF9931755D2F67BE8B280E82C374E4578B280E82C374E457213776806F6D15C2GJ1Z1Uw==; HMF_CI=7862b2e472f3a7e8a1448d834d8f81365ebffbf5b94b108bf6ae0c2e1a071ed4ca; HYB_SH=f042ed75e08e07354382c8419326ee4542")
	req, err := slingC.Request()
	if err != nil {
		xlog.S(ctx).Info("构建request失败：", err.Error())
		return ssq
	}
	xlog.S(ctx).Infof("Request:", req)
	_, err = slingC.Do(req, &ssq, nil)
	if err != nil {
		xlog.S(ctx).Info("发送请求失败:", err.Error())
		return ssq
	}
	xlog.S(ctx).Info("响应结果:", ssq)
	return ssq
}
