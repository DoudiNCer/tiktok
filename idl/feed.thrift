namespace go feed

include "favorite_api.thrift"

struct FeedRequest {
    1: i64 last_time
    2: string token
}

struct FeedResponse {
    1: i32 status_code
    2: string status_msg
    3: list<favorite_api.Video> video_list
    4: i64 next_time
}

service FeedService {
    FeedResponse QueryFeedList(1: FeedRequest req) (api.get="/douyin/feed/")
}