package events

import (
	"cqrs-es-example-go/domain"
	"fmt"
	esa "github.com/j5ik2o/event-store-adapter-go"
	"time"
)

type GroupChatCreated struct {
	id          string
	aggregateId domain.GroupChatId
	name        string
	seqNr       uint64
	occurredAt  uint64
}

func NewGroupChatCreated(id string, aggregateId domain.GroupChatId, name string, seqNr uint64) *GroupChatCreated {
	now := time.Now()
	occurredAt := uint64(now.UnixNano() / 1e6)
	return &GroupChatCreated{id, aggregateId, name, seqNr, occurredAt}
}

func (g *GroupChatCreated) GetId() string {
	return g.id
}

func (g *GroupChatCreated) GetTypeName() string {
	return "group-chat-created"
}

func (g *GroupChatCreated) GetAggregateId() esa.AggregateId {
	return &g.aggregateId
}

func (g *GroupChatCreated) GetName() string {
	return g.name
}

func (g *GroupChatCreated) GetSeqNr() uint64 {
	return g.seqNr
}

func (g *GroupChatCreated) IsCreated() bool {
	return true
}

func (g *GroupChatCreated) GetOccurredAt() uint64 {
	return g.occurredAt
}

func (g *GroupChatCreated) String() string {
	return fmt.Sprintf("%s{ id: %s, aggregateId: %s name: %s, seqNr: %d, occurredAt: %d}",
		g.GetTypeName(), g.id, g.aggregateId, g.name, g.seqNr, g.occurredAt)
}
