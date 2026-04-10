This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy Filter

Defines filters that the lifecycle policy uses to determine impacted resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetainAtLeast" : Integer,
  "Type" : String,
  "Unit" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  RetainAtLeast: Integer
  Type: String
  Unit: String
  Value: Integer

```

## Properties

`RetainAtLeast`

For age-based filters, this is the number of resources to keep on hand after the lifecycle
`DELETE` action is applied. Impacted resources are only deleted if you have more than
this number of resources. If you have fewer resources than this number, the impacted resource
is not deleted.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Filter resources based on either `age` or `count`.

_Required_: Yes

_Type_: String

_Allowed values_: `AGE | COUNT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

Defines the unit of time that the lifecycle policy uses to determine impacted
resources. This is required for age-based rules.

_Required_: No

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The number of units for the time period or for the count. For example, a value of
`6` might refer to six months or six AMIs.

###### Note

For count-based filters, this value represents the minimum number of resources
to keep on hand. If you have fewer resources than this number, the resource is
excluded from lifecycle actions.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExclusionRules

IncludeResources

All content copied from https://docs.aws.amazon.com/.
