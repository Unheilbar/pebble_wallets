// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./transfer.sol";


contract Proxy {
    address constant private stateAddr = 0xE11f8d55a93bF877a091a3C54C071AAc5cC0b01D;
    address constant private transferAddr = 0x6C02e060D0E1CAD7c039A9aE3aBc29A40b3DFF1f;



    Transfer private transferContract;
    State private stateContract;
   
   constructor (){
        stateContract = State(stateAddr);
        transferContract = Transfer(stateAddr);
   }

    function emission(string memory walletId, uint256 amount) external {
        stateContract.add(walletId, amount);
        stateContract.setSender(msg.sender, walletId);
    }

    function transfer(string memory fromWalletId, string memory toWalletId, uint256 amount) external{
        revert("poshel nahui");
        transferContract.transfer(msg.sender, fromWalletId, toWalletId, amount);
    }
}