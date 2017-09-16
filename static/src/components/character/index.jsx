import React, {Component} from 'react';
import Progress from '../progress';

const COLOR_EXP = '#607D8B';
const COLOR_STRESS = '#FF6D00';

export default class Character extends Component {
    render() {
        const {level, exp, stress, skills, name} = this.props.char;

        return (
            <div className='entity'>
                <div className='entity__name'>{name}</div>
                <div className='entity__info'>
                    <div className='entity__image' />
                    <div className='entity__row'>
                        <span className='entity__level'>{level}</span>
                        <Progress value={exp} color={COLOR_EXP} />
                    </div>
                    <div className='entity__row'>
                        <Progress value={stress} color={COLOR_STRESS} />
                    </div>
                    <div className='entity__row'>
                        <div className='entity__stat'>
                            <span className='icon-stat icon-stat_analyze' />
                            <div className='entity__stat-value'>{skills.analyze}</div>
                        </div>
                        <div className='entity__stat'>
                            <span className='icon-stat icon-stat_prog' />
                            <div className='entity__stat-value'>{skills.prog}</div>
                        </div>
                        <div className='entity__stat'>
                            <span className='icon-stat icon-stat_test' />
                            <div className='entity__stat-value'>{skills.test}</div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

