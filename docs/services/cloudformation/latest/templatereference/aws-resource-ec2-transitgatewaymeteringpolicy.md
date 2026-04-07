This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayMeteringPolicy

Describes a transit gateway metering policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayMeteringPolicy",
  "Properties" : {
      "MiddleboxAttachmentIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayMeteringPolicy
Properties:
  MiddleboxAttachmentIds:
    - String
  Tags:
    - Tag
  TransitGatewayId: String

```

## Properties

`MiddleboxAttachmentIds`

The IDs of the middlebox attachments associated with the metering policy.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags assigned to the transit gateway metering policy.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-transitgatewaymeteringpolicy-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway associated with the metering policy.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`State`

The state of the transit gateway metering policy.

`TransitGatewayMeteringPolicyId`

The ID of the transit gateway metering policy.

`UpdateEffectiveAt`

The date and time when the metering policy update becomes effective.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayConnectPeerConfiguration

Tag
