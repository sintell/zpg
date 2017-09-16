import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';
import {Route} from 'react-router-dom';

import routeMap from '../../modules/routePath';

import FormCharacter from '../../components/formCharacter';
import FormLogin from '../../components/formLogin';
import FormSignUp from '../../components/formSignUp';
import GamePlayContainer from '../../containers/gamePlay';
import Ratings from '../../components/ratings';

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
        const path = window.location.pathname;

        if (!props.loggedIn && path !== routeMap.signUp && path !== routeMap.ratings) {
            props.push(routeMap.signIn);
        } else if (props.loggedIn && !props.charExist && path !== routeMap.ratings) {
            props.push(routeMap.createChar);
        } else if (props.loggedIn && path !== routeMap.ratings) {
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
                <Route path={routeMap.ratings} component={Ratings} />
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
