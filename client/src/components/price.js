import React, { useState, useEffect } from 'react';
import styled from 'styled-components';

const PriceContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 50px;
  font-size: 24px;
  font-weight: bold;
  margin: 20px 0;
`;

const CurrentPrice = ({ socket }) => {
  const [currentPrice, setCurrentPrice] = useState('Loading...');

  // useEffect(() => {
  //   socket.on('priceUpdate', (price) => {
  //     setCurrentPrice(price);
  //   });

  //   return () => {
  //     socket.off('priceUpdate');
  //   };
  // }, [socket]);

  return (
    <PriceContainer>
      Current Price: ${currentPrice}
    </PriceContainer>
  );
};

export default CurrentPrice;
