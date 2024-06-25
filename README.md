### 食用
首先得到`Docker API version`并写入环境变量
```bash
API_VERSION=$(docker version | grep "API version" | head -1 | awk '{print $3}')
export DOCKER_API_VERSION=${API_VERSION}
```
之后是直接运行还是编译运行自行选择.
