// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc_queries

import (
	"encoding/json"
	"time"
)

type AuthUser struct {
	ID          int32     `json:"id"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"lastLogin"`
	IsSuperuser bool      `json:"isSuperuser"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	IsStaff     bool      `json:"isStaff"`
	IsActive    bool      `json:"isActive"`
	DateJoined  time.Time `json:"dateJoined"`
}

type AuthUserManagement struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"userID"`
	RateLimit int32     `json:"rateLimit"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChatLog struct {
	ID        int32           `json:"id"`
	Session   json.RawMessage `json:"session"`
	Question  json.RawMessage `json:"question"`
	Answer    json.RawMessage `json:"answer"`
	CreatedAt time.Time       `json:"createdAt"`
}

type ChatMessage struct {
	ID              int32           `json:"id"`
	Uuid            string          `json:"uuid"`
	ChatSessionUuid string          `json:"chatSessionUuid"`
	Role            string          `json:"role"`
	Content         string          `json:"content"`
	LlmSummary      string          `json:"llmSummary"`
	Score           float64         `json:"score"`
	UserID          int32           `json:"userID"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	CreatedBy       int32           `json:"createdBy"`
	UpdatedBy       int32           `json:"updatedBy"`
	IsDeleted       bool            `json:"isDeleted"`
	IsPin           bool            `json:"isPin"`
	TokenCount      int32           `json:"tokenCount"`
	Raw             json.RawMessage `json:"raw"`
}

type ChatModel struct {
	ID                     int32  `json:"id"`
	Name                   string `json:"name"`
	Label                  string `json:"label"`
	IsDefault              bool   `json:"isDefault"`
	Url                    string `json:"url"`
	ApiAuthHeader          string `json:"apiAuthHeader"`
	ApiAuthKey             string `json:"apiAuthKey"`
	UserID                 int32  `json:"userID"`
	EnablePerModeRatelimit bool   `json:"enablePerModeRatelimit"`
	MaxToken               int32  `json:"maxToken"`
	DefaultToken           int32  `json:"defaultToken"`
	OrderNumber            int32  `json:"orderNumber"`
	HttpTimeOut            int32  `json:"httpTimeOut"`
}

type ChatPrompt struct {
	ID              int32     `json:"id"`
	Uuid            string    `json:"uuid"`
	ChatSessionUuid string    `json:"chatSessionUuid"`
	Role            string    `json:"role"`
	Content         string    `json:"content"`
	Score           float64   `json:"score"`
	UserID          int32     `json:"userID"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatedBy       int32     `json:"createdBy"`
	UpdatedBy       int32     `json:"updatedBy"`
	IsDeleted       bool      `json:"isDeleted"`
	TokenCount      int32     `json:"tokenCount"`
}

type ChatSession struct {
	ID            int32     `json:"id"`
	UserID        int32     `json:"userID"`
	Uuid          string    `json:"uuid"`
	Topic         string    `json:"topic"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Active        bool      `json:"active"`
	Model         string    `json:"model"`
	MaxLength     int32     `json:"maxLength"`
	Temperature   float64   `json:"temperature"`
	TopP          float64   `json:"topP"`
	MaxTokens     int32     `json:"maxTokens"`
	N             int32     `json:"n"`
	SummarizeMode bool      `json:"summarizeMode"`
	Debug         bool      `json:"debug"`
}

type ChatSnapshot struct {
	ID           int32           `json:"id"`
	Uuid         string          `json:"uuid"`
	UserID       int32           `json:"userID"`
	Title        string          `json:"title"`
	Summary      string          `json:"summary"`
	Model        string          `json:"model"`
	Tags         json.RawMessage `json:"tags"`
	Session      json.RawMessage `json:"session"`
	Conversation json.RawMessage `json:"conversation"`
	CreatedAt    time.Time       `json:"createdAt"`
	Text         string          `json:"text"`
	SearchVector interface{}     `json:"searchVector"`
}

type JwtSecret struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Secret   string `json:"secret"`
	Audience string `json:"audience"`
	Lifetime int16  `json:"lifetime"`
}

type UserActiveChatSession struct {
	ID              int32     `json:"id"`
	UserID          int32     `json:"userID"`
	ChatSessionUuid string    `json:"chatSessionUuid"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type UserChatModelPrivilege struct {
	ID          int32     `json:"id"`
	UserID      int32     `json:"userID"`
	ChatModelID int32     `json:"chatModelID"`
	RateLimit   int32     `json:"rateLimit"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   int32     `json:"createdBy"`
	UpdatedBy   int32     `json:"updatedBy"`
}