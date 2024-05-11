import React from 'react';
import { Line } from 'react-chartjs-2';

const TradingChart = ({ chartData }) => {
    const data = {
        labels: chartData.map(data => data.time), // 차트의 x축
        datasets: [
            {
                label: 'Price',
                data: chartData.map(data => data.price), // 가격 데이터
                fill: false,
                borderColor: 'rgb(75, 192, 192)',
                tension: 0.1
            }
        ]
    };

    const options = {
        scales: {
            y: {
                beginAtZero: false
            }
        }
    };

    return <Line data={data} options={options} />;
}

export default TradingChart;
