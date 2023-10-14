// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;


contract Event {
   event TransferTokens(bytes32 fromWallet, bytes32 toWallet, uint256 amount);

   function emitTransfer(bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) public {
        emit TransferTokens(fromWalletId, toWalletId, amount);
    }
}


