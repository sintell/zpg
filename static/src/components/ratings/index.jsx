import React, {Component} from 'react';
import {connect} from 'react-redux';

import {fetchRating} from '../../modules/rating';

class Ratings extends Component {
    componentDidMount() {
        this.props.fetchRating();
    }

    renderItem(item, index) {
        return (
            <div>
                <div className='rating-row' key={item.id}>
                    <div className='rating-col rating-col_place'>{index + 1}</div>
                    <div className='rating-col rating-col_image'>
                        <span className={`entity__image entity__image_${(item.id % 3) + 1} entity__image_rating`} />
                    </div>
                    <div className='rating-col rating-col_name'>{item.name}</div>
                    <div className='rating-col rating-col_level'>{item.level}</div>
                    <div className='rating-col rating-col_place'>
                        <div className='rating-skills'>
                            <div className='rating-skill'>
                                <span className='icon-stat icon-stat_analyze' /><br />
                                1
                            </div>
                            <div className='rating-skill'>
                                <span className='icon-stat icon-stat_prog' /><br />
                                1
                            </div>
                            <div className='rating-skill'>
                                <span className='icon-stat icon-stat_test' /><br />
                                1
                            </div>
                        </div>
                    </div>
                </div>
                <div className='separate' />
            </div>
        );
    }

    render() {
        const {items} = this.props;

        return (
            <div className='rating-container'>
                <div className='rating-row rating-row_head'>
                    <div className='rating-col rating-col_place'>Место</div>
                    <div className='rating-col rating-col_char'>Персонаж</div>
                    <div className='rating-col rating-col_level'>Уровень</div>
                    <div className='rating-col rating-col_place'>Характеристики</div>
                </div>
                {items.map(this.renderItem)}
            </div>
        );
    }
}

export default connect((state) => ({
    items: state.rating.items,
}), {
    fetchRating,
})(Ratings);
