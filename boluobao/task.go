package boluobao

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/boluobaoapi"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

func NovelCompleteTas() {
	Task := boluobaoapi.Task{}
	Task.TaskSignInAPI()
	if TaskList := Task.TaskListAPI(); TaskList != nil {
		for _, data := range TaskList.Get("data").Array() {
			Task.TaskReceiveTaskAPI(data.Get("taskId").String())
		}
		for i, url := range []string{"user/tasks/4", "user/tasks/5", "user/tasks/17"} {
			fmt.Println("complete task No:", i+1)
			if url == "user/tasks/4" && Task.AccountId != "" {
				Task.TaskShareAPI(Task.AccountId)
			} else {
				Task.TaskListenTimeAPI()
				Task.TaskReadTimeAPI()
			}
			request.Post(url).AddString(Task.ListenData()).Request()
		}
	}
}
