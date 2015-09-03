import React from 'react';
import reqwest from 'reqwest';

import LoginForm from 'components/loginform';
import LoginMessages from 'components/loginmessages';

class LoginBox extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            messages: ["joho"]
        };
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
