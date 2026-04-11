This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection RedshiftCredentials

Amazon Redshift credentials of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretArn" : String,
  "UsernamePassword" : UsernamePassword
}

```

### YAML

```yaml

  SecretArn: String
  UsernamePassword:
    UsernamePassword

```

## Properties

`SecretArn`

The secret ARN of the Amazon Redshift credentials of a connection.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:secretsmanager:[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]:\d{12}:secret:.*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsernamePassword`

The username and password of the Amazon Redshift credentials of a connection.

_Required_: No

_Type_: [UsernamePassword](aws-properties-datazone-connection-usernamepassword.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PhysicalConnectionRequirements

RedshiftLineageSyncConfigurationInput

All content copied from https://docs.aws.amazon.com/.
