// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

import "./openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "./openzeppelin/contracts/access/Ownable.sol";
import "./openzeppelin/contracts/security/Pausable.sol";

contract NonVotedCapped is ERC20, ERC20Burnable, Ownable, Pausable {
    constructor(
        string memory name,
        string memory symbol,
        address _initialWallet,
        uint256 _supply
    ) ERC20(name, symbol) {
        _mint(_initialWallet, _supply * 10 ** decimals());
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount)
    internal
    whenNotPaused
    override
    {
        super._beforeTokenTransfer(from, to, amount);
    }
}