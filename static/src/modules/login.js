import axios from 'axios';
import createReducer from './createReducer';

const REQUEST_LOGIN = 'REQUEST_LOGIN';
const RECEIVE_LOGIN = 'RECEIVE_LOGIN';

const receivePosts = state => ({
    type: RECEIVE_LOGIN,
    payload: state,
});

const requestPosts = state => ({
    type: REQUEST_LOGIN,
    payload: state,
});

export const authUser = (authData) => (dispatch, getState) => {
    dispatch(requestPosts());

    return axios('account/login', {
        authData,
    })
        .then(({data}) => {
            dispatch(receivePosts(data));
        }, () => {
            dispatch({
                error: 'Ошибка бекенда',
            });
        });
};

export default createReducer({
    isFetching: false,
    loggedIn: true,
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
