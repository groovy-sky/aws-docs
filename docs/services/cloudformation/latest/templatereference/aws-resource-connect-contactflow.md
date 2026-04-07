This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ContactFlow

Specifies a flow for an Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ContactFlow",
  "Properties" : {
      "Content" : String,
      "Description" : String,
      "InstanceArn" : String,
      "Name" : String,
      "State" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ContactFlow
Properties:
  Content: String
  Description: String
  InstanceArn: String
  Name: String
  State: String
  Tags:
    - Tag
  Type: String

```

## Properties

`Content`

The content of the flow.

For more information, see [Amazon Connect Flow\
language](https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html) in the _Amazon Connect Administrator_
_Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the flow.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the flow.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the flow.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | ARCHIVED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-contactflow-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the flow. For descriptions of the available types, see [Choose a flow type](https://docs.aws.amazon.com/connect/latest/adminguide/create-contact-flow.html#contact-flow-types) in the
_Amazon Connect Administrator Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `CONTACT_FLOW | CUSTOMER_QUEUE | CUSTOMER_HOLD | CUSTOMER_WHISPER | AGENT_HOLD | AGENT_WHISPER | OUTBOUND_WHISPER | AGENT_TRANSFER | QUEUE_TRANSFER | CAMPAIGN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the flow name. For example:

`{ "Ref": "myFlowArn" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ContactFlowArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the flow.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Specify a flow resource

The following example specifies a flow resource for an Amazon Connect
instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a flow for an Amazon Connect instance
Resources:
  Flow:
    Type: 'AWS::Connect::ContactFlow'
    Properties:
      Name: ExampleFlow
      Description: flow created using cfn
      InstanceArn: arn:aws:connect:region-name:aws-account-id:instance/instance-arn
      Type: CONTACT_FLOW
      Content: ExampleFlow content(JSON) using Amazon Connect Flow Language.
      Tags:
        - Key: testkey
          Value: testValue
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::ApprovedOrigin

Tag
