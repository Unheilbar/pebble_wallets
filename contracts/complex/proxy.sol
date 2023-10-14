// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./transfer.sol";


contract Proxy {
    address constant private stateAddr = 0x1aEa632C29D2978A5C6336A3B8BFE9d737EB8fE3;
    address constant private transferAddr = 0x98aCaC3B9c77c934C12780a2852A959E674970A3;

    Transfer private transferContract;
    State private stateContract;
   
   constructor (){
        stateContract = State(stateAddr);
        transferContract = Transfer(transferAddr);
   }

    function emission(string memory walletId, uint256 amount) external {
        stateContract.add(walletId, amount);
        stateContract.setSender(msg.sender, walletId);
    }

    function transfer(string memory fromWalletId, string memory toWalletId, uint256 amount) external{
        transferContract.transfer(msg.sender, fromWalletId, toWalletId, amount);
    }
}