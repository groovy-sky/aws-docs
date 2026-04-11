This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SQS::QueueInlinePolicy

The `AWS::SQS::QueueInlinePolicy` resource associates one Amazon SQS
queue with one policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SQS::QueueInlinePolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "Queue" : String
    }
}

```

### YAML

```yaml

Type: AWS::SQS::QueueInlinePolicy
Properties:
  PolicyDocument: Json
  Queue: String

```

## Properties

`PolicyDocument`

A policy document that contains the permissions for the specified Amazon SQS
queues. For more information about Amazon SQS policies, see [Using\
custom policies with the Amazon SQS access policy language](../../../awssimplequeueservice/latest/sqsdeveloperguide/sqs-creating-custom-policies.md) in the _Amazon SQS Developer Guide_.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Queue`

The URLs of the queues to which you want to add the policy. You can use the `

          Ref
                        ` function to specify an `

          AWS::SQS::Queue
                        ` resource.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the URL of the Amazon SQS queue.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::SQS::QueuePolicy

All content copied from https://docs.aws.amazon.com/.
