This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign CollectionScheme

Specifies what data to collect and how often or when to collect it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionBasedCollectionScheme" : ConditionBasedCollectionScheme,
  "TimeBasedCollectionScheme" : TimeBasedCollectionScheme
}

```

### YAML

```yaml

  ConditionBasedCollectionScheme:
    ConditionBasedCollectionScheme
  TimeBasedCollectionScheme:
    TimeBasedCollectionScheme

```

## Properties

`ConditionBasedCollectionScheme`

Information about a collection scheme that uses a simple logical expression to
recognize what data to collect.

_Required_: No

_Type_: [ConditionBasedCollectionScheme](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-conditionbasedcollectionscheme.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeBasedCollectionScheme`

Information about a collection scheme that uses a time period to decide how often to
collect data.

_Required_: No

_Type_: [TimeBasedCollectionScheme](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-timebasedcollectionscheme.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTFleetWise::Campaign

ConditionBasedCollectionScheme
