import React, {Component} from 'react';

import Sleep from './sleep.mp3';

export default class sleepChar extends Component {
    render() {
        return (
            <div>
                <div className='sleep' title='Вы решили вздремнуть, чтобы снять стресс.' />
                <audio src={Sleep} autoPlay />
            </div>
        );
    }
}

