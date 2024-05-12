import React, { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import { useParams } from 'react-router-dom';
import { Web3Provider } from '@ethersproject/providers';
import tokenABI from '../abi/token_ABI';
import contractABI from '../abi/contract_ABI';
import TradeControls from './order';
import styled from 'styled-components';

const DepositSection = styled.section`
    padding: 10px;
    display: flex;
`;

const Dex = () => {
    const { contractAddress, tokenAddress } = useParams();
    const [contract, setContract] = useState(null);
    const [tokenContract, setTokenContract] = useState(null);
    const [etherAmount, setEtherAmount] = useState('');
    const [tokenAmount, setTokenAmount] = useState('');
    const [isTokenApproved, setIsTokenApproved] = useState(false);

    useEffect(() => {
        const provider = new Web3Provider(window.ethereum);
        const signer = provider.getSigner();
        const contractInstance = new ethers.Contract(contractAddress, contractABI, signer);
        const tokenInstance = new ethers.Contract(tokenAddress, tokenABI, signer);

        setContract(contractInstance);
        setTokenContract(tokenInstance);
    }, [contractAddress, tokenAddress]);

    const handleDepositEther = async () => {
        if (!etherAmount) return;
        try {
            const tx = await contract.depositETH({ value: ethers.utils.parseUnits(etherAmount, 'ether') });
            await tx.wait();
            console.log('Ether deposit successful');
            setEtherAmount('');
        } catch (error) {
            console.error('Ether deposit failed:', error);
        }
    };

    const handleApproveToken = async () => {
        if (!tokenAmount) return;
        try {
            const amountInWei = ethers.utils.parseUnits(tokenAmount, 'ether');
            const tx = await tokenContract.approve(contract.address, amountInWei);
            await tx.wait();
            console.log('Token approval successful');
            setIsTokenApproved(true);
        } catch (error) {
            console.error('Token approval failed:', error);
            setIsTokenApproved(false);
        }
    };

    const handleDepositToken = async () => {
        if (!tokenAmount || !isTokenApproved) return;
        try {
            const amountInWei = ethers.utils.parseUnits(tokenAmount, 'ether');
            const tx = await contract.depositToken(amountInWei);
            await tx.wait();
            console.log('Token deposit successful');
            setTokenAmount('');
            setIsTokenApproved(false); // Reset approval after deposit
        } catch (error) {
            console.error('Token deposit failed:', error);
        }
    };

    const handleBuyLimit = (amount, price) => {
        console.log('Buy limit order:', amount, price);
    };

    const handleSellLimit = (amount, price) => {
        console.log('Sell limit order:', amount, price);
    };

    // const handleBuyMarket = (amount) => {
    //     console.log('Buy market order:', amount);
    // };

    // const handleSellMarket = (amount) => {
    //     console.log('Sell market order:', amount);
    // };


    return (
        <div className="Dex">
            <DepositSection>
                <input
                    type="text"
                    value={etherAmount}
                    onChange={e => setEtherAmount(e.target.value)}
                    placeholder="Amount of ETH to deposit"
                />
                <button onClick={handleDepositEther}>Deposit ETH</button>
                <input
                    type="text"
                    value={tokenAmount}
                    onChange={e => setTokenAmount(e.target.value)}
                    placeholder="Amount of Tokens to deposit"
                />
                {!isTokenApproved ? (
                    <button onClick={handleApproveToken}>Approve Token</button>
                ) : (
                    <button onClick={handleDepositToken}>Deposit Token</button>
                )}
            </DepositSection>
            <TradeControls 
                tokenAddress={tokenAddress}
                onBuyLimit={handleBuyLimit} 
                onSellLimit={handleSellLimit} 
                // onBuyMarket={handleBuyMarket} 
                // onSellMarket={handleSellMarket} 
            />
        </div>
    );
};

export default Dex;
