import React, {Component} from 'react';
import {connect} from 'react-redux';

const COUNT_SCORE = 10;

class FormCharacter extends Component {
    render() {
        const companyName = 'Cисьски';

        return (
            <div className="cube">
                <div>
                    <h2 className="h2">Выбери имя своего персонажа.</h2>
                    <p>Выбирай мудро, потому-что ты не сможешь изменить имя позднее.</p>
                    <p><input className="input" type='text' placeholder='Мудро' autoFocus /></p>
                    <p>Вы будите работаетев компании {companyName}</p>
                </div>
                <div>
                    <h2 className="h2">Распределите очки характеристик:</h2>
                    <div className="param-item">
                        <span className="button-counter button-counter_minus"/>
                        <span className="param-item__name">Аналитика</span>
                        <span className="button-counter button-counter_plus"/>
                    </div>
                    <div className="param-item">
                        <span className="button-counter button-counter_minus"/>
                        <span className="param-item__name">Программирование</span>
                        <span className="button-counter button-counter_plus"/>
                    </div>
                    <div className="param-item">
                        <span className="button-counter button-counter_minus"/>
                        <span className="param-item__name">Тестирование</span>
                        <span className="button-counter button-counter_plus"/>
                    </div>
                    <p>Осталось очков <output>{COUNT_SCORE}</output></p>
                </div>
                <button className="button" type="button">Начинаем работать</button>
            </div>
        );
    }
}

export default connect((state) => ({
}), {
})(FormCharacter);
