import React from 'react';

const TradeHistory = ({ trades }) => {
    return (
        <div>
            <h3>Trade History</h3>
            <ul>
                {trades.map((trade, index) => (
                    <li key={index}>{`${trade.type} - ${trade.amount} @ ${trade.price}`}</li>
                ))}
            </ul>
        </div>
    );
}

export default TradeHistory;
