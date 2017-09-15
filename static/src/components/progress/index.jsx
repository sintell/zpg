import React, {Component} from 'react';


export default class Progress extends Component {
    render() {
        const { value, color } = this.props;

        return (
            <div className='progress'>
                <div className='progress__value' style={{width: value + '%', backgroundColor: color}} />
            </div>
        );
    }
}

