package api

import (
	"cqrs-es-example-go/pkg/command/domain/models"
	"cqrs-es-example-go/pkg/command/interfaceAdaptor/repository"
	"cqrs-es-example-go/pkg/command/interfaceAdaptor/validator"
	"cqrs-es-example-go/pkg/command/useCase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateGroupChatRequestBody struct {
	Name       string `json:"name"`
	ExecutorId string `json:"executor_id"`
}

type CreateGroupChatResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

// ---

type DeleteGroupChatRequestBody struct {
	GroupChatId string `json:"group_chat_id"`
	ExecutorId  string `json:"executor_id"`
}

type DeleteGroupChatResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

// ---

type RenameGroupChatRequestBody struct {
	GroupChatId string `json:"group_chat_id"`
	Name        string `json:"name"`
	ExecutorId  string `json:"executor_id"`
}

type RenameGroupChatResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

// ---

type AddMemberRequestBody struct {
	GroupChatId   string `json:"group_chat_id"`
	UserAccountId string `json:"user_account_id"`
	Role          string `json:"role"`
	ExecutorId    string `json:"executor_id"`
}

type AddMemberResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

// ---

type RemoveMemberRequestBody struct {
	GroupChatId   string `json:"group_chat_id"`
	UserAccountId string `json:"user_account_id"`
	ExecutorId    string `json:"executor_id"`
}

type RemoveMemberResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

// ---

type PostMessageRequestBody struct {
	GroupChatId   string `json:"group_chat_id"`
	Message       string `json:"message"`
	UserAccountId string `json:"user_account_id"`
	ExecutorId    string `json:"executor_id"`
}

type PostMessageResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
	MessageId   string `json:"message_id"`
}

// ---

type DeleteMessageRequestBody struct {
	GroupChatId   string `json:"group_chat_id"`
	MessageId     string `json:"message_id"`
	UserAccountId string `json:"user_account_id"`
	ExecutorId    string `json:"executor_id"`
}

type DeleteMessageResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

// ---

type GroupChatResponseErrorBody struct {
	Message string `json:"message"`
}

// ---

type GroupChatController struct {
	repository repository.GroupChatRepository
}

// NewGroupChatController は GroupChatController を生成します。
func NewGroupChatController(repository repository.GroupChatRepository) *GroupChatController {
	return &GroupChatController{
		repository,
	}
}

// CreateGroupChat はグループチャットを作成します。
func (g *GroupChatController) CreateGroupChat(c *gin.Context) {
	var jsonRequestBody CreateGroupChatRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatName, err := validator.ValidateGroupChatName(jsonRequestBody.Name).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.CreateGroupChat(groupChatName, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := CreateGroupChatResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

func (g *GroupChatController) DeleteGroupChat(c *gin.Context) {
	var jsonRequestBody DeleteGroupChatRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.DeleteGroupChat(&groupChatId, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := DeleteGroupChatResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

// RenameGroupChat はグループチャットをリネームします。
func (g *GroupChatController) RenameGroupChat(c *gin.Context) {
	var jsonRequestBody RenameGroupChatRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatName, err := validator.ValidateGroupChatName(jsonRequestBody.Name).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.RenameGroupChat(&groupChatId, groupChatName, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := RenameGroupChatResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

// AddMember はグループチャットにメンバーを追加します。
func (g *GroupChatController) AddMember(c *gin.Context) {
	var jsonRequestBody AddMemberRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	accountId, err := validator.ValidateUserAccountId(jsonRequestBody.UserAccountId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	role := models.StringToRole(jsonRequestBody.Role)

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.AddMember(&groupChatId, accountId, role, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := AddMemberResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

// RemoveMember はグループチャットからメンバーを削除します。
func (g *GroupChatController) RemoveMember(c *gin.Context) {
	var jsonRequestBody RemoveMemberRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	accountId, err := validator.ValidateUserAccountId(jsonRequestBody.UserAccountId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.RemoveMember(&groupChatId, accountId, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := RemoveMemberResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

// PostMessage はグループチャットにメッセージを投稿します。
func (g *GroupChatController) PostMessage(c *gin.Context) {
	var jsonRequestBody PostMessageRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	messageId := models.NewMessageId()

	senderId, err := validator.ValidateUserAccountId(jsonRequestBody.UserAccountId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	message, err := validator.ValidateMessage(messageId, jsonRequestBody.Message, senderId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.PostMessage(&groupChatId, message, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := PostMessageResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString(), MessageId: messageId.String()}
	c.JSON(http.StatusOK, response)
}

func (g *GroupChatController) DeleteMessage(c *gin.Context) {
	var jsonRequestBody DeleteMessageRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	messageId, err := validator.ValidateMessageId(jsonRequestBody.MessageId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		handleClientError(c, http.StatusBadRequest, err)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.DeleteMessage(&groupChatId, messageId, executorId)

	if err != nil {
		response := GroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := DeleteMessageResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

// handleClientError はクライアントエラーを処理します。
func handleClientError(c *gin.Context, statusCode int, err error) {
	response := GroupChatResponseErrorBody{Message: err.Error()}
	c.JSON(statusCode, response)
}
