import React, {Component} from 'react';


export default class Progress extends Component {
    render() {
        const { value, maxValue, color } = this.props;
        let percentage = Math.round((value / maxValue) * 100);

        return (
            <div className='progress'>
                <div className='progress__percentage' style={{width: percentage + '%', backgroundColor: color}} />
                <div className='progress__value'>{value}/{maxValue}</div>
            </div>
        );
    }
}

