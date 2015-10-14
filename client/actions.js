import Reflux from 'reflux';
import reqwest from 'reqwest';

const Actions = Reflux.createActions([
    // Auth actions
    'authenticate',

    // Game actions
    'updateGameList',

    // Modal actions
    'showModal',

    // Menu Actions
    'toggleMenu',
    'closeMenu',
]);

export default Actions;
