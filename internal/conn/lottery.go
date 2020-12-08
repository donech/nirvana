package conn

import (
	"context"
	"fmt"

	"github.com/dghubble/sling"

	"github.com/donech/tool/xlog"
)

const twoToneSphereHost = "http://www.cwl.gov.cn/cwl_admin/kjxx/findKjxx/forIssue?name=ssq&code=%s"
const supperLottoHost = "https://api.xinti.com/chart/QueryPrizeDetailInfo"

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

func (c *LotteryClient) GetTwoToneSphere(ctx context.Context, period string) string {
	url := fmt.Sprintf(twoToneSphereHost, period)
	xlog.S(ctx).Info("发起请求：", url)
	ssq := SSQResponse{}
	slingC := sling.New()
	slingC.Get(url).Add("referer", "go tes")
	slingC.Add("Cookie", "_Jo0OQK=36E2440B5D0808590F2B8E8E16280C78448693143B5B5E35C67FC6F8FBB046E6269F101AF90D1CAE075F58B53BEC6C5B9449B54F045CF741DCB2E11BDCE3495366DCF9931755D2F67BE8B280E82C374E4578B280E82C374E457213776806F6D15C2GJ1Z1Uw==; HMF_CI=7862b2e472f3a7e8a1448d834d8f81365ebffbf5b94b108bf6ae0c2e1a071ed4ca; HYB_SH=f042ed75e08e07354382c8419326ee4542")
	req, err := slingC.Request()
	if err != nil {
		xlog.S(ctx).Info("构建request失败：", err.Error())
		return ""
	}
	xlog.S(ctx).Infof("Request:", req)
	_, err = slingC.Do(req, &ssq, nil)
	if err != nil {
		xlog.S(ctx).Info("发送请求失败:", err.Error())
		return ""
	}
	xlog.S(ctx).Info("响应结果:", ssq)
	if ssq.State != 0 {
		return ""
	}
	return ssq.Result[0].Red + "|" + ssq.Result[0].Blue
}

type SupperLottoResponse struct {
	Date     string `json:"Date"`
	Token    string `json:"Token"`
	Code     int    `json:"Code"`
	GoAction int    `json:"GoAction"`
	Message  string `json:"Message"`
	Value    struct {
		BonusPoolList []struct {
			BonusIndex     int    `json:"BonusIndex"`
			CurrentSales   int    `json:"CurrentSales"`
			IssuseNumber   string `json:"IssuseNumber"`
			PrizeName      string `json:"PrizeName"`
			SingleWinBonus int    `json:"SingleWinBonus"`
			TotalBonusPool int    `json:"TotalBonusPool"`
			WinCount       int    `json:"WinCount"`
			WinNumber      string `json:"WinNumber"`
			StrPrizeDate   string `json:"StrPrizeDate"`
			CreateTime     string `json:"CreateTime"`
			WinConditions  string `json:"WinConditions"`
		} `json:"BonusPoolList"`
		MasterArticleList []struct {
			ArticleID    string  `json:"ArticleId"`
			ArticleTitle string  `json:"ArticleTitle"`
			Price        float64 `json:"Price"`
		} `json:"MasterArticleList"`
		ProficientSchemeList []struct {
			KeepBonusCount     int     `json:"KeepBonusCount"`
			MaxBonusCount      int     `json:"MaxBonusCount"`
			New10HitCount      int     `json:"New10HitCount"`
			SupportCount       int     `json:"SupportCount"`
			UserHeadURL        string  `json:"UserHeadUrl"`
			UserID             string  `json:"UserId"`
			UserName           string  `json:"UserName"`
			OrderID            string  `json:"OrderId"`
			OrderIndex         int     `json:"OrderIndex"`
			SchemePrice        float64 `json:"SchemePrice"`
			SchemeTitle        string  `json:"SchemeTitle"`
			EndTime            string  `json:"EndTime"`
			NewPublishGameCode string  `json:"NewPublishGameCode"`
		} `json:"ProficientSchemeList"`
		GameIssuseList []struct {
			IssuseNumber string `json:"IssuseNumber"`
		} `json:"GameIssuseList"`
		NextIssuseInfo struct {
			IssuseNumber string `json:"IssuseNumber"`
			StopTime     string `json:"StopTime"`
		} `json:"NextIssuseInfo"`
		PreIssuseInfo struct {
			AwardTime      string `json:"AwardTime"`
			StrPrizeDate   string `json:"StrPrizeDate"`
			CurrentSales   int    `json:"CurrentSales"`
			TotalBonusPool int    `json:"TotalBonusPool"`
			IssuseNumber   string `json:"IssuseNumber"`
			WinNumber      string `json:"WinNumber"`
		} `json:"PreIssuseInfo"`
	} `json:"Value"`
}

type SupperLottoReq struct {
	ClientSource int `json:"ClientSource"`
	Param        struct {
		GameCode     string `json:"GameCode"`
		IssuseNumber string `json:"IssuseNumber"`
	} `json:"Param"`
	Date  int64  `json:"Date"`
	Token string `json:"Token"`
	Sign  string `json:"Sign"`
}

func (c *LotteryClient) GetSupperLotto(ctx context.Context, period string) string {
	url := supperLottoHost
	xlog.S(ctx).Info("发起请求：", url)
	supperLotto := SupperLottoResponse{}
	slingC := sling.New()
	body := SupperLottoReq{
		ClientSource: 3,
		Param: struct {
			GameCode     string `json:"GameCode"`
			IssuseNumber string `json:"IssuseNumber"`
		}{GameCode: "DLT", IssuseNumber: period},
		Date:  1607411704948,
		Token: "",
		Sign:  "bc349b1bfb447b5fbbf0e765cc5d4c47",
	}
	req, err := slingC.Post(url).BodyJSON(body).Request()
	if err != nil {
		xlog.S(ctx).Info("构建request失败：", err.Error())
		return ""
	}
	xlog.S(ctx).Infof("Request:", req)
	_, err = slingC.Do(req, &supperLotto, nil)
	if err != nil {
		xlog.S(ctx).Info("发送请求失败:", err.Error())
		return ""
	}
	xlog.S(ctx).Info("响应结果:", supperLotto)
	return supperLotto.Value.PreIssuseInfo.WinNumber
}
