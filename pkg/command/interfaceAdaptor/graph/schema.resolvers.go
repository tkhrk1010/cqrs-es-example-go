package commandgraph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"cqrs-es-example-go/pkg/command/domain/models"
	writeapi "cqrs-es-example-go/pkg/command/interfaceAdaptor/graph/model"
	"cqrs-es-example-go/pkg/command/interfaceAdaptor/validator"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

// CreateGroupChat is the resolver for the createGroupChat field.
func (r *mutationRootResolver) CreateGroupChat(ctx context.Context, input writeapi.CreateGroupChatInput) (*writeapi.GroupChatResult, error) {
	var errors []error

	groupChatName, err := validator.ValidateGroupChatName(input.Name).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.CreateGroupChat(groupChatName, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.GroupChatResult{GroupChatID: event.GetAggregateId().AsString()}, nil
}

// DeleteGroupChat is the resolver for the deleteGroupChat field.
func (r *mutationRootResolver) DeleteGroupChat(ctx context.Context, input writeapi.DeleteGroupChatInput) (*writeapi.GroupChatResult, error) {
	var errors []error

	groupChatId, err := validator.ValidateGroupChatId(input.GroupChatID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.DeleteGroupChat(&groupChatId, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.GroupChatResult{GroupChatID: event.GetAggregateId().AsString()}, nil
}

// RenameGroupChat is the resolver for the renameGroupChat field.
func (r *mutationRootResolver) RenameGroupChat(ctx context.Context, input writeapi.RenameGroupChatInput) (*writeapi.GroupChatResult, error) {
	var errors []error

	groupChatId, err := validator.ValidateGroupChatId(input.GroupChatID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	groupChatName, err := validator.ValidateGroupChatName(input.Name).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.RenameGroupChat(&groupChatId, groupChatName, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.GroupChatResult{GroupChatID: event.GetAggregateId().AsString()}, nil
}

// AddMember is the resolver for the addMember field.
func (r *mutationRootResolver) AddMember(ctx context.Context, input writeapi.AddMemberInput) (*writeapi.GroupChatResult, error) {
	var errors []error

	groupChatId, err := validator.ValidateGroupChatId(input.GroupChatID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	accountId, err := validator.ValidateUserAccountId(input.UserAccountID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	role, err := models.StringToRole(input.Role.String())
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.AddMember(&groupChatId, accountId, role, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.GroupChatResult{GroupChatID: event.GetAggregateId().AsString()}, nil
}

// RemoveMember is the resolver for the removeMember field.
func (r *mutationRootResolver) RemoveMember(ctx context.Context, input writeapi.RemoveMemberInput) (*writeapi.GroupChatResult, error) {
	var errors []error

	groupChatId, err := validator.ValidateGroupChatId(input.GroupChatID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	userAccountId, err := validator.ValidateUserAccountId(input.UserAccountID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.RemoveMember(&groupChatId, userAccountId, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.GroupChatResult{GroupChatID: event.GetAggregateId().AsString()}, nil
}

// PostMessage is the resolver for the postMessage field.
func (r *mutationRootResolver) PostMessage(ctx context.Context, input writeapi.PostMessageInput) (*writeapi.MessageResult, error) {
	var errors []error

	groupChatId, err := validator.ValidateGroupChatId(input.GroupChatID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	messageId := models.NewMessageId()

	senderId, err := validator.ValidateUserAccountId(input.UserAccountID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	message, err := validator.ValidateMessage(messageId, input.Content, senderId).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.PostMessage(&groupChatId, message, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.MessageResult{GroupChatID: event.GetAggregateId().AsString(), MessageID: messageId.String()}, nil
}

// DeleteMessage is the resolver for the deleteMessage field.
func (r *mutationRootResolver) DeleteMessage(ctx context.Context, input writeapi.DeleteMessageInput) (*writeapi.GroupChatResult, error) {
	var errors []error

	groupChatId, err := validator.ValidateGroupChatId(input.GroupChatID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	messageId, err := validator.ValidateMessageId(input.MessageID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	executorId, err := validator.ValidateUserAccountId(input.ExecutorID).Get()
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			graphql.AddError(ctx, err)
		}
		return nil, nil
	}

	event, err := r.groupChatCommandProcessor.DeleteMessage(&groupChatId, messageId, executorId).Get()
	if err != nil {
		graphql.AddError(ctx, err)
		return nil, nil
	}

	return &writeapi.GroupChatResult{GroupChatID: event.GetAggregateId().AsString()}, nil
}

// HealthCheck is the resolver for the healthCheck field.
func (r *queryRootResolver) HealthCheck(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: HealthCheck - healthCheck"))
}

// MutationRoot returns MutationRootResolver implementation.
func (r *Resolver) MutationRoot() MutationRootResolver { return &mutationRootResolver{r} }

// QueryRoot returns QueryRootResolver implementation.
func (r *Resolver) QueryRoot() QueryRootResolver { return &queryRootResolver{r} }

type mutationRootResolver struct{ *Resolver }
type queryRootResolver struct{ *Resolver }
