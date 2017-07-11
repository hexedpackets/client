// @flow
import React from 'react'
import Header from './header/container'
import SearchHeader from '../search-header'
import * as ChatConstants from '../../constants/chat'
import * as SearchConstants from '../../constants/searchv3'

type Props = {
  inSearch: boolean,
  onChangeSearchText: (searchText: string) => void,
  searchText: string,
  selectedSearchId: ?SearchConstants.SearchResultId,
  selectedConversationIDKey: ?ChatConstants.ConversationIDKey,
  onUpdateSelectedSearchResult: (id: ?SearchConstants.SearchResultId) => void,
  sidePanelOpen: boolean,
  onToggleSidePanel: () => void,
  onBack: () => void,
  onAddNewParticipant: () => void,
  addNewParticipant: boolean,
}

export default (props: Props) =>
  props.inSearch
    ? <SearchHeader
        onChangeSearchText={props.onChangeSearchText}
        searchText={props.searchText}
        selectedConversationIDKey={props.selectedConversationIDKey}
        selectedSearchId={props.selectedSearchId}
        onUpdateSelectedSearchResult={props.onUpdateSelectedSearchResult}
        onAddNewParticipant={props.onAddNewParticipant}
        addNewParticipant={props.addNewParticipant}
      />
    : <Header
        sidePanelOpen={props.sidePanelOpen}
        onToggleSidePanel={props.onToggleSidePanel}
        onBack={props.onBack}
      />
