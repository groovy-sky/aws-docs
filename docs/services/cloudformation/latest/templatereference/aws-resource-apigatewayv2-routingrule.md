This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule

Represents a routing rule. When the incoming request to a domain name matches the conditions for a rule, API Gateway invokes a stage of a target API. Supported only for REST APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::RoutingRule",
  "Properties" : {
      "Actions" : [ Action, ... ],
      "Conditions" : [ Condition, ... ],
      "DomainNameArn" : String,
      "Priority" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::RoutingRule
Properties:
  Actions:
    - Action
  Conditions:
    - Condition
  DomainNameArn: String
  Priority: Integer

```

## Properties

`Actions`

The resulting action based on matching a routing rules condition. Only InvokeApi is supported.

_Required_: Yes

_Type_: Array of [Action](aws-properties-apigatewayv2-routingrule-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

The conditions of the routing rule.

_Required_: Yes

_Type_: Array of [Condition](aws-properties-apigatewayv2-routingrule-condition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainNameArn`

The ARN of the domain name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

The order in which API Gateway evaluates a rule. Priority is evaluated from the lowest value to the highest value. Rules can't have the same priority. Priority values 1-1,000,000 are supported.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the routing rule, such as `a1b2c3e`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RoutingRuleArn`

Represents an Amazon Resource Name (ARN).

`RoutingRuleId`

The identifier.

## Examples

### Routing rule creation

The following example creates a `RoutingRule` resource called `MyRoutingRule` that routes any incoming requests that have the header `x-version:beta` and the path `users` to the `prod` stage of REST API `abcd123`.

#### JSON

```json

{
  "Resources": {
    "MyRoutingRule": {
      "Type": "AWS::ApiGatewayV2::RoutingRule",
      "Properties": {
        "DomainNameArn": "arn:aws:apigateway:us-west-2::/domainnames/example.com",
        "Priority": 100,
        "Conditions": [
          {
            "MatchHeaders": {
              "AnyOf": [
                {
                  "Header": "x-version",
                  "ValueGlob": "beta"
                }
              ]
            }
          },
          {
            "MatchBasePaths": {
              "AnyOf": [
                "users"
              ]
            }
          }
        ],
        "Actions": [
          {
            "InvokeApi": {
              "ApiId": "abcd123",
              "Stage": "prod"
            }
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyRoutingRule:
    Type: 'AWS::ApiGatewayV2::RoutingRule'
    Properties:
      DomainNameArn: arn:aws:apigateway:us-west-2::/domainnames/example.com
      Priority: 100
      Conditions:
        - MatchHeaders:
            AnyOf:
              - Header: "x-version"
                ValueGlob: "beta"
        - MatchBasePaths:
            AnyOf:
              - "users"
      Actions:
        - InvokeApi:
            ApiId: "abcd123"
            Stage: "prod"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterConstraints

Action

All content copied from https://docs.aws.amazon.com/.
