#!/usr/bin/env bash

#hz cli 使用proto自动生成命令
source ~/.bash_profile

hz update -idl idl/user.proto
hz update -idl idl/publish.proto
hz update -idl idl/feed.proto
hz update -idl idl/favorite.proto
hz update -idl idl/comment.proto
hz update -idl idl/relation.proto
hz update -idl idl/message.proto
