import React from 'react';
import LoginForm from 'components/loginform';
import LoginMessages from 'components/loginmessages';

class LoginBox extends React.Component {
    render() {
        var messages = [];
        return (
            <div className="loginbox">
                <h2>Login</h2>
                <LoginForm />
                <LoginMessages messages={messages} />
            </div>
        );
    }
}

export default LoginBox;
