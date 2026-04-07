This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::TargetGroup Matcher

Specifies the HTTP codes that healthy targets must use when responding to an HTTP health
check.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GrpcCode" : String,
  "HttpCode" : String
}

```

### YAML

```yaml

  GrpcCode: String
  HttpCode: String

```

## Properties

`GrpcCode`

You can specify values between 0 and 99. You can specify multiple values (for example,
"0,1") or a range of values (for example, "0-5"). The default value is 12.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpCode`

For Application Load Balancers, you can specify values between 200 and 499, with the
default value being 200. You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299").

For Network Load Balancers, you can specify values between 200 and 599, with the
default value being 200-399. You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299").

For Gateway Load Balancers, this must be "200–399".

Note that when using shorthand syntax, some values such as commas need to be
escaped.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElasticLoadBalancingV2::TargetGroup

Tag
