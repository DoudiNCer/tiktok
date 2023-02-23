namespace go favorite_gorm

include 'follower_api.thrift'

struct Video {
    i64 id // 视频唯一标识
    follower_api.User author // 视频作者信息
    string play_url // 视频播放地址
    string cover_url // 视频封面地址
    i64 favorite_count // 视频的点赞总数
    i64 comment_count // 视频的评论总数
    bool is_favorite // true-已点赞，false-未点赞
    string title // 视频标题
}

struct FavoriteActionRequest {
    1: string token     (api.query="token")
    2: i64 video_id     (api.query="video_id")
    3: i32 action_type  (api.query="action_type")
}

struct FavoriteActionResponse {
    1: follower_api.Code status_code
    2: string status_msg
}

struct FavoriteListRequest {
    1: i64 user_id  (api.query="user_id", api.vd="regex('^[0-9]*$')")
    2: string token (api.query="token")
}

struct FavoriteListResponse {
    1: follower_api.Code status_code
    2: string status_msg
    3: list<Video> video_list
}

service favoriteService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req)   (api.get="/douyin/favorite/list/")
}