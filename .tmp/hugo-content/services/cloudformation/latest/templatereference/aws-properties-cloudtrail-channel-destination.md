This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Channel Destination

Contains information about the destination receiving events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Location" : String,
  "Type" : String
}

```

### YAML

```yaml

  Location: String
  Type: String

```

## Properties

`Location`

For channels used for a CloudTrail Lake integration, the location is the ARN of an event data store that receives events from a channel.
For service-linked channels, the location is the name of the AWS service.

_Required_: Yes

_Type_: String

_Pattern_: `(^[a-zA-Z0-9._/\-:]+$)`

_Minimum_: `3`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of destination for events arriving from a channel. For channels used for a CloudTrail Lake integration, the value is `EVENT_DATA_STORE`. For service-linked channels,
the value is `AWS_SERVICE`.

_Required_: Yes

_Type_: String

_Allowed values_: `EVENT_DATA_STORE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudTrail::Channel

Tag

All content copied from https://docs.aws.amazon.com/.
