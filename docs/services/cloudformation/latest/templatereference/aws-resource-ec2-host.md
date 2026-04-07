This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Host

Allocates a fully dedicated physical server for launching EC2 instances. Because the
host is fully dedicated for your use, it can help you address compliance requirements and
reduce costs by allowing you to use your existing server-bound software licenses. For more
information, see [Dedicated Hosts](../../../ec2/latest/userguide/dedicated-hosts-overview.md) in
the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::Host",
  "Properties" : {
      "AssetId" : String,
      "AutoPlacement" : String,
      "AvailabilityZone" : String,
      "HostMaintenance" : String,
      "HostRecovery" : String,
      "InstanceFamily" : String,
      "InstanceType" : String,
      "OutpostArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::Host
Properties:
  AssetId: String
  AutoPlacement: String
  AvailabilityZone: String
  HostMaintenance: String
  HostRecovery: String
  InstanceFamily: String
  InstanceType: String
  OutpostArn: String
  Tags:
    - Tag

```

## Properties

`AssetId`

The ID of the Outpost hardware asset on which the Dedicated Host is allocated.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoPlacement`

Indicates whether the host accepts any untargeted instance launches that match its
instance type configuration, or if it only accepts Host tenancy instance launches that
specify its unique host ID. For more information, see [Understanding auto-placement and affinity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-dedicated-hosts-work.html#dedicated-hosts-understanding) in the
_Amazon EC2 User Guide_.

Default: `off`

_Required_: No

_Type_: String

_Allowed values_: `on | off`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The Availability Zone in which to allocate the Dedicated Host.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostMaintenance`

Indicates whether host maintenance is enabled or disabled for the Dedicated
Host.

_Required_: No

_Type_: String

_Allowed values_: `on | off`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostRecovery`

Indicates whether to enable or disable host recovery for the Dedicated Host. Host
recovery is disabled by default. For more information, see [Host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html)
in the _Amazon EC2 User Guide_.

Default: `off`

_Required_: No

_Type_: String

_Allowed values_: `on | off`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceFamily`

The instance family supported by the Dedicated Host. For example,
`m5`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

Specifies the instance type to be supported by the Dedicated Hosts. If you specify an
instance type, the Dedicated Hosts support instances of the specified instance type
only.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutpostArn`

The Amazon Resource Name (ARN) of the AWS Outpost on which the
Dedicated Host is allocated.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the Dedicated Host.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-host-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the host ID, such as
`h-0ab123c45d67ef89`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HostId`

The ID of the host.

## Examples

### Allocate a Dedicated Host

The following example allocates a dedicated host for `c3.large`
instances in the `us-east-1a` Availability Zone.

#### JSON

```json

"Host" : {
  "Type" : "AWS::EC2::Host",
  "Properties" : {
    "AutoPlacement" : "on",
    "AvailabilityZone" : "us-east-1a",
    "InstanceType" : "c3.large"
  }
}
```

#### YAML

```yaml

Host:
  Type: AWS::EC2::Host
  Properties:
    AutoPlacement: on
    AvailabilityZone: us-east-1a
    InstanceType: c3.large
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::GatewayRouteTableAssociation

Tag
