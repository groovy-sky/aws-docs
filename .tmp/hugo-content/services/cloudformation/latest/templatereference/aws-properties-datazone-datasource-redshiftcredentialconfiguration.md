This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource RedshiftCredentialConfiguration

The details of the credentials required to access an Amazon Redshift cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretManagerArn" : String
}

```

### YAML

```yaml

  SecretManagerArn: String

```

## Properties

`SecretManagerArn`

The ARN of a secret manager for an Amazon Redshift cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[^:]*:secretsmanager:[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]:\d{12}:secret:.*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftClusterStorage

RedshiftRunConfigurationInput

All content copied from https://docs.aws.amazon.com/.
