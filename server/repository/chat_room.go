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

// newChatRoomFromChatRoomDTO は、chatRoomDTO から domain model の ChatRoom を作成する。
func newChatRoomFromChatRoomDTO(dto *ChatRoomDTO) *model.ChatRoom {
	return &model.ChatRoom{
		ID:        dto.ID,
		UserID:    dto.UserID,
		Title:     dto.Title,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

// ChatRoomRepository は、ChatRoom の Repository。
// 練習用の Project なので、Interface 等は定義しない。
type ChatRoomRepository struct {
	MockDB []*ChatRoomDTO
}

// IsExistID は、引数で渡された id を持つ ChatRoom が既に存在するかどうかを確認する。
func (r *ChatRoomRepository) IsExistID(id string) bool {
	room := r.GetChatRoomByID(id)

	return room != nil
}

// GetChatRoomByID は、ChatRoom を1件取得する。
func (r *ChatRoomRepository) GetChatRoomByID(id string) *model.ChatRoom {
	for _, dto := range r.MockDB {
		if dto.ID == id {
			return newChatRoomFromChatRoomDTO(dto)
		}
	}

	return nil
}

// GetChatRoomList は、ChatRoom の一覧を取得する。
func (r *ChatRoomRepository) GetChatRoomList() []*model.ChatRoom {
	n := len(r.MockDB)
	rooms := make([]*model.ChatRoom, n, n)
	for i, dto := range r.MockDB {
		rooms[i] = newChatRoomFromChatRoomDTO(dto)
	}

	return rooms
}

// GetChatRoomListByUserID は、指定された UserID を持つ ChatRoom の一覧を取得する。
func (r *ChatRoomRepository) GetChatRoomListByUserID(userID string) []*model.ChatRoom {
	var rooms []*model.ChatRoom
	for _, dto := range r.MockDB {
		if dto.UserID == userID {
			rooms = append(rooms, newChatRoomFromChatRoomDTO(dto))
		}
	}

	return rooms
}

// AddChatRoom は、ChatRoom を追加する。
func (r *ChatRoomRepository) AddChatRoom(chatRoom *model.ChatRoom) {
	c := newChatRoomDTOFromChatRoom(chatRoom)
	r.MockDB = append(r.MockDB, c)
}
