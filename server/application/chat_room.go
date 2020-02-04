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

// GetChatRoomList は、ChatRoom の一覧を取得する。
func (s ChatRoomApplicationService) GetChatRoomList() []*model.ChatRoom {
	return s.Repo.GetChatRoomList()
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
