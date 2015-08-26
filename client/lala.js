import React from 'react';

React.render(  
  <h1>Example</h1>,
  document.getElementById('content')
);

/*
import reqwest from 'reqwest';
import routes from 'routes';

import 'quiz.scss';

Router.run(routes, Router.HistoryLocation, (Root, state) => {
    React.render(<Root {...state}/>, document.getelementById('content'));
});

var LoginBox = React.createClass({displayName: 'LoginBox',
    render: function () {
        return (
            <div className="loginbox">
                <h1>Login</h1>
                <LoginForm />
                <LoginMessages />
            </div>
        );
    }
});

var LoginForm = React.createClass({displayName: 'LoginForm',
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
    render: function () {
        return (
            <form onSubmit={this.handleSubmit}>
                <input type="text" ref="email" placeholder="email" />
                <input type="text" ref="password" placeholder="pass" />
                <input type="submit" value="Login" />
            </form>
        );
    }
});

var LoginMessages = React.createClass({displayName: 'LoginMessages',
    render: function () {
        return (
            <ul className="loginbox-messages">
                <LoginMessage>Login failed</LoginMessage>
            </ul>
        );
    }
});

var LoginMessage = React.createClass({displayName: 'LoginMessage',
    render: function () {
        return (
            <li className="loginbox-message">{this.props.children}</li>
        );
    }
});

React.render(
        <LoginBox />,
        document.getElementById('content')
);
*/
