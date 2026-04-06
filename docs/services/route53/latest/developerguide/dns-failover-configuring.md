# Configuring DNS failover

When you have more than one resource performing the same function—for example, more than one HTTP server or mail server—you can
configure Amazon Route 53 to check the health of your resources and respond to DNS queries using only the healthy resources. For example, suppose your website,
example.com, is hosted on six servers, two each in three data centers around the world. You can configure Route 53 to check the health of those servers and
to respond to DNS queries for example.com using only the servers that are currently healthy.

Route 53 can check the health of your resources in both simple and complex configurations:

- In simple configurations, you create a group of records that all have the same name and type, such as a group of weighted records
with a type of A for example.com. You then configure Route 53 to check the health of the corresponding resources. Route 53 responds to
DNS queries based on the health of your resources. For more information, see
[How health checks work in simple Amazon Route 53 configurations](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-simple-configs.html).

- In more complex configurations, you create a tree of records that route traffic based on multiple criteria. For example, if latency for
your users is your most important criterion, then you might use latency alias records to route traffic to the region that provides the best latency.
The latency alias records might have weighted records in each region as the alias target. The weighted records might route traffic to EC2 instances
based on the instance type. As with a simple configuration, you can configure Route 53 to route traffic based on the health of your resources.
For more information, see [How health checks work in complex Amazon Route 53 configurations](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-complex-configs.html).

###### Topics

- [Task list for configuring DNS failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-how-to.html)

- [How health checks work in simple Amazon Route 53 configurations](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-simple-configs.html)

- [How health checks work in complex Amazon Route 53 configurations](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-complex-configs.html)

- [How Amazon Route 53 chooses records when health checking is configured](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-how-route-53-chooses-records.html)

- [Active-active and active-passive failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-types.html)

- [Configuring failover in a private hosted zone](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-private-hosted-zones.html)

- [How Amazon Route 53 averts failover problems](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-problems.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configuring router and
firewall rules for health checks

Task list for configuring DNS failover
