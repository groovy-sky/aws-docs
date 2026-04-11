This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CE::AnomalyMonitor ResourceTag

The tag structure that contains a tag key and value.

###### Note

Tagging is supported only for the following Cost Explorer resource types:
[`AnomalyMonitor`](../../../../reference/aws-cost-management/latest/apireference/api-anomalymonitor.md), [`AnomalySubscription`](../../../../reference/aws-cost-management/latest/apireference/api-anomalysubscription.md), [`CostCategory`](../../../../reference/aws-cost-management/latest/apireference/api-costcategory.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key that's associated with the tag.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!aws:).*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The value that's associated with the tag.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CE::AnomalyMonitor

AWS::CE::AnomalySubscription

All content copied from https://docs.aws.amazon.com/.
