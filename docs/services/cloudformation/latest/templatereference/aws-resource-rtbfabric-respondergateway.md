This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::ResponderGateway

Creates a responder gateway.

###### Important

A domain name or managed endpoint is required.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RTBFabric::ResponderGateway",
  "Properties" : {
      "Description" : String,
      "DomainName" : String,
      "ManagedEndpointConfiguration" : ManagedEndpointConfiguration,
      "Port" : Integer,
      "Protocol" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TrustStoreConfiguration" : TrustStoreConfiguration,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::RTBFabric::ResponderGateway
Properties:
  Description: String
  DomainName: String
  ManagedEndpointConfiguration:
    ManagedEndpointConfiguration
  Port: Integer
  Protocol: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  Tags:
    - Tag
  TrustStoreConfiguration:
    TrustStoreConfiguration
  VpcId: String

```

## Properties

`Description`

An optional description for the responder gateway.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9 ]+$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DomainName`

The domain name for the responder gateway.

_Required_: No

_Type_: String

_Pattern_: `^(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?)(?:\.(?:[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?))+$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ManagedEndpointConfiguration`

The configuration for the managed endpoint.

_Required_: No

_Type_: [ManagedEndpointConfiguration](aws-properties-rtbfabric-respondergateway-managedendpointconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Port`

The networking port to use.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Protocol`

The networking protocol to use.

_Required_: Yes

_Type_: String

_Allowed values_: `HTTP | HTTPS`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SecurityGroupIds`

The unique identifiers of the security groups.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SubnetIds`

The unique identifiers of the subnets.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

A map of the key-value pairs of the tag or tags to assign to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-rtbfabric-respondergateway-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustStoreConfiguration`

The configuration of the trust store.

_Required_: No

_Type_: [TrustStoreConfiguration](aws-properties-rtbfabric-respondergateway-truststoreconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VpcId`

The unique identifier of the Virtual Private Cloud (VPC).

_Required_: Yes

_Type_: String

_Minimum_: `5`

_Maximum_: `50`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Property description not available.

`CreatedTimestamp`

Property description not available.

`GatewayId`

Property description not available.

`ResponderGatewayStatus`

Property description not available.

`UpdatedTimestamp`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AutoScalingGroupsConfiguration

All content copied from https://docs.aws.amazon.com/.
