namespace go publish_gorm

include 'favorite_api.thrift'

struct publishListRequest {
    1: i64 user_id  (api.query="user_id")// 用户id
    2: string token    (api.query="token")// 用户鉴权token
}

struct publishListResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<favorite_api.Video> video_list // 用户发布的视频列表
}

struct publishActionRequest {
    1: string token (api.body="token")
//    2: binary data    (api.body="data")
    2: string title (api.body="title")
}

struct publishActionResponse {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

service publishService {
    publishListResponse publishList(1: publishListRequest req) (api.get="/douyin/publish/list/")
    publishActionResponse publishAction(1: publishActionRequest req)   (api.post="/douyin/publish/action/")
}