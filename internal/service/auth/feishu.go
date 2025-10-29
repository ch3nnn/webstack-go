package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/spf13/viper"
)

type FeishuService struct {
	config *viper.Viper
	jwt    *jwt.JWT
}

type FeishuConfig struct {
	AppID       string `mapstructure:"app_id"`
	AppSecret   string `mapstructure:"app_secret"`
	RedirectURL string `mapstructure:"redirect_url"`
	TenantKey   string `mapstructure:"tenant_key"`
}

type AuthConfig struct {
	CookieName     string `mapstructure:"cookie_name"`
	JWTExpireHours int    `mapstructure:"jwt_expire_hours"`
}

type FeishuAccessTokenResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken      string `json:"access_token"`
		TokenType        string `json:"token_type"`
		ExpiresIn        int    `json:"expires_in"`
		RefreshToken     string `json:"refresh_token"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		Scope            string `json:"scope"`
	} `json:"data"`
}

type FeishuUserInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name        string `json:"name"`
		AvatarThumb string `json:"avatar_thumb"`
		UnionID     string `json:"union_id"`
		OpenID      string `json:"open_id"`
		Email       string `json:"email"`
		TenantKey   string `json:"tenant_key"`
	} `json:"data"`
}

type FeishuAppAccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

func NewFeishuService(config *viper.Viper, jwt *jwt.JWT) *FeishuService {
	return &FeishuService{
		config: config,
		jwt:    jwt,
	}
}

// GenerateFeishuOAuthURL 生成飞书 OAuth 授权 URL
func (s *FeishuService) GenerateFeishuOAuthURL(state string) string {
	var feishuConfig FeishuConfig
	s.config.UnmarshalKey("feishu", &feishuConfig)

	params := url.Values{}
	params.Set("redirect_uri", feishuConfig.RedirectURL)
	params.Set("app_id", feishuConfig.AppID)
	params.Set("state", state)

	return fmt.Sprintf("https://open.feishu.cn/open-apis/authen/v1/index?%s", params.Encode())
}

// GetFeishuAppAccessToken 获取飞书应用访问令牌
func (s *FeishuService) GetFeishuAppAccessToken(ctx context.Context) (string, error) {
	var feishuConfig FeishuConfig
	s.config.UnmarshalKey("feishu", &feishuConfig)

	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

	requestBody := map[string]string{
		"app_id":     feishuConfig.AppID,
		"app_secret": feishuConfig.AppSecret,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResp FeishuAppAccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	if tokenResp.Code != 0 {
		return "", fmt.Errorf("获取应用访问令牌失败: %s", tokenResp.Msg)
	}

	return tokenResp.TenantAccessToken, nil
}

// GetFeishuUserInfo 获取飞书用户信息
func (s *FeishuService) GetFeishuUserInfo(ctx context.Context, code string) (*FeishuUserInfoResponse, error) {
	appAccessToken, err := s.GetFeishuAppAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	url := "https://open.feishu.cn/open-apis/authen/v1/access_token"

	requestBody := map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", appAccessToken))

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userResp FeishuUserInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, err
	}

	if userResp.Code != 0 {
		return nil, fmt.Errorf("获取用户信息失败: %s", userResp.Msg)
	}

	return &userResp, nil
}

// GenerateJWTToken 生成 JWT 令牌
func (s *FeishuService) GenerateJWTToken(userInfo *FeishuUserInfoResponse) (string, error) {
	var authConfig AuthConfig
	s.config.UnmarshalKey("auth", &authConfig)

	expireTime := time.Now().Add(time.Duration(authConfig.JWTExpireHours) * time.Hour)

	// 使用固定的用户ID 1 来生成JWT，避免影响现有系统
	return s.jwt.GenToken(1, expireTime)
}
