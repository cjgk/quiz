import Reflux from 'reflux';
import Request from 'reqwest';
import Actions from 'actions';
import Cookie from 'js-cookie';

const authSkel = {
    id: null,
    authenticated: false,
    name: null,
    email: null
};

let auth = Object.assign({}, authSkel);

var LoginStore = Reflux.createStore({
    listenables: Actions,

    onAuthenticate(loginData) {
        Request({
            url: '/api/1.0/sessions',
            method: 'post',
            data: loginData,

            success: function (resp) {
                let user = JSON.parse(resp);

                auth = Object.assign(auth, user);
                auth.authenticated = true;

                this.storeInSession();

                this.trigger(auth);
            }.bind(this),

            error: function (err) {
                console.log(err);
            }.bind(this)
        });

    },

    getStatus() {
        if (this.isStoredInSession()) {
           return this.getFromSession();
        }
        return auth;
    },

    isStoredInSession() {
        if (null !== sessionStorage.getItem('id')) {
            return true;
        }

        return false;
    },

    getFromSession() {
        let sessAuth = {
            id:            parseInt(sessionStorage.getItem('id')),
            authenticated: sessionStorage.getItem('authenticated') === "true" ? true : false,
            name:          sessionStorage.getItem('name'),
            email:         sessionStorage.getItem('email')
        };

        auth = Object.assign(auth, sessAuth);

        return auth;
    },

    storeInSession() {
        sessionStorage.setItem('id', auth.id);
        sessionStorage.setItem('authenticated', auth.authenticated);
        sessionStorage.setItem('name', auth.name);
        sessionStorage.setItem('email', auth.email);
    }

});

export default LoginStore;
