import Reflux from 'reflux';
import Request from 'reqwest';
import Actions from 'actions';

const gameSkel = {
    id: null,
    created: null,
    status: null,
    name: null
};

let gameList = [];

let GameStore = Reflux.createStore({
    listenables: [Actions],
    
    onUpdateGameList() {
        Request({
            url: '/games',
            method: 'get',

            success: function (resp) {
                let games = JSON.parse(resp);

                for (let i = 0; i < games.length; i++) {
                    let game = games[i];
                    gameList.push(game);
                }

                this.trigger(gameList);
            }.bind(this),

            error: function (err) {
                console.log(err);
            }.bind(this)
        });
    }

});

export default GameStore;
