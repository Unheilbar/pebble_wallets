// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;


contract State {
    mapping (bytes32 => uint256) balances;
    mapping (address => bytes32) senders;
    mapping (bytes32 => uint256) monthAmounts;
    mapping (bytes32 => uint8) walletStatusMap;
    mapping (bytes32 => uint8) walletTypesMap;
    mapping (bytes32 => bool) walletAvailableTransferTypes;

    // user with walletType 0 can transfer to user with wallet types 1 and 2 which is abi encoded 0x0000000000000000000000000000000000000000000000000000000000000001;
    constructor() {
        // first 16 bytes encode fromWallet, last 16 bytes toWallet
        walletAvailableTransferTypes[0x0000000000000000000000000000000000100000000000000000000000000001]=true;
    }

    function add(bytes32 walletId, uint256 amount) public innerCall {
        balances[walletId] = balances[walletId] + amount;
    }

    function addMonth(bytes32 walletId, uint256 amount) public innerCall {
        monthAmounts[walletId] = monthAmounts[walletId] + amount;
    }

    function sub(bytes32 walletId, uint256 amount) public innerCall {
        balances[walletId] = balances[walletId] - amount;
    }

    function setSender(address sender ,bytes32 walletId) public innerCall {
        senders[sender] = walletId;
    }


    function balance(bytes32 walletId) external view returns(uint256){
        return balances[walletId];
    }

    function monthAmount (bytes32 walletId) external view returns(uint256){
        return monthAmounts[walletId];
    }

    function getSender(address sender) external view returns(bytes32) {
        return senders[sender];
    }
    
    function getWalletStatus(bytes32 walletId) external view returns(uint8) {
        return walletStatusMap[walletId];
    }

    function getWalletType(bytes32 walletId) external view returns(uint8) {
        return walletTypesMap[walletId];
    }

     function getWalletStatusI(bytes32 walletId) internal view returns(uint8) {
        return walletStatusMap[walletId];
    }

    function getWalletTypeI(bytes32 walletId) internal view returns(uint8) {
        return walletTypesMap[walletId];
    }

    function getAvailableTransferTypes (bytes32 fromWalletId, bytes32 toWalletId) external view returns(bool) {
        return walletAvailableTransferTypes[keccak256(abi.encodePacked(getWalletTypeI(fromWalletId), getWalletTypeI(toWalletId)))];
    }

    modifier innerCall virtual{
        require(msg.sender!=tx.origin);
        _;
    }
}