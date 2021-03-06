import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';
import {Route} from 'react-router-dom';
import styled from 'styled-components';

import routeMap from '../../modules/routePath';
import {receiveLogin} from '../../modules/login';
import {receiveChar} from  '../../modules/character';

import FormCharacter from '../../components/formCharacter';
import FormLogin from '../../components/formLogin';
import FormSignUp from '../../components/formSignUp';
import GamePlayContainer from '../../containers/gamePlay';
import Ratings from '../../components/ratings';

const Exit = styled.div`
    position: fixed;
    bottom: 20px;
    right: 10px;
`;

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

    logOut = () => {
        this.props.receiveLogin(false);
        this.props.receiveChar(false);
        this.props.push(routeMap.home);
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
                {this.props.loggedIn ?
                    <Exit onClick={this.logOut}>
                        <span className='button button-s'>Выйти</span>
                    </Exit> : ''}
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
    receiveLogin,
    receiveChar,
})(App);
