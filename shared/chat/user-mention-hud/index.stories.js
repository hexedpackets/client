// @flow
import React from 'react'
import * as I from 'immutable'
import {MentionRowRenderer, MentionHud} from '.'
import {compose, withState} from 'recompose'
import {Box, Button, Input, ButtonBar} from '../../common-adapters'
import {createPropProvider, storiesOf, action} from '../../stories/storybook'
import {globalStyles} from '../../styles'

const provider = createPropProvider({
  MentionHud: props => ({
    ...props,
    following: I.Set(),
    you: 'chris',
    data: props.users
      .map((u, i) => ({
        username: u.username,
        fullName: u.fullName,
        key: u.username,
      }))
      .filter(u => {
        return (
          u.username.toLowerCase().indexOf(props.filter) >= 0 ||
          u.fullName.toLowerCase().indexOf(props.filter) >= 0
        )
      })
      .map((u, i) => ({...u, selected: i === props.selectedIndex})),
  }),
})

const UpDownFilterHoc = compose(
  withState('upCounter', 'setUpCounter', 0),
  withState('downCounter', 'setDownCounter', 0),
  withState('filter', 'setFilter', ''),
  Component => props => (
    <Box style={globalStyles.flexBoxColumn}>
      <Component upCounter={props.upCounter} downCounter={props.downCounter} filter={props.filter} />
      <ButtonBar>
        <Button label="Up" type="Primary" onClick={() => props.setUpCounter(n => n + 1)} />
        <Button label="Down" type="Primary" onClick={() => props.setDownCounter(n => n + 1)} />
      </ButtonBar>
      <Input onChangeText={props.setFilter} hintText="Filter" />
    </Box>
  )
)

const load = () => {
  storiesOf('Chat/Heads up Display', module)
    .addDecorator(provider)
    .add('Mention Row', () => (
      <Box style={{width: 240}}>
        <MentionRowRenderer
          following={I.Set()}
          you="chris"
          username="trex"
          fullName="T. Bone Rexasaurus"
          key="trex"
          selected={false}
          onClick={action('onClick')}
          onHover={action('onHover')}
        />
        <MentionRowRenderer
          following={I.Set()}
          you="chris"
          username="marcopolo"
          fullName="Marco Munizaga"
          key="marcopolo"
          selected={true}
          onClick={action('onClick')}
          onHover={action('onHover')}
        />
        <MentionRowRenderer
          following={I.Set()}
          you="chris"
          username="missingno"
          fullName="MissingNo"
          key="missingno"
          selected={false}
          onClick={action('onClick')}
          onHover={action('onHover')}
        />
      </Box>
    ))
    .add('Mention Hud', () => {
      const Hud = UpDownFilterHoc(({upCounter, downCounter, filter}) => (
        <Box style={{height: 100, width: 240, ...globalStyles.flexBoxColumn}}>
          <MentionHud
            users={[{username: 'marcopolo', fullName: 'Marco Munizaga'}, {username: 'trex', fullName: ''}]}
            onPickUser={action('onPickUser')}
            onSelectUser={action('onSelectUser')}
            selectUpCounter={upCounter}
            selectDownCounter={downCounter}
            pickSelectedUserCounter={0}
            filter={filter}
            style={{flex: 1}}
          />
        </Box>
      ))
      return <Hud />
    })
}

export default load
