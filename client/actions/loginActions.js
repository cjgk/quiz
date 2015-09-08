import flux from 'dispatchers/appDispatcher'
import {createActions} from 'alt/utils/decorators'

@createActions(flux)
class LoginActions {
    constructor() {
        this.generateActions('authenticate');
    }
}

export default LoginActions;
