/* Generated by ts-generator ver. 0.0.8 */
/* tslint:disable */

import { Contract, Signer } from 'ethers';
import { Provider } from 'ethers/providers';

import { ArbChannel } from './ArbChannel';

export class ArbChannelFactory {
    static connect(address: string, signerOrProvider: Signer | Provider): ArbChannel {
        return new Contract(address, _abi, signerOrProvider) as ArbChannel;
    }
}

const _abi = [
    {
        constant: true,
        inputs: [],
        name: 'challengeManager',
        outputs: [
            {
                name: '',
                type: 'address',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [
            {
                name: 'validator',
                type: 'address',
            },
        ],
        name: 'currentDeposit',
        outputs: [
            {
                name: '',
                type: 'uint256',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'validatorCount',
        outputs: [
            {
                name: '',
                type: 'uint16',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'getState',
        outputs: [
            {
                name: '',
                type: 'uint8',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_players',
                type: 'address[2]',
            },
            {
                name: '_rewards',
                type: 'uint128[2]',
            },
        ],
        name: 'completeChallenge',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_assertPreHash',
                type: 'bytes32',
            },
        ],
        name: 'initiateChallenge',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'vm',
        outputs: [
            {
                name: 'machineHash',
                type: 'bytes32',
            },
            {
                name: 'pendingHash',
                type: 'bytes32',
            },
            {
                name: 'inbox',
                type: 'bytes32',
            },
            {
                name: 'asserter',
                type: 'address',
            },
            {
                name: 'escrowRequired',
                type: 'uint128',
            },
            {
                name: 'deadline',
                type: 'uint64',
            },
            {
                name: 'sequenceNum',
                type: 'uint64',
            },
            {
                name: 'gracePeriod',
                type: 'uint32',
            },
            {
                name: 'maxExecutionSteps',
                type: 'uint32',
            },
            {
                name: 'state',
                type: 'uint8',
            },
            {
                name: 'inChallenge',
                type: 'bool',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_preconditionHash',
                type: 'bytes32',
            },
            {
                name: '_afterHash',
                type: 'bytes32',
            },
            {
                name: '_numSteps',
                type: 'uint32',
            },
            {
                name: '_tokenTypes',
                type: 'bytes21[]',
            },
            {
                name: '_messageData',
                type: 'bytes',
            },
            {
                name: '_messageTokenNums',
                type: 'uint16[]',
            },
            {
                name: '_messageAmounts',
                type: 'uint256[]',
            },
            {
                name: '_messageDestinations',
                type: 'address[]',
            },
            {
                name: '_logsAccHash',
                type: 'bytes32',
            },
        ],
        name: 'confirmDisputableAsserted',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'terminateAddress',
        outputs: [
            {
                name: '',
                type: 'address',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'exitAddress',
        outputs: [
            {
                name: '',
                type: 'address',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'activatedValidators',
        outputs: [
            {
                name: '',
                type: 'uint16',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'owner',
        outputs: [
            {
                name: '',
                type: 'address',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: false,
        inputs: [],
        name: 'activateVM',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_fields',
                type: 'bytes32[4]',
            },
            {
                name: '_numSteps',
                type: 'uint32',
            },
            {
                name: '_timeBounds',
                type: 'uint64[2]',
            },
            {
                name: '_tokenTypes',
                type: 'bytes21[]',
            },
            {
                name: '_messageDataHash',
                type: 'bytes32[]',
            },
            {
                name: '_messageTokenNums',
                type: 'uint16[]',
            },
            {
                name: '_messageAmounts',
                type: 'uint256[]',
            },
            {
                name: '_messageDestinations',
                type: 'address[]',
            },
        ],
        name: 'pendingDisputableAssert',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'escrowRequired',
        outputs: [
            {
                name: '',
                type: 'uint256',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: false,
        inputs: [],
        name: 'ownerShutdown',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: true,
        inputs: [],
        name: 'globalInbox',
        outputs: [
            {
                name: '',
                type: 'address',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        inputs: [
            {
                name: '_vmState',
                type: 'bytes32',
            },
            {
                name: '_gracePeriod',
                type: 'uint32',
            },
            {
                name: '_maxExecutionSteps',
                type: 'uint32',
            },
            {
                name: '_escrowRequired',
                type: 'uint128',
            },
            {
                name: '_owner',
                type: 'address',
            },
            {
                name: '_challengeManagerAddress',
                type: 'address',
            },
            {
                name: '_globalInboxAddress',
                type: 'address',
            },
            {
                name: '_validatorKeys',
                type: 'address[]',
            },
        ],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'constructor',
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: false,
                name: 'unanHash',
                type: 'bytes32',
            },
            {
                indexed: false,
                name: 'sequenceNum',
                type: 'uint64',
            },
        ],
        name: 'PendingUnanimousAssertion',
        type: 'event',
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: false,
                name: 'sequenceNum',
                type: 'uint64',
            },
        ],
        name: 'ConfirmedUnanimousAssertion',
        type: 'event',
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: false,
                name: 'unanHash',
                type: 'bytes32',
            },
        ],
        name: 'FinalizedUnanimousAssertion',
        type: 'event',
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: false,
                name: 'fields',
                type: 'bytes32[3]',
            },
            {
                indexed: false,
                name: 'asserter',
                type: 'address',
            },
            {
                indexed: false,
                name: 'timeBounds',
                type: 'uint64[2]',
            },
            {
                indexed: false,
                name: 'tokenTypes',
                type: 'bytes21[]',
            },
            {
                indexed: false,
                name: 'numSteps',
                type: 'uint32',
            },
            {
                indexed: false,
                name: 'lastMessageHash',
                type: 'bytes32',
            },
            {
                indexed: false,
                name: 'logsAccHash',
                type: 'bytes32',
            },
            {
                indexed: false,
                name: 'amounts',
                type: 'uint256[]',
            },
        ],
        name: 'PendingDisputableAssertion',
        type: 'event',
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: false,
                name: 'newState',
                type: 'bytes32',
            },
            {
                indexed: false,
                name: 'logsAccHash',
                type: 'bytes32',
            },
        ],
        name: 'ConfirmedDisputableAssertion',
        type: 'event',
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: false,
                name: 'challenger',
                type: 'address',
            },
        ],
        name: 'InitiatedChallenge',
        type: 'event',
    },
    {
        constant: false,
        inputs: [],
        name: 'increaseDeposit',
        outputs: [],
        payable: true,
        stateMutability: 'payable',
        type: 'function',
    },
    {
        constant: true,
        inputs: [
            {
                name: 'validator',
                type: 'address',
            },
        ],
        name: 'isListedValidator',
        outputs: [
            {
                name: '',
                type: 'bool',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: true,
        inputs: [
            {
                name: '_validators',
                type: 'address[]',
            },
        ],
        name: 'isValidatorList',
        outputs: [
            {
                name: '',
                type: 'bool',
            },
        ],
        payable: false,
        stateMutability: 'view',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_afterHash',
                type: 'bytes32',
            },
            {
                name: '_newInbox',
                type: 'bytes32',
            },
            {
                name: '_tokenTypes',
                type: 'bytes21[]',
            },
            {
                name: '_messageData',
                type: 'bytes',
            },
            {
                name: '_messageTokenNums',
                type: 'uint16[]',
            },
            {
                name: '_messageAmounts',
                type: 'uint256[]',
            },
            {
                name: '_messageDestinations',
                type: 'address[]',
            },
            {
                name: '_logsAccHash',
                type: 'bytes32',
            },
            {
                name: '_signatures',
                type: 'bytes',
            },
        ],
        name: 'finalizedUnanimousAssert',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_unanRest',
                type: 'bytes32',
            },
            {
                name: '_tokenTypes',
                type: 'bytes21[]',
            },
            {
                name: '_messageTokenNums',
                type: 'uint16[]',
            },
            {
                name: '_messageAmounts',
                type: 'uint256[]',
            },
            {
                name: '_sequenceNum',
                type: 'uint64',
            },
            {
                name: '_logsAccHash',
                type: 'bytes32',
            },
            {
                name: '_signatures',
                type: 'bytes',
            },
        ],
        name: 'pendingUnanimousAssert',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
    {
        constant: false,
        inputs: [
            {
                name: '_afterHash',
                type: 'bytes32',
            },
            {
                name: '_newInbox',
                type: 'bytes32',
            },
            {
                name: '_tokenTypes',
                type: 'bytes21[]',
            },
            {
                name: '_messageData',
                type: 'bytes',
            },
            {
                name: '_messageTokenNums',
                type: 'uint16[]',
            },
            {
                name: '_messageAmounts',
                type: 'uint256[]',
            },
            {
                name: '_messageDestinations',
                type: 'address[]',
            },
        ],
        name: 'confirmUnanimousAsserted',
        outputs: [],
        payable: false,
        stateMutability: 'nonpayable',
        type: 'function',
    },
];