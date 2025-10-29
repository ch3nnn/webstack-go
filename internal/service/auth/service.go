package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/spf13/viper"
)

type Service interface {
	GenerateOAuthURL(ctx context.Context) (string, string, error)                // 返回 URL 和 state
	HandleOAuthCallback(ctx context.Context, code, state string) (string, error) // 返回 JWT token
	ValidateTenant(ctx context.Context, tenantKey string) bool
}

type service struct {
	feishuService *FeishuService
	config        *viper.Viper
}

func NewService(config *viper.Viper, jwt *jwt.JWT) Service {
	return &service{
		feishuService: NewFeishuService(config, jwt),
		config:        config,
	}
}

// GenerateOAuthURL 生成 OAuth 授权 URL
func (s *service) GenerateOAuthURL(ctx context.Context) (string, string, error) {
	// 生成随机 state
	stateBytes := make([]byte, 16)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", "", err
	}
	state := hex.EncodeToString(stateBytes)

	url := s.feishuService.GenerateFeishuOAuthURL(state)
	return url, state, nil
}

// HandleOAuthCallback 处理 OAuth 回调
func (s *service) HandleOAuthCallback(ctx context.Context, code, state string) (string, error) {
	// 获取用户信息
	userInfo, err := s.feishuService.GetFeishuUserInfo(ctx, code)
	if err != nil {
		return "", fmt.Errorf("获取飞书用户信息失败: %w", err)
	}

	// 验证租户（如果配置了）
	var feishuConfig FeishuConfig
	s.config.UnmarshalKey("feishu", &feishuConfig)
	if feishuConfig.TenantKey != "" && userInfo.Data.TenantKey != feishuConfig.TenantKey {
		return "", fmt.Errorf("未授权用户禁止登录")
	}

	// 生成 JWT token
	token, err := s.feishuService.GenerateJWTToken(userInfo)
	if err != nil {
		return "", fmt.Errorf("生成 JWT 令牌失败: %w", err)
	}

	return token, nil
}

// ValidateTenant 验证租户
func (s *service) ValidateTenant(ctx context.Context, tenantKey string) bool {
	var feishuConfig FeishuConfig
	s.config.UnmarshalKey("feishu", &feishuConfig)

	// 如果没有配置租户限制，则允许所有用户
	if feishuConfig.TenantKey == "" {
		return true
	}

	return tenantKey == feishuConfig.TenantKey
}
