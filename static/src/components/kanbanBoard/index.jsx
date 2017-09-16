import React, {Component} from 'react';

export default class KanbanBoardComponent extends Component {
    renderProjectsTokanbanBoard(projects, type) {
        return projects.map((item) => {
            return (
                <div className={`sticker sticker_${type}`} key={item.id}>{item.name}</div>
            );
        });
    }

    render() {
        const {todoProjects, analyzeProjects, progProjects, testProjects, releasedProjects} = this.props;

        return (
            <div className='board'>
                <div className='board__col'>
                    <div className='board__title board__title_todo'>Todo</div>
                    <div className='board__content'>
                        {this.renderProjectsTokanbanBoard(todoProjects, 'todo')}
                    </div>
                </div>
                <div className='board__col'>
                    <div className='board__title board__title_analyze'>Analyze</div>
                    <div className='board__content'>
                        {this.renderProjectsTokanbanBoard(analyzeProjects, 'analyze')}
                    </div>
                </div>
                <div className='board__col'>
                    <div className='board__title board__title_prog'>In progress</div>
                    <div className='board__content'>
                        {this.renderProjectsTokanbanBoard(progProjects, 'prog')}
                    </div>
                </div>
                <div className='board__col'>
                    <div className='board__title board__title_test'>Testing</div>
                    <div className='board__content'>
                        {this.renderProjectsTokanbanBoard(testProjects, 'test')}
                    </div>
                </div>
                <div className='board__col'>
                    <div className='board__title board__title_released'>Released</div>
                    <div className='board__content'>
                        {this.renderProjectsTokanbanBoard(releasedProjects, 'released')}
                    </div>
                </div>
            </div>
        );
    }
}

