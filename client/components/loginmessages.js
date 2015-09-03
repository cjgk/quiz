import React from 'react';
import LoginMessage from 'components/loginmessage';

class LoginMessages extends React.Component {

    render() {
        var messages = [];
        this.props.messages.forEach(function(message) {
            messages.push(<LoginMessage>{message}</LoginMessage>);
        }.bind(this));

        return (
            <ul className="loginbox-messages">
                {messages}
            </ul>
        );
    }
}

export default LoginMessages;
