import React, {Component} from 'react';
import {connect} from 'react-redux';

import {authUser} from '../../modules/login';


class FormLogin extends Component {
    state = {
        login: '',
        password: '',
    };

    submit = (event) => {
        event.preventDefault();
        this.props.authUser();
    }

    changeInput = (event) => {
        let state = {};
        state[event.target.name] = event.target.value;
        this.setState(state);
    }

    render() {
        console.log(this.props)
        return (
            <form method='post' onSubmit={this.submit}>
                <fieldset>
                    <label htmlFor='login'>
                        <input
                            type='text'
                            name='login'
                            placeholder='Логин'
                            value={this.state.login}
                            onChange={this.changeInput} />
                    </label>
                    <label htmlFor='password'>
                        <input
                            type='password'
                            name='password'
                            placeholder='Пароль'
                            value={this.state.password}
                            onChange={this.changeInput} />
                    </label>
                    <button type='submit'>Войти</button>
                </fieldset>
            </form>
        );
    }
}

export default connect((state) => ({
    isFetching: state.login.isFetching,
}), {
    authUser,
})(FormLogin);
