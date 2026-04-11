This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::UsagePlan

The `AWS::ApiGateway::UsagePlan` resource creates a usage plan for deployed APIs. A usage plan sets a target for the throttling and quota limits on individual client API keys. For more information, see [Creating and Using API Usage Plans in Amazon API Gateway](../../../apigateway/latest/developerguide/api-gateway-api-usage-plans.md) in the _API Gateway Developer Guide_.

In some cases clients can exceed the targets that you set. Don’t rely on usage plans to control costs. Consider using [AWS Budgets](../../../cost-management/latest/userguide/budgets-managing-costs.md) to monitor costs
and [AWS WAF](../../../waf/latest/developerguide/waf-chapter.md) to manage API requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::UsagePlan",
  "Properties" : {
      "ApiStages" : [ ApiStage, ... ],
      "Description" : String,
      "Quota" : QuotaSettings,
      "Tags" : [ Tag, ... ],
      "Throttle" : ThrottleSettings,
      "UsagePlanName" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::UsagePlan
Properties:
  ApiStages:
    - ApiStage
  Description: String
  Quota:
    QuotaSettings
  Tags:
    - Tag
  Throttle:
    ThrottleSettings
  UsagePlanName: String

```

## Properties

`ApiStages`

The associated API stages of a usage plan.

_Required_: No

_Type_: Array of [ApiStage](aws-properties-apigateway-usageplan-apistage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of a usage plan.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Quota`

The target maximum number of permitted requests per a given unit time interval.

_Required_: No

_Type_: [QuotaSettings](aws-properties-apigateway-usageplan-quotasettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-apigateway-usageplan-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Throttle`

A map containing method level throttling information for API stage in a usage plan.

_Required_: No

_Type_: [ThrottleSettings](aws-properties-apigateway-usageplan-throttlesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsagePlanName`

The name of a usage plan.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the usage plan ID, such as `abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID for the usage plan. For example: `abc123`.

## Examples

### Create usage plan

The following examples create a usage plan for the Prod API stage, with a quota of 5000 requests per month and a rate limit of 100 requests per second.

#### JSON

```json

{
    "usagePlan": {
        "Type": "AWS::ApiGateway::UsagePlan",
        "Properties": {
            "ApiStages": [
                {
                    "ApiId": {
                        "Ref": "MyRestApi"
                    },
                    "Stage": {
                        "Ref": "Prod"
                    }
                }
            ],
            "Description": "Customer ABC's usage plan",
            "Quota": {
                "Limit": 5000,
                "Period": "MONTH"
            },
            "Throttle": {
                "BurstLimit": 200,
                "RateLimit": 100
            },
            "UsagePlanName": "Plan_ABC"
        }
    }
}
```

#### YAML

```yaml

usagePlan:
  Type: 'AWS::ApiGateway::UsagePlan'
  Properties:
    ApiStages:
      - ApiId: !Ref MyRestApi
        Stage: !Ref Prod
    Description: Customer ABC's usage plan
    Quota:
      Limit: 5000
      Period: MONTH
    Throttle:
      BurstLimit: 200
      RateLimit: 100
    UsagePlanName: Plan_ABC

```

## See also

- [usageplan:create](../../../apigateway/latest/api/api-createusageplan.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ApiStage

All content copied from https://docs.aws.amazon.com/.
