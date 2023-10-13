// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./transfer.sol";


contract Proxy {
    address constant private stateAddr = 0xE0f5206BBD039e7b0592d8918820024e2a7437b9;
    address constant private transferAddr = 0xE0f5206BBD039e7b0592d8918820024e2a7437b9;


    Transfer private transferContract;
    State private stateContract;
   
   constructor (){
        stateContract = State(stateAddr);
        transferContract = Transfer(stateAddr);
   }

    function emission(string memory walletId, uint256 amount) public {
        stateContract.add(walletId, amount);
        stateContract.setSender(msg.sender, walletId);
    }

    function transfer(string memory fromWalletId, string memory toWalletId, uint256 amount) public{
        transferContract.transfer(msg.sender, fromWalletId, toWalletId, amount);
    }
}