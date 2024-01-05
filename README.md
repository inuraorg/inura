<div align="center">
  <br />
  <br />
  <a href="https://inurascan.io"><img alt="Inura" src="./inuralogo-black 1.png" width=400></a>
  <br />
  <h3><a href="https://inurascan.io">Inura</a> is Ethereum, scaled.</h3>
  <br />
</div>

## What is Inura?

[Inura](https://www.inurascan.io/) is a blockchain built on the Optimism rollup, inheriting the powerful OP Stack. As a project dedicated to advancing Ethereum's capabilities, Inura focuses on scaling technology and fostering global collaboration in decentralized economies and governance systems. The Inura Collective, driving this initiative, develops open-source software to address key challenges in the broader cryptocurrency ecosystem. Guided by the principle of **impact=profit**, Inura rewards contributors proportionately, aiming to redefine incentives and positively impact the world. Explore our repository to engage with essential components of the OP Stack, contributing to the collaborative advancement of the Ethereum ecosystem.

## Directory Structure

<pre>
├── <a href="./docs">docs</a>: A collection of documents including audits and post-mortems
├── <a href="./inura-bindings">inura-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./inura-batcher">inura-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./inura-bootnode">inura-bootnode</a>: Standalone inura-node discovery bootnode
├── <a href="./inura-chain-ops">inura-chain-ops</a>: State surgery utilities
├── <a href="./inura-challenger">inura-challenger</a>: Dispute game challenge agent
├── <a href="./inura-e2e">inura-e2e</a>: End-to-End testing of all bedrock components in Go
├── <a href="./inura-heartbeat">inura-heartbeat</a>: Heartbeat monitor service
├── <a href="./inura-node">inura-node</a>: rollup consensus-layer client
├── <a href="./inura-preimage">inura-preimage</a>: Go bindings for Preimage Oracle
├── <a href="./inura-program">inura-program</a>: Fault proof program
├── <a href="./inura-proposer">inura-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./inura-service">inura-service</a>: Common codebase utilities
├── <a href="./inura-wheel">inura-wheel</a>: Database utilities
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-ts">contracts-ts</a>: ABI and Address constants
│   ├── <a href="./packages/contracts-bedrock">contracts-bedrock</a>: Bedrock smart contracts
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Inura easier
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Inura
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
└── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade
</pre>

## Branching Model

### Active Branches

| Branch          | Status                                                                           |
| --------------- | -------------------------------------------------------------------------------- |
| [master](https://github.com/inuraorg/inura/tree/master/)                   | Accepts PRs from `develop` when intending to deploy to production.                  |
| [develop](https://github.com/inuraorg/inura/tree/develop/)                 | Accepts PRs that are compatible with `master` OR from `release/X.X.X` branches.                    |
| release/X.X.X                                                                          | Accepts PRs for all changes, particularly those not backwards compatible with `develop` and `master`. |

### Overview

This repository generally follows [this Git branching model](https://nvie.com/posts/a-successful-git-branching-model/).
Please read the linked post if you're planning to make frequent PRs into this repository.


## How to Contribute

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/inuraorg/inura/contribute) to find something fun to work on!

<br/>

## License

Code forked from [`optimism`](https://github.com/inuraorg/inura) under the name [`optimism`](https://github.com/inuraorg/inura) is licensed under the [GNU GPLv3](https://gist.github.com/kn9ts/cbe95340d29fc1aaeaa5dd5c059d2e60) in accordance with the [original license](https://github.com/inuraorg/inura/blob/master/COPYING).