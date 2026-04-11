This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase RedshiftServerlessAuthConfiguration

Specifies configurations for authentication to a Redshift Serverless. Specify the type of authentication to use in the `type` field and include the corresponding field. If you specify IAM authentication, you don't need to include another field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "UsernamePasswordSecretArn" : String
}

```

### YAML

```yaml

  Type: String
  UsernamePasswordSecretArn: String

```

## Properties

`Type`

The type of authentication to use.

_Required_: Yes

_Type_: String

_Allowed values_: `IAM | USERNAME_PASSWORD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UsernamePasswordSecretArn`

The ARN of an Secrets Manager secret for authentication.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):secretsmanager:[a-z0-9-]{1,20}:([0-9]{12}|):secret:[a-zA-Z0-9!/_+=.@-]{1,512}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftQueryEngineStorageConfiguration

RedshiftServerlessConfiguration

All content copied from https://docs.aws.amazon.com/.
