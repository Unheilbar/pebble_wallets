// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;


contract State {
    mapping (string => uint256) balances;
    mapping (address => string) senders;
    mapping (string => uint256) monthAmounts;
    mapping (string => uint8) walletStatusMap;
    mapping (string => uint8) walletTypesMap;
    mapping (uint8 => uint8[]) walletAvailableTransferTypes;

    constructor() {
        walletAvailableTransferTypes[0].push(0);
    }

    function add(string memory walletId, uint256 amount) public innerCall {
        balances[walletId] = balances[walletId] + amount;
    }

    function addMonth(string memory walletId, uint256 amount) public innerCall {
        monthAmounts[walletId] = monthAmounts[walletId] + amount;
    }

    function sub(string memory walletId, uint256 amount) public innerCall {
        balances[walletId] = balances[walletId] - amount;
    }

    function setSender(address sender ,string memory walletId) public innerCall {
        senders[sender] = walletId;
    }


    function balance(string memory walletId) external view returns(uint256){
        return balances[walletId];
    }

    function monthAmount (string memory walletId) external view returns(uint256){
        return monthAmounts[walletId];
    }

    function getSender(address sender) external view returns(string memory) {
        return senders[sender];
    }
    
    function getWalletStatus(string memory walletId) external view returns(uint8) {
        return walletStatusMap[walletId];
    }

    function getWalletType(string memory walletId) external view returns(uint8) {
        return walletTypesMap[walletId];
    }

    function getAvailableTransferTypes (uint8 walletType) external view returns(uint8[] memory) {
        return walletAvailableTransferTypes[walletType];
    }

    modifier innerCall virtual{
        require(msg.sender!=tx.origin);
        _;
    }
}