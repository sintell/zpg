import React, {Component} from 'react';
import {connect} from 'react-redux';
import {push} from 'react-router-redux';


import {signUp} from '../../modules/signup';
import routePath from '../../modules/routePath';


class FormSignUp extends Component {
    state = {
        login: '',
        password: '',
    };

    submit = (event) => {
        event.preventDefault();
        this.props.signUp(this.state);
    }

    changeInput = (event) => {
        let state = {};
        state[event.target.name] = event.target.value;
        this.setState(state);
    }

    signIn = () => {
        this.props.push(routePath.signIn);
    }

    render() {
        return (
            <div className='popup popup_login'>
                <form method='post' onSubmit={this.submit}>
                    <div className='cube'>
                        <div className='form-padding'>
                            <h2 className='h2'>Регистрация</h2>
                            <div className='login-grid'>
                                <div className='form-row'>
                                    <label className='form-label'>Логин:</label>
                                    <input
                                        required
                                        className='input'
                                        type='text'
                                        name='login'
                                        value={this.state.login}
                                        onChange={this.changeInput} />

                                </div>
                                <div className='form-row'>
                                    <label className='form-label'>Пароль:</label>
                                    <input
                                        required
                                        className='input'
                                        type='password'
                                        name='password'
                                        value={this.state.password}
                                        onChange={this.changeInput} />
                                </div>
                            </div>
                            <div className='separate'></div>
                            <div className='form-actions'>
                                <div className='form-row'>
                                    <button className='button' type='submit'>Зарегистрироваться</button>
                                </div>
                                <div className='form-row'>
                                    или
                                </div>
                                <span className='link link_switch' onClick={this.signIn}>Войти</span>
                            </div>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}

export default connect((state) => ({
    isFetching: state.signup.isFetching,
}), {
    signUp,
    push,
})(FormSignUp);
