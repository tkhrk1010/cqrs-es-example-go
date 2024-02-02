package repository

import (
	"cqrs-es-example-go/pkg/command/domain/events"
	"cqrs-es-example-go/pkg/command/domain/models"
	"fmt"
	esa "github.com/j5ik2o/event-store-adapter-go"
)

func EventConverter(m map[string]interface{}) (esa.Event, error) {
	eventId := m["id"].(string)
	groupChatId, err := models.ConvertGroupChatIdFromJSON(m["aggregate_id"].(map[string]interface{})).Get()
	if err != nil {
		return nil, err
	}
	groupChatName, err := models.ConvertGroupChatNameFromJSON(m["name"].(map[string]interface{})).Get()
	if err != nil {
		return nil, err
	}
	members := models.ConvertMembersFromJSON(m["members"].(map[string]interface{}))
	executorId, err := models.ConvertUserAccountIdFromJSON(m["executor_id"].(map[string]interface{})).Get()
	if err != nil {
		return nil, err
	}
	seqNr := uint64(m["seq_nr"].(float64))
	occurredAt := uint64(m["occurred_at"].(float64))
	switch m["type_name"].(string) {
	case "GroupChatCreated":
		event := events.NewGroupChatCreatedFrom(
			eventId,
			groupChatId,
			groupChatName,
			members,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	case "GroupChatDeleted":
		event := events.NewGroupChatDeletedFrom(
			eventId,
			groupChatId,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	case "GroupChatRenamed":
		name, err := models.NewGroupChatName(m["name"].(string)).Get()
		if err != nil {
			return nil, err
		}
		event := events.NewGroupChatRenamedFrom(
			eventId,
			groupChatId,
			name,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	case "GroupChatMemberAdded":
		memberObj := m["member"].(map[string]interface{})
		memberId, err := models.ConvertMemberIdFromJSON(memberObj["member_id"].(map[string]interface{})).Get()
		if err != nil {
			return nil, err
		}
		userAccountId, err := models.ConvertUserAccountIdFromJSON(memberObj["user_account_id"].(map[string]interface{})).Get()
		if err != nil {
			return nil, err
		}
		role := models.Role(memberObj["role"].(int))
		member := models.NewMember(memberId, userAccountId, role)
		event := events.NewGroupChatMemberAddedFrom(
			eventId,
			groupChatId,
			member,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	case "GroupChatMemberRemoved":
		userAccountId, err := models.ConvertUserAccountIdFromJSON(m["user_account_id"].(map[string]interface{})).Get()
		if err != nil {
			return nil, err
		}
		event := events.NewGroupChatMemberRemovedFrom(
			eventId,
			groupChatId,
			userAccountId,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	case "GroupChatMessagePosted":
		message, err := models.ConvertMessageFromJSON(m["message"].(map[string]interface{})).Get()
		if err != nil {
			return nil, err
		}
		event := events.NewGroupChatMessagePostedFrom(
			eventId,
			groupChatId,
			message,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	case "GroupChatMessageDeleted":
		messageId := models.ConvertMessageIdFromJSON(m["message_id"].(map[string]interface{}))
		event := events.NewGroupChatMessageDeletedFrom(
			eventId,
			groupChatId,
			messageId,
			seqNr,
			executorId,
			occurredAt,
		)
		return &event, nil
	default:
		return nil, fmt.Errorf("unknown event type: %s", m["type_name"].(string))
	}
}
