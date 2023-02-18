namespace go follower_gorm

enum Code {
     Success         = 0
     ParamInvalid    = 1
     DBErr           = 2
     RTErr           = 3
     TokenErr        = 4
}

struct User {
    i64 id // 用户id
    string name // 用户名称
    i64 follow_count // 关注总数
    i64 follower_count // 粉丝总数
    bool is_follow // true-已关注，false-未关注
}

struct CreateFollowerRequest{
    1: string token       (api.query="token")
    2: i64 to_user_id  (api.query="to_user_id")
    3: i32 action_type (api.query="action_type")
}

struct CreateFollowerResponse{
    1: Code status_code
    2: string status_msg
}

struct QueryFollowListRequest {
    1: string user_id    (api.query="user_id") // 用户id
    2: string token   (api.query="token")// 用户鉴权token
}

struct QueryFollowListResponse {
    1: Code status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<User> user_list // 用户信息列表
}

service FollowerService {
   CreateFollowerResponse CreateFollower(1:CreateFollowerRequest req)(api.post="/douyin/relation/action/")
   QueryFollowListResponse QueryFollowList(1:QueryFollowListRequest req)(api.get="/douyin/relation/follow/list/")
}