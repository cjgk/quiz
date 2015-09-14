import React from 'react';

let GameListItem = React.createClass({
    render() {
        return (
            <li>{this.props.name}</li>
        );
    }
});

export default GameListItem;
