import React, {Component} from 'react';
import {connect} from 'react-redux';

import {createChar} from '../../modules/character';

const COUNT_SCORE = 10;
const MIN_COUNT = 1;

class FormCharacter extends Component {
    state = {
        name: 'Мудро',
        prog: MIN_COUNT,
        test: MIN_COUNT,
        analyze: MIN_COUNT,
        countScore: COUNT_SCORE - MIN_COUNT * 3,
    };

    submit = (event) => {
        event.preventDefault();
        this.props.createChar(this.state);
    }

    changeInput = (event) => {
        this.setState({
            name: event.target.value,
        });
    }

    changeAnalize = (event) => {
        let state = {};
        const stateKey = event.target.getAttribute('name');

        if (event.target.getAttribute('type') === 'plus' &&
            this.state.countScore <= COUNT_SCORE &&
            this.state.countScore > 0) {
            state[stateKey] = this.state[stateKey] + 1;
            state.countScore = this.state.countScore - 1;
            this.setState(state);
        } else if (event.target.getAttribute('type') === 'minus' && this.state[stateKey] > MIN_COUNT) {
            state[stateKey] = this.state[stateKey] - 1;
            state.countScore = this.state.countScore + 1;
            this.setState(state);
        }
    }

    render() {
        const companyName = 'Большие Тити';

        return (
            <form className='popup popup_reg' onSubmit={this.submit}>
                <div className='cube'>
                    <div className='form-padding'>
                        <div className='create-char'>
                            <h2 className='h2'>Выбери имя своего персонажа.</h2>
                            <p className='create-char__subtitle'>
                                Выбирай мудро, потому-что ты не сможешь изменить имя позднее.
                            </p>
                            <p>
                                <input
                                    className='input'
                                    type='text'
                                    value={this.state.name}
                                    onChange={this.changeInput}
                                    autoFocus />
                            </p>
                            <p>Вы будете работаеть в компании "{companyName}"</p>
                            <div className='separate'/>
                            <h2 className='h2'>Распределите очки характеристик:</h2>
                            <div className='params-list'>
                                <div className='params-list__item'>
                                    <div className='param-item'>
                                        <span
                                            className='button-counter button-counter_minus'
                                            type='minus'
                                            name='analyze'
                                            onClick={this.changeAnalize} />
                                        <span className='param-item__name'>
                                            Аналитика
                                            <span className='param-value'>{this.state.analyze}</span>
                                        </span>
                                        <span
                                            className='button-counter button-counter_plus'
                                            type='plus'
                                            name='analyze'
                                            onClick={this.changeAnalize} />
                                    </div>
                                    <div className='param-item'>
                                        <span
                                            className='button-counter button-counter_minus'
                                            type='minus'
                                            name='prog'
                                            onClick={this.changeAnalize} />
                                        <span className='param-item__name'>
                                            Программирование
                                            <span className='param-value'>{this.state.prog}</span>
                                        </span>
                                        <span
                                            className='button-counter button-counter_plus'
                                            type='plus'
                                            name='prog'
                                            onClick={this.changeAnalize} />
                                    </div>
                                    <div className='param-item'>
                                        <span
                                            className='button-counter button-counter_minus'
                                            type='minus'
                                            name='test'
                                            onClick={this.changeAnalize} />
                                        <span className='param-item__name'>
                                            Тестирование
                                            <span className='param-value'>{this.state.test}</span>
                                        </span>
                                        <span
                                            className='button-counter button-counter_plus'
                                            type='plus'
                                            name='test'
                                            onClick={this.changeAnalize} />
                                    </div>
                                </div>
                                <div className='params-list__item'>
                                    <p>Осталось очков:</p>
                                    <span className='param-value param-value_big'>{this.state.countScore}</span>
                                </div>
                            </div>
                            <div className='separate' />
                            <div className='form-actions'>
                                <button className='button' type='submit'>Начинаем работать</button>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        );
    }
}

export default connect((state) => ({
    isFetching: state.char.isFetching,
}), {
    createChar,
})(FormCharacter);
