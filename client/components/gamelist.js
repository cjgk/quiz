import React from 'react';
import {Link} from 'react-router';
import Actions from 'actions';
import mui from 'material-ui';

let GameList = React.createClass({
    onAddGame(e) {
        e.preventDefault();
        Actions.showModal('game');
    },

    render() {
        let styles = {
            root: {
                width: 300,
                overflow: "auto",
            },
            head: {
                margin: 20,
                marginBottom: 0,
                display: "flex",
                flexDirection: "row",
                alignItems: "center"
            },
            titleBox: {
                flex: 1,
            },
            button: {
                flexGrow: 0,
                flexShrink: 1,
                flexBasis: "auto"
            },
            title: {
                margin: 0,
                padding: 0
            }
        };

        let games = [];
        this.props.games.forEach(game => {
            games.push(<mui.ListItem key={game.id} primaryText={game.name}/>);
        });

        return (
                <mui.Paper style={styles.root}>
                    <div style={styles.head}>
                        <div style={styles.titleBox}>
                            <h2 style={styles.title}>Your games</h2>
                        </div>
                        <div style={styles.button}>
                            <mui.FloatingActionButton
                                label="Add game" 
                                mini={true}
                                secondary={true}
                                onClick={this.onAddGame}>
                                <mui.FontIcon className="material-icons">add</mui.FontIcon> 
                            </mui.FloatingActionButton>
                        </div>
                    </div>
                    <mui.List>
                        {games}
                    </mui.List>
                </mui.Paper>
        );
    }
});

export default GameList;
