  
pragma solidity 0.5.9;


/**
* Generic Registry Interface, used in Frontier for Candidates and Voters
* Potentially enabling Any king of registry (Centralized or decentralized, like TCR)
**/

contract Registry {


    event _WhiteList(address _whiteListedAccount); 
    event _Remove(address _removedAccount);
    event _Application(bytes32 indexed listingHash, uint deposit, string data, address indexed applicant);

    /**
    * @dev Adds a new account to the registry
    * @param _accountToWhiteList account to be added to the registry
    **/

    function whiteList(address _accountToWhiteList) public;

    /**
    * @dev Removes an account from the registry
    * @param _accountToRemove account to be removed from the Registry
    **/

    function remove(address _accountToRemove) public;


    /**
    *  @dev Returns true when a listing Hash is already Whitelisted
    *  @param _accountChecked identifier queried listing
    */

    function isWhitelisted(address _accountChecked) view public returns (bool whitelisted);

    
}