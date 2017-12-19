// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/chat_ui.avdl

package chat1

import (
	"errors"
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type UIPagination struct {
	Next     string `codec:"next" json:"next"`
	Previous string `codec:"previous" json:"previous"`
	Num      int    `codec:"num" json:"num"`
	Last     bool   `codec:"last" json:"last"`
}

func (o UIPagination) DeepCopy() UIPagination {
	return UIPagination{
		Next:     o.Next,
		Previous: o.Previous,
		Num:      o.Num,
		Last:     o.Last,
	}
}

type UIMessageValid struct {
	MessageID             MessageID      `codec:"messageID" json:"messageID"`
	Ctime                 gregor1.Time   `codec:"ctime" json:"ctime"`
	OutboxID              *string        `codec:"outboxID,omitempty" json:"outboxID,omitempty"`
	MessageBody           MessageBody    `codec:"messageBody" json:"messageBody"`
	SenderUsername        string         `codec:"senderUsername" json:"senderUsername"`
	SenderDeviceName      string         `codec:"senderDeviceName" json:"senderDeviceName"`
	SenderDeviceType      string         `codec:"senderDeviceType" json:"senderDeviceType"`
	Superseded            bool           `codec:"superseded" json:"superseded"`
	SenderDeviceRevokedAt *gregor1.Time  `codec:"senderDeviceRevokedAt,omitempty" json:"senderDeviceRevokedAt,omitempty"`
	AtMentions            []string       `codec:"atMentions" json:"atMentions"`
	ChannelMention        ChannelMention `codec:"channelMention" json:"channelMention"`
	ChannelNameMentions   []string       `codec:"channelNameMentions" json:"channelNameMentions"`
}

func (o UIMessageValid) DeepCopy() UIMessageValid {
	return UIMessageValid{
		MessageID: o.MessageID.DeepCopy(),
		Ctime:     o.Ctime.DeepCopy(),
		OutboxID: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.OutboxID),
		MessageBody:      o.MessageBody.DeepCopy(),
		SenderUsername:   o.SenderUsername,
		SenderDeviceName: o.SenderDeviceName,
		SenderDeviceType: o.SenderDeviceType,
		Superseded:       o.Superseded,
		SenderDeviceRevokedAt: (func(x *gregor1.Time) *gregor1.Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SenderDeviceRevokedAt),
		AtMentions: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.AtMentions),
		ChannelMention: o.ChannelMention.DeepCopy(),
		ChannelNameMentions: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ChannelNameMentions),
	}
}

type UIMessageOutbox struct {
	State       OutboxState  `codec:"state" json:"state"`
	OutboxID    string       `codec:"outboxID" json:"outboxID"`
	MessageType MessageType  `codec:"messageType" json:"messageType"`
	Body        string       `codec:"body" json:"body"`
	Ctime       gregor1.Time `codec:"ctime" json:"ctime"`
	Ordinal     float64      `codec:"ordinal" json:"ordinal"`
}

func (o UIMessageOutbox) DeepCopy() UIMessageOutbox {
	return UIMessageOutbox{
		State:       o.State.DeepCopy(),
		OutboxID:    o.OutboxID,
		MessageType: o.MessageType.DeepCopy(),
		Body:        o.Body,
		Ctime:       o.Ctime.DeepCopy(),
		Ordinal:     o.Ordinal,
	}
}

type MessageUnboxedState int

const (
	MessageUnboxedState_VALID       MessageUnboxedState = 1
	MessageUnboxedState_ERROR       MessageUnboxedState = 2
	MessageUnboxedState_OUTBOX      MessageUnboxedState = 3
	MessageUnboxedState_PLACEHOLDER MessageUnboxedState = 4
)

func (o MessageUnboxedState) DeepCopy() MessageUnboxedState { return o }

var MessageUnboxedStateMap = map[string]MessageUnboxedState{
	"VALID":       1,
	"ERROR":       2,
	"OUTBOX":      3,
	"PLACEHOLDER": 4,
}

