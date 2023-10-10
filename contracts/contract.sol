// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Storage {

    uint256 number;
    mapping (string => uint256) m;
    mapping (address => uint256) g;
    mapping (address => uint256) balances;
    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number = num;
    }

    function transfer(address from, address to, uint256 amount) public {
        require(balances[from] - amount >= 0);
        balances[from] = balances[from] - amount;
        balances[to] = balances[to] + amount;
    }       

    function setBalance(string memory s,uint256 n) external{
        m[s] = m[s]+n;
    }

    function setBalanceA(address a, uint256 n) public {
        g[a]=n;
    }

    /**
     * @dev Return value 
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256){
        return number;
    }
}