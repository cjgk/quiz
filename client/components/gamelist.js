import React from 'react';

import GameListItem from 'components/gamelistitem';

let GameList = React.createClass({
    render() {
        let games = [];
        this.props.games.forEach(game => {
            games.push(<GameListItem name={game.name} key={game.id}/>);
        });

        return (
            <ul>
                {games}
            </ul>
        );
    }
});

export default GameList;
