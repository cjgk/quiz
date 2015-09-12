import React from 'react';
import {Route, DefaultRoute} from 'react-router';

import App from 'components/app';
import LoginBox from 'components/loginbox';
import Home from 'components/home';

const routes = (
    <Route handler={App}>
        <DefaultRoute handler={LoginBox} />
        <Route name="home" handler={Home} />
    </Route>
);

export default routes;
