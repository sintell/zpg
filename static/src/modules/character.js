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

export const createChar = (charData) => (dispatch, getState) => {
    dispatch(requestChar());
    console.log(charData);

    return axios({
        method: 'post',
        url: 'game/create_char',
        data: charData,
        responseType: 'json',
    })
        .then(({data}) => {
            dispatch(receiveChar({
                char: charData,
            }));

            window.localStorage.setItem('ZPGchar', JSON.stringify(charData));
        }, (e) => {
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

const char = window.localStorage.getItem('ZPGchar');

export default createReducer({
    isFetching: false,
    char: char && JSON.parse(char),
}, {
    [REQUEST_CHAR]: state => ({
        ...state,
        isFetching: true,
    }),
    [RECEIVE_CHAR]: (state, action) => ({
        ...state,
        isFetching: false,
        char: action.payload,
    }),
});
