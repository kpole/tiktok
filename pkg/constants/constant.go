package constants

const (
	// NoteTableName           = "note"
	UserTableName    = "users"
	FollowsTableName = "follows"
	VideosTableName  = "videos"
	// SecretKey               = "secret key"
	// IdentityKey             = "id"
	// Total                   = "total"
	// Notes                   = "notes"
	// NoteID                  = "note_id"
	// ApiServiceName          = "demoapi"
	// NoteServiceName         = "demonote"
	// UserServiceName         = "demouser"
	MySQLDefaultDSN = "douyin:douyin123@tcp(127.0.0.1:18000)/douyin?charset=utf8&parseTime=True&loc=Local"
	// EtcdAddress             = "127.0.0.1:2379"
	// CPURateLimit    float64 = 80.0
	// DefaultLimit            = 10

	VideoFeedCount = 30

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
