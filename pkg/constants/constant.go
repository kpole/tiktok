package constants

const (
	UserTableName      = "users"
	FollowsTableName   = "follows"
	VideosTableName    = "videos"
	MessageTableName   = "messages"
	FavoritesTableName = "likes"
	CommentTableName   = "comments"

	MySQLDefaultDSN = "douyin:douyin123@tcp(127.0.0.1:18000)/douyin?charset=utf8&parseTime=True&loc=Local"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

	RedisAddr     = "localhost:6379"
	RedisPassword = ""

	MinioEndPoint        = "localhost:18001"
	MinioAccessKeyID     = "douyin"
	MinioSecretAccessKey = "douyin123"
	MiniouseSSL          = false
	MinioVideoBucketName = "videobucket"
	MinioImgBucketName   = "imagebucket"

	TestSign       = "测试账号！ offer"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)
