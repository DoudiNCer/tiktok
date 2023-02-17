namespace go friend_talk_message_gorm

enum Code {
     Success         = 0
     ParamInvalid    = 1
     DBErr           = 2
     RTErr           = 3
}

struct Message{
    i64 id
    i64 to_user_id
    i64 from_user_id
    string content
    string create_time
}
struct GetChatMessageRequest{
    1: string token
    2: i64 to_user_id
}

struct GetChatMessageResponse{
    1: Code status_code
    2: string status_msg
    3: list<Message> message_list
}

struct PostMessageActionRequest{
    1: string token
    2: i64 to_user_id
    3: i32 action_type
    4: string content
}

struct PostMessageActionResponse{
    1: Code status_code
    2: string status_msg
}

service FriendTalkService{
    GetChatMessageResponse getChatMessage(1:GetChatMessageRequest req)(api.get="/douyin/message/chat/")
    PostMessageActionResponse postMessageAction(1:PostMessageActionRequest req)(api.post="/douyin/message/action/")
}