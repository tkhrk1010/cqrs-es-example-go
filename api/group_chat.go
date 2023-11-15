package api

import (
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

type RenameGroupChatRequestBody struct {
	GroupChatId string `json:"group_chat_id"`
	Name        string `json:"name"`
	ExecutorId  string `json:"executor_id"`
}

type RenameGroupChatResponseSuccessBody struct {
	GroupChatId string `json:"group_chat_id"`
}

type CreateGroupChatResponseErrorBody struct {
	Message string `json:"message"`
}

type GroupChatController struct {
	repository repository.GroupChatRepository
}

func NewGroupChatController(repository repository.GroupChatRepository) *GroupChatController {
	return &GroupChatController{
		repository,
	}
}

func (g *GroupChatController) CreateGroupChat(c *gin.Context) {
	var jsonRequestBody CreateGroupChatRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	groupChatName, err := validator.ValidateGroupChatName(jsonRequestBody.Name).Get()
	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.CreateGroupChat(groupChatName, executorId, executorId)

	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := CreateGroupChatResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}

func (g *GroupChatController) RenameGroupChat(c *gin.Context) {
	var jsonRequestBody RenameGroupChatRequestBody

	if err := c.ShouldBindJSON(&jsonRequestBody); err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	groupChatId, err := validator.ValidateGroupChatId(jsonRequestBody.GroupChatId).Get()
	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	groupChatName, err := validator.ValidateGroupChatName(jsonRequestBody.Name).Get()
	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	executorId, err := validator.ValidateUserAccountId(jsonRequestBody.ExecutorId).Get()
	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	commandProcessor := useCase.NewGroupChatCommandProcessor(g.repository)
	event, err := commandProcessor.RenameGroupChat(groupChatId, groupChatName, executorId)

	if err != nil {
		response := CreateGroupChatResponseErrorBody{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := CreateGroupChatResponseSuccessBody{GroupChatId: event.GetAggregateId().AsString()}
	c.JSON(http.StatusOK, response)
}
