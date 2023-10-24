// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Lsm
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Lsm {
   mapping(bytes32=>uint128) balance;

   function transfer(bytes32 from, bytes32 to, uint128 amount) external {
        balance[from] -= amount;
        balance[to] += amount;
   }

   function emission(bytes32 id, uint128 amount) external {
        balance[id]+=amount;
   }
}