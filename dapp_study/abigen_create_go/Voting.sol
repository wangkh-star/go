// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting{
    //投票池
    mapping(address =>  uint256 )  public votes;
    //参与人
    address[] public voters; 

   //投票池
    mapping(address =>  bool )  public hasVotes;
    function vote(address voteTo)  public {
        require(!hasVotes[msg.sender]," you are already vote");
        votes[voteTo]=votes[voteTo]+1;
        voters.push(msg.sender);
        hasVotes[msg.sender]=true;
    }

    function getHasVotes(address voteTo)  view  public  returns (bool){
       return  hasVotes[voteTo];
    }
    //获取投票数
    function getVotes(address voteTo)  view  public  returns (uint256){
       return  votes[voteTo];
    }
    // 清空记录数组
   function resetVotes()  public {
    for (uint256 i = 0; i < voters.length; i++) {
            address voter = voters[i];
            delete votes[voter];
            delete hasVotes[voter];
        }
         delete voters;
    }

}