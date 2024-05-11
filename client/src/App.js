import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import MetamaskConnection from './components/metamask';
import DeployPage from './components/token';
import DEX from './components/dex';
import './App.css';

function App() {
  return (
    <Router>
        <Routes>
          <Route path="/" element={<MetamaskConnection />} />
          <Route path="/token" element={<DeployPage />} />
          <Route path="/dex/contract/:contractAddress/pair/:tokenAddress" element={<DEX />} />
        </Routes>
    </Router>
  );
}

export default App;
