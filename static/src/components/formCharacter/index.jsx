import React, {Component} from 'react';
import {connect} from 'react-redux';

const COUNT_SCORE = 10;

class FormCharacter extends Component {
    render() {
        const companyName = 'Большие Тити';

        return (
            <div className="popup popup_reg">
                <div className="cube">
                    <div className="form-padding">
                        <div class="create-char">
                            <h2 className="h2">Выбери имя своего персонажа.</h2>
                            <p className="create-char__subtitle">
                                Выбирай мудро, потому-что ты не сможешь изменить имя позднее.
                            </p>
                            <p><input className="input" type='text' placeholder='Мудро' autoFocus /></p>
                            <p>Вы будете работаеть в компании "{companyName}"</p>
                            <div className="separate"/>

                            <h2 className="h2">Распределите очки характеристик:</h2>
                            <div className="params-list">
                                <div className="params-list__item">
                                    <div className="param-item">
                                        <span className="button-counter button-counter_minus"/>
                                        <span className="param-item__name">
                                            Аналитика
                                            <span className="param-value">1</span>
                                        </span>
                                        <span className="button-counter button-counter_plus"/>
                                    </div>
                                    <div className="param-item">
                                        <span className="button-counter button-counter_minus"/>
                                        <span className="param-item__name">
                                            Программирование
                                            <span className="param-value">1</span>
                                        </span>
                                        <span className="button-counter button-counter_plus"/>
                                    </div>
                                    <div className="param-item">
                                        <span className="button-counter button-counter_minus"/>
                                        <span className="param-item__name">
                                            Тестирование
                                            <span className="param-value">1</span>
                                        </span>
                                        <span className="button-counter button-counter_plus"/>
                                    </div>
                                </div>
                                <div className="params-list__item">
                                    <p>Осталось очков:</p>
                                    <span className="param-value param-value_big">{COUNT_SCORE}</span>
                                </div>
                            </div>
                            <div className="separate"/>
                            <div className="form-actions">
                                <button className="button" type="button">Начинаем работать</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default connect((state) => ({
}), {
})(FormCharacter);
