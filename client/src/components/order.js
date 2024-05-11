import React, { useState } from 'react';
import styled from 'styled-components';

const ControlSection = styled.section`
    padding: 10px;
    display: flex;    
    align-items: center;
`;


const TradeControls = ({ onBuy, onSell }) => {
    const [amount, setAmount] = useState('');

    return (
        <ControlSection>
            <input
                type="number"
                value={amount}
                onChange={e => setAmount(e.target.value)}
                placeholder="Amount"
            />
            <button onClick={() => onBuy(amount)}>Buy</button>
            <button onClick={() => onSell(amount)}>Sell</button>
        </ControlSection>
    );
}

export default TradeControls;