var MessageUnboxedStateRevMap = map[MessageUnboxedState]string{
	1: "VALID",
	2: "ERROR",
	3: "OUTBOX",
	4: "PLACEHOLDER",
}

func (e MessageUnboxedState) String() string {
	if v, ok := MessageUnboxedStateRevMap[e]; ok {
		return v
	}
	return ""
}

type UIMessage struct {
	State__       MessageUnboxedState        `codec:"state" json:"state"`
	Valid__       *UIMessageValid            `codec:"valid,omitempty" json:"valid,omitempty"`
	Error__       *MessageUnboxedError       `codec:"error,omitempty" json:"error,omitempty"`
	Outbox__      *UIMessageOutbox           `codec:"outbox,omitempty" json:"outbox,omitempty"`
	Placeholder__ *MessageUnboxedPlaceholder `codec:"placeholder,omitempty" json:"placeholder,omitempty"`
}

func (o *UIMessage) State() (ret MessageUnboxedState, err error) {
	switch o.State__ {
	case MessageUnboxedState_VALID:
		if o.Valid__ == nil {
			err = errors.New("unexpected nil value for Valid__")
			return ret, err
		}
	case MessageUnboxedState_ERROR:
		if o.Error__ == nil {
			err = errors.New("unexpected nil value for Error__")
			return ret, err
		}
	case MessageUnboxedState_OUTBOX:
		if o.Outbox__ == nil {
			err = errors.New("unexpected nil value for Outbox__")
			return ret, err
		}
	case MessageUnboxedState_PLACEHOLDER:
		if o.Placeholder__ == nil {
			err = errors.New("unexpected nil value for Placeholder__")
			return ret, err
		}
	}
	return o.State__, nil
}

func (o UIMessage) Valid() (res UIMessageValid) {
	if o.State__ != MessageUnboxedState_VALID {
		panic("wrong case accessed")
	}
	if o.Valid__ == nil {
		return
	}
	return *o.Valid__
}

func (o UIMessage) Error() (res MessageUnboxedError) {
	if o.State__ != MessageUnboxedState_ERROR {
		panic("wrong case accessed")
	}
	if o.Error__ == nil {
		return
	}
	return *o.Error__
}

func (o UIMessage) Outbox() (res UIMessageOutbox) {
	if o.State__ != MessageUnboxedState_OUTBOX {
		panic("wrong case accessed")
	}
	if o.Outbox__ == nil {
		return
	}
	return *o.Outbox__
}

func (o UIMessage) Placeholder() (res MessageUnboxedPlaceholder) {
	if o.State__ != MessageUnboxedState_PLACEHOLDER {
		panic("wrong case accessed")
	}
	if o.Placeholder__ == nil {
		return
	}
	return *o.Placeholder__
}

func NewUIMessageWithValid(v UIMessageValid) UIMessage {
	return UIMessage{
		State__: MessageUnboxedState_VALID,
		Valid__: &v,
	}
}

func NewUIMessageWithError(v MessageUnboxedError) UIMessage {
	return UIMessage{
		State__: MessageUnboxedState_ERROR,
		Error__: &v,
	}
}

func NewUIMessageWithOutbox(v UIMessageOutbox) UIMessage {
	return UIMessage{
		State__:  MessageUnboxedState_OUTBOX,
		Outbox__: &v,
	}
}

func NewUIMessageWithPlaceholder(v MessageUnboxedPlaceholder) UIMessage {
	return UIMessage{
		State__:       MessageUnboxedState_PLACEHOLDER,
		Placeholder__: &v,
	}
}

