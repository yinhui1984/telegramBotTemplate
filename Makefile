all:build

tidy:
	go mod tidy && go mod vendor

build:
	go build -o ./bin/BasicTelBot ./


run:build
	## if `heroku maintenance` output is off, don't run the bot
	-heroku maintenance | grep -q "on" && heroku local worker
	-heroku maintenance | grep -q "off" && echo "heroku maintenance is off, please switch it on"


login:
	heroku login

push:maintenance_off
	go mod tidy && go mod vendor
	git add -A . && git commit -m "auto commit" && git push heroku main && echo "push success, the changes will take effect in about 1 minute "

log:
	heroku logs --tail


#https://devcenter.heroku.com/articles/maintenance-mode
#维护模式, 暂时不能访问heroku的服务, 以便使用本地服务进行调试
maintenance_on:
	heroku maintenance:on
	heroku ps:scale worker=0
	heroku maintenance

#关闭维护模式, 恢复访问heroku的服务
maintenance_off:
	heroku maintenance:off
	heroku ps:scale worker=1
	heroku maintenance