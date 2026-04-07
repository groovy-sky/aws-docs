This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::NetworkSettings

This resource specifies network settings that can be associated with a web portal. Once
associated with a web portal, network settings define how streaming instances will connect
with your specified VPC.

The VPC must have default tenancy. VPCs with dedicated tenancy are not supported.

For availability consideration, you must have at least two subnets created in two different Availability Zones. WorkSpaces Secure Browser is available in a subset of the Availability Zones for each supported Region. For more information, see [Supported Availability Zones](https://docs.aws.amazon.com/workspaces-web/latest/adminguide/availability-zones.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::NetworkSettings",
  "Properties" : {
      "SecurityGroupIds" : [ String, ... ],
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::NetworkSettings
Properties:
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`SecurityGroupIds`

One or more security groups used to control access from streaming instances to your VPC.

_Pattern_: `^[\w+\-]+$`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `128 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The subnets in which network interfaces are created to connect streaming instances to your VPC. At least two of these subnets must be in different availability zones.

_Pattern_: `^subnet-([0-9a-f]{8}|[0-9a-f]{17})$`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 2`

_Maximum_: `32 | 3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the network settings resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-networksettings-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The VPC that streaming instances will connect to.

_Pattern_: `^vpc-[0-9a-z]*$`

_Required_: Yes

_Type_: String

_Pattern_: `^vpc-[0-9a-z]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that this network settings is associated with.

`NetworkSettingsArn`

The ARN of the network settings.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
