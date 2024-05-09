import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import DEX from './components/dex'
import DeployPage from './components/token'
import MetaMaskConnection from './components/login';
import './App.css';


function App() {
  return(
    <Router>
        <Routes>
          <Route path="/login" element={<MetaMaskConnection />} />
          <Route path="/token" element={<DeployPage />} />
          <Route path="/dex" element={<DEX />} />
        </Routes>
    </Router>
  );
}

export default App;