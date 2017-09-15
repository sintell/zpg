import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';

import login from './login';
import signup from './signup';
import char from './character';

export default combineReducers({
    routing: routerReducer,
    login,
    signup,
    char,
});
