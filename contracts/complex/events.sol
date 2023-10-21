// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;


contract Event {
   event TransferTokens(bytes32 fromWallet, bytes32 toWallet, uint128 fromBalance, uint128 toBalance);

   function emitTransfer(bytes32 fromWalletId, bytes32 toWalletId, uint128 fromBalance, uint128 toBalance) public {
        emit TransferTokens(fromWalletId, toWalletId, fromBalance, toBalance);
    }
}


