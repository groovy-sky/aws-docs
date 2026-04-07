# Failover routing

Failover routing lets you route traffic to a resource when the resource is healthy or to a different resource when the
first resource is unhealthy. The primary and secondary records can route traffic to anything from an Amazon S3 bucket
that is configured as a website to a complex tree of records. For more information, see
[Active-passive failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-types.html#dns-failover-types-active-passive).

You can use failover routing policy for records in a private hosted zone.

For information about values that you specify when you use the failover routing policy to create records, see the following topics:

- [Values specific for failover records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values-failover.html)

- [Values specific for failover alias records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values-failover-alias.html)

- [Values that are common for all routing policies](resource-record-sets-values-shared.md)

- [Values that are common for alias records for all routing policies](resource-record-sets-values-alias-common.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Simple routing

Geolocation routing
