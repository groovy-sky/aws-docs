This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection AmazonQPropertiesInput

The Amazon Q properties of the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthMode" : String,
  "IsEnabled" : Boolean,
  "ProfileArn" : String
}

```

### YAML

```yaml

  AuthMode: String
  IsEnabled: Boolean
  ProfileArn: String

```

## Properties

`AuthMode`

The authentication mode of the connection's Amazon Q properties.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsEnabled`

Specifies whether Amazon Q is enabled for the connection.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileArn`

The profile ARN of the connection's Amazon Q properties.

_Required_: No

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:[a-z0-9\-]+:[a-z0-9\-]*:[0-9]*:.*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataZone::Connection

AthenaPropertiesInput

All content copied from https://docs.aws.amazon.com/.
