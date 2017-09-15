import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';

import FormCharacter from '../../components/formCharacter';


class App extends Component {
    render() {
        return (
            <div>
                <FormCharacter />
            </div>
        );
    }
}

export default connect(state => ({
    routePath: state.routing.location.pathname,
}), {
    push,
})(App);
