import styled from 'styled-components';
import plusIMG from './plus.png';
import plusHoverIMG from './plus-hover.png';
import minusIMG from './minus.png';
import minusHoverIMG from './minus-hover.png';


const ButtonCounter = styled.div`
    background: url(${props => props.type === 'plus' ? plusIMG : minusIMG}) no-repeat 0 0;
    width: 30px;
    height: 30px;
    display: inline-block;
    vertical-align: middle;
    cursor: pointer;
    
    &:hover {
        background-image: url(${props => props.type === 'plus' ? plusHoverIMG : minusHoverIMG});
    }
`;

export default ButtonCounter;
