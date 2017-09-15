import React, {Component} from 'react';
import styled from 'styled-components';

const LogOutput = styled.div`
        width: 100%;
        height: 200px;
        padding: 5px;
        font-size: 14px;
        box-sizing: border-box;
        color: #39ff00;
        font-family: 'Ubuntu Mono', monospace;
`;

export default class LogOutputComponent extends Component {
    render() {
        return (
            <div className='cube' style={{margin: '0 30px', background: '#2f2d2e'}}>
                <div className='board'>
                    <LogOutput>
                        <p>1. Ударила молния</p>
                        <p>2. Приступил к работе с проектом блабла</p>
                    </LogOutput>
                </div>
            </div>
        );
    }
}

