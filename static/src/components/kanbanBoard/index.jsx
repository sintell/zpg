import React, {Component} from 'react';


export default class KanbanBoardComponent extends Component {
    renderProjectsTokanbanBoard(projects, color) {
        return projects.map((item) => {
            return (
                <div className={`sticker sticker_${color}`} key={item.id}>{item.name}</div>
            );
        });
    }

    render() {
        const {todoProjects, analyzeProjects, progProjects, testProjects, releasedProjects} = this.props;
        const COLORS = {
            todo: 'yellow',
            analyze: 'red',
            prog: 'blue',
            test: 'green',
            released: 'orange',
        };

        return (
            <div className='cube'>
                <div className='board'>
                    <div className='board__col'>
                        <div className='board__title board__title_todo'>Todo</div>
                        <div className='board__content'>
                            {this.renderProjectsTokanbanBoard(todoProjects, COLORS.todo)}
                        </div>
                    </div>
                    <div className='board__col'>
                        <div className='board__title board__title_analyze'>Analyze</div>
                        <div className='board__content'>
                            {this.renderProjectsTokanbanBoard(analyzeProjects, COLORS.analyze)}
                        </div>
                    </div>
                    <div className='board__col'>
                        <div className='board__title board__title_prog'>In progress</div>
                        <div className='board__content'>
                            {this.renderProjectsTokanbanBoard(progProjects, COLORS.prog)}
                        </div>
                    </div>
                    <div className='board__col'>
                        <div className='board__title board__title_test'>Testing</div>
                        <div className='board__content'>
                            {this.renderProjectsTokanbanBoard(testProjects, COLORS.test)}
                        </div>
                    </div>
                    <div className='board__col'>
                        <div className='board__title board__title_released'>Released</div>
                        <div className='board__content'>
                            {this.renderProjectsTokanbanBoard(releasedProjects, COLORS.released)}
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

