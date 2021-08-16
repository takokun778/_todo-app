package main

import (
	"client/pb/task"
	"client/pb/user"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:19003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userClient := user.NewUserServiceClient(conn)

	userRegisterRequest := new(user.RegisterRequest)
	userRegisterRequest.FirstName = "玲"
	userRegisterRequest.LastName = "大園"
	userRegisterRequest.Birthday = &date.Date{
		Year:  2001,
		Month: 4,
		Day:   18,
	}

	userRegisterResponse, err := userClient.Register(ctx, userRegisterRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(userRegisterResponse.User.Id + userRegisterResponse.User.LastName + userRegisterResponse.User.FirstName + userRegisterResponse.User.Birthday.String() + strconv.FormatInt(userRegisterResponse.User.Age, 10))

	userGetRequest := new(user.GetRequest)
	userGetRequest.Id = userRegisterResponse.User.Id

	userGetResponse, err := userClient.Get(ctx, userGetRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(userGetResponse.User.Id + userGetResponse.User.LastName + userGetResponse.User.FirstName + userGetResponse.User.Birthday.String() + strconv.FormatInt(userGetResponse.User.Age, 10))

	userUpdateRequest := new(user.UpdateRequest)
	userUpdateRequest.Id = userGetResponse.User.Id
	userUpdateRequest.FirstName = "れい"
	userUpdateRequest.LastName = "おおぞの"
	userUpdateRequest.Birthday = &date.Date{
		Year:  2001,
		Month: 4,
		Day:   18,
	}

	userUpdateResponse, err := userClient.Update(ctx, userUpdateRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(userUpdateResponse.User.Id + userUpdateResponse.User.LastName + userUpdateResponse.User.FirstName + userUpdateResponse.User.Birthday.String() + strconv.FormatInt(userUpdateResponse.User.Age, 10))

	// userTerminateRequest := new(user.TerminateRequest)
	// userTerminateRequest.Id = userUpdateResponse.User.Id

	// userTerminateResponse, err := userClient.Terminate(ctx, userTerminateRequest)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// log.Println(userTerminateResponse.User.Id + userTerminateResponse.User.LastName + userTerminateResponse.User.FirstName + userTerminateResponse.User.Birthday.String() + strconv.FormatInt(userTerminateResponse.User.Age, 10))

	userRegisterRequest = new(user.RegisterRequest)
	userRegisterRequest.FirstName = "麗奈"
	userRegisterRequest.LastName = "守屋"
	userRegisterRequest.Birthday = &date.Date{
		Year:  2000,
		Month: 1,
		Day:   2,
	}

	userRegisterResponse, err = userClient.Register(ctx, userRegisterRequest)
	if err != nil {
		log.Println(err.Error())
	}

	taskClient := task.NewTaskServiceClient(conn)

	TaskTest(ctx, taskClient, userRegisterResponse.User.Id)

}

func TaskTest(ctx context.Context, taskClient task.TaskServiceClient, userId string) {

	fmt.Println("タスク作成1")
	taskCreateRequest := new(task.CreateRequest)
	taskCreateRequest.CreatorId = userId
	taskCreateRequest.Title = "タイトル1"
	taskCreateRequest.Detail = "詳細1"
	taskCreateRequest.Deadline = timestamppb.New(time.Now().Add(6 * time.Hour).UTC())

	taskCreateResponse, err := taskClient.Create(ctx, taskCreateRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskCreateResponse.Task.Id + taskCreateResponse.Task.CreatorId + taskCreateResponse.Task.Title + taskCreateResponse.Task.Detail + taskCreateResponse.Task.Status.String())
	log.Println(taskCreateResponse.Task.CreatedAt.AsTime().String() + taskCreateResponse.Task.UpdatedAt.AsTime().String() + taskCreateResponse.Task.Deadline.AsTime().String())

	fmt.Println("タスク作成2")
	taskCreateRequest2 := new(task.CreateRequest)
	taskCreateRequest2.CreatorId = userId
	taskCreateRequest2.Title = "タイトル2"
	taskCreateRequest2.Detail = "詳細2"

	taskCreateResponse2, err := taskClient.Create(ctx, taskCreateRequest2)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskCreateResponse2.Task.Id + taskCreateResponse2.Task.CreatorId + taskCreateResponse2.Task.Title + taskCreateResponse2.Task.Detail + taskCreateResponse2.Task.Status.String())
	log.Println(taskCreateResponse2.Task.CreatedAt.AsTime().String() + taskCreateResponse2.Task.UpdatedAt.AsTime().String() + taskCreateResponse2.Task.Deadline.AsTime().String())

	fmt.Println("タスク作成3")
	taskCreateRequest3 := new(task.CreateRequest)
	taskCreateRequest3.CreatorId = userId
	taskCreateRequest3.Title = "タイトル3"
	taskCreateRequest3.Deadline = timestamppb.New(time.Now().Add(10 * time.Hour).UTC())

	taskCreateResponse3, err := taskClient.Create(ctx, taskCreateRequest3)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskCreateResponse3.Task.Id + taskCreateResponse3.Task.CreatorId + taskCreateResponse3.Task.Title + taskCreateResponse3.Task.Detail + taskCreateResponse3.Task.Status.String())
	log.Println(taskCreateResponse3.Task.CreatedAt.AsTime().String() + taskCreateResponse3.Task.UpdatedAt.AsTime().String() + taskCreateResponse3.Task.Deadline.AsTime().String())

	fmt.Println("タスク作成4")
	taskCreateRequest4 := new(task.CreateRequest)
	taskCreateRequest4.CreatorId = userId
	taskCreateRequest4.Title = "タイトル4"
	taskCreateRequest4.Detail = "詳細4"

	taskCreateResponse4, err := taskClient.Create(ctx, taskCreateRequest4)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskCreateResponse4.Task.Id + taskCreateResponse4.Task.CreatorId + taskCreateResponse4.Task.Title + taskCreateResponse4.Task.Detail + taskCreateResponse4.Task.Status.String())
	log.Println(taskCreateResponse4.Task.CreatedAt.AsTime().String() + taskCreateResponse4.Task.UpdatedAt.AsTime().String() + taskCreateResponse4.Task.Deadline.AsTime().String())

	fmt.Println("タスク取得")
	taskGetRequest := new(task.GetRequest)
	taskGetRequest.Id = taskCreateResponse.Task.Id

	taskGetResponse, err := taskClient.Get(ctx, taskGetRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskGetResponse.Task.Id + taskGetResponse.Task.CreatorId + taskGetResponse.Task.Title + taskGetResponse.Task.Detail + taskGetResponse.Task.Status.String())
	log.Println(taskGetResponse.Task.CreatedAt.AsTime().String() + taskGetResponse.Task.UpdatedAt.AsTime().String() + taskGetResponse.Task.Deadline.AsTime().String())

	fmt.Println("タスク完了")
	taskDoneRequest := new(task.DoneRequest)
	taskDoneRequest.Id = taskCreateResponse2.Task.Id

	taskDoneResponse, err := taskClient.Done(ctx, taskDoneRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskDoneResponse.Task.Id + taskDoneResponse.Task.CreatorId + taskDoneResponse.Task.Title + taskDoneResponse.Task.Detail + taskDoneResponse.Task.Status.String())
	log.Println(taskDoneResponse.Task.CreatedAt.AsTime().String() + taskDoneResponse.Task.UpdatedAt.AsTime().String() + taskDoneResponse.Task.Deadline.AsTime().String())

	fmt.Println("タスク未完了")
	taskUndoneRequest := new(task.UndoneRequest)
	taskUndoneRequest.Id = taskCreateResponse.Task.Id

	taskUndoneResponse, err := taskClient.Undone(ctx, taskUndoneRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskUndoneResponse.Task.Id + taskUndoneResponse.Task.CreatorId + taskUndoneResponse.Task.Title + taskUndoneResponse.Task.Detail + taskUndoneResponse.Task.Status.String())
	log.Println(taskUndoneResponse.Task.CreatedAt.AsTime().String() + taskUndoneResponse.Task.UpdatedAt.AsTime().String() + taskUndoneResponse.Task.Deadline.AsTime().String())

	fmt.Println("タスク更新")
	taskUpdateRequest := new(task.UpdateRequest)
	taskUpdateRequest.Id = taskCreateResponse.Task.Id
	taskUpdateRequest.Title = "更新"
	taskUpdateRequest.Detail = "更新"
	taskUpdateRequest.Deadline = timestamppb.New(time.Now().Add(10 * time.Hour).UTC())

	taskUpdateResponse, err := taskClient.Update(ctx, taskUpdateRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskUpdateResponse.Task.Id + taskUpdateResponse.Task.CreatorId + taskUpdateResponse.Task.Title + taskUpdateResponse.Task.Detail + taskUpdateResponse.Task.Status.String())
	log.Println(taskUpdateResponse.Task.CreatedAt.AsTime().String() + taskUpdateResponse.Task.UpdatedAt.AsTime().String() + taskUpdateResponse.Task.Deadline.AsTime().String())

	fmt.Println("タスク一覧取得")
	taskListRequest := new(task.ListRequest)
	taskListRequest.UserId = userId

	taskListResponse, err := taskClient.List(ctx, taskListRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(len(taskListResponse.Task))
	for _, task := range taskListResponse.Task {
		log.Println(task.CreatedAt.AsTime().String())
		log.Println(task.UpdatedAt.AsTime().String())
	}

	fmt.Println("タスク削除")
	taskDeleteRequest := new(task.DeleteRequest)
	taskDeleteRequest.Id = taskCreateResponse4.Task.Id

	taskDeleteResponse, err := taskClient.Delete(ctx, taskDeleteRequest)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(taskDeleteResponse.Task.Id)

	fmt.Println("タスク一覧取得")
	taskListRequest2 := new(task.ListRequest)
	taskListRequest2.UserId = userId

	taskListResponse2, err := taskClient.List(ctx, taskListRequest2)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(len(taskListResponse2.Task))
}
