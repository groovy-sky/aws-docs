This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kinesis::ResourcePolicy

Attaches a resource-based policy to a data stream or registered consumer. If you are using an identity other than the root user of
the AWS account that owns the resource, the calling identity must have the `PutResourcePolicy` permissions on the
specified Kinesis Data Streams resource and belong to the owner's account in order to use this operation.
If you don't have `PutResourcePolicy` permissions, Amazon Kinesis Data Streams returns a `403 Access Denied error`.
If you receive a `ResourceNotFoundException`, check to see if you passed a valid stream or consumer resource.

Request patterns can be one of the following:

- Data stream pattern: `arn:aws.*:kinesis:.*:\d{12}:.*stream/\S+`

- Consumer pattern: `^(arn):aws.*:kinesis:.*:\d{12}:.*stream\/[a-zA-Z0-9_.-]+\/consumer\/[a-zA-Z0-9_.-]+:[0-9]+`

For more information, see [Controlling Access to Amazon Kinesis Data Streams Resources Using IAM](https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Kinesis::ResourcePolicy",
  "Properties" : {
      "ResourceArn" : String,
      "ResourcePolicy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::Kinesis::ResourcePolicy
Properties:
  ResourceArn: String
  ResourcePolicy: Json

```

## Properties

`ResourceArn`

Returns the Amazon Resource Name (ARN) of the resource-based policy.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws.*:kinesis:.*:\d{12}:stream/\S+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourcePolicy`

This is the description for the resource policy.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

## Examples

### Resource policy that gives write access to a specific data stream

This policy allows Account12345 to perform the `DescribeStreamSummary`,
`ListShards`, `PutRecord`, and `PutRecords`
operations on the specified stream `datastreamABC`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Kinesis Data Streams resource policy example",
  "Resources": {
    "SampleStreamPolicy": {
      "Type": "AWS::Kinesis::ResourcePolicy",
      "Properties": {
        "ResourceArn": "arn:aws:kinesis:us-east-2:123456789012:stream/datastreamABC",
        "ResourcePolicy": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "WriteStatement",
              "Effect": "Allow",
              "Principal": {
                "AWS": "arn:aws:iam::123456789012:root"
              },
              "Action": [
                "kinesis:DescribeStreamSummary",
                "kinesis:ListShards",
                "kinesis:PutRecord",
                "kinesis:PutRecords"
              ],
              "Resource": "arn:aws:kinesis:us-east-2:123456789012:stream/datastreamABC"
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

SampleResourcePolicy:
  Type: AWS::Kinesis::ResourcePolicy
  Properties:
    ResourceArn: arn:aws:kinesis:us-east-2:123456789012:stream/datastreamABC
    ResourcePolicy:
      Version: 2012-10-17
      Statement:
        - Action:
            - 'kinesis:DescribeStreamSummary'
            - 'kinesis:ListShards'
            - 'kinesis:PutRecord'
            - 'kinesis:PutRecords'
          Effect: Allow
          Resource:
            - 'arn:aws:kinesis:us-east-2:123456789012:stream/datastreamABC'
          Principal:
            AWS: 'arn:aws:iam::123456789012:root'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Kinesis

AWS::Kinesis::Stream
