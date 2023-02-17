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
1. 成功：查看1001的所有粉丝。其中1004关注了1002。
```json
{
    "status_code": 0,
    "status_msg": "Success",
    "user_list": [
        {
            "id": 1000,
            "name": "1",
            "follow_count": 3,
            "follower_count": 0,
            "is_follow": false
        },
        {
            "id": 1002,
            "name": "3",
            "follow_count": 1,
            "follower_count": 1,
            "is_follow": true
        },
        {
            "id": 1004,
            "name": "5",
            "follow_count": 2,
            "follower_count": 1,
            "is_follow": false
        }
    ]
```
2. followerList为空的情况：查看1000的粉丝列表。
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