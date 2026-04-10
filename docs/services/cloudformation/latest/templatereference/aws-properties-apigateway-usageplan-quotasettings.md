This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::UsagePlan QuotaSettings

`QuotaSettings` is a property of the [AWS::ApiGateway::UsagePlan](../userguide/aws-resource-apigateway-usageplan.md) resource that specifies a target for the maximum number of requests users can make to your REST APIs.

In some cases clients can exceed the targets that you set. Don’t rely on usage plans to control costs. Consider using [AWS Budgets](../../../cost-management/latest/userguide/budgets-managing-costs.md) to monitor costs
and [AWS WAF](../../../waf/latest/developerguide/waf-chapter.md) to manage API requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Limit" : Integer,
  "Offset" : Integer,
  "Period" : String
}

```

### YAML

```yaml

  Limit: Integer
  Offset: Integer
  Period: String

```

## Properties

`Limit`

The target maximum number of requests that can be made in a given time period.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Offset`

The number of requests subtracted from the given limit in the initial time period.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".

_Required_: No

_Type_: String

_Allowed values_: `DAY | WEEK | MONTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [UsagePlan](../../../apigateway/latest/api/api-usageplan.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiStage

Tag

All content copied from https://docs.aws.amazon.com/.
