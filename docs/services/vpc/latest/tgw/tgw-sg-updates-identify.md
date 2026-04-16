---
title: "Identify AWS Transit Gateway referenced security groups"
---

# Identify AWS Transit Gateway referenced security groups

To determine if your security group is being referenced in the rules of a security
group in a VPC attached to the same transit gateway, use one of the following commands.

- [describe-security-group-references](../../../cli/latest/reference/ec2/describe-security-group-references.md) (AWS CLI)

- [Get-EC2SecurityGroupReference](../../../powershell/latest/reference/items/get-ec2securitygroupreference.md) (AWS Tools for Windows PowerShell)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Update security group inbound rules

Remove stale security group rules

All content copied from https://docs.aws.amazon.com/.
