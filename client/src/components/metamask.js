import { Web3Provider } from '@ethersproject/providers';
import React, { useEffect, useState } from 'react';
import ReCAPTCHA from 'react-google-recaptcha';

const MetamaskConnection = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isHuman, setIsHuman] = useState(false);
    const [showCaptcha, setShowCaptcha] = useState(false);

    useEffect(() => {
        checkNetwork();
    }, []);

    const onCaptchaChange = (value) => {
        console.log("Captcha value:", value);
        if (value) {
            setIsHuman(true);
            setVerifiedUser();
        }
    };

    const checkNetwork = async () => {
        if (window.ethereum) {
            try {
                const provider = new Web3Provider(window.ethereum);
                await provider.send("eth_requestAccounts", []);
                const network = await provider.getNetwork();

                if (network.chainId !== 11155111) { // Sepolia 테스트넷의 Chain ID
                    switchToSepolia();
                } else {
                    authenticateUser();
                }
            } catch (error) {
                console.error("Error checking network:", error);
            }
        } else {
            alert("Please install MetaMask!");
        }
    };

    const switchToSepolia = async () => {
        try {
            await window.ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0xaa36a7' }] // Sepolia 테스트넷의 Chain ID in hexadecimal
            });
        } catch (switchError) {
            console.error("Error switching to Sepolia:", switchError);
            alert("Error switching network. Please add Sepolia Testnet manually in MetaMask if not already done.");
        }
    };

    const authenticateUser = async () => {
        const accounts = await window.ethereum.request({ method: 'eth_accounts' });
        const walletAddress = accounts[0];
        
        setShowCaptcha(true);
        fetch('http://localhost:5555/api/auth', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ walletAddress })
        })
        .then(response => response.json())
        .then(data => {
            if (data.IsNewUser) {
                console.log("Please complete CAPTCHA to verify you're a human.");
                setShowCaptcha(true);
            } else {
                console.log(data)
                // 신규유저가 아니지만, 캡차를 통한 인간인증을 안한사람이 있을 수 있으므로, data.isVerified를 추가적으로 확인해야함
                // setIsLoggedIn(true);
                console.log("Login successful.");
            }
        })
        .catch(error => console.error('Error during authentication:', error));
    };

    const setVerifiedUser = async () => {
        const accounts = await window.ethereum.request({ method: 'eth_accounts' });
        const walletAddress = accounts[0];

        fetch('http://localhost:5555/api/verifyHuman', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ walletAddress, isHuman })
        })
        .then(response => response.json())
        .then(() => {
            setIsLoggedIn(true);
            console.log("Human verification successful and user logged in.");
        })
        .catch(error => console.error('Error verifying human:', error));
    };

    return (
        <div>
            {isLoggedIn ? (
                <p>You are logged in.</p>
            ) : (
                <>
                    <button onClick={checkNetwork}>Connect to MetaMask</button>
                    {showCaptcha && (
                        <ReCAPTCHA
                            sitekey={process.env.REACT_APP_CAPTCHA_KEY}
                            onChange={onCaptchaChange}
                        />
                    )}
                </>
            )}
        </div>
    );
};

export default MetamaskConnection;
