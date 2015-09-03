import React from 'react';
import reqwest from 'reqwest';

class LoginForm extends React.Component {
    
    handleSubmit = evt => {
        evt.preventDefault()

        var email = this.refs.email.getDOMNode().value;
        var password = this.refs.password.getDOMNode().value;
        var loginData = {email: email, password: password};

        this.props.onLoginSubmit(loginData);
    }

    render () {
        return (
            <form onSubmit={this.handleSubmit}>
                <input type="text" ref="email" placeholder="email" />
                <input type="text" ref="password" placeholder="pass" />
                <input type="submit" value="Login" />
            </form>
        );
    }
}

export default LoginForm;
