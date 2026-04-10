This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster JwtKey

The JWT key stored in AWS Secrets Manager for Slurm REST API authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretArn" : String,
  "SecretVersion" : String
}

```

### YAML

```yaml

  SecretArn: String
  SecretVersion: String

```

## Properties

`SecretArn`

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the JWT key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretVersion`

The version of the AWS Secrets Manager secret containing the JWT key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JwtAuth

Networking

All content copied from https://docs.aws.amazon.com/.
