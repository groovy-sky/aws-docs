This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput SecretsManagerEncryptionKeyConfiguration

The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  RoleArn: String
  SecretArn: String

```

## Properties

`RoleArn`

The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::[0-9]{12}:role/[a-zA-Z0-9_+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The ARN of the AWS Secrets Manager secret used for transit encryption.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):secretsmanager:[a-z0-9-]+:[0-9]{12}:secret:[a-zA-Z0-9/_+=.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RtpRouterInputConfiguration

SrtCallerRouterInputConfiguration

All content copied from https://docs.aws.amazon.com/.
