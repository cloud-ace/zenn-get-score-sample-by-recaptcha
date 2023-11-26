package re_capcha

import (
	"context"
	"fmt"
	"time"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	recaptchapb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
)

type Client struct {
	ProjectID  string // Google CloudのプロジェクトID
	SiteKey    string // reCAPTCHAのサイトキー
	outputPath string // 検証のため、reCAPTCHAの結果を保存するパス
}

func NewClient(projectID string, siteKey string, output string) Client {
	return Client{ProjectID: projectID, SiteKey: siteKey, outputPath: output}
}

func (c Client) CreateAssessment(ctx context.Context, token string, recaptchaAction string) (*recaptchapb.RiskAnalysis, error) {
	// reCAPTCHA クライアントを作成する。
	client, err := recaptcha.NewClient(ctx)
	if err != nil {
		fmt.Println("Error creating reCAPTCHA client", err.Error())
		return nil, err
	}
	defer client.Close()

	// 追跡するイベントのプロパティを設定する。
	event := &recaptchapb.Event{
		Token:   token,
		SiteKey: c.SiteKey,
	}

	assessment := &recaptchapb.Assessment{
		Event: event,
	}

	// 評価リクエストを作成する。
	request := &recaptchapb.CreateAssessmentRequest{
		Assessment: assessment,
		Parent:     fmt.Sprintf("projects/%s", c.ProjectID),
	}

	response, err := client.CreateAssessment(
		ctx,
		request)

	if err != nil {
		fmt.Printf("Error calling CreateAssessment: %v", err.Error())
		return nil, err
	}

	// トークンが有効かどうかを確認する。
	if !response.TokenProperties.Valid {
		fmt.Printf("The CreateAssessment() call failed because the token was invalid for the following reasons: %v",
			response.TokenProperties.InvalidReason)
		return nil, err
	}

	// 想定どおりのアクションが実行されたかどうかを確認する。
	if response.TokenProperties.Action != recaptchaAction {
		fmt.Printf("The action attribute in your reCAPTCHA tag does not match the action you are expecting to score")
		return nil, err
	}

	// リスクスコアと理由を出力
	// 評価の解釈の詳細については、以下を参照:
	// https://cloud.google.com/recaptcha-enterprise/docs/interpret-assessment
	fmt.Println(fmt.Sprintf("score: %v. risk: %v", response.RiskAnalysis.Score, response.RiskAnalysis.Reasons))

	// 検証のため、ローカルファイルへ結果保存
	row := [][]string{
		{time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("%v", response.RiskAnalysis.Score), fmt.Sprintf("%v", response.RiskAnalysis.Reasons)},
	}

	WriteLocalFile(c.outputPath, row)

	return response.RiskAnalysis, nil
}
