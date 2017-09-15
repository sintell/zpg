import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';


import {authUser} from '../../modules/login';
import routePath from '../../modules/routePath';


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

    signUp = () => {
        this.props.push(routePath.signUp);
    }

    render() {
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
                <p onClick={this.signUp}>Создать нового персонажа</p>
            </form>
        );
    }
}

export default connect((state) => ({
    isFetching: state.login.isFetching,
}), {
    authUser,
    push,
})(FormLogin);
