namespace go follower_gorm

struct CreateFollowerRequest{
    1: string token       (api.query="token")
    2: string to_user_id  (api.query="to_user_id")
    3: string action_type (api.query="action_type")
}

struct CreateFollowerResponse{
    1: i64 status_code
    2: string status_msg
}

service FollowerService {
   CreateFollowerResponse CreateFollower(1:CreateFollowerRequest req)(api.post="/douyin/relation/action/")
}