This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SQS::QueuePolicy

The `AWS::SQS::QueuePolicy` type applies a policy to Amazon SQS queues.
For an example snippet, see [Declaring an\
Amazon SQS policy](../userguide/quickref-iam.md#scenario-sqs-policy) in the _CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SQS::QueuePolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "Queues" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SQS::QueuePolicy
Properties:
  PolicyDocument: Json
  Queues:
    - String

```

## Properties

`PolicyDocument`

A policy document that contains the permissions for the specified Amazon SQS
queues. For more information about Amazon SQS policies, see [Using\
custom policies with the Amazon SQS access policy language](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html) in the _Amazon SQS Developer Guide_.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Queues`

The URLs of the queues to which you want to add the policy. You can use the `

          Ref
                        ` function to specify an `

          AWS::SQS::Queue
                        ` resource.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Id`

The provider-assigned unique ID for this managed resource.

## Examples

### Amazon SQS Queue Policy

The following sample is a queue policy that allows AWS account
111122223333 to send and receive messages on queue queue2. You add the policy to the
resources section of your template.

#### JSON

```json

"SampleSQSPolicy" : {
  "Type" : "AWS::SQS::QueuePolicy",
  "Properties" : {
    "Queues" :  ["https://sqs:us-east-2.amazonaws.com/444455556666/queue2"],
    "PolicyDocument": {
      "Statement":[{
        "Action":["SQS:SendMessage", "SQS:ReceiveMessage"],
        "Effect":"Allow",
        "Resource": "arn:aws:sqs:us-east-2:444455556666:queue2",
        "Principal": {
          "AWS": [
            "111122223333"]
        }
      }]
    }
  }
}

```

#### YAML

```yaml

SampleSQSPolicy:
  Type: AWS::SQS::QueuePolicy
  Properties:
    Queues:
      - "https://sqs:us-east-2.amazonaws.com/444455556666/queue2"
    PolicyDocument:
      Statement:
        -
          Action:
            - "SQS:SendMessage"
            - "SQS:ReceiveMessage"
          Effect: "Allow"
          Resource: "arn:aws:sqs:us-east-2:444455556666:queue2"
          Principal:
            AWS:
              - "111122223333"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SQS::QueueInlinePolicy

Next
