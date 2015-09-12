import React from 'react';  
import Reflux from 'reflux';
import {RouteHandler, Link, Navigation, State} from 'react-router';

import LoginStore from 'stores/loginstore';

let App = React.createClass({
    mixins: [
        Reflux.listenTo(LoginStore, 'onAuthUpdate'),
        Navigation,
        State
    ],

    getInitialState() {
        return {
            auth: LoginStore.getStatus()
        }
    },

    onAuthUpdate(auth) {
        this.setState({auth: auth});
    },

    sendToHomeIfAuthenticated() {
        if (this.getPath() == '/' && this.state.auth.authenticated) {
            this.replaceWith('home');
        }
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
            </div>
        );
    }
});

export default App;
