import React, {Component} from 'react';


export default class LogOutputComponent extends Component {
    renderMessages(messages) {
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
            WAS_EVENT: {
                name: 'Событие',
                color: '#90CAF9',
            },
            USE_EFFECT: {
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
        };

        return messages.map((item, index) => {
            return (
                <div key={index}>
                    <span className='logs__time'>11:05 </span>
                    <span className='logs__type' style={{color: eventTypes[item.type].color}}>
                        [{eventTypes[item.type].name}]:&nbsp;
                    </span>
                    {item.text}
                </div>
            );
        });
    }

    render() {
        let messages = [
            {
                text: 'Приступил к работе с проектом блабла',
                type: 'NEW_PROJECT',
            },
            {
                text: 'Приступил к работе с проектом блабла',
                type: 'CHANGE_PROJECT_STATUS',
            },
            {
                text: 'Приступил к работе с проектом блабла',
                type: 'COMPLETE_PROJECT',
            },
            {
                text: 'Приступил к работе с проектом блабла',
                type: 'USE_EFFECT',
            },
        ];

        return (
            <div className='logs-wrapper'>
                <div className='cube' style={{background: '#2f2d2e'}}>
                    <div className='logs'>
                        {this.renderMessages(messages)}
                    </div>
                </div>
            </div>
        );
    }
}

