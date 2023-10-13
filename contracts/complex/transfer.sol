// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./events.sol";

contract Transfer {
    address constant private stateAddr = 0xE11f8d55a93bF877a091a3C54C071AAc5cC0b01D;
    address constant private eventsAddr = 0x6027946B05e7ab6Ef245093622AB18eaD5453877;

    State private stateContract;
    Event private eventContract;
    uint256 monthAmount;
    constructor() {
        stateContract = State(stateAddr);
        eventContract = Event(eventsAddr);
        monthAmount = 10000000;
    }


    function transfer(address origin, string memory fromWalletId, string memory toWalletId, uint256 amount) public {
        string memory walletFromOrigin = stateContract.getSender(origin);

        require(keccak256(bytes(walletFromOrigin)) == keccak256(bytes(fromWalletId)), "origin doesnt match wallet"); //check correct sender
        require(stateContract.balance(fromWalletId) - amount >= 0, "not enough balance"); // balance availability
        require(stateContract.monthAmount(fromWalletId) + amount < monthAmount, "limit exceeded"); // check month amount
        require(stateContract.getWalletStatus(fromWalletId)!=2, "from wallet status failed");
        require(stateContract.getWalletStatus(toWalletId)!=2, "to wallet status failed");
        
        uint8 fromWalletType = stateContract.getWalletType(fromWalletId);
        uint8 toWalletType = stateContract.getWalletType(toWalletId);
        uint8[] memory canTransferTo = stateContract.getAvailableTransferTypes(fromWalletType);
        bool ok = false;
        for (uint i=0;i<canTransferTo.length;i++){
            if (canTransferTo[i] == toWalletType) {
                ok = true;
                break;
            }
        }
        require(ok, "unavailable transfer type");

        stateContract.add(toWalletId, amount);
        stateContract.sub(toWalletId, amount);
        stateContract.addMonth(toWalletId, amount);

        eventContract.emitTransfer(fromWalletId, toWalletId, amount);
    }
}