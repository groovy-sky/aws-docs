This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxONTAP ManagedSecretConfig

Specifies configuration information for a DataSync-managed secret, such as an
authentication token or set of credentials that DataSync uses to access a specific
transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretArn" : String
}

```

### YAML

```yaml

  SecretArn: String

```

## Properties

`SecretArn`

Specifies the ARN for an AWS Secrets Manager secret.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):secretsmanager:[a-z-0-9]+:[0-9]{12}:secret:.*|)$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomSecretConfig

NFS

All content copied from https://docs.aws.amazon.com/.
