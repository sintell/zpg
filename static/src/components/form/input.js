import styled from 'styled-components';
import inputIMG from './input.png';


const Input = styled.input`
    border: 0;
    background: url(${inputIMG}) no-repeat 0 0;
    padding: 0 15px;
    width: 200px;
    height: 36px;
    line-height: 36px;
    box-sizing: border-box;
    
    &:focus {
        outline: none;
    }
`;

export default Input;
