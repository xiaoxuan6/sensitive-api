#!/bin/bash

DIR="/root/sensitive"
SERVICE_FILE="/etc/systemd/system/sensitive_api.service"
NGINX_FILE="/etc/nginx/sites-enabled/sensitive_api"

# 根据操作系统和处理器架构选择下载的文件名
if [[ "$(uname -s)" == "Linux" ]]; then
  if [[ "$(uname -m)" == "x86_64" ]]; then
    file_name="linux_x86_64.tar.gz"
  else
    file_name="linux_arm64.tar.gz"
  fi
elif [[ "$(uname -s)" == "Darwin" ]]; then
  if [[ "$(uname -m)" == "x86_64" ]]; then
    file_name="darwin_x86_64.tar.gz"
  else
    file_name="darwin_arm64.tar.gz"
  fi
else
  echo "Unsupported platform: $(uname -s) $(uname -m)"
  exit 1
fi

addSystemd() {
  echo
  echo "create systemctl sensitive_api service"
  echo
  cat <<EOL | sudo tee "$SERVICE_FILE"
[Unit]
Description=sensitive api Service
After=network.target

[Service]
ExecStart=$DIR/sensitive_api
WorkingDirectory=$DIR
Type=simple
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOL

  # 重新加载 systemd 配置
  systemctl daemon-reload
  systemctl start sensitive_api.service
  systemctl enable sensitive_api.service
}

createConfOption() {
  echo "1、创建新文件 $NGINX_FILE 并写入："
  cat <<EOL | tee
server {
listen 9211;
location / {
  proxy_pass http://127.0.0.1:9210;
  proxy_set_header Host \$host;
  proxy_set_header X-Real-IP \$remote_addr;
  proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
  proxy_set_header X-Forwarded-Proto \$scheme;
}
}
EOL
  echo "2、查看配置文件是否有效：'nginx -t'"
  echo "3、重载 nginx 配置文件：'systemctl reload nginx'"
  echo "4、输出如下、执行成功！"
  echo "已启动服务 sensitive_api"
  echo "地址：http://127.0.0.1"
  echo "端口：9211"
  echo "接口文档：https://github.com/xiaoxuan6/sensitive-api?tab=readme-ov-file#demo"

}

nginxProxy() {
  if [ -x "$(command -v nginx)" ]; then
    local nginxFile="/etc/nginx/sites-enabled"
    if [ -d "$nginxFile" ]; then
      cat <<EOL | sudo tee "$NGINX_FILE"
server {
  listen 9211;
  location / {
    proxy_pass http://127.0.0.1:9210;
    proxy_set_header Host \$host;
    proxy_set_header X-Real-IP \$remote_addr;
    proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto \$scheme;
  }
}
EOL
      nginx -t
      systemctl reload nginx
    else
      echo "nginx 配置文件夹 sites-enabled 不存在！"
      echo "查找 nginx 配置文件夹所在位置并执行如下："
      createConfOption
      exit 1
    fi
  else
    echo "未安装 nginx 请安装之后继续执行如下："
    createConfOption
    exit 1
  fi
}

install() {
  mkdir -p "$DIR"
  cd "$DIR" || exit

  URL=$(curl -s https://api.github.com/repos/xiaoxuan6/sensitive-api/releases/latest | grep "browser_download_url" | grep "tar.gz" | cut -d '"' -f 4 | grep "$file_name")
  FILENAME=$(echo "$URL" | cut -d '/' -f 9)

  curl -L -O "$URL"
  tar xf "$FILENAME"
  rm "$FILENAME"

  chmod +x "$DIR"/sensitive_api
  addSystemd
  echo
  echo "sensitive_api 已安装！"
  echo

  # 配置 nginx 反向代理
  nginxProxy
  echo
  echo "已启动服务 sensitive_api"
  echo "地址：http://127.0.0.1"
  echo "端口：9211"
  echo "接口文档：https://github.com/xiaoxuan6/sensitive-api?tab=readme-ov-file#demo"
}

remove() {
  systemctl stop sensitive_api.service
  systemctl disable sensitive_api.service

  if [ -f "$SERVICE_FILE" ]; then
    rm "$SERVICE_FILE"
    echo "服务文件 $SERVICE_FILE 已删除"
  else
    echo "服务文件 $SERVICE_FILE 不存在！"
  fi

  systemctl daemon-reload

  if [ -d "$DIR" ]; then
    rm -rf "$DIR"
    echo "文件夹 $DIR 已删除"
  else
    echo "文件夹 $DIR 不存在！"
  fi

  echo "sensitive_api 已卸载！"

  rm "$NGINX_FILE"
  systemctl reload nginx
}

case "$1" in
install)
  install
  ;;
remove)
  remove
  ;;
*)
  echo "Not found $1 option"
  echo "Usage: $0 {install|remove}"
  echo ""
  exit 1
  ;;
esac
