package traqclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/traP-jp/h24s_10/model"
	"github.com/traPtitech/go-traq"
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

	GroupDetail struct {
		Id          string                  // グループUUID
		Name        string                  // グループ名
		Description string                  // グループ説明
		Type        string                  // グループタイプ
		Icon        string                  // グループアイコンUUID
		Members     []UserGroupMemberDetail // グループメンバーの配列
		CreatedAt   time.Time               // 作成日時
		UpdatedAt   time.Time               // 更新日時
		Admins      []string                // グループ管理者のUUIDの配列
	}

	UserGroupMember struct {
		Id   string // ユーザーUUID
		Role string // ユーザーの役割
	}

	UserGroupMemberDetail struct {
		Id          string // ユーザーUUID
		Role        string // ユーザーの役割
		Name        string
		DisplayName string
	}
)

func (c *Client) GetUserGroups(ctx context.Context) ([]Group, error) {
	ctx = context.WithValue(ctx, traq.ContextAccessToken, ACCESS_TOKEN)
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

func (c *Client) GetUserGroupsMap(ctx context.Context) (map[string]Group, error) {
	groupMap := make(map[string]Group)
	groupList, err := c.GetUserGroups(ctx)
	if err != nil {
		return nil, err
	}
	for _, group := range groupList {
		groupMap[group.Name] = group
	}
	return groupMap, nil
}

func (c *Client) CreateUserGroup(ctx context.Context, name string, description string, groupType string, adminID string, participants []model.Participant) (GroupDetail, error) {
	fmt.Println("CreateUserGroup")
	if len(participants) == 0 || len(participants) == 1 {
		log.Println("No participants")
		return GroupDetail{}, nil
	}
	fmt.Println("participants: ", participants)

	ctx = context.WithValue(ctx, traq.ContextAccessToken, ACCESS_TOKEN)
	postUserGroupRequest := *traq.NewPostUserGroupRequest(
		name,
		description,
		groupType,
	)

	fmt.Println("postUserGroupRequest: ", postUserGroupRequest)

	resp, _, err := c.apiClient.GroupApi.CreateUserGroup(ctx).PostUserGroupRequest(postUserGroupRequest).Execute()
	if err != nil {
		fmt.Printf("create user group error: %v", err)
		return GroupDetail{}, err
	}

	fmt.Println("resp: ", resp)

	admin, err := c.GetUser(ctx, adminID)
	if err != nil {
		fmt.Printf("get user error: %v", err)
		return GroupDetail{}, err
	}

	postUserGroupAdminRequest := *traq.NewPostUserGroupAdminRequest(admin.Id)
	_, err = c.apiClient.GroupApi.AddUserGroupAdmin(ctx, resp.Id).PostUserGroupAdminRequest(postUserGroupAdminRequest).Execute()
	if err != nil {
		fmt.Printf("add user group admin error: %v", err)
		return GroupDetail{}, err
	}

	// fmt.Println("userGroupAdmin: ", postUserGroupAdminRequest)

	userGroupMembers := make([]UserGroupMemberDetail, 0, len(participants))
	for _, participant := range participants {
		user, err := c.GetUser(ctx, participant.TraQID)
		if err != nil {
			fmt.Printf("get user error: %v", err)
			return GroupDetail{}, err
		}
		userGroupMember := *traq.NewUserGroupMember(user.Id, "")
		_, err = c.apiClient.GroupApi.AddUserGroupMember(ctx, resp.Id).UserGroupMember(userGroupMember).Execute()
		if err != nil {
			fmt.Printf("add user group member error: %v", err)
			return GroupDetail{}, err
		}
		userGroupMembers = append(userGroupMembers, UserGroupMemberDetail{
			Id:          participant.TraQID,
			Role:        "一般人",
			Name:        user.Name,
			DisplayName: user.DisplayName,
		})
		fmt.Println("userGroupMember: ", userGroupMember)
	}
	return GroupDetail{
		Id:          resp.Id,
		Name:        resp.Name,
		Description: resp.Description,
		Type:        resp.Type,
		Icon:        resp.Icon,
		Members:     userGroupMembers,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
		Admins:      resp.Admins,
	}, nil
}
