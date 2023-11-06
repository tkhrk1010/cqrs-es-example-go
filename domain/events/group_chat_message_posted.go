package events

import (
	"cqrs-es-example-go/domain/models"
	"fmt"
	esa "github.com/j5ik2o/event-store-adapter-go"
	"github.com/oklog/ulid/v2"
	"time"
)

type GroupChatMessagePosted struct {
	id          string
	aggregateId *models.GroupChatId
	message     *models.Message
	seqNr       uint64
	executorId  *models.UserAccountId
	occurredAt  uint64
}

func NewGroupChatMessagePosted(aggregateId *models.GroupChatId, message *models.Message, seqNr uint64, executorId *models.UserAccountId) *GroupChatMessagePosted {
	id := ulid.Make().String()
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return &GroupChatMessagePosted{id, aggregateId, message, seqNr, executorId, occurredAt}
}

func NewGroupChatMessagePostedFrom(id string, aggregateId *models.GroupChatId, message *models.Message, seqNr uint64, executorId *models.UserAccountId, occurredAt uint64) *GroupChatMessagePosted {
	return &GroupChatMessagePosted{id, aggregateId, message, seqNr, executorId, occurredAt}
}

func (g *GroupChatMessagePosted) GetId() string {
	return g.id
}

func (g *GroupChatMessagePosted) GetTypeName() string {
	return "GroupChatMessagePosted"
}

func (g *GroupChatMessagePosted) GetAggregateId() esa.AggregateId {
	return g.aggregateId
}

func (g *GroupChatMessagePosted) GetSeqNr() uint64 {
	return g.seqNr
}

func (g *GroupChatMessagePosted) GetMessage() *models.Message {
	return g.message
}

func (g *GroupChatMessagePosted) GetExecutorId() *models.UserAccountId {
	return g.executorId
}

func (g *GroupChatMessagePosted) IsCreated() bool {
	return false
}

func (g *GroupChatMessagePosted) GetOccurredAt() uint64 {
	return g.occurredAt
}

func (g *GroupChatMessagePosted) String() string {
	return fmt.Sprintf("%s{ id: %s, aggregateId: %s seqNr: %d, occurredAt: %d}",
		g.GetTypeName(), g.id, g.aggregateId, g.seqNr, g.occurredAt)
}
