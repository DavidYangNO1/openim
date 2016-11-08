package aliim

import (
	"encoding/json"
	"fmt"
	"time"
)

type ImUserInfo struct {
	Userid   string `json:"userid"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IconUrl  string `json:"icon_url"`
}

type UidSucc struct {
	Uids []string `json:"string"`
}
type UidFail struct {
	Uids []string `json:"string"`
}
type FailMsg struct {
	FailMsg []string `json:"string"`
}
type UserAddResponse struct {
	UidSucc UidSucc `json:"uid_succ"`
	UidFail UidFail `json:"uid_fail"`
	FailMsg FailMsg `json:"fail_msg"`
}

type DeleteMsg struct {
	Msg []string `json:"string"`
}
type UserDeleteResponse struct {
	DeleteMsg DeleteMsg `json:"result"`
}

func getCommonParams() map[string]string {

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	return params
}

// 导入用户
func SendAddUsers(imUserInfos []ImUserInfo) (success bool, response string) {

	for _, user := range imUserInfos {
		if user.Userid == "" || user.Password == "" {
			return false, "userid or password is required"
		}
	}
	params := getCommonParams()
	params["method"] = OpenImUserAdd

	result, err := json.Marshal(imUserInfos)
	if err != nil {
		return false, err.Error()
	}
	params["userinfos"] = string(result)

	succ, resData := IMPost(params)

	if succ == false {
		return false, response
	}

	type Result struct {
		Result UserAddResponse `json:"openim_users_add_response"`
	}

	var resultResponse Result
	err = json.Unmarshal(resData, &resultResponse)
	if err != nil {
		//fmt.Println("err   " + err.Error())
		return false, err.Error()
	}

	fmt.Println("resData   " + string(resData))
	failMsg := resultResponse.Result.FailMsg
	if len(failMsg.FailMsg) <= 0 {
		return true, "add success"
	}
	return false, failMsg.FailMsg[0]
}

func SendDeleteUsers(userids string) (success bool, response string) {
	if userids == "" {
		return false, "userid is required"
	}
	params := getCommonParams()
	params["method"] = OpenImUserDelete
	params["userids"] = userids

	succ, resData := IMPost(params)
	//fmt.Println("resData " + string(resData))
	if succ == false {
		return false, response
	}

	type Result struct {
		UserDeleteResponse UserDeleteResponse `json:"openim_users_delete_response"`
	}
	fmt.Println("resData   " + string(resData))
	var resultResponse Result
	err := json.Unmarshal(resData, &resultResponse)
	if err != nil {
		return false, err.Error()
	}
	return true, "ok"
}

func SendUpdateUsers(imUserInfos []ImUserInfo) (success bool, response string) {

	for _, user := range imUserInfos {
		if user.Userid == "" {
			return false, "userid is required"
		}
	}
	params := getCommonParams()
	params["method"] = OpenImUserUpdate

	result, err := json.Marshal(imUserInfos)
	if err != nil {
		return false, err.Error()
	}
	params["userinfos"] = string(result)

	succ, resData := IMPost(params)

	if succ == false {
		return false, response
	}

	type Result struct {
		Result UserAddResponse `json:"openim_users_update_response"`
	}
	fmt.Println("resData   " + string(resData))
	var resultResponse Result
	err = json.Unmarshal(resData, &resultResponse)
	if err != nil {
		//fmt.Println("err   " + err.Error())
		return false, err.Error()
	}
	failMsg := resultResponse.Result.FailMsg
	if len(failMsg.FailMsg) <= 0 {
		return true, "update success"
	}
	return false, failMsg.FailMsg[0]

}
