import React, {Component} from 'react';


export default class Progress extends Component {
    render() {
        const { value, maxValue, color, title } = this.props;
        let percentage = Math.round((value / maxValue) * 100);
        percentage = percentage > 0 ? percentage : 0;

        return (
            <div className='progress' title={title || ''}>
                <div className='progress__percentage' style={{width: percentage + '%', backgroundColor: color}} />
                <div className='progress__value'>{value > 0 ? value : 0}/{maxValue}</div>
            </div>
        );
    }
}

