# 阿里百川openim

## 功能
- `通过` [添加好友](taobao.openim.users.add)
- `通过` [更新好友](taobao.openim.users.update)
- `通过` [删除好友](taobao.openim.users.delet)

## 环境
   mac golang

## 使用

``` c++
1：配置

	func configIM() {
		freeAliim.AppKey = freeFile.FlAppInfo.IMAppkey
		freeAliim.AppSecret = freeFile.FlAppInfo.IMAppSecret
		freeAliim.IsPro = freeFile.FlAppInfo.IsPro
		freeAliim.UseHTTP = freeFile.FlAppInfo.UseHTTP
	}

2：传参数

	imUserinfo := freeAliim.ImUserInfo{
		Userid:   userid,
		Password: pwd,
		Name:     name,
		IconUrl:  iconUrl,
	}

3： 调用封装的接口

	succ, resposne := freeAliim.SendAddUsers(imUserinfos)
	if succ == false {
		
	} else {
		
	}
```