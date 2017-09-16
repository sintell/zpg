import axios from 'axios';
import createReducer from './createReducer';

const REQUEST_CHAR = 'REQUEST_CHAR';
const RECEIVE_CHAR = 'RECEIVE_CHAR';

export const receiveChar = state => ({
    type: RECEIVE_CHAR,
    payload: state.char,
});

const requestChar = state => ({
    type: REQUEST_CHAR,
    payload: state,
});

export const createChar = (charData) => (dispatch) => {
    dispatch(requestChar());

    return axios({
        method: 'post',
        url: 'game/create_char',
        data: charData,
        responseType: 'json',
    })
        .then(({data}) => {
            dispatch(receiveChar({
                char: data,
            }));

            window.localStorage.setItem('ZPGchar', JSON.stringify(data));
        }, (e) => {
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

const char = window.localStorage.getItem('ZPGchar');

export default createReducer({
    isFetching: false,
    char: char && JSON.parse(char).char,
    projects: char && JSON.parse(char).projects,
}, {
    [REQUEST_CHAR]: state => ({
        ...state,
        isFetching: true,
    }),
    [RECEIVE_CHAR]: (state, action) => ({
        ...state,
        isFetching: false,
        char: action.payload && action.payload.char,
        projects: action.payload && action.payload.projects,
    }),
});
