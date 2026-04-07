This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::BatchScramSecret

Represents a secret stored in the AWS Secrets Manager that can be used to authenticate with a cluster using a user name and a password.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::BatchScramSecret",
  "Properties" : {
      "ClusterArn" : String,
      "SecretArnList" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MSK::BatchScramSecret
Properties:
  ClusterArn: String
  SecretArnList:
    - String

```

## Properties

`ClusterArn`

The Amazon Resource Name (ARN) that uniquely identifies the cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecretArnList`

List of Amazon Resource Name (ARN)s of Secrets Manager secrets.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic function, `Ref` returns the secret stored in the Secrets Manager.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Managed Streaming for Apache Kafka

AWS::MSK::Cluster
