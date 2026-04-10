This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ResourceConfiguration

Creates a resource configuration. A resource configuration defines a specific resource. You
can associate a resource configuration with a service network or a VPC endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ResourceConfiguration",
  "Properties" : {
      "AllowAssociationToSharableServiceNetwork" : Boolean,
      "CustomDomainName" : String,
      "DomainVerificationId" : String,
      "GroupDomain" : String,
      "Name" : String,
      "PortRanges" : [ String, ... ],
      "ProtocolType" : String,
      "ResourceConfigurationAuthType" : String,
      "ResourceConfigurationDefinition" : ResourceConfigurationDefinition,
      "ResourceConfigurationGroupId" : String,
      "ResourceConfigurationType" : String,
      "ResourceGatewayId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ResourceConfiguration
Properties:
  AllowAssociationToSharableServiceNetwork: Boolean
  CustomDomainName: String
  DomainVerificationId: String
  GroupDomain: String
  Name: String
  PortRanges:
    - String
  ProtocolType: String
  ResourceConfigurationAuthType: String
  ResourceConfigurationDefinition:
    ResourceConfigurationDefinition
  ResourceConfigurationGroupId: String
  ResourceConfigurationType: String
  ResourceGatewayId: String
  Tags:
    - Tag

```

## Properties

`AllowAssociationToSharableServiceNetwork`

Specifies whether the resource configuration can be associated with a sharable service
network.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDomainName`

The custom domain name.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainVerificationId`

The domain verification ID.

_Required_: No

_Type_: String

_Pattern_: `^dv-[a-fA-F0-9]{17}$`

_Minimum_: `20`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupDomain`

(GROUP) The group domain for a group resource configuration. Any domains that you create for the child resource are subdomains of the group domain. Child resources inherit the verification status of the domain.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the resource configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!rcfg-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortRanges`

(SINGLE, GROUP, CHILD) The TCP port ranges that a consumer can use to access a resource
configuration (for example: 1-65535). You can separate port ranges using commas (for example:
1,2,22-30).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolType`

(SINGLE, GROUP) The protocol accepted by the resource configuration.

_Required_: No

_Type_: String

_Allowed values_: `TCP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceConfigurationAuthType`

The auth type for the resource configuration.

_Required_: No

_Type_: String

_Allowed values_: `NONE | AWS_IAM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceConfigurationDefinition`

Identifies the resource configuration in one of the following ways:

- **Amazon Resource Name (ARN)** \- Supported resource-types
that are provisioned by AWS services, such as RDS databases, can be identified
by their ARN.

- **Domain name** \- Any domain name that is publicly
resolvable.

- **IP address** \- For IPv4 and IPv6, only IP addresses in the
VPC are supported.

_Required_: No

_Type_: [ResourceConfigurationDefinition](aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceConfigurationGroupId`

The ID of the group resource configuration.

_Required_: No

_Type_: String

_Pattern_: `^rcfg-[0-9a-z]{17}$`

_Minimum_: `22`

_Maximum_: `22`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceConfigurationType`

The type of resource configuration. A resource configuration can be one of the following
types:

- **SINGLE** \- A single resource.

- **GROUP** \- A group of resources. You must create a group
resource configuration before you create a child resource configuration.

- **CHILD** \- A single resource that is part of a group
resource configuration.

- **ARN** \- An AWS resource.

_Required_: Yes

_Type_: String

_Allowed values_: `GROUP | CHILD | SINGLE | ARN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceGatewayId`

The ID of the resource gateway.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the resource configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-resourceconfiguration-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the resource
configuration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resource configuration.

`Id`

The ID of the resource configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WeightedTargetGroup

DnsResource

All content copied from https://docs.aws.amazon.com/.
