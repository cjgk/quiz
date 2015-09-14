import Reflux from 'reflux';
import reqwest from 'reqwest';

const Actions = Reflux.createActions([
    'authenticate',
    'updateGameList'
]);

export default Actions;
