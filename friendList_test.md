## followerList 测试说明

| 用户id      | 粉丝id |
| ----------- | ----------- |
| 1001        | 1000        |
| 1001        | 1002        |
| 1001        | 1004        |
| 1002        | 1004        |
| 1003        | 1000        |
| 1004        | 1001        |
| 1005        | 1000        |

（当前登录用户为1004。）

1. 存在好友：查看1004的好友列表
```json
{
    "status_code": 0,
    "status_msg": "Success",
    "user_list": [
        {
            "friend": {
                "id": 1001,
                "name": "2",
                "follow_count": 1,
                "follower_count": 3,
                "is_follow": true
            },
            "message": "this is a message",
            "msgType": -1
        }
    ]
}
```
2. 不存在好友：查看1005的好友列表
```json
{
    "status_code": 0,
    "status_msg": "Success",
    "user_list": null
}
```
3. 用户不存在：查找1007的粉丝列表。（疑问：是否需要报错）
```json
{
    "status_code": 0,
    "status_msg": "Success",
    "user_list": null
}
```