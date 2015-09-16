import Reflux from 'reflux';
import Actions from 'actions';

let modalSkel = {
    type: null
}

let modal = Object.assign({}, modalSkel);

let ModalStore = Reflux.createStore({
    listenables: [Actions],

    getDefaultState() {
        return modalSkel;
    },

    onShowModal(type) {
        modal.type = type;
        this.trigger(modal);
    }
});

export default ModalStore;
