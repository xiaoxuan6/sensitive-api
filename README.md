# sensitive-api
敏感词过滤，查找、替换，仅限用于合法的、积极向上的敏感词过滤使用，严禁用于从事违反法律法规、危害国家、危害人民、不道德的活动！！！

# Docker 

```docker
docker run --name sensitive-api -p 9210:9210 -d ghcr.io/xiaoxuan6/sensitive-api:latest
```

<details>
<summary><b> Shell </b></summary>

支持自定义文件，文件所在 `/root/sensitive/dict`, 自定义文件无需任何操作，程序会自动加载

## install

```shell
bash -c "$(curl -L https://raw.githubusercontent.com/xiaoxuan6/sensitive-api/main/sensitive-api.sh)" @ install
```

## remove

```shell
bash -c "$(curl -L https://raw.githubusercontent.com/xiaoxuan6/sensitive-api/main/sensitive-api.sh)" @ remove
```
</details>

## Demo

### 查找 - findall

```shell
curl -X POST -d "{\"content\":\"测试色情\"}" http://127.0.0.1:9210/sensitive/findall -H "content-type:application/json"
```

输出：

```shell
{"code":200,"data":["色情"],"msg":"OK"}
```

### 替换 - replace 

```shell
curl -X POST -d "{\"content\":\"欲女测试色情\"}" http://127.0.0.1:9210/sensitive/replace -H "content-type:application/json"
```

输出：

```shell
{"code":200,"data":"**测试色情","msg":"OK"}
```

### 过滤 - filter

```shell
curl -X POST -d "{\"content\":\"欲女测试色情\"}" http://127.0.0.1:9210/sensitive/filter -H "content-type:application/json"
```

输出：

```shell
{"code":200,"data":"测试","msg":"OK"}
```