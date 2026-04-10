This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary VPCConfig

If this canary is to test an endpoint in a VPC, this structure contains
information about the subnet and security groups of the VPC endpoint.
For more information, see [Running a Canary in a VPC](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-vpc.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ipv6AllowedForDualStack" : Boolean,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ],
  "VpcId" : String
}

```

### YAML

```yaml

  Ipv6AllowedForDualStack: Boolean
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  VpcId: String

```

## Properties

`Ipv6AllowedForDualStack`

Set this to `true` to allow outbound IPv6 traffic on VPC canaries that are connected to dual-stack subnets. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The IDs of the security groups for this canary.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The IDs of the subnets where this canary is to run.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC where this canary is to run.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualReference

AWS::Synthetics::Group

All content copied from https://docs.aws.amazon.com/.
