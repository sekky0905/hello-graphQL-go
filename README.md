# hello-graphQL-go
Simple sample of graphQL


## Query Sample 

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
