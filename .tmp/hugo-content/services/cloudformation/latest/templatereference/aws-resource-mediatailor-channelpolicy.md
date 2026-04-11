This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::ChannelPolicy

Specifies an IAM policy for the channel. IAM policies are used to control access to your channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaTailor::ChannelPolicy",
  "Properties" : {
      "ChannelName" : String,
      "Policy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::MediaTailor::ChannelPolicy
Properties:
  ChannelName: String
  Policy: Json

```

## Properties

`ChannelName`

The name of the channel associated with this Channel Policy.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The IAM policy for the channel. IAM policies are used to control access to your channel.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeShiftConfiguration

AWS::MediaTailor::LiveSource

All content copied from https://docs.aws.amazon.com/.
