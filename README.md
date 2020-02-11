# hello-graphQL-go
Simple sample of graphQL


## 仕様

### List 系

- チャットルームの一覧を返すこと
- チャットルーム内のコメントの一覧を返すこと

## Create 系

- ユーザーを新規作成できること
    - 練習なのでパスワード等の設定はなし
- チャットルームを新規作成できること
- コメントを新規作成できること

## Update 系

- コメントを編集できること

## Delete 系

- コメントを削除できること


## クエリのサンプル

```graphql
mutation createUser {
  createUser(input: { name: "name" }) {
    id
    name
  }
}

mutation createChatRoom {
  createChatRoom(input: { userID: "user_{uuid for user_id}", title: "title" }) {
    id
    title
    createdBy {
      id
      name
    }
  }
}

mutation createComment {
  createComment(input: { content: "content", chatRoomID: "chat_room_{uuid for chat_room}", userID: "user_{uuid for user_id}" }) {
    id
    content
    postedBy {
      id
      name
    }
    createdAt
  }
}


query GetUsers {
  users {
    id
    name
  }
}

query GetChatRooms {
  chatRooms {
    id
    title
  }
}

query GetComments {
  commentsByUser(userID: "user_{uuid for user_id}") {
    id
    content
  }
}
```
