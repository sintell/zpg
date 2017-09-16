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
    if (event) {
        window.history.go();
    }
};

class App extends Component {
    componentDidMount() {
        this.route(this.props);
    }
    componentWillUpdate(nextProps) {
        this.route(nextProps);
    }

    route(props) {
        if (!props.loggedIn && window.location.pathname !== routeMap.signUp) {
            props.push(routeMap.signIn);
        } else if (props.loggedIn && !props.charExist) {
            props.push(routeMap.createChar);
        } else if (props.loggedIn) {
            props.push(routeMap.game);
        }
    }

    render() {
        return (
            <div>
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
    charExist: state.charData.char,
}), {
    push,
})(App);
