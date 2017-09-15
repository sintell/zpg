import React, {Component} from 'react';
import styled from 'styled-components';

import Progress from '../progress';

const CharacterWrapper = styled.div`
        display: flex;
`;

const CharacterImg = styled.div`
        width: 5vw;
        height: 5vw;
        background: red;
`;

const CharacterSkills = styled.div`
        padding: 10px;
        border: 1px solid blanchedalmond;
        flex: 1;
`;

const CharacterSkill = styled.div`
        display: flex;
        margin-bottom: 10px;
`;

const CurrentProject = styled.div`
        padding: 10px;
        flex: 2;
        background: green;
`;

export default class Character extends Component {
    renderActiveProject(project) {
        if (!project) {
            return null;
        }

        return (
            <div>
                <p>{project.name}</p>
                <p>{project.desc}</p>
                <p>Статус:</p>
                <p>
                    Программирование: ({project.process.prog}/{project.req_values.prog})
                    <Progress max='100' value={(project.process.prog / project.req_values.prog) * 100} />
                </p>
                <p>
                    Аналитика: ({project.process.analyze}/{project.req_values.analyze})
                    <Progress max='100' value={(project.process.analyze / project.req_values.analyze) * 100} />
                </p>
                <p>
                    Тестирование: ({project.process.test}/{project.req_values.test})
                    <Progress max='100' value={(project.process.test / project.req_values.test) * 100} />
                </p>
            </div>
        );
    }

    render() {
        const {char, activeProject} = this.props;

        return (
            <CharacterWrapper>
                <CharacterImg />
                <CharacterSkills>
                    <CharacterSkill>Уровень: {char.level}</CharacterSkill>
                    <CharacterSkill>Программирование: {char.skills.prog}</CharacterSkill>
                    <CharacterSkill>Аналитика: {char.skills.analyze}</CharacterSkill>
                    <CharacterSkill>Тестирование: {char.skills.test}</CharacterSkill>
                    <CharacterSkill>Стресс: {char.stress} <Progress max='100' value={char.stress} /></CharacterSkill>
                    <CharacterSkill>Опыт: {char.exp} <Progress max='100' value={char.exp} /></CharacterSkill>
                </CharacterSkills>
                <CurrentProject>
                    <CharacterSkill>Проект в работе:</CharacterSkill>
                    {this.renderActiveProject(activeProject)}
                </CurrentProject>
            </CharacterWrapper>
        );
    }
}

