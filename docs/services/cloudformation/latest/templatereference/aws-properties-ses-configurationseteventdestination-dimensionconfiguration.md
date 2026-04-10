This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSetEventDestination DimensionConfiguration

An object that defines the dimension configuration to use when you send email events
to Amazon CloudWatch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultDimensionValue" : String,
  "DimensionName" : String,
  "DimensionValueSource" : String
}

```

### YAML

```yaml

  DefaultDimensionValue: String
  DimensionName: String
  DimensionValueSource: String

```

## Properties

`DefaultDimensionValue`

The default value of the dimension that is published to Amazon CloudWatch if you don't provide the
value of the dimension when you send an email. This value has to meet the following
criteria:

- Can only contain ASCII letters (a–z, A–Z), numbers (0–9),
underscores (\_), or dashes (-), at signs (@), and periods (.).

- It can contain no more than 255 characters.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,256}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionName`

The name of an Amazon CloudWatch dimension associated with an email sending metric. The name has to
meet the following criteria:

- It can only contain ASCII letters (a–z, A–Z), numbers (0–9),
underscores (\_), or dashes (-).

- It can contain no more than 255 characters.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_:-]{1,256}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionValueSource`

The location where the Amazon SES API v2 finds the value of a dimension to publish to
Amazon CloudWatch. To use the message tags that you specify using an
`X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail`
or `SendRawEmail` API, choose `messageTag`. To use your own email
headers, choose `emailHeader`. To use link tags, choose
`linkTag`.

_Required_: Yes

_Type_: String

_Allowed values_: `messageTag | emailHeader | linkTag`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchDestination

EventBridgeDestination

All content copied from https://docs.aws.amazon.com/.
