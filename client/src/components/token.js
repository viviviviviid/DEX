import React, { useState } from 'react';

const DeployPage = ({ onSubmit }) => {
    const [tokenName, setTokenName] = useState('');
    const [tokenAmount, setTokenAmount] = useState('');

    function handleSubmit(event) {
        event.preventDefault();
        onSubmit(tokenName, tokenAmount);
    }

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" value={tokenName} onChange={e => setTokenName(e.target.value)} placeholder="Token Name" />
            <input type="number" value={tokenAmount} onChange={e => setTokenAmount(e.target.value)} placeholder="Amount" />
            <button type="submit">Issue Token</button>
        </form>
    );
}
export default DeployPage;
