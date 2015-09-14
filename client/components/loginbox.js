import Reflux from 'reflux';
import React from 'react';

import Actions from 'actions';
import LoginStore from 'stores/loginstore';

import LoginForm from 'components/loginform';
import LoginMessages from 'components/loginmessages';

const LoginBox = React.createClass({
    
    mixins: [Reflux.connect(LoginStore)],

    handleSubmit(loginData) {
        Actions.authenticate(loginData);
    },

    render() {
        return (
            <div className="loginbox">
                <h2>Login</h2>
                <LoginForm onLoginSubmit={this.handleSubmit} />
                <LoginMessages messages={[]} />
            </div>
        );
    }
});

export default LoginBox;
