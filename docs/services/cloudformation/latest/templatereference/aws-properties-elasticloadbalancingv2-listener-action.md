This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener Action

Specifies an action for a listener rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticateCognitoConfig" : AuthenticateCognitoConfig,
  "AuthenticateOidcConfig" : AuthenticateOidcConfig,
  "FixedResponseConfig" : FixedResponseConfig,
  "ForwardConfig" : ForwardConfig,
  "JwtValidationConfig" : JwtValidationConfig,
  "Order" : Integer,
  "RedirectConfig" : RedirectConfig,
  "TargetGroupArn" : String,
  "Type" : String
}

```

### YAML

```yaml

  AuthenticateCognitoConfig:
    AuthenticateCognitoConfig
  AuthenticateOidcConfig:
    AuthenticateOidcConfig
  FixedResponseConfig:
    FixedResponseConfig
  ForwardConfig:
    ForwardConfig
  JwtValidationConfig:
    JwtValidationConfig
  Order: Integer
  RedirectConfig:
    RedirectConfig
  TargetGroupArn: String
  Type: String

```

## Properties

`AuthenticateCognitoConfig`

\[HTTPS listeners\] Information for using Amazon Cognito to authenticate users. Specify only
when `Type` is `authenticate-cognito`.

_Required_: No

_Type_: [AuthenticateCognitoConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listener-authenticatecognitoconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticateOidcConfig`

\[HTTPS listeners\] Information about an identity provider that is compliant with OpenID
Connect (OIDC). Specify only when `Type` is `authenticate-oidc`.

_Required_: No

_Type_: [AuthenticateOidcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixedResponseConfig`

\[Application Load Balancer\] Information for creating an action that returns a custom HTTP
response. Specify only when `Type` is `fixed-response`.

_Required_: No

_Type_: [FixedResponseConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listener-fixedresponseconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardConfig`

Information for creating an action that distributes requests among multiple target
groups. Specify only when `Type` is `forward`.

If you specify both `ForwardConfig` and `TargetGroupArn`, you can
specify only one target group using `ForwardConfig` and it must be the same
target group specified in `TargetGroupArn`.

_Required_: No

_Type_: [ForwardConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwtValidationConfig`

\[HTTPS listeners\] Information for validating JWT access tokens in client requests.
Specify only when `Type` is `jwt-validation`.

_Required_: No

_Type_: [JwtValidationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listener-jwtvalidationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Order`

The order for the action. This value is required for rules with multiple actions. The
action with the lowest value for order is performed first.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedirectConfig`

\[Application Load Balancer\] Information for creating a redirect action. Specify only when
`Type` is `redirect`.

_Required_: No

_Type_: [RedirectConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listener-redirectconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupArn`

The Amazon Resource Name (ARN) of the target group. Specify only when `Type` is
`forward` and you want to route to a single target group. To route to multiple
target groups, you must use `ForwardConfig` instead.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of action.

_Required_: Yes

_Type_: String

_Allowed values_: `forward | authenticate-oidc | authenticate-cognito | redirect | fixed-response | jwt-validation`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElasticLoadBalancingV2::Listener

AuthenticateCognitoConfig
