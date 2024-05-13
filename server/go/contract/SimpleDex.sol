// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract SimpleDex is Ownable, ReentrancyGuard {
    IERC20 public tradeToken;

    struct Order {
        uint256 id;
        address user;
        uint256 amount;
        uint256 price;
        bool isBuy;
    }

    uint256 public nextOrderId = 0;
    mapping(uint256 => Order) public orders;
    mapping(address => uint256) public ethBalances;
    mapping(address => uint256) public tokenBalances;

    event OrderPlaced(uint256 indexed orderId, address indexed user, uint256 amount, uint256 price, bool isBuy);
    event OrderExecuted(uint256 indexed buyOrderId, uint256 indexed sellOrderId, uint256 amount, uint256 price);
    event OrderCancelled(uint256 indexed orderId);

    constructor(address _tradeToken) Ownable(msg.sender) {
        tradeToken = IERC20(_tradeToken);
    }

    function depositETH() public payable {
        require(msg.value > 0, "Deposit amount must bigger than zero");
        ethBalances[msg.sender] += msg.value;
    }

    function withdrawETH(uint256 amount) public {
        require(amount <= ethBalances[msg.sender], "Insufficient balance");
        ethBalances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }

    function depositToken(uint256 amount) public {
        require(tradeToken.transferFrom(msg.sender, address(this), amount), "transfer failed");
        tokenBalances[msg.sender] += amount;
    }

    function withdrawToken(uint256 amount) public {
        require(amount <= tokenBalances[msg.sender], "Insufficient balance");
        tokenBalances[msg.sender] -= amount;
        require(tradeToken.transfer(msg.sender, amount), "transfer failed");
    }

    function placeOrder(uint256 amount, uint256 price, bool isBuy) public {
        require(amount > 0 && price > 0, "Amount and price must biiggerr than zero");
        require(isBuy ? ethBalances[msg.sender] >= amount * price : tokenBalances[msg.sender] >= amount, "Insufficient balance for order");

        uint256 orderId = nextOrderId++;
        orders[orderId] = Order({
            id: orderId,
            user: msg.sender,
            amount: amount,
            price: price,
            isBuy: isBuy
        });

        emit OrderPlaced(orderId, msg.sender, amount, price, isBuy);
    }

    function cancelOrder(uint256 orderId) public {
        require(orders[orderId].user == msg.sender, "only deployer");
        require(orders[orderId].id != 0, "order not exist");

        delete orders[orderId];
        emit OrderCancelled(orderId);
    }

    function executeOrder(uint256 buyOrderId, uint256 sellOrderId) public onlyOwner {
        Order storage buyOrder = orders[buyOrderId];
        Order storage sellOrder = orders[sellOrderId];

        require(buyOrder.isBuy && !sellOrder.isBuy, "order types do not match");
        require(buyOrder.price >= sellOrder.price, "Price mismatch");
        uint256 amount = min(buyOrder.amount, sellOrder.amount);

        ethBalances[buyOrder.user] -= amount * sellOrder.price;
        tokenBalances[sellOrder.user] -= amount;
        ethBalances[sellOrder.user] += amount * sellOrder.price;
        tokenBalances[buyOrder.user] += amount;

        emit OrderExecuted(buyOrderId, sellOrderId, amount, sellOrder.price);

        updateOrderAfterExecution(buyOrder, amount);
        updateOrderAfterExecution(sellOrder, amount);
    }

    function updateOrderAfterExecution(Order storage order, uint256 amount) internal {
        if (order.amount == amount) {
            delete orders[order.id];
        } else {
            order.amount -= amount;
        }
    }

    function getBalances(address _user) public view returns (uint256 balanceETH, uint256 balanceToken) {
        return (ethBalances[_user], tokenBalances[_user]);
    }

    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }
}
