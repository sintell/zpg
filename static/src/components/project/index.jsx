import React, {Component} from 'react';
import Progress from '../progress/index';


export default class Project extends Component {
    render() {
        const {project} = this.props;
        const COLOR_ANALYZE = '#E91E63';
        const COLOR_PROG = '#2196F3';
        const COLOR_TEST = '#4CAF50';

        return (
            <div className='project'>
                <div className='project__name'>{project.name}</div>
                <div className='project__info'>
                    <div className='project__image' />
                    <div className='project__stats'>
                        <div className='project__stat'>
                            <div className='project__stat-name'>Аналитика:</div>
                            <Progress
                                value={(project.process.analyze / project.req_values.analyze) * 100}
                                color={COLOR_ANALYZE} />
                        </div>
                        <div className='project__stat'>
                            <div className='project__stat-name'>Программирование:</div>
                            <Progress
                                value={(project.process.prog / project.req_values.prog) * 100}
                                color={COLOR_PROG} />
                        </div>
                        <div className='project__stat'>
                            <div className='project__stat-name'>Тестирование:</div>
                            <Progress
                                value={(project.process.test / project.req_values.test) * 100}
                                color={COLOR_TEST} />
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
