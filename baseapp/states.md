# States of Baseapp


Baseapp is constructed of many states that can be accessed in modules. Currently all these states are handled in the context, but it is not well defined which states should be set where and which are not needed. This document is meant to define where states are available and which should be used, if they are set which block do they adhere to. 

## PrepareProposal

PrepareProposal is only run by proposer, it will set certain values on the context to be used through the execution phases. 

The PrepareProposal Request provides 6 fields: 

* CommitInfo
    * This commitInfo is of the last block, with comet since there is the notion of delayed execution this commitInfo is of h-1 which is not finalized.
* Misbehavior
    * Misbehavior contains the evidence that is provided in BeginBlock. While it is safe to use this evidence to execute state transitions, the new state should not be committed to state until the commit phase. 
* Height
    * Height of the proposed block
* Timestamp
    * Time of the proposed block
* Proposer Address
    * Validator address of the proposing block
* NextValidatorsHash
    * The validator hash of the next validtor set   


## ProcessProposal

ProcessProposal is run by all validators once they have received a proposed block. 

ProcessProposal Request contains 7 fields:

* CommitInfo (proposed)
    * This commitInfo is of the last block, with comet since there is the notion of delayed execution this commitInfo is of h-1 which is not finalized, which means it does not have finalization.
* Hash
    * Hash of all of the fields in the proposed block
* Misbehavior
    * Misbehavior contains the evidence that was provided in BeginBlock previously. While it is safe to use this evidence to execute state transitions, the new state should not be committed to state until the commit phase. 
* Height
    * The current height of the chains
* Timestamp
    * time of the proposed block
* Proposer Address
    * proposer address of the block
* NextValidatorsHash
    * The validator hash of the next validtor set  

## VoteExtensions

## FinalizeBlock

Finalize Block combined BeginBlock, DeliverTx and EndBlock into a single request/response.

The FinalizeBlock request has 7 fields:

* CommitInfo (proposed)
    * This commitInfo is of the last block, with comet since there is the notion of delayed execution this commitInfo is of h-1 *which is finalized at this point*.
* Hash
    * Hash of all of the fields in the proposed block
* Misbehavior
    * Misbehavior contains the evidence that was provided in BeginBlock previously. At this point the block is finalized and it is safe to commit the evidence based state transitions to state. 
* Height
    * The current height of the chains
* Timestamp
    * time of the proposed block
* Proposer Address
    * proposer address of the block
* NextValidatorsHash
    * The validator hash of the next validtor set  