func (o UIMessage) DeepCopy() UIMessage {
	return UIMessage{
		State__: o.State__.DeepCopy(),
		Valid__: (func(x *UIMessageValid) *UIMessageValid {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Valid__),
		Error__: (func(x *MessageUnboxedError) *MessageUnboxedError {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Error__),
		Outbox__: (func(x *UIMessageOutbox) *UIMessageOutbox {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Outbox__),
		Placeholder__: (func(x *MessageUnboxedPlaceholder) *MessageUnboxedPlaceholder {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Placeholder__),
	}
}

type UIMessages struct {
	Messages   []UIMessage   `codec:"messages" json:"messages"`
	Pagination *UIPagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

func (o UIMessages) DeepCopy() UIMessages {
	return UIMessages{
		Messages: (func(x []UIMessage) []UIMessage {
			if x == nil {
				return nil
			}
			var ret []UIMessage
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Messages),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
	}
}

type UnverifiedInboxUIItemMetadata struct {
	ChannelName       string   `codec:"channelName" json:"channelName"`
	Headline          string   `codec:"headline" json:"headline"`
	Snippet           string   `codec:"snippet" json:"snippet"`
	WriterNames       []string `codec:"writerNames" json:"writerNames"`
	ResetParticipants []string `codec:"resetParticipants" json:"resetParticipants"`
}

func (o UnverifiedInboxUIItemMetadata) DeepCopy() UnverifiedInboxUIItemMetadata {
	return UnverifiedInboxUIItemMetadata{
		ChannelName: o.ChannelName,
		Headline:    o.Headline,
		Snippet:     o.Snippet,
		WriterNames: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.WriterNames),
		ResetParticipants: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ResetParticipants),
	}
}

type UnverifiedInboxUIItem struct {
	ConvID        string                         `codec:"convID" json:"convID"`
	Name          string                         `codec:"name" json:"name"`
	Visibility    keybase1.TLFVisibility         `codec:"visibility" json:"visibility"`
	Status        ConversationStatus             `codec:"status" json:"status"`
	MembersType   ConversationMembersType        `codec:"membersType" json:"membersType"`
	MemberStatus  ConversationMemberStatus       `codec:"memberStatus" json:"memberStatus"`
	TeamType      TeamType                       `codec:"teamType" json:"teamType"`
	Notifications *ConversationNotificationInfo  `codec:"notifications,omitempty" json:"notifications,omitempty"`
	Time          gregor1.Time                   `codec:"time" json:"time"`
	Version       ConversationVers               `codec:"version" json:"version"`
	MaxMsgID      MessageID                      `codec:"maxMsgID" json:"maxMsgID"`
	LocalMetadata *UnverifiedInboxUIItemMetadata `codec:"localMetadata,omitempty" json:"localMetadata,omitempty"`
}

func (o UnverifiedInboxUIItem) DeepCopy() UnverifiedInboxUIItem {
	return UnverifiedInboxUIItem{
		ConvID:       o.ConvID,
		Name:         o.Name,
		Visibility:   o.Visibility.DeepCopy(),
		Status:       o.Status.DeepCopy(),
		MembersType:  o.MembersType.DeepCopy(),
		MemberStatus: o.MemberStatus.DeepCopy(),
		TeamType:     o.TeamType.DeepCopy(),
		Notifications: (func(x *ConversationNotificationInfo) *ConversationNotificationInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Notifications),
		Time:     o.Time.DeepCopy(),
		Version:  o.Version.DeepCopy(),
		MaxMsgID: o.MaxMsgID.DeepCopy(),
		LocalMetadata: (func(x *UnverifiedInboxUIItemMetadata) *UnverifiedInboxUIItemMetadata {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.LocalMetadata),
	}
}

type UnverifiedInboxUIItems struct {
	Items      []UnverifiedInboxUIItem `codec:"items" json:"items"`
	Pagination *UIPagination           `codec:"pagination,omitempty" json:"pagination,omitempty"`
	Offline    bool                    `codec:"offline" json:"offline"`
}

func (o UnverifiedInboxUIItems) DeepCopy() UnverifiedInboxUIItems {
	return UnverifiedInboxUIItems{
		Items: (func(x []UnverifiedInboxUIItem) []UnverifiedInboxUIItem {
			if x == nil {
				return nil
			}
			var ret []UnverifiedInboxUIItem
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Items),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
		Offline: o.Offline,
	}
}

type InboxUIItem struct {
	ConvID            string                        `codec:"convID" json:"convID"`
	IsEmpty           bool                          `codec:"isEmpty" json:"isEmpty"`
	Name              string                        `codec:"name" json:"name"`
	Snippet           string                        `codec:"snippet" json:"snippet"`
	Channel           string                        `codec:"channel" json:"channel"`
	Headline          string                        `codec:"headline" json:"headline"`
	SnippetMessage    *UIMessage                    `codec:"snippetMessage,omitempty" json:"snippetMessage,omitempty"`
	Visibility        keybase1.TLFVisibility        `codec:"visibility" json:"visibility"`
	Participants      []string                      `codec:"participants" json:"participants"`
	FullNames         map[string]string             `codec:"fullNames" json:"fullNames"`
	ResetParticipants []string                      `codec:"resetParticipants" json:"resetParticipants"`
	Status            ConversationStatus            `codec:"status" json:"status"`
	MembersType       ConversationMembersType       `codec:"membersType" json:"membersType"`
	MemberStatus      ConversationMemberStatus      `codec:"memberStatus" json:"memberStatus"`
	TeamType          TeamType                      `codec:"teamType" json:"teamType"`
	Time              gregor1.Time                  `codec:"time" json:"time"`
	Notifications     *ConversationNotificationInfo `codec:"notifications,omitempty" json:"notifications,omitempty"`
	CreatorInfo       *ConversationCreatorInfoLocal `codec:"creatorInfo,omitempty" json:"creatorInfo,omitempty"`
	Version           ConversationVers              `codec:"version" json:"version"`
	MaxMsgID          MessageID                     `codec:"maxMsgID" json:"maxMsgID"`
	FinalizeInfo      *ConversationFinalizeInfo     `codec:"finalizeInfo,omitempty" json:"finalizeInfo,omitempty"`
	Supersedes        []ConversationMetadata        `codec:"supersedes" json:"supersedes"`
	SupersededBy      []ConversationMetadata        `codec:"supersededBy" json:"supersededBy"`
}

func (o InboxUIItem) DeepCopy() InboxUIItem {
	return InboxUIItem{
		ConvID:   o.ConvID,
		IsEmpty:  o.IsEmpty,
		Name:     o.Name,
		Snippet:  o.Snippet,
		Channel:  o.Channel,
		Headline: o.Headline,
		SnippetMessage: (func(x *UIMessage) *UIMessage {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SnippetMessage),
		Visibility: o.Visibility.DeepCopy(),
		Participants: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Participants),
		FullNames: (func(x map[string]string) map[string]string {
			if x == nil {
				return nil
			}
			ret := make(map[string]string)
			for k, v := range x {
				kCopy := k
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.FullNames),
		ResetParticipants: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ResetParticipants),
		Status:       o.Status.DeepCopy(),
		MembersType:  o.MembersType.DeepCopy(),
		MemberStatus: o.MemberStatus.DeepCopy(),
		TeamType:     o.TeamType.DeepCopy(),
		Time:         o.Time.DeepCopy(),
		Notifications: (func(x *ConversationNotificationInfo) *ConversationNotificationInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Notifications),
		CreatorInfo: (func(x *ConversationCreatorInfoLocal) *ConversationCreatorInfoLocal {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.CreatorInfo),
		Version:  o.Version.DeepCopy(),
		MaxMsgID: o.MaxMsgID.DeepCopy(),
		FinalizeInfo: (func(x *ConversationFinalizeInfo) *ConversationFinalizeInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.FinalizeInfo),
		Supersedes: (func(x []ConversationMetadata) []ConversationMetadata {
			if x == nil {
				return nil
			}
			var ret []ConversationMetadata
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Supersedes),
		SupersededBy: (func(x []ConversationMetadata) []ConversationMetadata {
			if x == nil {
				return nil
			}
			var ret []ConversationMetadata
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.SupersededBy),
	}
}

type InboxUIItems struct {
	Items      []InboxUIItem `codec:"items" json:"items"`
	Pagination *UIPagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
	Offline    bool          `codec:"offline" json:"offline"`
}

func (o InboxUIItems) DeepCopy() InboxUIItems {
	return InboxUIItems{
		Items: (func(x []InboxUIItem) []InboxUIItem {
			if x == nil {
				return nil
			}
			var ret []InboxUIItem
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Items),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
		Offline: o.Offline,
	}
}

type ChatAttachmentUploadOutboxIDArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	OutboxID  OutboxID `codec:"outboxID" json:"outboxID"`
}

type ChatAttachmentUploadStartArg struct {
	SessionID        int           `codec:"sessionID" json:"sessionID"`
	Metadata         AssetMetadata `codec:"metadata" json:"metadata"`
	PlaceholderMsgID MessageID     `codec:"placeholderMsgID" json:"placeholderMsgID"`
}

type ChatAttachmentUploadProgressArg struct {
	SessionID     int   `codec:"sessionID" json:"sessionID"`
	BytesComplete int64 `codec:"bytesComplete" json:"bytesComplete"`
	BytesTotal    int64 `codec:"bytesTotal" json:"bytesTotal"`
}

type ChatAttachmentUploadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentPreviewUploadStartArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	Metadata  AssetMetadata `codec:"metadata" json:"metadata"`
}

type ChatAttachmentPreviewUploadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentDownloadStartArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentDownloadProgressArg struct {
	SessionID     int   `codec:"sessionID" json:"sessionID"`
	BytesComplete int64 `codec:"bytesComplete" json:"bytesComplete"`
	BytesTotal    int64 `codec:"bytesTotal" json:"bytesTotal"`
}

type ChatAttachmentDownloadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatInboxUnverifiedArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Inbox     string `codec:"inbox" json:"inbox"`
}

type ChatInboxConversationArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Conv      string `codec:"conv" json:"conv"`
}

type ChatInboxFailedArg struct {
	SessionID int                    `codec:"sessionID" json:"sessionID"`
	ConvID    ConversationID         `codec:"convID" json:"convID"`
	Error     ConversationErrorLocal `codec:"error" json:"error"`
}

type ChatThreadCachedArg struct {
	SessionID int     `codec:"sessionID" json:"sessionID"`
	Thread    *string `codec:"thread,omitempty" json:"thread,omitempty"`
}

type ChatThreadFullArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Thread    string `codec:"thread" json:"thread"`
}

type ChatConfirmChannelDeleteArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Channel   string `codec:"channel" json:"channel"`
}

type ChatUiInterface interface {
	ChatAttachmentUploadOutboxID(context.Context, ChatAttachmentUploadOutboxIDArg) error
	ChatAttachmentUploadStart(context.Context, ChatAttachmentUploadStartArg) error
	ChatAttachmentUploadProgress(context.Context, ChatAttachmentUploadProgressArg) error
	ChatAttachmentUploadDone(context.Context, int) error
	ChatAttachmentPreviewUploadStart(context.Context, ChatAttachmentPreviewUploadStartArg) error
	ChatAttachmentPreviewUploadDone(context.Context, int) error
	ChatAttachmentDownloadStart(context.Context, int) error
	ChatAttachmentDownloadProgress(context.Context, ChatAttachmentDownloadProgressArg) error
	ChatAttachmentDownloadDone(context.Context, int) error
	ChatInboxUnverified(context.Context, ChatInboxUnverifiedArg) error
	ChatInboxConversation(context.Context, ChatInboxConversationArg) error
	ChatInboxFailed(context.Context, ChatInboxFailedArg) error
	ChatThreadCached(context.Context, ChatThreadCachedArg) error
	ChatThreadFull(context.Context, ChatThreadFullArg) error
	ChatConfirmChannelDelete(context.Context, ChatConfirmChannelDeleteArg) (bool, error)
}

func ChatUiProtocol(i ChatUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.chatUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"chatAttachmentUploadOutboxID": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadOutboxIDArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadOutboxIDArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadOutboxIDArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadOutboxID(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentUploadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadStart(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatAttachmentUploadProgress": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadProgressArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadProgressArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadProgressArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadProgress(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatAttachmentUploadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentUploadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentUploadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentUploadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentUploadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentPreviewUploadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentPreviewUploadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentPreviewUploadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentPreviewUploadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentPreviewUploadStart(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatAttachmentPreviewUploadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentPreviewUploadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentPreviewUploadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentPreviewUploadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentPreviewUploadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentDownloadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadStart(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentDownloadProgress": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadProgressArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadProgressArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadProgressArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadProgress(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatAttachmentDownloadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatInboxUnverified": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxUnverifiedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxUnverifiedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxUnverifiedArg)(nil), args)
						return
					}
					err = i.ChatInboxUnverified(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatInboxConversation": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxConversationArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxConversationArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxConversationArg)(nil), args)
						return
					}
					err = i.ChatInboxConversation(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatInboxFailed": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxFailedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxFailedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxFailedArg)(nil), args)
						return
					}
					err = i.ChatInboxFailed(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatThreadCached": {
				MakeArg: func() interface{} {
					ret := make([]ChatThreadCachedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatThreadCachedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatThreadCachedArg)(nil), args)
						return
					}
					err = i.ChatThreadCached(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatThreadFull": {
				MakeArg: func() interface{} {
					ret := make([]ChatThreadFullArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatThreadFullArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatThreadFullArg)(nil), args)
						return
					}
					err = i.ChatThreadFull(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatConfirmChannelDelete": {
				MakeArg: func() interface{} {
					ret := make([]ChatConfirmChannelDeleteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatConfirmChannelDeleteArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatConfirmChannelDeleteArg)(nil), args)
						return
					}
					ret, err = i.ChatConfirmChannelDelete(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ChatUiClient struct {
	Cli rpc.GenericClient
}

func (c ChatUiClient) ChatAttachmentUploadOutboxID(ctx context.Context, __arg ChatAttachmentUploadOutboxIDArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentUploadOutboxID", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentUploadStart(ctx context.Context, __arg ChatAttachmentUploadStartArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatAttachmentUploadStart", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatAttachmentUploadProgress(ctx context.Context, __arg ChatAttachmentUploadProgressArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatAttachmentUploadProgress", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatAttachmentUploadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentUploadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentUploadDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentPreviewUploadStart(ctx context.Context, __arg ChatAttachmentPreviewUploadStartArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatAttachmentPreviewUploadStart", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatAttachmentPreviewUploadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentPreviewUploadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentPreviewUploadDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentDownloadStart(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentDownloadStartArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadStart", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentDownloadProgress(ctx context.Context, __arg ChatAttachmentDownloadProgressArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatAttachmentDownloadProgress", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatAttachmentDownloadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentDownloadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatInboxUnverified(ctx context.Context, __arg ChatInboxUnverifiedArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatInboxUnverified", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatInboxConversation(ctx context.Context, __arg ChatInboxConversationArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatInboxConversation", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatInboxFailed(ctx context.Context, __arg ChatInboxFailedArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatInboxFailed", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatThreadCached(ctx context.Context, __arg ChatThreadCachedArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatThreadCached", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatThreadFull(ctx context.Context, __arg ChatThreadFullArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatThreadFull", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatConfirmChannelDelete(ctx context.Context, __arg ChatConfirmChannelDeleteArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatConfirmChannelDelete", []interface{}{__arg}, &res)
	return
}
