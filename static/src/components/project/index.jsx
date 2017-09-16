import React, {Component} from 'react';
import Progress from '../progress/index';


export default class Project extends Component {
    render() {
        const {project} = this.props;
        const COLOR_ANALYZE = '#E91E63';
        const COLOR_PROG = '#2196F3';
        const COLOR_TEST = '#4CAF50';

        return (
            <div className='entity'>
                <div className='entity__name'>{project.name}</div>
                <div className='entity__info entity__info_project'>
                    <div className='entity__image entity__image_project' />
                    <div className='entity__row'>
                        <div className='entity__stat-project'>
                            <span className='icon-stat icon-stat_analyze' />
                        </div>
                        <Progress
                            value={(project.progress.analyze / project.req_values.analyze) * 100}
                            color={COLOR_ANALYZE} />
                    </div>
                    <div className='entity__row'>
                        <div className='entity__stat-project'>
                            <span className='icon-stat icon-stat_prog' />
                        </div>
                        <Progress
                            value={(project.progress.prog / project.req_values.prog) * 100}
                            color={COLOR_PROG} />
                    </div>
                    <div className='entity__row'>
                        <div className='entity__stat-project'>
                            <span className='icon-stat icon-stat_test' />
                        </div>
                        <Progress
                            value={(project.progress.test / project.req_values.test) * 100}
                            color={COLOR_TEST} />
                    </div>
                </div>
            </div>
        );
    }
}

