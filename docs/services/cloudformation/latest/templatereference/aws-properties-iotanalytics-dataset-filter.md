This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset Filter

Information which is used to filter message data, to segregate it according to the time
frame in which it arrives.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeltaTime" : DeltaTime
}

```

### YAML

```yaml

  DeltaTime:
    DeltaTime

```

## Properties

`DeltaTime`

Used to limit data to that which has arrived since the last execution of the action.

_Required_: No

_Type_: [DeltaTime](aws-properties-iotanalytics-dataset-deltatime.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeltaTimeSessionWindowConfiguration

GlueConfiguration

All content copied from https://docs.aws.amazon.com/.
