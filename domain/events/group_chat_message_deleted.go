package events

import (
	"cqrs-es-example-go/domain/models"
	"fmt"
	esa "github.com/j5ik2o/event-store-adapter-go"
	"github.com/oklog/ulid/v2"
	"time"
)

type GroupChatMessageDeleted struct {
	id          string
	aggregateId *models.GroupChatId
	messageId   *models.MessageId
	seqNr       uint64
	executorId  *models.UserAccountId
	occurredAt  uint64
}

func NewGroupChatMessageDeleted(aggregateId *models.GroupChatId, messageId *models.MessageId, seqNr uint64, executorId *models.UserAccountId) *GroupChatMessageDeleted {
	id := ulid.Make().String()
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return &GroupChatMessageDeleted{id, aggregateId, messageId, seqNr, executorId, occurredAt}
}

func NewGroupChatMessageDeletedFrom(id string, aggregateId *models.GroupChatId, messageId *models.MessageId, seqNr uint64, executorId *models.UserAccountId, occurredAt uint64) *GroupChatMessageDeleted {
	return &GroupChatMessageDeleted{id, aggregateId, messageId, seqNr, executorId, occurredAt}
}

func (g *GroupChatMessageDeleted) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"Id":          g.id,
		"AggregateId": g.aggregateId.ToJSON(),
		"MessageId":   g.messageId.ToJSON(),
		"SeqNr":       g.seqNr,
		"ExecutorId":  g.executorId.ToJSON(),
		"OccurredAt":  g.occurredAt,
	}
}

func (g *GroupChatMessageDeleted) GetId() string {
	return g.id
}

func (g *GroupChatMessageDeleted) GetTypeName() string {
	return "GroupChatMessageDeleted"
}

func (g *GroupChatMessageDeleted) GetAggregateId() esa.AggregateId {
	return g.aggregateId
}

func (g *GroupChatMessageDeleted) GetSeqNr() uint64 {
	return g.seqNr
}

func (g *GroupChatMessageDeleted) GetMessageId() *models.MessageId {
	return g.messageId
}

func (g *GroupChatMessageDeleted) GetExecutorId() *models.UserAccountId {
	return g.executorId
}

func (g *GroupChatMessageDeleted) IsCreated() bool {
	return false
}

func (g *GroupChatMessageDeleted) GetOccurredAt() uint64 {
	return g.occurredAt
}

func (g *GroupChatMessageDeleted) String() string {
	return fmt.Sprintf("%s{ id: %s, aggregateId: %s seqNr: %d, occurredAt: %d}",
		g.GetTypeName(), g.id, g.aggregateId, g.seqNr, g.occurredAt)
}
