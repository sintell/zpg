import React, {Component} from 'react';
import {connect} from 'react-redux';

import {fetchRating} from '../../modules/rating';

class Ratings extends Component {
    componentDidMount() {
        this.props.fetchRating();
    }

    renderItem(item, index) {
        return (
            <tr key={item.id}>
                <td className='rating-table__col-place'>{index + 1}</td>
                <td width='60'>
                    <span className={`entity__image entity__image_${(item.id % 3) + 1} entity__image_rating`} />
                </td>
                <td>{item.name}</td>
                <td className='rating-table__col-level'>{item.level}</td>
                <td>
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
                </td>
            </tr>
        );
    }

    render() {
        const {items} = this.props;

        return (
            <div className='rating-container'>
                <table className='rating-table'>
                    <thead>
                        <tr>
                            <th width='90'>Место</th>
                            <th colSpan='2'>Персонаж</th>
                            <th>Уровень</th>
                            <th>Характеристики</th>
                        </tr>
                    </thead>
                    <tbody>
                        {items.map(this.renderItem)}
                    </tbody>
                </table>
            </div>
        );
    }
}

export default connect((state) => ({
    items: state.rating.items,
}), {
    fetchRating,
})(Ratings);
