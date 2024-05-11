import { Web3Provider } from '@ethersproject/providers';
import React, { useEffect, useState } from 'react';
import ReCAPTCHA from 'react-google-recaptcha';

const MetamaskConnection = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isHuman, setIsHuman] = useState(false);
    const [showCaptcha, setShowCaptcha] = useState(false);
    const [sig, setSig] = useState("");

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
    
        // 메시지 서명 요청
        const message = "Plz Sign";
        const signer = new Web3Provider(window.ethereum).getSigner();
        const signature = await signer.signMessage(message);
        setSig(signature)
    
        // 서명과 주소를 서버로 전송
        fetch('http://localhost:5555/api/auth', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ "wallet_address": walletAddress, "signature": signature })
        })
        .then(response => response.json())
        .then(data => {
            if (data.IsNewUser) {
                console.log("CAPTCHA");
                setShowCaptcha(true);
            } else {
                if (data.signature !== signature) { // DB에 저장된 서명값과 다를경우 반려
                    alert("Warning. Signature is different");
                }

                if (data.IsVerified) {
                    setIsLoggedIn(true);
                    setShowCaptcha(false);
                    console.log("Login successful.");
                } else {
                    setShowCaptcha(true);
                }
            }
        })
        .catch(error => console.error('Error during authentication:', error));
    };

    const setVerifiedUser = async () => {
        const accounts = await window.ethereum.request({ method: 'eth_accounts' });
        const walletAddress = accounts[0];

        fetch('http://localhost:5555/api/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ "wallet_address": walletAddress, "isHuman": isHuman, "signature": sig })
        })
        .then(response => console.log(response))
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
