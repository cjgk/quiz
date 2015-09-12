import React from 'react';

let LoginForm = React.createClass({
    getInitialState() {
        return {
            email: "user1@example.com",
            password: "password1"
        }
    },

    handleSubmit(evt) {
        evt.preventDefault()

        var loginData = this.state;

        this.props.onLoginSubmit(loginData);
    },

    render () {
        return (
            <form onSubmit={this.handleSubmit}>
                <input
                    type="text" 
                    ref="email" 
                    placeholder="email" 
                    onChange={(e) => this.setState({email: e.target.value.trim()})} />
                <input 
                    type="text" 
                    ref="password" 
                    placeholder="pass" 
                    onChange={(e) => this.setState({password: e.target.value.trim()})} />
                <input type="submit" value="Login" />
            </form>
        );
    }
});

export default LoginForm;
