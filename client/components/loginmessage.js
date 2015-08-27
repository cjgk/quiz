import React from 'react';

class LoginMessage extends React.Component {
    render () {
        return (
            <li className="loginbox-message">{this.props.children}</li>
        );
    }
}

export default LoginMessage;
