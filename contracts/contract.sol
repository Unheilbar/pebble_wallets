// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Storage {
    enum Roles{
        NONE,
        PRO,
        AMATEUR
    }
    uint256 number;
    uint256 monthAmount;
    mapping (string => uint256) m;
    mapping (address => uint256) g;
    mapping (string => uint256) balances;
    mapping (string => uint256) monthAmounts;
    mapping (string => Roles) userRoles;

    constructor() {
        monthAmount = 1000000;
    }
    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number = num;
    }

    function transfer(string memory from, string  memory to, uint256 amount) public {
        require(balances[from] - amount >= 0);
        require(monthAmounts[from] + amount < monthAmount);
        require(userRoles[from]!=Roles.AMATEUR && userRoles[to]!=Roles.PRO);
        balances[from] = balances[from] - amount;
        balances[to] = balances[to] + amount;
        monthAmounts[from] = monthAmounts[from] + amount;
        number = amount;
    }       

    function setBalance(string memory s,uint256 n) public{
        balances[s] = balances[s] + n;
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