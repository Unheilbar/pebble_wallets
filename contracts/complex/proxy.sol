// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./transfer.sol";


contract Proxy {
    address constant private stateAddr = 0x6027946B05e7ab6Ef245093622AB18eaD5453877;
    address constant private transferAddr = 0x5dBC355B93DD7A0C0D759Fd6a7859d2610219221;

    Transfer private transferContract;
    State private stateContract;
   
   constructor (){
        stateContract = State(stateAddr);
        transferContract = Transfer(transferAddr);
   }

    function emission(address senderId, bytes32 walletId, uint256 amount) external {
        stateContract.add(walletId, amount);
        stateContract.setSender(senderId, walletId);
    }

    function transfer(bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) external{
        transferContract.transfer(msg.sender, fromWalletId, toWalletId, amount);
    }
}