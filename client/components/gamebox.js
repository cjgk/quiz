import React from 'react';
import reqwest from 'reqwest';

class GameBox extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            games: []
        };
    }

    render() {
        return (
            <div className="gamebox">
                <h2>Your games</h2>
                <GameList games={this.state.games} />
            </div>
        );
    }
}

export default GameBox;
