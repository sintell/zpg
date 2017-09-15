import React, {Component} from 'react';
import styled from 'styled-components';

const KanbanBoard = styled.div`
        display: flex;
        border-top: 1px dotted black;
        border-bottom: 1px dotted black;
        border-rigth: 1px dotted black;
        margin-bottom: 25px;
`;

const KanbanBoardCol = styled.div`
        width: 20%;
        padding: 10px;
        border-left: 1px dotted black;
        box-sizing: border-box;
`;

const KanbanBoardProject = styled.div`
        width: 100%;
        background: red;
`;

export default class KanbanBoardComponent extends Component {
    renderProjectsTokanbanBoard(projects) {
        return projects.map((item) => {
            return (
                <KanbanBoardProject key={item.id}>
                    <p>{item.name}</p>
                    <p>{item.desc}</p>
                </KanbanBoardProject>
            );
        });
    }

    render() {
        const {todoProjects, analyzeProjects, progProjects, testProjects, releasedProjects} = this.props;

        return (
            <KanbanBoard>
                <KanbanBoardCol>
                    <p>Todo</p>
                    {this.renderProjectsTokanbanBoard(todoProjects)}
                </KanbanBoardCol>
                <KanbanBoardCol>
                    <p>Аналитика</p>
                    {this.renderProjectsTokanbanBoard(analyzeProjects)}
                </KanbanBoardCol>
                <KanbanBoardCol>
                    <p>Программирование</p>
                    {this.renderProjectsTokanbanBoard(progProjects)}
                </KanbanBoardCol>
                <KanbanBoardCol>
                    <p>Тестирование</p>
                    {this.renderProjectsTokanbanBoard(testProjects)}
                </KanbanBoardCol>
                <KanbanBoardCol>
                    <p>Релиз</p>
                    {this.renderProjectsTokanbanBoard(releasedProjects)}
                </KanbanBoardCol>
            </KanbanBoard>
        );
    }
}

