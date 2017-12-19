// @flow
import * as I from 'immutable'
import * as RPCChatTypes from './flow-types-chat'

// TODO put back
// export opaque type ConversationIDKey: string = string
export type ConversationIDKey = string
export const stringToConversationIDKey = (s: string): ConversationIDKey => s

// TODO put back
// export opaque type MessageID: string = string
export type MessageID = string
export const stringToMessageID = (s: string): MessageID => s

// TODO
type Message = any

type MembershipType = 'active' | 'youArePreviewing' | 'youAreReset'
type TeamType = 'small' | 'big' | 'adhoc'
export type LoadingState = 'untrusted' | 'unboxing' | 'unboxed'

// Metadata about a conversation. We keep the messages sep. since we update these at different times
export type _ConversationMeta = {
  id: ConversationIDKey,
  inboxVersion: number,
  isMuted: boolean,
  loadingState: LoadingState,
  membershipType: MembershipType,
  notificationSettings: ?RPCChatTypes.ConversationNotificationInfo,
  participants: I.Set<string>,
  resetParticipants: I.Set<string>,
  supersededBy: ?ConversationIDKey,
  supersedes: ?ConversationIDKey,
  teamType: TeamType,
}

export type ConversationMeta = I.RecordOf<_ConversationMeta>

export type _State = {
  metaMap: I.Map<ConversationIDKey, ConversationMeta>,
  messageMap: I.Map<ConversationIDKey, I.Map<MessageID, Message>>,
  messageIDsLsit: I.Map<ConversationIDKey, I.List<MessageID>>,
}

export type State = I.RecordOf<_State>
