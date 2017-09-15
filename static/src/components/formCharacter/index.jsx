import React, {Component} from 'react';
import {connect} from 'react-redux';

import Cube from '../cube/cube';

import { H2 } from '../typography/typography';
import Input from '../form/input';
import Button from '../form/button';
import ButtonCounter from '../form/buttonCounter';

import { ParamItem, ParamName } from './styles';

const COUNT_SCORE = 10;

class FormCharacter extends Component {
    render() {
        const companyName = 'Cисьски';

        return (
            <Cube>
                <div>
                    <H2>Выбери имя своего персонажа.</H2>
                    <p>Выбирай мудро, потому-что ты не сможешь изменить имя позднее.</p>
                    <p><Input type='text' placeholder='Мудро' autoFocus /></p>
                    <p>Вы будите работаетев компании {companyName}</p>
                </div>
                <div>
                    <h2>Распределите очки характеристик:</h2>
                    <div>
                        <ParamItem>
                            <ButtonCounter type={'minus'}/>
                            <ParamName>Аналитика</ParamName>
                            <ButtonCounter type={'plus'}/>
                        </ParamItem>
                        <ParamItem>
                            <ButtonCounter type={'minus'}/>
                            <ParamName>Программирование</ParamName>
                            <ButtonCounter type={'plus'}/>
                        </ParamItem>
                        <ParamItem>
                            <ButtonCounter type={'minus'}/>
                            <ParamName>Тестирование</ParamName>
                            <ButtonCounter type={'plus'}/>
                        </ParamItem>
                    </div>
                    <p>Осталось очков <output>{COUNT_SCORE}</output></p>
                </div>
                <Button>Go</Button>
            </Cube>
        );
    }
}

export default connect((state) => ({
}), {
})(FormCharacter);
