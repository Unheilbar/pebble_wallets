// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;


contract Event {
   event TransferTokens(string fromWallet, string toWallet, uint256 amount);

   function emitTransfer(string memory fromWalletId, string memory toWalletId, uint256 amount) public {
        emit TransferTokens(fromWalletId, toWalletId, amount);
    }
}


