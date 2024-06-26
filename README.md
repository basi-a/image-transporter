## IMAGE-TRANSPORTER
镜像搬运工
### 食用
显而易见哈，运行这个的用户`docker`组的权限还是要有的.

首先得到`Docker API version`并写入环境变量:
```bash
API_VERSION=$(docker version | grep "API version" | head -1 | awk '{print $3}')
export DOCKER_API_VERSION=${API_VERSION}
```
之后是直接运行还是编译运行自行选择.
### 配置存放
- 一般配置记录于`$HOME/.config/image-transporter/config.json`
- 敏感配置记录于`sqlite`, 默认为`$HOME/.config/image-transporter/config.db`, 当然数据库文件存放位置可以到一般配置中修改
### TODO
[x] 镜像拉取、查询、保存等基础实现

[] 基础API

[] 云服务对接API

[] 前端界面