import React, {Component} from 'react';
import styled from 'styled-components';

const LogOutput = styled.div`
        height: 200px;
        padding: 5px;
        border: 4px solid red;
        font-size: 14px;
`;

export default class LogOutputComponent extends Component {
    render() {
        return (
            <LogOutput>
                <p>1. Ударила молния</p>
                <p>2. Приступил к работе с проектом блабла</p>
            </LogOutput>
        );
    }
}

