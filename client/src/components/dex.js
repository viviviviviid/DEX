import React, { useState, useEffect } from 'react';
import { parseUnits, Contract } from 'ethers';
import { useParams } from 'react-router-dom';
import { Web3Provider } from '@ethersproject/providers';
import tokenABI from '../abi/token_ABI';
import contractABI from '../abi/contract_ABI';
import { ethers } from 'ethers';

const TokenDeposit = () => {
    const { contractAddress, tokenAddress } = useParams();
    const [contract, setContract] = useState(null);
    const [tokenContract, setTokenContract] = useState(null);
    const [etherAmount, setEtherAmount] = useState('');
    const [tokenAmount, setTokenAmount] = useState('');
    const [isTokenApproved, setIsTokenApproved] = useState(false);

    // 컨트랙트 초기화
    useEffect(() => {
        const provider = new Web3Provider(window.ethereum);
        const signer = provider.getSigner();
        const contractInstance = new Contract(contractAddress, contractABI, signer);
        const tokenInstance = new Contract(tokenAddress, tokenABI, signer);

        setContract(contractInstance);
        setTokenContract(tokenInstance);
    }, [contractAddress, tokenAddress]);

    // 이더 예치
    const handleDepositEther = async () => {
        if (!etherAmount) return;
        try {
            const tx = await contract.depositETH({
                value: ethers.utils.parseUnits(etherAmount, 'ether')
            });
            const receipt = await tx.wait();
            console.log('Ether deposit successful:', receipt);
            setEtherAmount('');
        } catch (error) {
            console.error('Ether deposit failed:', error);
        }
    };
    
    const handleApproveToken = async () => {
        if (!tokenAmount) return;
        try {
            const amountInWei = ethers.utils.parseUnits(tokenAmount, 'ether');
            await tokenContract.approve(contract.address, amountInWei);
            console.log('Token approval successful for', amountInWei.toString());
            setIsTokenApproved(true);
        } catch (error) {
            console.error('Token approval failed:', error);
        }
    };

    // 토큰 예치
    const handleDepositToken = async () => {
        if (!tokenAmount || !isTokenApproved) return;
        try {
            const amountInWei = ethers.utils.parseUnits(tokenAmount, 'ether');
            const tx = await contract.depositToken(amountInWei);
            const txResult = await tx.wait(); 
            const eventLogs = txResult.events;
            console.log('Token deposit successful:', eventLogs);
            setIsTokenApproved(false);
            setTokenAmount('');
        } catch (error) {
            console.error('Token deposit failed:', error);
        }
    };

    return (
        <div className="TokenDeposit">
            <div>
                <input
                    type="text"
                    value={etherAmount}
                    onChange={e => setEtherAmount(e.target.value)}
                    placeholder="Amount of ETH to deposit"
                />
                <button onClick={handleDepositEther}>Deposit ETH</button>
            </div>
            <div>
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
            </div>
        </div>
    );
}

export default TokenDeposit;
