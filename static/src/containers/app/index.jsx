import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';
import {Route} from 'react-router-dom';

import routeMap from '../../modules/routePath';

import FormCharacter from '../../components/formCharacter';
import FormLogin from '../../components/formLogin';
import FormSignUp from '../../components/formSignUp';
import GamePlayContainer from '../../containers/gamePlay';

window.onpopstate = function(event) {
    console.log(event);
    if (event) {
        window.history.go();
    }
};

class App extends Component {
    render() {
        if (this.props.loggedIn && !this.props.char) {
            return <FormCharacter />;
        }

        if (this.props.loggedIn) {
            return <GamePlayContainer />;
        }

        return (
            <div>
                <Route path={routeMap.home} component={FormLogin} />
                <Route path={routeMap.createChar} component={FormCharacter} />
                <Route path={routeMap.game} component={GamePlayContainer} />
                <Route path={routeMap.signUp} component={FormSignUp} />
                <Route path={routeMap.signIn} component={FormLogin} />
            </div>
        );
    }
}

export default connect(state => ({
    routePath: state.routing.location.pathname,
    loggedIn: state.login.loggedIn,
    char: state.char.char,
}), {
    push,
})(App);
