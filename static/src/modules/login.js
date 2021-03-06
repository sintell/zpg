import axios from 'axios';
import createReducer from './createReducer';

import {receiveChar} from './character';

const REQUEST_LOGIN = 'REQUEST_LOGIN';
const RECEIVE_LOGIN = 'RECEIVE_LOGIN';

export const receiveLogin = state => {
    if (state.token) {
        window.localStorage.setItem('ZPGtoken', state.token);
        axios.defaults.headers.common['Authorization'] = `Bearer ${state.token}`;
    } else {
        window.localStorage.removeItem('ZPGtoken');
        window.localStorage.removeItem('ZPGchar');
    }

    return {
        type: RECEIVE_LOGIN,
        payload: state,
    };
};

const requestLogin = state => ({
    type: REQUEST_LOGIN,
    payload: state,
});

export const authUser = (authData) => (dispatch) => {
    dispatch(requestLogin());

    return axios({
        method: 'post',
        url: 'signin',
        data: authData,
        responseType: 'json',
    })
        .then(({data}) => {
            dispatch(receiveLogin(data));

            if (data.state) {
                dispatch(receiveChar(data.state));
            }
        }, (e) => {
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

export default createReducer({
    isFetching: false,
    loggedIn: window.localStorage.getItem('ZPGtoken'),
}, {
    [REQUEST_LOGIN]: state => ({
        ...state,
        isFetching: true,
    }),
    [RECEIVE_LOGIN]: (state, action) => ({
        ...state,
        isFetching: false,
        loggedIn: action.payload,
    }),
});
