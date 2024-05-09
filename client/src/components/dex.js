import React, { useEffect, useState } from 'react';

function DEX() {
    const [price, setPrice] = useState(null);

    useEffect(() => {
        const socket = new WebSocket('wss://yourserver.com/path');
        socket.onmessage = function(event) {
            const data = JSON.parse(event.data);
            setPrice(data.price);
        };
        return () => socket.close();
    }, []);

    return (
        <div>
            <h1>Current Token Price: {price}</h1>
            {/* 여기에 주문 창, 주문 리스트 등 추가 */}
        </div>
    );
}

export default DEX;
