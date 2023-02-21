namespace go friend_list_gorm

enum Code {
     Success         = 0
     ParamInvalid    = 1
     DBErr           = 2
     RTErr           = 3
}

struct GetFriendListRequest{
    1: string user_id       (api.query="user_id")
    2: string token         (api.query="token")
}

struct FriendUser{
    i64 id // 用户ID
    string name // 用户名
    i64 follow_count // 关注总数
    i64 follower_count // 粉丝总数
    bool is_follow // true-已关注，false-未关注
    string avatar // 用户头像Url
    string background_image // 用户个人页顶部大图
    string signature //个人简介
    i64 total_favorited //获赞数量
    i64 work_count //作品数量
    i64 favorite_count //点赞数量
    string message // 和该好友的最新聊天消息
    i64 msgType //  message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

struct GetFriendListResponse{
    1: Code status_code
    2: string status_msg
    3: list<FriendUser> user_list
}

service FriendService{
    GetFriendListResponse GetFriendList(1:GetFriendListRequest req)(api.get="/douyin/relation/friend/list/")
}