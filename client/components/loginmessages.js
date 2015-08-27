import React from 'react';
import LoginMessage from 'components/loginmessage';

class LoginMessages extends React.Component {
    constructor() {
        super();

        this.state = {
            messages: []
        };
    }

    render() {
        this.props.messages.forEach(function(message) {
            messages.push(<LoginMessage>{message}</LoginMessage>);
        }.bind(this));

        return (
            <ul className="loginbox-messages">
                {this.state.messages}
            </ul>
        );
    }
}

export default LoginMessages;
