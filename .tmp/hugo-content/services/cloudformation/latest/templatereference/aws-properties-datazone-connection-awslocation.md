This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection AwsLocation

The location of a project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessRole" : String,
  "AwsAccountId" : String,
  "AwsRegion" : String,
  "IamConnectionId" : String
}

```

### YAML

```yaml

  AccessRole: String
  AwsAccountId: String
  AwsRegion: String
  IamConnectionId: String

```

## Properties

`AccessRole`

The access role of a connection.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccountId`

The account ID of a connection.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsRegion`

The Region of a connection.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z]{4,10}-\d$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamConnectionId`

The IAM connection ID of a connection.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationCodeProperties

BasicAuthenticationCredentials

All content copied from https://docs.aws.amazon.com/.
