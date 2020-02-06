package application

import (
	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
	"github.com/sekky0905/hello-graphQL-go/server/repository"
	"golang.org/x/xerrors"
)

// ChatRoomApplicationService は、ChatRoom の ApplicationService。
type ChatRoomApplicationService struct {
	Repo *repository.ChatRoomRepository
}

// GetChatRoomByID は、ChatRoom を1件取得する。
func (s ChatRoomApplicationService) GetChatRoomByID(id string) *model.ChatRoom {
	return s.Repo.GetChatRoomByID(id)
}

// GetChatRoomList は、ChatRoom の一覧を取得する。
func (s ChatRoomApplicationService) GetChatRoomList() []*model.ChatRoom {
	return s.Repo.GetChatRoomList()
}

// GetChatRoomListByUserID は、指定された UserID を持つ ChatRoom の一覧を取得する。
func (s ChatRoomApplicationService) GetChatRoomListByUserID(userID string) []*model.ChatRoom {
	return s.Repo.GetChatRoomListByUserID(userID)
}

// CreateChatRoom は、ChatRoom を作成する。
func (s ChatRoomApplicationService) CreateChatRoom(userID, title string) (*model.ChatRoom, error) {
	comment, err := model.NewChatRoom(userID, title)
	if err != nil {
		return nil, xerrors.Errorf("failed to generate comment: %w", err)
	}
	// 普通だと諸々の処理を行う。
	s.Repo.AddChatRoom(comment)
	return comment, nil
}
