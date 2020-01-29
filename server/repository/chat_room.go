package repository

import (
	"time"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

// ChatRoomDTO は、ChatRoom を表す。
type ChatRoomDTO struct {
	ID        string
	UserID    string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// newChatRoomDTOFromChatRoom は、domain model の ChatRoom から ChatRoomDTO を作成する。
func newChatRoomDTOFromChatRoom(chatRoom *model.ChatRoom) *ChatRoomDTO {
	return &ChatRoomDTO{
		ID:        chatRoom.ID,
		UserID:    chatRoom.UserID,
		Title:     chatRoom.Title,
		CreatedAt: chatRoom.CreatedAt,
		UpdatedAt: chatRoom.UpdatedAt,
	}
}

// ChatRoomRepository は、ChatRoom の Repository。
// 練習用の Project なので、Interface 等は定義しない。
type ChatRoomRepository struct {
	MockDB []*ChatRoomDTO
}

// AddChatRoom は、ChatRoom を追加する。
func (r *ChatRoomRepository) AddChatRoom(chatRoom *model.ChatRoom) {
	c := newChatRoomDTOFromChatRoom(chatRoom)
	r.MockDB = append(r.MockDB, c)
}
