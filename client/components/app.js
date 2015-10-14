import React from 'react';  
import Reflux from 'reflux';
import {RouteHandler, Link, Navigation, State} from 'react-router';
import injectTapEventPlugin from 'react-tap-event-plugin';

import UiThemeManager from 'material-ui/lib/styles/theme-manager';
import Colors from 'material-ui/lib/styles/colors';

import Actions from 'actions';
import LoginStore from 'stores/loginstore';
import ModalStore from 'stores/modalstore';
import MenuStore from 'stores/menustore';

import AppBar from 'material-ui/lib/app-bar';
import LeftNav from 'material-ui/lib/left-nav';
import Modal from 'components/modal';
import GameForm from 'components/gameform';

let ThemeManager = new(UiThemeManager);
injectTapEventPlugin();

let menuItems = [{route: "home", text: "Home"}];

let App = React.createClass({
    childContextTypes: {
        muiTheme: React.PropTypes.object
    },

    mixins: [
        Reflux.listenTo(LoginStore, 'onAuthUpdate'),
        Reflux.listenTo(ModalStore, 'onModalUpdate'),
        Reflux.listenTo(MenuStore, 'onMenuUpdate'),
        Navigation,
        State
    ],

    getChildContext() {
        return {
            muiTheme: ThemeManager.getCurrentTheme()
        }
    },

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

    onMenuUpdate(menuState) {
        this.refs.leftMenu.toggle();
    },

    sendToHomeIfAuthenticated() {
        if (this.getPath() == '/' && this.state.auth.authenticated) {
            this.replaceWith('home');
        }
    },

    getModal(modal) {
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

    onToggleMenu(e) {
        Actions.toggleMenu();
    },

    componentDidUpdate() {
        this.sendToHomeIfAuthenticated();
    },

    componentDidMount() {
        this.sendToHomeIfAuthenticated();
    },

    componentWillMount() {
        ThemeManager.setPalette({
            accent1Color: Colors.deepOrange500
        });
    },

    render() {
        return (
            <div>
                <AppBar 
                    title="Quiz"
                    showMenuIconButton={this.state.auth.authenticated}
                    onLeftIconButtonTouchTap={this.onToggleMenu}/>

                <LeftNav
                    ref="leftMenu"
                    docked={false}
                    menuItems={menuItems}/>

                <div className="inside">
                    <RouteHandler/>
                </div>

                {this.getModal(this.state.modal)}
            </div>
        );
    }
});

export default App;
