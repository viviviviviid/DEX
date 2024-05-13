import React, { useState } from 'react';
import styled from 'styled-components';
import { ethers } from 'ethers';
import contractABI from '../abi/contract_ABI';

const ControlSection = styled.section`
    padding: 10px;
    display: flex;
    flex-direction: column;
`;

const InfoText = styled.div`
    margin-top: 10px;
`;

const TradeControls = ({contractAddress, tokenAddress, onBuyLimit, onSellLimit }) => {
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
    
            const reqBody = JSON.stringify({
                "wallet_address": userAddress,
                "token_address": tokenAddress,
                amount: formattedAmount,
                price: formattedPrice,
                "is_buy": isBuy,
                signature,
            });
    
            const signResponse = await fetch('http://localhost:5555/api/sign', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: reqBody,
            });
    
            const signData = await signResponse.json();
            console.log(signData);
    
            if (signResponse.ok && signData.okayToProceed) {
              
                const contract = new ethers.Contract(contractAddress, contractABI, signer);
                const tx = await contract.placeOrder(
                    formattedAmount.toString(),
                    ethers.utils.parseUnits(formattedPrice.toString(), 'ether'),
                    isBuy
                );
    
                const orderResponse = await fetch('http://localhost:5555/api/order', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: reqBody, 
                });
    
                const orderData = await orderResponse.json();
                console.log(orderData);
    
                if (orderResponse.ok) {
                    await tx.wait();
                    setStatus(`Order placed on-chain and recorded: ${tx.hash}`);
                } else {
                    setStatus(`Order placed on-chain but failed to record: ${orderData.message}`);
                }
            } else {
                setStatus(`Failed to place order: ${signData.message}`);
            }
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
