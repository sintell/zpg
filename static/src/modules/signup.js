import axios from 'axios';
import createReducer from './createReducer';

import {receiveLogin} from './login';

const REQUEST_SIGNUP = 'REQUEST_SIGNUP';
const RECEIVE_SIGNUP = 'RECEIVE_SIGNUP';
const ERROR_SIGNUP = 'ERROR_SIGNUP';

const receiveSignUp = state => ({
    type: RECEIVE_SIGNUP,
    payload: state,
});

const requestSignUp = state => ({
    type: REQUEST_SIGNUP,
    payload: state,
});

const errorSignUp = state => ({
    type: ERROR_SIGNUP,
    payload: state,
});

export const signUp = (signUpData) => (dispatch, getState) => {
    dispatch(requestSignUp());

    return axios({
        method: 'post',
        url: 'signup',
        data: signUpData,
        responseType: 'json',
    })
        .then(({data}) => {
            dispatch(receiveSignUp(data));
            dispatch(receiveLogin(data));

            // if (data.token) {
            //     window.localStorage.setItem('ZPGtoken', data.token);
            // }
        }, (e) => {
            // dispatch(errorSignUp({message: e.message}));
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

export default createReducer({
    isFetching: false,
    loggedIn: false,
    error: false,
}, {
    [REQUEST_SIGNUP]: state => ({
        ...state,
        isFetching: true,
    }),
    [RECEIVE_SIGNUP]: (state, action) => ({
        ...state,
        isFetching: false,
        loggedIn: action.payload,
    }),
    [ERROR_SIGNUP]: (state, action) => ({
        ...state,
        isFetching: false,
        error: action.payload,
    }),
});
