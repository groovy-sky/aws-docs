This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset DatasetContentDeliveryRule

When dataset contents are created, they are delivered to destination specified
here.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : DatasetContentDeliveryRuleDestination,
  "EntryName" : String
}

```

### YAML

```yaml

  Destination:
    DatasetContentDeliveryRuleDestination
  EntryName: String

```

## Properties

`Destination`

The destination to which dataset contents are delivered.

_Required_: Yes

_Type_: [DatasetContentDeliveryRuleDestination](aws-properties-iotanalytics-dataset-datasetcontentdeliveryruledestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntryName`

The name of the dataset content delivery rules entry.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerAction

DatasetContentDeliveryRuleDestination

All content copied from https://docs.aws.amazon.com/.
