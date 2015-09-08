import flux from 'dispatchers/appDispatcher';
import {createStore, bind} from 'alt/utils/decorators';
import actions from 'actions/loginActions';

@createStore(flux)
class LoginStore {
    constructor(props) {
        this.authenticated = false;
        this.user = null;
    }

    @bind(actions.authenticate)
    authenticate(username) {
        this.user = username;
    }
}

export default LoginStore;
