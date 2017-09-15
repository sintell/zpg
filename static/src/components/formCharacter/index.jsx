import React, {Component} from 'react';
import {connect} from 'react-redux';

const COUNT_SCORE = 10;

class FormCharacter extends Component {
    render() {
        const companyName = 'Cисьски';

        return (
            <div>
                <fieldset>
                    <h2>Выбери имя своего персонажа. Выбирай мудро, потому-что ты не сможешь изменить имя позднее.</h2>
                    <p><input type='text' placeholder='Мудро' autoFocus /></p>
                    <p>Вы будите работаетев компании {companyName}</p>
                </fieldset>
                <fieldset>
                    <h2>Распределите очки характеристик:</h2>
                    <ol>
                        <li>
                            <label>
                                <input type='number' name='development' />
                                <span>Программирование</span>
                            </label>
                        </li>
                        <li>
                            <input type='number' name='analytics' />
                            <span>Аналитика</span>
                        </li>
                        <li>
                            <input type='number' name='communion' />
                            <span>Общение</span>
                        </li>
                    </ol>

                    <p>Осталось очков <output>{COUNT_SCORE}</output></p>
                </fieldset>
                <button>Go</button>
            </div>
        );
    }
}

export default connect((state) => ({
}), {
})(FormCharacter);
