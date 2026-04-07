This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SNS::TopicPolicy

The `AWS::SNS::TopicPolicy` resource associates Amazon SNS topics
with a policy. For an example snippet, see [Declaring\
an Amazon SNS policy](../userguide/quickref-iam.md#scenario-sns-policy) in the _CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SNS::TopicPolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "Topics" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SNS::TopicPolicy
Properties:
  PolicyDocument: Json
  Topics:
    - String

```

## Properties

`PolicyDocument`

A policy document that contains permissions to add to the specified SNS topics.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topics`

The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You
can use the `
                            Ref
                        ` function to specify an `
                            AWS::SNS::Topic
                        ` resource.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The provider-assigned unique ID for this managed resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SNS::TopicInlinePolicy

Next
