#!/usr/bin/env sh


#  https://devcenter.heroku.com/articles/getting-started-with-go#deploy-the-app

# 创建程序时, 运行一次这个脚本即可

git init && git add -A . && git commit -m "Initial commit"


heroku create basic-telegram-app  && git push heroku main

