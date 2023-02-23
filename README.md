## Tiktok

### 1.Setup Basic Dependence

#### 1.1 本地环境
[Download ffmpeg package](https://ffmpeg.org/download.html) && **add ffmpeg to system path or user path**

```shell
docker-compose up
```
#### 1.2 服务器docker环境
```shell
docker build -t tiktok:latest -f ./docker-build/Dockerfile .

cd docker-build 
docker-compose up -d
```