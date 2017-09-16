import React, {Component} from 'react';
import {connect} from 'react-redux';

import {fetchRating} from '../../modules/rating';

class Ratings extends Component {
    componentDidMount() {
        this.props.fetchRating();
    }

    renderItem(item) {
        return (
            <div key={item.id}>
                <p>{item.name}</p>
                <p>Уровень: {item.level}</p>
            </div>
        );
    }

    render() {
        const {items} = this.props;

        return (
            <div>
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
