import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';


class App extends Component {
    render() {
        return (
            <div>
                ЦПГ
            </div>
        );
    }
}

export default connect(state => ({
    routePath: state.routing.location.pathname,
}), {
    push,
})(App);
