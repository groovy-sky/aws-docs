This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::Instance

A complex type that contains information about an instance that AWS Cloud Map creates when you
submit a `RegisterInstance` request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceDiscovery::Instance",
  "Properties" : {
      "InstanceAttributes" : Json,
      "InstanceId" : String,
      "ServiceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceDiscovery::Instance
Properties:
  InstanceAttributes: Json
  InstanceId: String
  ServiceId: String

```

## Properties

`InstanceAttributes`

A string map that contains the following information for the service that you specify in
`ServiceId`:

- The attributes that apply to the records that are defined in the service.

- For each attribute, the applicable value.

Supported attribute keys include the following:

AWS\_ALIAS\_DNS\_NAME

If you want AWS Cloud Map to create a Route 53 alias record that routes
traffic to an Elastic Load Balancing load balancer, specify the DNS name that is
associated with the load balancer. For information about how to get the DNS name,
see [AliasTarget->DNSName](../../../../reference/route53/latest/apireference/api-aliastarget.md#Route53-Type-AliasTarget-DNSName) in the _Route 53 API Reference_.

Note the following:

- The configuration for the service that is specified by
`ServiceId` must include settings for an `A`
record, an `AAAA` record, or both.

- In the service that is specified by `ServiceId`, the value of
`RoutingPolicy` must be `WEIGHTED`.

- If the service that is specified by `ServiceId` includes
`HealthCheckConfig` settings, AWS Cloud Map will
create the health check, but it won't associate the health check with the
alias record.

- Auto naming currently doesn't support creating alias records that route
traffic to AWS resources other than ELB load
balancers.

- If you specify a value for `AWS_ALIAS_DNS_NAME`, don't specify
values for any of the `AWS_INSTANCE` attributes.

AWS\_EC2\_INSTANCE\_ID

_HTTP namespaces only._ The Amazon EC2 instance ID for the
instance. The `AWS_INSTANCE_IPV4` attribute contains the primary
private IPv4 address. When creating resources with a type of [AWS::ServiceDiscovery::Instance](../userguide/aws-resource-servicediscovery-instance.md), if the
`AWS_EC2_INSTANCE_ID` attribute is specified, the only other
attribute that can be specified is `AWS_INIT_HEALTH_STATUS`. After the
resource has been created, the `AWS_INSTANCE_IPV4` attribute contains
the primary private IPv4 address.

AWS\_INIT\_HEALTH\_STATUS

If the service configuration includes `HealthCheckCustomConfig`,
when creating resources with a type of [AWS::ServiceDiscovery::Instance](../userguide/aws-resource-servicediscovery-instance.md) you can optionally use
`AWS_INIT_HEALTH_STATUS` to specify the initial status of the custom
health check, `HEALTHY` or `UNHEALTHY`. If you don't specify
a value for `AWS_INIT_HEALTH_STATUS`, the initial status is
`HEALTHY`. This attribute can only be used when creating resources
and will not be seen on existing resources.

AWS\_INSTANCE\_CNAME

If the service configuration includes a `CNAME` record, the domain
name that you want Route 53 to return in response to DNS queries, for example,
`example.com`.

This value is required if the service specified by `ServiceId`
includes settings for an `CNAME` record.

AWS\_INSTANCE\_IPV4

If the service configuration includes an `A` record, the IPv4
address that you want Route 53 to return in response to DNS queries, for example,
`192.0.2.44`.

This value is required if the service specified by `ServiceId`
includes settings for an `A` record. If the service includes settings
for an `SRV` record, you must specify a value for
`AWS_INSTANCE_IPV4`, `AWS_INSTANCE_IPV6`, or both.

AWS\_INSTANCE\_IPV6

If the service configuration includes an `AAAA` record, the IPv6
address that you want Route 53 to return in response to DNS queries, for example,
`2001:0db8:85a3:0000:0000:abcd:0001:2345`.

This value is required if the service specified by `ServiceId`
includes settings for an `AAAA` record. If the service includes
settings for an `SRV` record, you must specify a value for
`AWS_INSTANCE_IPV4`, `AWS_INSTANCE_IPV6`, or both.

AWS\_INSTANCE\_PORT

If the service includes an `SRV` record, the value that you want
Route 53 to return for the port.

If the service includes `HealthCheckConfig`, the port on the
endpoint that you want Route 53 to send requests to.

This value is required if you specified settings for an `SRV` record
or a Route 53 health check when you created the service.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

An identifier that you want to associate with the instance. Note the following:

- If the service that's specified by `ServiceId` includes settings for an
`SRV` record, the value of `InstanceId` is automatically included as
part of the value for the `SRV` record. For more information, see [DnsRecord >\
Type](../../../cloud-map/latest/api/api-dnsrecord.md#cloudmap-Type-DnsRecord-Type).

- You can use this value to update an existing instance.

- To register a new instance, you must specify a value that's unique among instances that
you register by using the same service.

- If you specify an existing `InstanceId` and `ServiceId`, AWS Cloud Map
updates the existing DNS records, if any. If there's also an existing health check, AWS Cloud Map
deletes the old health check and creates a new one.

###### Note

The health check isn't deleted immediately, so it will still appear for a while if you
submit a `ListHealthChecks` request, for example.

###### Note

Do not include sensitive information in `InstanceId` if the namespace is
discoverable by public DNS queries and any `Type` member of `DnsRecord`
for the service contains `SRV` because the `InstanceId` is discoverable by
public DNS queries.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z_/:.@-]+$`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceId`

The ID or Amazon Resource Name (ARN) of the service that you want to use for settings for the instance. For
services created in a shared namespace, specify the service ARN. For more information about
shared namespaces, see [Cross-account AWS Cloud Map namespace sharing](../../../cloud-map/latest/dg/sharing-namespaces.md) in the
_AWS Cloud Map Developer Guide_.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the value of `Id` for the instance, such as
`i-abcd1234`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`InstanceId`

The ID of an instance that matches the values that you specified in the request.

## Examples

- [Use custom attributes to register resources](#aws-resource-servicediscovery-instance--examples--Use_custom_attributes_to_register_resources)

- [Provide IP addresses for an instance](#aws-resource-servicediscovery-instance--examples--Provide_IP_addresses_for_an_instance)

### Use custom attributes to register resources

The following example registers application resources like a DynamoDB
table and Lambda functions as instances using the table and function
names as custom attributes. The Lambda functions also have an
`action` custom attribute used to distinguish the read function from
the write function. The DynamoDB table is registered using the service that
has the ID `srv-e3anhexample0003`. The read and write Lambda
functions are registered using the service that has the ID
`srv-e4anhexample0004`. This example is based on the tutorial [Learn how to use AWS Cloud Map service discovery with custom\
attributes](../../../cloud-map/latest/dg/tutorial-private-namespace.md) in the _AWS Cloud Map Developer Guide_.

#### JSON

```json

{
  "Resources": {
    "ReadInstance": {
      "Type": "AWS::ServiceDiscovery::Instance",
      "Properties": {
        "InstanceAttributes": {
          "functionname": "readfunction",
          "action": "read"
        },
        "InstanceId": "read-instance",
        "ServiceId": "srv-e4anhexample0004"
      }
    },
    "WriteInstance": {
      "Type": "AWS::ServiceDiscovery::Instance",
      "Properties": {
        "InstanceAttributes": {
          "functionname": "writefunction",
          "action": "write"
        },
        "InstanceId": "write-instance",
        "ServiceId": "srv-e4anhexample0004"
      }
    },
    "DataInstance": {
      "Type": "AWS::ServiceDiscovery::Instance",
      "Properties": {
        "InstanceAttributes": {
          "tablename": "cloudmap"
        },
        "InstanceId": "data-instance",
        "ServiceId": "srv-e3anhexample0003"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  ReadInstance:
    Type: AWS::ServiceDiscovery::Instance
    Properties:
      InstanceAttributes:
        functionname: readfunction
        action: read
      InstanceId: read-instance
      ServiceId: srv-e4anhexample0004
  WriteInstance:
    Type: AWS::ServiceDiscovery::Instance
    Properties:
      InstanceAttributes:
        functionname: writefunction
        action: write
      InstanceId: write-instance
      ServiceId: srv-e4anhexample0004
  DataInstance:
    Type: AWS::ServiceDiscovery::Instance
    Properties:
      InstanceAttributes:
        tablename: cloudmap
      InstanceId: data-instance
      ServiceId: srv-e3anhexample0003
```

### Provide IP addresses for an instance

The following example provides IPv4 and IPV6 IP addresses for the instance that
has an ID of `i-abcd1234`. The instance was registered using the service
that has an ID of `srv-e4anhexample0004`.

#### JSON

```json

{
    "Resources": {
        "Instance": {
            "Type": "AWS::ServiceDiscovery::Instance",
            "Properties": {
                "InstanceAttributes": {
                    "AWS_INSTANCE_IPV4": "192.0.2.44",
                    "AWS_INSTANCE_IPV6": "2001:0db8:85a3:0000:0000:abcd:0001:2345"
                },
                "InstanceId": "i-abcd1234",
                "ServiceId": "srv-e4anhexample0004"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  Instance:
    Type: AWS::ServiceDiscovery::Instance
    Properties:
      InstanceAttributes:
        AWS_INSTANCE_IPV4: 192.0.2.44
        AWS_INSTANCE_IPV6: 2001:0db8:85a3:0000:0000:abcd:0001:2345
      InstanceId: i-abcd1234
      ServiceId: srv-e4anhexample0004
```

## See also

- [RegisterInstance](../../../cloud-map/latest/api/api-registerinstance.md) in the _AWS Cloud Map API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ServiceDiscovery::PrivateDnsNamespace

All content copied from https://docs.aws.amazon.com/.
