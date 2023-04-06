
# A

```sh
#!/bin/bash

# 本行也是注释
echo 'hello world' # 井号后面的部分是注释

echo "全部参数：" $@
echo "命令行参数数量："$#
echo '$0 = ' $0
echo '$1 = ' $1
echo '$2 = ' $2
echo '$3 = ' $3
```

## date\time

```sh
date

```

## 交叉编译

go env -w GOOS=windows

set http_proxy=<http://127.0.0.1:7890>
set https_proxy=<http://127.0.0.1:7890>
