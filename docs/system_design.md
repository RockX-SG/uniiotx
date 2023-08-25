# Design of IoTeX Liquid Staking Contracts
To-Do: This document is awaiting perfection.
## Introduction

## Backgrounds

## System Architecture and Context
This project is based on the IoTex network. It entails multiple collaborative roles and modules where smart contracts are crucial.

![architecture_context](./architecture_context.png) <br>

## Business Lifecycle
Investors first need to deposit any amount of IOTXs for uniIOTXs. After that they have two ways to get their assets back:
1. Visit the exchange to swap uniIOTXs for IOTXs. There's no limit on the amount of IOTXs.
2. Submit a redemption request until the principal and rewards are fully claimed. 
The permissible quantity of IOTXs for redemption is governed by the protocol preset by contracts. 

The manager's fee is subtracted from the shared rewards of the investors, a process that is also regulated by contracts.

![business_lifecycle](./business_lifecycle.png) <br>

## Business Proces: Depositing IOTXs
![business_process_deposit](./business_process_deposit.png) <br>

## Business Proces: Redeeming IOTXs
![business_process_redeem](./business_process_redeem.png) <br>

## References
