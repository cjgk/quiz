import React from 'react';
import {Route, DefaultRoute} from 'react-router';

import Main from 'components/main';
import LoginBox from 'components/loginbox';


const routes = (
    <Route handler={Main}>
        <DefaultRoute handler={LoginBox}/>
    </Route>
);

export default routes;
