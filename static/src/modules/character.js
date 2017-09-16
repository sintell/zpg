import axios from 'axios';
import createReducer from './createReducer';

const REQUEST_CHAR = 'REQUEST_CHAR';
const RECEIVE_CHAR = 'RECEIVE_CHAR';
const SLEEP_CHAR = 'SLEEP_CHAR';
export const TICK_MS = 500;

export const receiveChar = state => ({
    type: RECEIVE_CHAR,
    payload: state.char,
});

const requestChar = state => ({
    type: REQUEST_CHAR,
    payload: state,
});

const sleepChar = state => ({
    type: SLEEP_CHAR,
    payload: state,
});

function getCharDataInLocalStorage() {
    let char = window.localStorage.getItem('ZPGchar');
    char = char && JSON.parse(char);
    return char;
}

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

            data.createData = new Date();

            window.localStorage.setItem('ZPGchar', JSON.stringify(data));
        }, (e) => {
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

export const fetchChar = () => (dispatch) => {
    const char = getCharDataInLocalStorage();

    if (char && char.char.createData - new Date() < TICK_MS) {
        return;
    }

    dispatch(requestChar());

    return axios({
        method: 'get',
        url: 'game/state',
        responseType: 'json',
    })
        .then(({data}) => {
            console.log(data);
            dispatch(receiveChar({
                char: data,
            }));

            window.localStorage.setItem('ZPGchar', JSON.stringify(data));
        }, (e) => {
            // window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

export const sendToSleep = () => (dispatch) => {
    return axios({
        method: 'post',
        url: 'game/rest',
        responseType: 'json',
    })
        .then(({data}) => {
            dispatch(sleepChar({
                sleep: true,
            }));
        }, (e) => {
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

const char = getCharDataInLocalStorage();

export default createReducer({
    isFetching: false,
    char: char && char.char,
    projects: char && char.projects,
    sleep: false,
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
    [SLEEP_CHAR]: (state, action) => ({
        ...state,
        isFetching: false,
        sleep: action.payload,
    }),
});
