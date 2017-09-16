import React, {Component} from 'react';
import {connect} from 'react-redux';

import Character from '../../components/character';
import Project from '../../components/project';
import KanbanBoard from '../../components/kanbanBoard';
import LogOutput from '../../components/logOutput';

import {fetchChar, TICK_MS} from '../../modules/character';

class GamePlayContainer extends Component {
    componentDidMount() {
        this.tickTimeoutID = window.setInterval(() => {
            console.log(new Date());
            this.props.fetchChar();
        }, TICK_MS);
    }

    componentWillUnmount() {
        window.clearTimeout(this.tickTimeoutID);
    }

    render() {
        const {char, projects} = this.props;
        if (!char) {
            return null;
        }

        const todoProjects = projects.filter(item => (item.status === 'TODO'));
        const progProjects = projects.filter(item => (item.status === 'PROG'));
        const testProjects = projects.filter(item => (item.status === 'TESTING'));
        const analyzeProjects = projects.filter(item => (item.status === 'ANALYZE'));
        const releasedProjects = projects.filter(item => (item.status === 'RELEASED'));

        const activeProject = projects.reduce((result, item) => {
            if (item.active) {
                return item;
            }

            return result;
        }, false);

        return (
            <div className='container'>
                <div className='game'>
                    <div className='game-info'>
                        <Character char={char} />
                        { activeProject ? <Project project={activeProject} /> : ''}
                    </div>
                    <div className='game-board'>
                        <KanbanBoard
                            todoProjects={todoProjects}
                            progProjects={progProjects}
                            testProjects={testProjects}
                            analyzeProjects={analyzeProjects}
                            releasedProjects={releasedProjects} />
                    </div>
                    <div className='game-logs'>
                        <LogOutput />
                    </div>
                </div>
            </div>
        );
    }
}

export default connect((state) => ({
    char: state.charData.char,
    projects: state.charData.projects,
}), {
    fetchChar,
})(GamePlayContainer);
