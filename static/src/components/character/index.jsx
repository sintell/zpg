import React, {Component} from 'react';
import Progress from '../progress';


export default class Character extends Component {
    render() {
        const {char} = this.props;
        const COLOR_EXP = '#607D8B';
        const COLOR_STRESS = '#FF6D00';

        return (
            <div className='entity'>
                <div className='entity__name'>Имя персонажа</div>
                <div className='entity__info'>
                    <div className='entity__image' />
                    <div className='entity__row'>
                        <span className='entity__level'>{char.level}</span>
                        <Progress value={char.exp} color={COLOR_EXP} />
                    </div>
                    <div className='entity__row'>
                        <Progress value={char.stress} color={COLOR_STRESS} />
                    </div>
                    <div className='entity__row'>
                        <div className='entity__stat'>
                            <span className='icon-stat icon-stat_analyze' />
                            <div className='entity__stat-value'>{char.skills.analyze}</div>
                        </div>
                        <div className='entity__stat'>
                            <span className='icon-stat icon-stat_prog' />
                            <div className='entity__stat-value'>{char.skills.prog}</div>
                        </div>
                        <div className='entity__stat'>
                            <span className='icon-stat icon-stat_test' />
                            <div className='entity__stat-value'>{char.skills.test}</div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

