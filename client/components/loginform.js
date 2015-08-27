import React from 'react';

class LoginForm extends React.Component {
    /*
    handleSubmit: function (e) {
        e.preventDefault();

        var email = this.refs.email.getDOMNode().value;
        var password = this.refs.password.getDOMNode().value;
        var data = {email: email, password: password};

        reqwest({
            url: '/sessions',
            method: 'post',
            data: data,

            success: function (resp) {
                console.log(resp);
            }
        });

        console.log(data);

    },
    */
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
