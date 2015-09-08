import React from 'react';
import reqwest from 'reqwest';

import connectToStores from 'alt/utils/connectToStores';
import LoginStore from 'stores/loginStore';

import LoginForm from 'components/loginform';
import LoginMessages from 'components/loginmessages';

@connectToStores
class LoginBox extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            messages: []
        };

    }

    static getStores(props) {
        return [LoginStore];
    }

    static getPropsFromStores(props) {
        return LoginStore.getState();
    }

    handleLoginSubmit = (loginData) => {
        reqwest({
            url: '/sessions',
            method: 'post',
            data: loginData,

            success: function (resp) {
                console.log("success",resp);
            }.bind(this),

            error: function (err) {
                this.setState({messages: [err.responseText]});
            }.bind(this)
        });
    }

    render() {
        return (
            <div className="loginbox">
                <h2>Login</h2>
                <LoginForm onLoginSubmit={this.handleLoginSubmit} />
                <LoginMessages messages={this.state.messages} />
            </div>
        );
    }
}

export default LoginBox;
