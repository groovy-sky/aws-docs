This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::Subscriber Source

Sources are logs and events generated from a single system that match a specific event class in the Open Cybersecurity Schema Framework (OCSF) schema. Amazon Security Lake can collect logs and events from a variety of sources, including natively supported AWS services and third-party custom sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsLogSource" : AwsLogSource,
  "CustomLogSource" : CustomLogSource
}

```

### YAML

```yaml

  AwsLogSource:
    AwsLogSource
  CustomLogSource:
    CustomLogSource

```

## Properties

`AwsLogSource`

The natively supported AWS service which is used a Amazon Security Lake source to collect logs and events from.

_Required_: No

_Type_: [AwsLogSource](aws-properties-securitylake-subscriber-awslogsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomLogSource`

The custom log source AWS which is used a Amazon Security Lake source to collect logs and events from.

_Required_: No

_Type_: [CustomLogSource](aws-properties-securitylake-subscriber-customlogsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomLogSource

SubscriberIdentity

All content copied from https://docs.aws.amazon.com/.
