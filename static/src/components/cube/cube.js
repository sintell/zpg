import styled from 'styled-components';


const borderWidth = 10;

const Cube = styled.div`
    background: #EEEEEE;
    boxx-sizingL border-box;
    width: 600px;
    padding: 20px;
    margin: 50px auto 0;
    position: relative; 
    
    &:before, &:after {
        content: '';
        position: absolute;
    }
    
    &:before {
        background-color: #FAFAFA;
        width: ${borderWidth}px;
        height: 100%;
        left: 100%;
        top: ${borderWidth / 2}px;
        transform: skewY(45deg);
    }
    
    &:after {
        background-color: #E0E0E0;
        width: 100%;
        height: ${borderWidth}px;
        left: ${borderWidth / 2}px;
        top: 100%;
        transform: skewX(45deg);
    }
    
`;

export default Cube;
