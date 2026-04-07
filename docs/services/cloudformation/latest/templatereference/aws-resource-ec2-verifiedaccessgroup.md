This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessGroup

An AWS Verified Access group is a collection of AWS Verified Access endpoints who's associated applications have
similar security requirements. Each instance within a Verified Access group shares an Verified Access policy. For
example, you can group all Verified Access instances associated with "sales" applications together and
use one common Verified Access policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VerifiedAccessGroup",
  "Properties" : {
      "Description" : String,
      "PolicyDocument" : String,
      "PolicyEnabled" : Boolean,
      "SseSpecification" : SseSpecification,
      "Tags" : [ Tag, ... ],
      "VerifiedAccessInstanceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VerifiedAccessGroup
Properties:
  Description: String
  PolicyDocument: String
  PolicyEnabled: Boolean
  SseSpecification:
    SseSpecification
  Tags:
    - Tag
  VerifiedAccessInstanceId: String

```

## Properties

`Description`

A description for the AWS Verified Access group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDocument`

The Verified Access policy document.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyEnabled`

The status of the Verified Access policy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SseSpecification`

The options for additional server side encryption.

_Required_: No

_Type_: [SseSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessgroup-ssespecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessgroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifiedAccessInstanceId`

The ID of the AWS Verified Access instance.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Verified Access group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The creation time.

`LastUpdatedTime`

The last updated time.

`Owner`

The ID of the AWS account that owns the group.

`VerifiedAccessGroupArn`

The ARN of the Verified Access group.

`VerifiedAccessGroupId`

The ID of the Verified Access group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

SseSpecification
