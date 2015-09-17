import React from 'react';
import {Link} from 'react-router';
import Actions from 'actions';
import GameListItem from 'components/gamelistitem';

let GameList = React.createClass({
    onAddGame(e) {
        e.preventDefault();
        Actions.showModal('game');
    },

    render() {
        let games = [];
        this.props.games.forEach(game => {
            games.push(<GameListItem name={game.name} key={game.id}/>);
        });

        return (
            <div>
                <h2>Your games</h2>
                <Link to={`/gameadd`} onClick={this.onAddGame}>Add game</Link>
                <ul>
                    {games}
                </ul>
            </div>
        );
    }
});

export default GameList;
