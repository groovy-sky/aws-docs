This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Channel CustomerManagedS3

Used to store channel data in an S3 bucket that you manage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "KeyPrefix" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  Bucket: String
  KeyPrefix: String
  RoleArn: String

```

## Properties

`Bucket`

The name of the S3 bucket in which channel data is stored.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9.\-_]*$`

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPrefix`

(Optional) The prefix used to create the keys of the channel data objects. Each object in
an S3 bucket has a key that is its unique identifier within the bucket (each object in a
bucket has exactly one key). The prefix must end with a forward slash (/).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9!_.*'()/{}:-]*/$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3
resources.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChannelStorage

RetentionPeriod

All content copied from https://docs.aws.amazon.com/.
