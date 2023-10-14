// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

import "./state.sol";
import "./events.sol";

contract Transfer {
    address constant private stateAddr = 0x1aEa632C29D2978A5C6336A3B8BFE9d737EB8fE3;
    address constant private eventsAddr = 0x94a562Ef266F41D4AC4b125c1C2a5aAf7E952467;

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
        require(amount!=0, "invalid amount");
        require(stateContract.balance(fromWalletId) > amount , "not enough balance"); // balance availability
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