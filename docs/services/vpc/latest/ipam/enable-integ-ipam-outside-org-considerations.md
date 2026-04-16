---
title: "Considerations and limitations"
---

# Considerations and limitations

This section contains considerations and limitations for integrating IPAM with accounts outside of your organization:

- When you share a resource discovery with another account, the only data that is exchanged is IP address and account status monitoring data. You can view this data before sharing using the [get-ipam-discovered-resource-cidrs](../../../cli/latest/reference/ec2/get-ipam-discovered-resource-cidrs.md) and [get-ipam-discovered-accounts](../../../cli/latest/reference/ec2/get-ipam-discovered-accounts.md) CLI commands or [GetIpamDiscoveredResourceCidrs](../../../../reference/awsec2/latest/apireference/api-getipamdiscoveredresourcecidrs.md) and [GetIpamDiscoveredAccounts](../../../../reference/awsec2/latest/apireference/api-getipamdiscoveredaccounts.md) APIs. For resource discoveries that monitor resources across an organization, no organization data (such as the names of Organizational Units in your organization) are shared.

- When you create a resource discovery, the resource discovery monitors all visible resources in the owner account. If the owner account is a third-party service AWS account that creates resources for multiple of their own customers, those resources will be discovered by the resource discovery. If the third-party AWS service account shares the resource discovery with an end-user AWS account, the end-user will have visibility into the resources of the other customers of the third-party AWS service. For that reason, the third-party AWS service should exercise caution creating and sharing resource discoveries or use a separate AWS account for each customer.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Integrate IPAM with accounts outside of your organization

Process overview

All content copied from https://docs.aws.amazon.com/.
