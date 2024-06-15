package traqclient

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/traPtitech/go-traq"
)

type (
	User struct {
		Id          string    // ユーザーUUID
		Name        string    // ユーザー名
		DisplayName string    // ユーザー表示名
		IconFileId  string    // アイコンファイルUUID
		Bot         bool      // BOTかどうか
		UpdatedAt   time.Time // 更新日時
	}

	UserMap map[string]User
)

func (c *Client) GetUsers(ctx context.Context) ([]User, error) {
	ctx = context.WithValue(ctx, traq.ContextAccessToken, ACCESS_TOKEN)
	resp, _, err := c.apiClient.UserApi.GetUsers(ctx).Execute()
	if err != nil {
		log.Printf("get user error: %v", err)
		return nil, err
	}
	var UserList []User
	for _, user := range resp {
		UserList = append(UserList, User{
			Id:          user.Id,
			Name:        user.Name,
			DisplayName: user.DisplayName,
			IconFileId:  user.IconFileId,
			Bot:         user.Bot,
			UpdatedAt:   user.UpdatedAt,
		})
	}
	return UserList, nil
}

func (c *Client) GetUsersMap(ctx context.Context) (UserMap, error) {
	ctx = context.WithValue(ctx, traq.ContextAccessToken, ACCESS_TOKEN)
	resp, _, err := c.apiClient.UserApi.GetUsers(ctx).Execute()
	if err != nil {
		log.Printf("get user map error: %v", err)
		return nil, err
	}
	UserMap := make(UserMap)
	for _, user := range resp {
		UserMap[user.Id] = User{
			Id:          user.Id,
			Name:        user.Name,
			DisplayName: user.DisplayName,
			IconFileId:  user.IconFileId,
			Bot:         user.Bot,
			UpdatedAt:   user.UpdatedAt,
		}
	}
	return UserMap, nil
}

var ErrUserNotFound = errors.New("user not found")

// GetUser gets a user information by user ID.
// If the user is not found, it returns [ErrUserNotFound].
func (c *Client) GetUser(ctx context.Context, userID string) (User, error) {
	ctx = context.WithValue(ctx, traq.ContextAccessToken, ACCESS_TOKEN)
	user, res, err := c.apiClient.UserApi.GetUser(ctx, userID).Execute()
	if err != nil {
		return User{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return User{}, ErrUserNotFound
	}
	return User{
		Id:          user.Id,
		Name:        user.Name,
		DisplayName: user.DisplayName,
		IconFileId:  user.IconFileId,
		Bot:         user.Bot,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
