import React, {Component} from 'react';

const eventTypes = {
    NEW_PROJECT: {
        name: 'Новый проект',
        color: '#FFEB3B',
    },
    CHANGE_PROJECT_STATUS: {
        name: 'Измен статус проекта',
        color: '#E91E63',
    },
    COMPLETE_PROJECT: {
        name: 'Проект закончен',
        color: '#4CAF50',
    },
    NEW_EVENT: {
        name: 'Событие',
        color: '#90CAF9',
    },
    NEW_EFFECT: {
        name: 'Эффект',
        color: '#2196F3',
    },
    CHANGE_STRESS: {
        name: 'Стресс',
        color: '#ff6d00',
    },
    LEVEL_UP: {
        name: 'Level Up',
        color: '#CFD8DC',
    },
    EFFECT_EXPIRES: {
        name: 'Эффект закончился',
        color: '#2196F3',
    },
    FORCED_REST: {
        name: 'Пошел спать',
        color: '#999',
    },
};

var formatter = new Intl.DateTimeFormat("ru", {
    hour: "2-digit",
    minute: "2-digit",
});

export default class LogOutputComponent extends Component {
    renderMessages(messages) {
        return messages.map((item) => {
            if (!item.desc) {
                return null;
            }

            var time = formatter.format(new Date(item.timestamp));

            return (
                <div key={item.order}>
                    <span className='logs__time'>{time}</span>
                    <span className='logs__type'
                          style={{color: eventTypes[item.event_type] && eventTypes[item.event_type].color}}>
                        [{eventTypes[item.event_type].name}]:&nbsp;
                    </span>
                    {item.desc}
                </div>
            );
        });
    }

    render() {
        const messages = this.props.messages;

        return (
            <div style={{background: '#2f2d2e', height: '100%'}}>
                <div className='logs'>
                    {this.renderMessages(messages)}
                </div>
            </div>
        );
    }
}

