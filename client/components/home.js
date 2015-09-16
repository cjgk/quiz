import Reflux from 'reflux';
import React from 'react';

import Actions from 'actions';
import GameStore from 'stores/gamestore';

import GameList from 'components/gamelist';

let Home = React.createClass({ 
    mixins: [Reflux.listenTo(GameStore, 'onGamesUpdate')],

    getInitialState() {
        return {
            games: []
        }
    },

    onGamesUpdate(games) {
        this.setState({
            games: games
        });
    },

    componentDidMount() {
        Actions.updateGameList();
    },

    render() {
        return (
            <div className="home">
                <GameList games={this.state.games} />
            </div>
        );
    }
});

export default Home;
