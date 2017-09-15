import React from 'react';
import {render} from 'react-dom';
import {Provider} from 'react-redux';
import {ConnectedRouter} from 'react-router-redux';
import getOrCreateStore, {history} from './store';
import axios from 'axios';

import App from './containers/app';

import './index.css';

const store = getOrCreateStore();

axios.defaults.baseURL = 'http://10.208.10.15:1323/';
axios.defaults.headers.common['Authorization'] = `Bearer ${window.localStorage.getItem('ZPGtoken')}`;

render(
    <Provider store={store}>
        <ConnectedRouter history={history}>
            <App />
        </ConnectedRouter>
    </Provider>,
    document.querySelector('#root'),
);
