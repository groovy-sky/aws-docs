This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSetEventDestination PinpointDestination

An object that defines a Amazon Pinpoint destination for email events. You can use Amazon Pinpoint events
to create attributes in Amazon Pinpoint projects. You can use these attributes to create segments
for your campaigns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationArn" : String
}

```

### YAML

```yaml

  ApplicationArn: String

```

## Properties

`ApplicationArn`

The Amazon Resource Name (ARN) of the Amazon Pinpoint project that you want to send email
events to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseDestination

SnsDestination

All content copied from https://docs.aws.amazon.com/.
