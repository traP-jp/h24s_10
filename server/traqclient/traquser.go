package traqclient

import (
	"context"
	"log"
	"time"
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
