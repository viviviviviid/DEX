import React, { useState } from 'react';
import styled from 'styled-components';
import { ethers } from 'ethers';

const ControlSection = styled.section`
    padding: 10px;
    display: flex;
    flex-direction: column;
`;

const InfoText = styled.div`
    margin-top: 10px;
`;

const TradeControls = ({tokenAddress, onBuyLimit, onSellLimit }) => {
    const [amount, setAmount] = useState('');
    const [price, setPrice] = useState('');
    const [status, setStatus] = useState('');

    const formattedAmount = parseInt(amount);
    const formattedPrice = parseFloat(price);

    const handleTransaction = async (isBuy) => {
        try {
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            await provider.send("eth_requestAccounts", []); 
            const signer = provider.getSigner();
            const userAddress = await signer.getAddress();
            const message = "Plz Sign";
            const signature = await signer.signMessage(message);

            const response = await fetch('http://localhost:5555/api/order', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    "wallet_address": userAddress,
                    "token_address": tokenAddress,
                    amount: formattedAmount,
                    price: formattedPrice,
                    "order_type": isBuy,
                    signature,
                }),
            });

            const data = await response.json();
            setStatus(`Order processed: ${data.message}`);
        } catch (error) {
            console.error('Error processing order:', error);
            setStatus('Error processing order. Check the console.');
        }
    };

    return (
        <ControlSection>
            <div>
                <input
                    type="number"
                    value={amount}
                    onChange={e => setAmount(e.target.value)}
                    placeholder="Amount"
                />
                <input
                    type="number"
                    value={price}
                    onChange={e => setPrice(e.target.value)}
                    placeholder="Price (for limit orders)"
                />
                <button onClick={() => handleTransaction(true)}>Buy Limit</button>
                <button onClick={() => handleTransaction(false)}>Sell Limit</button>
            </div>
            <InfoText>
                {amount && price && `Total ETH: ${parseFloat(amount) * parseFloat(price).toFixed(4)}`}
                {status && <p>{status}</p>}
            </InfoText>
        </ControlSection>
    );
}

export default TradeControls;
