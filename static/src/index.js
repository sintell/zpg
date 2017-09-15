import React from 'react';
import {render} from 'react-dom';
import {Provider} from 'react-redux';
import {ConnectedRouter} from 'react-router-redux';
import getOrCreateStore, {history} from './store';
import axios from 'axios';

import App from './containers/app';

import './index.css';

const store = getOrCreateStore();

axios.defaults.baseURL = 'url/';

render(
    <Provider store={store}>
        <ConnectedRouter history={history}>
            <App />
        </ConnectedRouter>
    </Provider>,
    document.querySelector('#root'),
);
