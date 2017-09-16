import axios from 'axios';
import createReducer from './createReducer';

const REQUEST_RATING = 'REQUEST_RATING';
const RECEIVE_RATING = 'RECEIVE_RATING';

const receiveRating = state => ({
    type: RECEIVE_RATING,
    payload: state,
});

const requestRating = state => ({
    type: REQUEST_RATING,
    payload: state,
});

export const fetchRating = () => (dispatch) => {
    dispatch(requestRating());

    return axios({
        method: 'get',
        url: '/ratings',
        responseType: 'json',
    })
        .then(({data}) => {
            dispatch(receiveRating(data));
        }, (e) => {
            window.alert(`Сорян, что-то упало. ${e.message}`);
        });
};

export default createReducer({
    isFetching: false,
    items: [],
}, {
    [REQUEST_RATING]: state => ({
        ...state,
        isFetching: true,
    }),
    [RECEIVE_RATING]: (state, action) => ({
        ...state,
        isFetching: false,
        items: action.payload,
    }),
});
