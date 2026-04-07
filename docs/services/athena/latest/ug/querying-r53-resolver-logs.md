# Query Amazon Route 53 resolver query logs

You can create Athena tables for your Amazon Route 53 Resolver query logs and query them from
Athena.

Route 53 Resolver query logging is for logging of DNS queries made by resources within a VPC,
on-premises resources that use inbound resolver endpoints, queries that use an outbound
Resolver endpoint for recursive DNS resolution, and queries that use Route 53 Resolver DNS
firewall rules to block, allow, or monitor a domain list. For more information about
Resolver query logging, see [Resolver query\
logging](../../../route53/latest/developerguide/resolver-query-logs.md) in the _Amazon Route 53 Developer Guide_. For
information about each of the fields in the logs, see [Values that appear\
in resolver query logs](../../../route53/latest/developerguide/resolver-query-logs-format.md) in the _Amazon Route 53 Developer Guide_.

###### Topics

- [Create the table for resolver query logs](https://docs.aws.amazon.com/athena/latest/ug/querying-r53-resolver-logs-creating-the-table.html)

- [Use partition projection](https://docs.aws.amazon.com/athena/latest/ug/querying-r53-resolver-logs-partitioning-example.html)

- [Example queries](https://docs.aws.amazon.com/athena/latest/ug/querying-r53-resolver-logs-example-queries.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Network Load Balancer

Create the table
