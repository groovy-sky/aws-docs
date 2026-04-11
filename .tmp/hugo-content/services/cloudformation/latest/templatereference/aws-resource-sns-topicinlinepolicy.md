This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SNS::TopicInlinePolicy

The `AWS::SNS::TopicInlinePolicy` resource associates one Amazon SNS
topic with one policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SNS::TopicInlinePolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "TopicArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SNS::TopicInlinePolicy
Properties:
  PolicyDocument: Json
  TopicArn: String

```

## Properties

`PolicyDocument`

A policy document that contains permissions to add to the specified Amazon SNS
topic.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the topic to which you want to add the policy.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the topic for which the policy was added.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::SNS::TopicPolicy

All content copied from https://docs.aws.amazon.com/.
