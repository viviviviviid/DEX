import { ethers } from 'ethers';
import { Web3Provider } from '@ethersproject/providers';
import React, { useEffect } from 'react';

function MetaMaskConnection() {
    useEffect(() => {
        checkNetwork();
    }, []);

    async function checkNetwork() {
        if (window.ethereum) {
            try {
                const provider = new Web3Provider(window.ethereum);
                await provider.send("eth_requestAccounts", []);
                const network = await provider.getNetwork();

                if (network.chainId !== 11155111) { // Sepolia 테스트넷의 Chain ID
                    console.log("not chain id")
                    switchToSepolia();
                } else {
                    console.log("You are already connected to the Sepolia network!");
                }
            } catch (error) {
                console.error("Error checking network:", error);
            }
        } else {
            alert("Please install MetaMask!");
        }
    }
    async function switchToSepolia() {
        try {
            await window.ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }], // Sepolia 테스트넷의 Chain ID in hexadecimal
            });
        } catch (error) {
            console.error("Error switching to Sepolia:", error);
            alert("Error switching network. Please add Sepolia Testnet manually in MetaMask if not already done.");
        }
    }
    
}    

export default MetaMaskConnection;
