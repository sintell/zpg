import React, {Component} from 'react';
import Progress from '../progress';


export default class Character extends Component {
    render() {
        const {char} = this.props;
        const COLOR_EXP = '#607D8B';
        const COLOR_STRESS = '#FF6D00';

        return (
            <div className='char'>
                <div className='char__name'>Имя персонажа</div>
                <div className='char__info'>
                    <div className='char__avatar' />
                    <div className='char__stats'>
                        <div className='char__stat'>
                            <span className='char__level'>{char.level}</span>
                            <Progress value={char.exp} color={COLOR_EXP} />
                        </div>
                        <div className='char__stat'>
                            <Progress value={char.stress} color={COLOR_STRESS} />
                        </div>
                        <div className='char__stat'>Программирование: {char.skills.prog}</div>
                        <div className='char__stat'>Аналитика: {char.skills.analyze}</div>
                        <div className='char__stat'>Тестирование: {char.skills.test}</div>
                    </div>
                </div>
            </div>
        );
    }
}

