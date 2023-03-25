package boluobaoapi

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/request"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"time"
)

type Task struct {
	AccountId string
}

func (task *Task) Data(entityType int) string {
	return fmt.Sprintf(`{'seconds': 3605, 'readingDate': '%v', 'entityType': %v}`, task.GetDay(), entityType)
}
func (task *Task) ReadDate() string {
	return task.Data(2)
}
func (task *Task) ListenData() string {
	return task.Data(3)
}

func (task *Task) GetDay() string {
	return time.Now().Format("2006-01-02")
}

func (task *Task) TaskListenTimeAPI() *gjson.Result {
	return VerifyAPI(request.Put("user/readingtime").AddString(task.ListenData()).Json())
}

func (task *Task) TaskReadTimeAPI() *gjson.Result {
	return VerifyAPI(request.Put("user/readingtime").AddString(task.ReadDate()).Json())
}

func (task *Task) TaskReceiveTaskAPI(TaskId string) *gjson.Result {
	params := map[string]string{"seconds": "3605", "readingDate": task.GetDay(), "entityType": "3"}
	return VerifyAPI(request.Post("user/tasks/" + TaskId).Data(params).Json())
}

func (task *Task) TaskSignInAPI() *gjson.Result {
	//var SignIn Template.SignIn
	// "user/signInfo" is old api, now it's "user/newSignInfo"
	FormData := fmt.Sprintf(`{'signDate': '%v'}`, task.GetDay())
	return VerifyAPI(request.Put("user/newSignInfo").AddString(FormData).Json())
}

func (task *Task) TaskListAPI() *gjson.Result {
	return VerifyAPI(request.Get("user/tasks").Data(map[string]string{
		"taskCategory": "1",
		"package":      "com.sfacg",
		"deviceToken":  uuid.New().String(),
		"page":         "0",
		"size":         "20",
	}).Json())
}

func (task *Task) TaskShareAPI(account_id string) {
	// This interaction logic is so bad, I will refactor it, but now I will leave it like this.
	// I will refactor it when I have time.
	var change bool
	if request.Request.App {
		change = true
		request.Request.App = false
	} else {
		change = false
	}
	request.Put("user/tasks?taskId=4&userId=" + account_id).AddString(`{"env": 0}`).Request()

	if change {
		request.Request.App = true // change back to app
	}
}
