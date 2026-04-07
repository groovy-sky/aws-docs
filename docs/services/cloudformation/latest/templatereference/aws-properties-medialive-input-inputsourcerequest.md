This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input InputSourceRequest

Settings that apply only if the input is a pull type of input.

The parent of this entity is Input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PasswordParam" : String,
  "Url" : String,
  "Username" : String
}

```

### YAML

```yaml

  PasswordParam: String
  Url: String
  Username: String

```

## Properties

`PasswordParam`

The password parameter that holds the password for accessing the upstream system. The password parameter applies
only if the upstream system requires credentials.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

For a pull input, the URL where MediaLive pulls the source content from.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The user name to connect to the upstream system. The user name applies only if the upstream system requires
credentials.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InputSdpLocation

InputVpcRequest
