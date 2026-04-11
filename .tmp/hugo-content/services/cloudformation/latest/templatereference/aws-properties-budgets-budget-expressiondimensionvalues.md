This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget ExpressionDimensionValues

Contains the specifications for the filters to use for your request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "MatchOptions" : [ String, ... ],
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  MatchOptions:
    - String
  Values:
    - String

```

## Properties

`Key`

The name of the dimension that you want to filter on.

_Required_: No

_Type_: String

_Allowed values_: `AZ | INSTANCE_TYPE | LINKED_ACCOUNT | LINKED_ACCOUNT_NAME | OPERATION | PURCHASE_TYPE | REGION | SERVICE | SERVICE_CODE | USAGE_TYPE | USAGE_TYPE_GROUP | RECORD_TYPE | OPERATING_SYSTEM | TENANCY | SCOPE | PLATFORM | SUBSCRIPTION_ID | LEGAL_ENTITY_NAME | INVOICING_ENTITY | DEPLOYMENT_OPTION | DATABASE_ENGINE | CACHE_ENGINE | INSTANCE_TYPE_FAMILY | BILLING_ENTITY | RESERVATION_ID | RESOURCE_ID | RIGHTSIZING_TYPE | SAVINGS_PLANS_TYPE | SAVINGS_PLAN_ARN | PAYMENT_OPTION | RESERVATION_MODIFIED | TAG_KEY | COST_CATEGORY_NAME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchOptions`

The match options that you can use to filter your results. You can specify only one of these
values in the array.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The metadata values you can specify to filter upon, so that the results all match at least
one of the specified values.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Expression

HistoricalOptions

All content copied from https://docs.aws.amazon.com/.
