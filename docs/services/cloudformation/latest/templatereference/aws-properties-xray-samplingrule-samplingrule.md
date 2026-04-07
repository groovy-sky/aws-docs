This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::SamplingRule SamplingRule

A sampling rule that services use to decide whether to instrument a request. Rule
fields can match properties of the service, or properties of a request. The service can ignore
rules that don't match its properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : {Key: Value, ...},
  "FixedRate" : Number,
  "Host" : String,
  "HTTPMethod" : String,
  "Priority" : Integer,
  "ReservoirSize" : Integer,
  "ResourceARN" : String,
  "RuleARN" : String,
  "RuleName" : String,
  "SamplingRateBoost" : SamplingRateBoost,
  "ServiceName" : String,
  "ServiceType" : String,
  "URLPath" : String,
  "Version" : Integer
}

```

### YAML

```yaml

  Attributes:
    Key: Value
  FixedRate: Number
  Host: String
  HTTPMethod: String
  Priority: Integer
  ReservoirSize: Integer
  ResourceARN: String
  RuleARN: String
  RuleName: String
  SamplingRateBoost:
    SamplingRateBoost
  ServiceName: String
  ServiceType: String
  URLPath: String
  Version: Integer

```

## Properties

`Attributes`

Matches attributes derived from the request.

_Map Entries:_ Maximum number of 5 items.

_Key Length Constraints:_ Minimum length of 1. Maximum length of 32.

_Value Length Constraints:_ Minimum length of 1. Maximum length of 32.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixedRate`

The percentage of matching requests to instrument, after the reservoir is
exhausted.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

Matches the hostname from a request URL.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTPMethod`

Matches the HTTP method of a request.

_Required_: Yes

_Type_: String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority of the sampling rule.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `9999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReservoirSize`

A fixed number of matching requests to instrument per second, prior to applying the
fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceARN`

Matches the ARN of the AWS resource on which the service runs.

_Required_: Yes

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleARN`

The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.

###### Note

Specifying a sampling rule by name is recommended, as specifying by
ARN will be deprecated in future.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the sampling rule. Specify a rule by either name or ARN, but not both.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamplingRateBoost`

Specifies the multiplier applied to the base sampling rate.
This boost allows you to temporarily increase sampling without changing the rule's configuration.

_Required_: No

_Type_: [SamplingRateBoost](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-xray-samplingrule-samplingrateboost.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

Matches the `name` that the service uses to identify itself in segments.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceType`

Matches the `origin` that the service uses to identify its type in segments.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URLPath`

Matches the path from a request URL.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the sampling rule. `Version` can only be set when creating a new sampling rule.

_Required_: Conditional

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SamplingRateBoost

Tag
