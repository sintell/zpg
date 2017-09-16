import React, {Component} from 'react';
import {connect} from 'react-redux';

import Progress from '../progress';

import {sendToSleep} from '../../modules/character';

const COLOR_EXP = '#607D8B';
const COLOR_STRESS = '#FF6D00';

class Character extends Component {
    sleep = () => {
        this.props.sendToSleep();
    }

    render() {
        const {id, level, exp, stress, skills, name, company} = this.props.char;
        const maxExp = 100;
        const maxStress = 100;
        let charImageIndex = (id % 3) + 1;

        return (
            <div className='entity'>
                <div className='entity__name'>{name} ({company})</div>
                <div className='entity__info'>
                    <div className={`entity__image entity__image_${charImageIndex}`} />
                    <div className='entity__row'>
                        <span className='entity__level'>{level}</span>
                        <Progress value={exp} maxValue={maxExp} color={COLOR_EXP} />
                    </div>
                    <div className='entity__row'>
                        <Progress value={stress} maxValue={maxStress} color={COLOR_STRESS} />
                        <span
                            className='sleep-button'
                            title='Пойти поспать, чтобы снять стресс'
                            onClick={this.sleep} />
                    </div>
                    <div className='entity__row'>
                        <div className='entity__stat'>
                            <span
                                className='icon-stat icon-stat_analyze'
                                title='Аналитика – влияет на скорость выполнения задач по аналитике' />
                            <div className='entity__stat-value'>{skills.analyze}</div>
                        </div>
                        <div className='entity__stat'>
                            <span
                                className='icon-stat icon-stat_prog'
                                title='Программирование – влияет на скорость выполнения задач по программированию' />
                            <div className='entity__stat-value'>{skills.prog}</div>
                        </div>
                        <div className='entity__stat'>
                            <span
                                className='icon-stat icon-stat_test'
                                title='Тестирование – влияет на скорость тестирования' />
                            <div className='entity__stat-value'>{skills.test}</div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default connect(state => ({
    sleep: state.charData.sleep,
}), {
    sendToSleep,
})(Character);
