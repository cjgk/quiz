import Reflux from 'reflux';
import Actions from 'actions';

let menuSkel = {
    isOpen: false
}

let menu = Object.assign({}, menuSkel);

let MenuStore = Reflux.createStore({
    listenables: [Actions],

    getDefaultState() {
        return menuSkel;
    },

    onToggleMenu() {
        menu.isOpen = !menu.isOpen;
        this.trigger(menu);
    }
});

export default MenuStore;
