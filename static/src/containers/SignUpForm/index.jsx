import React, {Component} from 'react';
import {connect} from 'react-redux';

import {fetchBoxPlotStatuses} from '../../modules/boxPlotStatusesData';
import format from '../../modules/formatValues';
import {notify} from '../../modules/notificationManager';
import t from '../../modules/translations';

import BoxPlot from '../../components/BoxPlot';

class SignUpFormContainer extends Component {
    componentDidMount() {
        this.props.fetchBoxPlotStatuses();
    }
    
    handleEjectionMaxClick = (datum) => {
        if (datum.values.ejectionKeys.max) {
            window.open(`https://jira.hh.ru/browse/${datum.values.ejectionKeys.max}`);
        }
    };
    
    handleEjectionMinClick = (datum) => {
        if (datum.values.ejectionKeys.min) {
            window.open(`https://jira.hh.ru/browse/${datum.values.ejectionKeys.min}`);
        }
    };
    
    handlerHoverBar = (datum) => {
        if (datum) {
            const message = `${datum.title}: (max: ${format.toFixed(datum.values.max)},
                                              min: ${format.toFixed(datum.values.min)},
                                              median: ${format.toFixed(datum.values.median)},
                                              ${t('quartiles.25')}: ${format.toFixed(datum.values.quartiles.min)},
                                              ${t('quartiles.75')}: ${format.toFixed(datum.values.quartiles.max)},
                                              ${t('boxplot.count')}: ${datum.values.count})`;
            this.props.notify({message, show: true});
        } else {
            this.props.notify({message: '', show: false});
        }
    }
    
    render() {
        const {items, dashboard, isFetching} = this.props;
        
        return (
            <BoxPlot
                items={items}
                isFetching={isFetching}
                dashboard={dashboard}
                axesProps={axesProps}
                handleEjectionMaxClick={this.handleEjectionMaxClick}
                handleEjectionMinClick={this.handleEjectionMinClick}
                handlerHoverBar={this.handlerHoverBar} />
        );
    }
}

export default connect((state) => ({
    items: state.boxPlotStatusesData.items,
    isFetching: state.boxPlotStatusesData.isFetching,
}), {
    fetchBoxPlotStatuses,
    notify,
})(BoxPlotTeamContainer);
