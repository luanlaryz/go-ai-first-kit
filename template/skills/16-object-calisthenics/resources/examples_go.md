# Object Calisthenics in Go: Before/After

## 1) One indentation level + no else

Before:
```go
func (u *SendMessageUseCase) Execute(ctx context.Context, in SendMessageInput) error {
    if in.{{DOMAIN_ACTOR_TITLE}}ID != "" {
        if in.ConversationID != "" {
            if in.Body != "" {
                return u.repo.Save(ctx, in)
            } else {
                return ErrEmptyBody
            }
        } else {
            return ErrInvalidConversationID
        }
    } else {
        return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
    }
}
```

After:
```go
func (u *SendMessageUseCase) Execute(ctx context.Context, in SendMessageInput) error {
    if in.{{DOMAIN_ACTOR_TITLE}}ID == "" {
        return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
    }
    if in.ConversationID == "" {
        return ErrInvalidConversationID
    }
    if in.Body == "" {
        return ErrEmptyBody
    }
    return u.repo.Save(ctx, in)
}
```

## 2) Strong types over primitives

Before:
```go
func BuildStreamKey({{DOMAIN_ACTOR}}ID, conversationID string) string {
    return {{DOMAIN_ACTOR}}ID + ":" + conversationID
}
```

After:
```go
type {{DOMAIN_ACTOR_TITLE}}ID string
type ConversationID string

func BuildStreamKey({{DOMAIN_ACTOR}}ID {{DOMAIN_ACTOR_TITLE}}ID, conversationID ConversationID) string {
    return string({{DOMAIN_ACTOR}}ID) + ":" + string(conversationID)
}
```

## 3) Limit arguments with input struct

Before:
```go
func NewCommand({{DOMAIN_ACTOR}}ID, requestID, conversationID, body, channel, locale string) Command
```

After:
```go
type NewCommandInput struct {
    {{DOMAIN_ACTOR_TITLE}}ID       {{DOMAIN_ACTOR_TITLE}}ID
    RequestID      RequestID
    ConversationID ConversationID
    Body           string
    Channel        string
    Locale         string
}

func NewCommand(in NewCommandInput) Command
```

## 4) Encapsulate collections with invariants

Before:
```go
type Conversation struct {
    Messages []Message
}
```

After:
```go
type Messages struct {
    items []Message
}

func (m *Messages) Append(msg Message) error {
    if msg.Seq <= 0 {
        return ErrInvalidSeq
    }
    m.items = append(m.items, msg)
    return nil
}

type Conversation struct {
    Messages Messages
}
```
