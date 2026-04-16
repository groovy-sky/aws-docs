---
title: "Understand traffic mirror filter concepts"
---

# Understand traffic mirror filter concepts

A _traffic mirror filter_ is a set of inbound and outbound rules that
determines which traffic is copied from the traffic mirror source and sent to the traffic mirror
target. You can also choose to mirror certain network services traffic, including Amazon DNS.
When you add network services traffic, all traffic (inbound and outbound) related to that
network service is mirrored.

We evaluate traffic mirror filter rules from the lowest value to the highest value. The
first rule that matches the traffic determines whether the traffic is mirrored. If you don't add
any rules, then no traffic is mirrored.

For example, in the following set of filter rules, rule 10 ensures that SSH traffic from my
network to my VPC is not mirrored and rule 20 mirrors all other IPv4 TCP traffic.

NumberRule actionProtocolSource port rangeDestination port rangeSource CIDR blockDestination CIDR block10rejectTCP (6)22`my-network``vpc-cidr`20acceptTCP (6)0.0.0.0/00.0.0.0/0

In the following set of filter rules, rule 10 mirrors HTTPS traffic from all IPv4 addresses
and rule 20 mirrors HTTPS traffic from all IPv6 addresses.

NumberRule actionProtocolSource port rangeDestination port rangeSource CIDR blockDestination CIDR block10acceptTCP (6)4430.0.0.0/00.0.0.0/020acceptTCP (6)443::/0::/0

Note that if you don't add outbound rules, then no outbound traffic is mirrored.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Targets

Sessions

All content copied from https://docs.aws.amazon.com/.
