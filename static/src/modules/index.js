import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';

import login from './login';
import signup from './signup';
import charData from './character';
import rating from './rating';

export default combineReducers({
    routing: routerReducer,
    login,
    signup,
    charData,
    rating,
});
