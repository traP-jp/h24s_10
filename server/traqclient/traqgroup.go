package traqclient

import (
	"context"
	"log"
	"time"
)

type (
	Group struct {
		Id          string            // グループUUID
		Name        string            // グループ名
		Description string            // グループ説明
		Type        string            // グループタイプ
		Icon        string            // グループアイコンUUID
		Members     []UserGroupMember // グループメンバーの配列
		CreatedAt   time.Time         // 作成日時
		UpdatedAt   time.Time         // 更新日時
		Admins      []string          // グループ管理者のUUIDの配列
	}

	UserGroupMember struct {
		Id   string //	ユーザーUUID
		Role string //	ユーザーの役割
	}
)

func (c *Client) GetUserGroups(ctx context.Context) ([]Group, error) {
	resp, _, err := c.apiClient.GroupApi.GetUserGroups(ctx).Execute()
	if err != nil {
		log.Printf("get user groups error: %v", err)
		return nil, err
	}
	var GroupList []Group
	for _, group := range resp {
		memberList := make([]UserGroupMember, 0, len(group.Members))
		for _, member := range group.Members {
			memberList = append(memberList, UserGroupMember{
				Id:   member.Id,
				Role: member.Role,
			})
		}
		GroupList = append(GroupList, Group{
			Id:          group.Id,
			Name:        group.Name,
			Description: group.Description,
			Type:        group.Type,
			Icon:        group.Icon,
			Members:     memberList,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
			Admins:      group.Admins,
		})
	}
	return GroupList, nil
}
