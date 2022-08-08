# telegramBotTemplate
 电报机器人模板


 主要是将Message和Query以事件的形式发出.
 业务代码只需要在commandsHandler中添加

 比如:
 1. touch commandsHandler/myCommandHander.go

 2. `myCommandHander.go`中编写

 ```go
 func (c myCommandHander) RegisterUpdaterEvent() {
	events.UpdaterEventInstance().Register(func(event *events.UpdaterEvent, arg interface{}) {

		// if the type of msg is *tgbotapi.Message
		if msg, ok := arg.(*tgbotapi.Message); ok {
			switch msg.Command() {
			case myCommandName: // "/xxx"
				//当机器人收到 /xxx 时调用处理函数
			}
		}

		// if the type of msg is *tgbotapi.CallbackQuery
		if query, ok := arg.(*tgbotapi.CallbackQuery); ok {
			switch query.Data {
                // /当机器人收到相应query时调用处理函数
			}
		}
	})
}
 ```

 3.  在 `main`中
 ```go
 	my := commandsHandler.myCommandHander{}
	my.RegisterUpdaterEvent()
```



>NOTE: 在部署和测试之前, 修改.evn文件中的token, token可以从botFather中申请
