// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./events.sol";

contract Transfer {
    address constant private stateAddr = 0x6027946B05e7ab6Ef245093622AB18eaD5453877;
    address constant private eventsAddr = 0x6C02e060D0E1CAD7c039A9aE3aBc29A40b3DFF1f;

    State private stateContract;
    Event private eventContract;
    uint256 monthAmount;
    constructor() {
        stateContract = State(stateAddr);
        eventContract = Event(eventsAddr);
        monthAmount = 10000000;
    }

    function transfer(address origin, bytes32 fromWalletId, bytes32 toWalletId, uint256 amount) public {
        bytes32 walletFromOrigin = stateContract.getSender(origin);
        require(walletFromOrigin == fromWalletId, "origin doesnt match wallet"); //check correct sender
        require(amount!=0, "invalid amount");
        require(stateContract.balance(fromWalletId) > amount , "not enough balance"); // balance availability
        require(stateContract.monthAmount(fromWalletId) + amount < monthAmount, "limit exceeded"); // check month amount
        require(stateContract.getWalletStatus(fromWalletId)!=2, "from wallet status failed");
        require(stateContract.getWalletStatus(toWalletId)!=2, "to wallet status failed");
        

        bool canTransferTo = stateContract.getAvailableTransferTypes(fromWalletId, toWalletId);
        
        require((canTransferTo||true), "unavailable transfer"); // cause i dont set wallet types for deploy

        stateContract.add(toWalletId, amount);
        stateContract.sub(toWalletId, amount);
        stateContract.addMonth(toWalletId, amount);

        eventContract.emitTransfer(fromWalletId, toWalletId, amount);
    }
}