This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::Service DnsConfig

A complex type that contains information about the Amazon Route 53 DNS records that you want
AWS Cloud Map to create when you register an instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DnsRecords" : [ DnsRecord, ... ],
  "NamespaceId" : String,
  "RoutingPolicy" : String
}

```

### YAML

```yaml

  DnsRecords:
    - DnsRecord
  NamespaceId: String
  RoutingPolicy: String

```

## Properties

`DnsRecords`

An array that contains one `DnsRecord` object for each Route 53 DNS record
that you want AWS Cloud Map to create when you register an instance.

###### Important

The record type of a service can't be updated directly and can only be changed by
deleting the service and recreating it with a new `DnsConfig`.

_Required_: Yes

_Type_: Array of [DnsRecord](aws-properties-servicediscovery-service-dnsrecord.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceId`

_Use_
_NamespaceId in [Service](../../../cloud-map/latest/api/api-service.md) instead._

The ID of the namespace to use for DNS configuration.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoutingPolicy`

The routing policy that you want to apply to all Route 53 DNS records that AWS Cloud Map creates
when you register an instance and specify this service.

###### Note

If you want to use this service to register instances that create alias records, specify
`WEIGHTED` for the routing policy.

You can specify the following values:

MULTIVALUE

If you define a health check for the service and the health check is healthy, Route 53
returns the applicable value for up to eight instances.

For example, suppose that the service includes configurations for one `A`
record and a health check. You use the service to register 10 instances. Route 53 responds to DNS
queries with IP addresses for up to eight healthy instances. If fewer than eight instances are
healthy, Route 53 responds to every DNS query with the IP addresses for all of the healthy
instances.

If you don't define a health check for the service, Route 53 assumes that all instances are
healthy and returns the values for up to eight instances.

For more information about the multivalue routing policy, see [Multivalue\
Answer Routing](../../../route53/latest/developerguide/routing-policy.md#routing-policy-multivalue) in the _Route 53 Developer Guide_.

WEIGHTED

Route 53 returns the applicable value from one randomly selected instance from among the
instances that you registered using the same service. Currently, all records have the same
weight, so you can't route more or less traffic to any instances.

For example, suppose that the service includes configurations for one `A`
record and a health check. You use the service to register 10 instances. Route 53 responds to DNS
queries with the IP address for one randomly selected instance from among the healthy
instances. If no instances are healthy, Route 53 responds to DNS queries as if all of the
instances were healthy.

If you don't define a health check for the service, Route 53 assumes that all instances are
healthy and returns the applicable value for one randomly selected instance.

For more information about the weighted routing policy, see [Weighted\
Routing](../../../route53/latest/developerguide/routing-policy.md#routing-policy-weighted) in the _Route 53 Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `MULTIVALUE | WEIGHTED`

_Update requires_: Updates are not supported.

## See also

- [Return values](../userguide/aws-resource-servicediscovery-service.md#aws-resource-servicediscovery-service-return-values) in the topic [AWS::ServiceDiscovery::Service](../userguide/aws-resource-servicediscovery-service.md)

- [DnsConfig](../../../cloud-map/latest/api/api-dnsconfig.md) in the _AWS Cloud Map API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceDiscovery::Service

DnsRecord

All content copied from https://docs.aws.amazon.com/.
