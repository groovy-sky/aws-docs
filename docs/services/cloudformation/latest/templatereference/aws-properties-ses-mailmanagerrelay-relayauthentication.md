This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRelay RelayAuthentication

Authentication for the relay destination server—specify the secretARN where the SMTP
credentials are stored, or specify an empty NoAuthentication structure if the relay
destination server does not require SMTP credential authentication.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NoAuthentication" : Json,
  "SecretArn" : String
}

```

### YAML

```yaml

  NoAuthentication: Json
  SecretArn: String

```

## Properties

`NoAuthentication`

Keep an empty structure if the relay destination server does not require SMTP credential authentication.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The ARN of the secret created in secrets manager where the relay server's SMTP credentials are stored.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc):secretsmanager:[a-z0-9-]+:\d{12}:secret:[a-zA-Z0-9/_+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::MailManagerRelay

Tag

All content copied from https://docs.aws.amazon.com/.
