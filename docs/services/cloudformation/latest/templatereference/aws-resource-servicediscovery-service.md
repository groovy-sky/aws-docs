This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::Service

A complex type that contains information about the specified service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceDiscovery::Service",
  "Properties" : {
      "Description" : String,
      "DnsConfig" : DnsConfig,
      "HealthCheckConfig" : HealthCheckConfig,
      "HealthCheckCustomConfig" : HealthCheckCustomConfig,
      "Name" : String,
      "NamespaceId" : String,
      "ServiceAttributes" : Json,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceDiscovery::Service
Properties:
  Description: String
  DnsConfig:
    DnsConfig
  HealthCheckConfig:
    HealthCheckConfig
  HealthCheckCustomConfig:
    HealthCheckCustomConfig
  Name: String
  NamespaceId: String
  ServiceAttributes: Json
  Tags:
    - Tag
  Type: String

```

## Properties

`Description`

The description of the service.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsConfig`

A complex type that contains information about the Route 53 DNS records that you want
AWS Cloud Map to create when you register an instance.

###### Important

The record types of a service can only be changed by deleting the service and recreating it
with a new `Dnsconfig`.

_Required_: No

_Type_: [DnsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicediscovery-service-dnsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckConfig`

_Public DNS and HTTP namespaces only._ A complex type that contains
settings for an optional health check. If you specify settings for a health check, AWS Cloud Map
associates the health check with the records that you specify in `DnsConfig`.

For information about the charges for health checks, see [Amazon Route 53 Pricing](https://aws.amazon.com/route53/pricing).

_Required_: No

_Type_: [HealthCheckConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicediscovery-service-healthcheckconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckCustomConfig`

A complex type that contains information about an optional custom health check.

###### Important

If you specify a health check configuration, you can specify either
`HealthCheckCustomConfig` or `HealthCheckConfig` but not both.

_Required_: No

_Type_: [HealthCheckCustomConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicediscovery-service-healthcheckcustomconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the service.

_Required_: No

_Type_: String

_Pattern_: `((?=^.{1,127}$)^([a-zA-Z0-9_][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9_]|[a-zA-Z0-9])(\.([a-zA-Z0-9_][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9_]|[a-zA-Z0-9]))*$)|(^\.$)`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NamespaceId`

The ID or Amazon Resource Name (ARN) of the namespace that you want to use to create the service. For
namespaces shared with your AWS account, specify the namespace ARN. For more information
about shared namespaces, see [Cross-account AWS Cloud Map namespace sharing](https://docs.aws.amazon.com/cloud-map/latest/dg/sharing-namespaces.html) in the
_AWS Cloud Map Developer Guide_.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceAttributes`

A complex type that contains information about attributes associated with a specific
service.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the service. Each tag consists of a key and an optional value, both of
which you define. Tag keys can have a maximum character length of 128 characters, and tag
values can have a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicediscovery-service-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: Updates are not supported.

`Type`

If present, specifies that the service instances are only discoverable using the
`DiscoverInstances` API operation. No DNS records is registered for the service
instances. The only valid value is `HTTP`.

_Required_: No

_Type_: String

_Allowed values_: `HTTP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the value of `Id` for the service, such as
`srv-e4anhexample0004`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the service.

`Id`

The ID of the service.

`Name`

The name that you assigned to the service.

## Examples

- [Create a service](#aws-resource-servicediscovery-service--examples--Create_a_service)

- [Specify attributes at the service level](#aws-resource-servicediscovery-service--examples--Specify_attributes_at_the_service_level)

### Create a service

The following example creates a service based on a public DNS namespace. The
service includes settings for Amazon Route 53 A and AAAA records that have a routing
policy of `WEIGHTED`. It also includes a Route 53 health check.

#### JSON

```json

{
  "Resources": {
    "Service": {
      "Type": "AWS::ServiceDiscovery::Service",
      "Properties": {
        "Description": "Service based on a public DNS namespace",
        "DnsConfig": {
          "DnsRecords": [
            {
              "Type": "A",
              "TTL": 60
            },
            {
              "Type": "AAAA",
              "TTL": 60
            }
          ],
          "RoutingPolicy": "WEIGHTED"
        },
        "HealthCheckConfig": {
          "FailureThreshold": 3,
          "ResourcePath": "/",
          "Type": "HTTPS"
        },
        "Name": "example-public-DNS-service",
        "NamespaceId": "ns-e4anhexample0004"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  Service:
    Type: AWS::ServiceDiscovery::Service
    Properties:
      Description: Service based on a public DNS namespace
      DnsConfig:
        DnsRecords:
          - Type: A
            TTL: 60
          - Type: AAAA
            TTL: 60
        RoutingPolicy: WEIGHTED
      HealthCheckConfig:
        FailureThreshold: 3
        ResourcePath: /
        Type: HTTPS
      Name: example-public-DNS-service
      NamespaceId: ns-e4anhexample0004
```

### Specify attributes at the service level

This example specifies the port that the client application should use to
communicate with the service, as a service attribute. The example also specifies the service
revision and that 100% of traffic should be directed to the revision.

#### JSON

```json

{
  "Resources": {
    "Service": {
      "Type": "AWS::ServiceDiscovery::Service",
      "Properties": {
        "Description": "Service based on a public DNS namespace",
        "DnsConfig": {
          "DnsRecords": [
            {
              "Type": "A",
              "TTL": 60
            },
            {
              "Type": "AAAA",
              "TTL": 60
            }
          ],
          "RoutingPolicy": "WEIGHTED"
        },
        "ServiceAttributes": {
          "Port": "8080",
          "Weight": "{'revision-A': '100'}"
        },
        "HealthCheckConfig": {
          "FailureThreshold": 3,
          "ResourcePath": "/",
          "Type": "HTTPS"
        },
        "Name": "example-public-DNS-service",
        "NamespaceId": "ns-e4anhexample0004"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  Service:
    Type: AWS::ServiceDiscovery::Service
    Properties:
      Description: Service based on a public DNS namespace
      DnsConfig:
        DnsRecords:
          - Type: A
            TTL: 60
          - Type: AAAA
            TTL: 60
        RoutingPolicy: WEIGHTED
      HealthCheckConfig:
        FailureThreshold: 3
        ResourcePath: /
        Type: HTTPS
      ServiceAttributes:
        Port: 8080
        Weight: {'revision-A': '100'}
      Name: example-public-DNS-service
      NamespaceId: ns-e4anhexample0004
```

## See also

- [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html) in the _AWS Cloud Map API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DnsConfig
