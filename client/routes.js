import React from 'react';
import {Route, DefaultRoute} from 'react-router';

import App from 'components/app';
import LoginBox from 'components/loginbox';
import GameBox from 'components/gamebox';

const routes = (
    <Route handler={App}>
        <DefaultRoute handler={LoginBox} />
    </Route>
);

export default routes;
