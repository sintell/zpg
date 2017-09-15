import React, {Component} from 'react';
import {connect} from 'react-redux';

import Character from '../../components/character';
import KanbanBoard from '../../components/kanbanBoard';
import LogOutput from '../../components/logOutput';

const char = {
    level: 10,
    exp: 25,
    skills: {
        prog: 20,
        test: 30,
        analyze: 10,
    },
    stress: 20,
};

const projects = [
    {
        id: 1,
        active: false,
        name: 'Тестовое имя проекта',
        desc: 'тестовое описание проекта бла бла бла бла бла бла',
        req_values: {
            prog: 20,
            test: 30,
            analyze: 10,
        },
        process: {
            prog: 10,
            test: 0,
            analyze: 0,
        },
        status: 'todo',
    },
    {
        id: 2,
        active: true,
        name: 'Тестовое имя проекта 2',
        desc: 'тестовое описание проекта бла бла бла бла бла бла 2',
        req_values: {
            prog: 40,
            test: 30,
            analyze: 10,
        },
        process: {
            prog: 10,
            test: 20,
            analyze: 3,
        },
        status: 'prog',
    },
    {
        id: 3,
        active: false,
        name: 'Тестовое имя проекта 3',
        desc: 'тестовое описание проекта бла бла бла бла бла бла 3',
        req_values: {
            prog: 40,
            test: 30,
            analyze: 10,
        },
        process: {
            prog: 10,
            test: 30,
            analyze: 3,
        },
        status: 'todo',
    },
];

class GamePlayContainer extends Component {
    componentDidMount() {
        //fetch
    }


    render() {
        // const {isFetching, char} = this.props;

        const todoProjects = projects.filter(item => (item.status === 'todo'));
        const progProjects = projects.filter(item => (item.status === 'prog'));
        const testProjects = projects.filter(item => (item.status === 'test'));
        const analyzeProjects = projects.filter(item => (item.status === 'analyze'));
        const releasedProjects = projects.filter(item => (item.status === 'released'));

        const activeProject = projects.reduce((item, array) => {
            if (item.active) {
                return item;
            }

            return array;
        });

        return (
            <div>
                <Character char={char} activeProject={activeProject} />
                <KanbanBoard
                    todoProjects={todoProjects}
                    progProjects={progProjects}
                    testProjects={testProjects}
                    analyzeProjects={analyzeProjects}
                    releasedProjects={releasedProjects} />
                <p>Действия:</p>
                <LogOutput />
            </div>
        );
    }
}

export default connect((state) => ({
    // isFetching: state.gamePlay.data.isFetching,
}), {
})(GamePlayContainer);
