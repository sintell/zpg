import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';
import {Route} from 'react-router-dom';

import routeMap from '../../modules/routePath';

import FormCharacter from '../../components/formCharacter';
import FormLogin from '../../components/formLogin';

window.onpopstate = function() {
    window.history.go();
};

class App extends Component {
    render() {
        return (
            <div>
                <Route path={routeMap.home} component={FormLogin} />
                <Route path={routeMap.signUp} component={FormCharacter} />
            </div>
        );
    }
}

export default connect(state => ({
    routePath: state.routing.location.pathname,
    loggedIn: state.login.loggedIn,
}), {
    push,
})(App);
