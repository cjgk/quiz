import React from 'react';  
import Reflux from 'reflux';
import {RouteHandler, Link, Navigation, State} from 'react-router';

import LoginStore from 'stores/loginstore';
import ModalStore from 'stores/modalstore';

import Modal from 'components/modal';
import GameForm from 'components/gameform';

let App = React.createClass({
    mixins: [
        Reflux.listenTo(LoginStore, 'onAuthUpdate'),
        Reflux.listenTo(ModalStore, 'onModalUpdate'),
        Navigation,
        State
    ],

    getInitialState() {
        return {
            auth: LoginStore.getStatus(),
            modal: ModalStore.getDefaultState()
        }
    },

    onAuthUpdate(auth) {
        this.setState({auth: auth});
    },

    onModalUpdate(modalState) {
        this.setState({modal: modalState});
    },


    sendToHomeIfAuthenticated() {
        if (this.getPath() == '/' && this.state.auth.authenticated) {
            this.replaceWith('home');
        }
    },

    getModal(modal) {
        console.log(modal);
        if (modal.type == null) {
            return null;
        }

        let modalInner;
        switch (modal.type) {
            case 'game':
                modalInner = <GameForm/>;
                break;
        }

        return (
            <Modal><i>Hej!</i></Modal>
        );
    },

    componentDidUpdate() {
        this.sendToHomeIfAuthenticated();
    },

    componentDidMount() {
        this.sendToHomeIfAuthenticated();
    },

    render() {
        return (
            <div>
               <h1>Quiz</h1>
                <RouteHandler/>

                {this.getModal(this.state.modal)}
            </div>
        );
    }
});

export default App;
