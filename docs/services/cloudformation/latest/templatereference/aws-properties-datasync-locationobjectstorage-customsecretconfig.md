This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationObjectStorage CustomSecretConfig

Specifies configuration information for a customer-managed Secrets Manager secret where
a storage location credentials is stored in Secrets Manager as plain text (for authentication
token, secret key, or password) or as binary (for Kerberos keytab). This configuration includes
the secret ARN, and the ARN for an IAM role that provides access to the secret.

###### Note

You can use either `CmkSecretConfig` or `CustomSecretConfig` to
provide credentials for a `CreateLocation` request. Do not provide both
parameters for the same request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretAccessRoleArn" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  SecretAccessRoleArn: String
  SecretArn: String

```

## Properties

`SecretAccessRoleArn`

Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to
access the secret specified for `SecretArn`.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*|)$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

Specifies the ARN for an AWS Secrets Manager secret.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):secretsmanager:[a-z-0-9]+:[0-9]{12}:secret:.*|)$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CmkSecretConfig

ManagedSecretConfig

All content copied from https://docs.aws.amazon.com/.
