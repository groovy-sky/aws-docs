This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSetEventDestination DimensionConfiguration

An array of objects that define the dimensions to use when you send email events to Amazon
CloudWatch.

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

The default value of the dimension that is published to Amazon CloudWatch if you don't
provide the value of the dimension when you send an email. This value has to meet the
following criteria:

- It can only contain ASCII letters (a–z, A–Z), numbers (0–9),
underscores (\_), or dashes (-).

- It can contain no more than 256 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionName`

The name of an Amazon CloudWatch dimension associated with an email sending metric. The
name has to meet the following criteria:

- It can only contain ASCII letters (a–z, A–Z), numbers (0–9),
underscores (\_), or dashes (-).

- It can contain no more than 256 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionValueSource`

The location where Amazon Pinpoint finds the value of a dimension to publish to Amazon
CloudWatch. Acceptable values: `MESSAGE_TAG`, `EMAIL_HEADER`, and
`LINK_TAG`.

If you want Amazon Pinpoint to use the message tags that you specify using an
`X-SES-MESSAGE-TAGS` header or a parameter to the `SendEmail` API,
choose `MESSAGE_TAG`. If you want Amazon Pinpoint to use your own email headers,
choose `EMAIL_HEADER`. If you want Amazon Pinpoint to use tags that are specified
in your links, choose `LINK_TAG`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchDestination

EventDestination

All content copied from https://docs.aws.amazon.com/.
