This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::SamplingRule

Use the `AWS::XRay::SamplingRule` resource to specify a sampling rule, which controls sampling behavior for instrumented applications.
Include a `SamplingRule` entity to create or update a sampling rule.

###### Note

`SamplingRule.Version` can only be set when creating a sampling rule. Updating the version
will cause the update to fail.

Services retrieve rules with [GetSamplingRules](../../../xray/latest/api/api-getsamplingrules.md), and evaluate each rule in ascending
order of _priority_ for each request. If a rule matches, the service records a trace, borrowing it from the reservoir size. After 10 seconds, the service
reports back to X-Ray with [GetSamplingTargets](../../../xray/latest/api/api-getsamplingtargets.md) to get updated versions of
each in-use rule. The updated rule contains a trace quota that the service can use instead of borrowing from the reservoir.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::XRay::SamplingRule",
  "Properties" : {
      "SamplingRule" : SamplingRule,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::XRay::SamplingRule
Properties:
  SamplingRule:
    SamplingRule
  Tags:
    - Tag

```

## Properties

`SamplingRule`

The sampling rule to be created or updated.

_Required_: No

_Type_: [SamplingRule](aws-properties-xray-samplingrule-samplingrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-xray-samplingrule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the sampling rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RuleARN`

The sampling rule ARN that was created or updated.

## Examples

- [Create sampling rule](#aws-resource-xray-samplingrule--examples--Create_sampling_rule)

- [Update sampling rule](#aws-resource-xray-samplingrule--examples--Update_sampling_rule)

### Create sampling rule

This example creates a new sampling rule called MySamplingRule.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09T00:00:00.000Z",
   "Resources": {
        "MySamplingRuleResource": {
         "Type": "AWS::XRay::SamplingRule",
         "Properties": {
            "SamplingRule": {
               "RuleName": "MySamplingRule",
               "ResourceARN": "*",
               "Priority": 2,
               "FixedRate": 0.05,
               "ReservoirSize": 50,
               "ServiceName": "MyServiceName",
               "ServiceType": "MyServiceType",
               "Host": "MyHost",
               "HTTPMethod": "GET",
               "URLPath": "*",
               "Version": 1
            }
         }
      }
   }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
   MySamplingRuleResource:
      Type: AWS::XRay::SamplingRule
      Properties:
         SamplingRule:
            RuleName: MySamplingRule
            ResourceARN: "*"
            Priority: 2
            FixedRate: 0.05
            ReservoirSize: 50
            ServiceName: "MyServiceName"
            ServiceType: "MyServiceType"
            Host: "MyHost"
            HTTPMethod: "GET"
            URLPath: "*"
            Version: 1

```

### Update sampling rule

This example updates an existing sampling rule called MySamplingRule.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09T00:00:00.000Z",
   "Resources": {
        "MySamplingRuleResource": {
         "Type": "AWS::XRay::SamplingRule",
         "Properties": {
            "SamplingRule": {
               "RuleName": "MySamplingRule",
               "ResourceARN": "*",
               "Priority": 1,
               "FixedRate": 0.07,
               "ReservoirSize": 20,
               "ServiceName": "MyServiceName",
               "ServiceType": "MyServiceType",
               "Host": "MyHost",
               "HTTPMethod": "GET",
               "URLPath": "*"
            }
         }
      }
   }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
   MySamplingRuleResource:
      Type: AWS::XRay::SamplingRule
      Properties:
         SamplingRule:
            RuleName: MySamplingRule
            ResourceARN: "*"
            Priority: 1
            FixedRate: 0.07
            ReservoirSize: 20
            ServiceName: "MyServiceName"
            ServiceType: "MyServiceType"
            Host: "MyHost"
            HTTPMethod: "GET"
            URLPath: "*"

```

## See also

- [Configuring sampling rules in the X-Ray console](../../../xray/latest/devguide/xray-console-sampling.md)

- [Using sampling rules with the X-Ray API](../../../xray/latest/devguide/xray-api-sampling.md)

- [CreateSamplingRule](../../../xray/latest/api/api-createsamplingrule.md) action in the X-Ray API Reference

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::XRay::ResourcePolicy

SamplingRateBoost

All content copied from https://docs.aws.amazon.com/.
