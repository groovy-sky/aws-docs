This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::RemediationConfiguration

An object that represents the details about the remediation configuration that includes the remediation action, parameters, and data to execute the action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::RemediationConfiguration",
  "Properties" : {
      "Automatic" : Boolean,
      "ConfigRuleName" : String,
      "ExecutionControls" : ExecutionControls,
      "MaximumAutomaticAttempts" : Integer,
      "Parameters" : {Key: Value, ...},
      "ResourceType" : String,
      "RetryAttemptSeconds" : Integer,
      "TargetId" : String,
      "TargetType" : String,
      "TargetVersion" : String
    }
}

```

### YAML

```yaml

Type: AWS::Config::RemediationConfiguration
Properties:
  Automatic: Boolean
  ConfigRuleName: String
  ExecutionControls:
    ExecutionControls
  MaximumAutomaticAttempts: Integer
  Parameters:
    Key: Value
  ResourceType: String
  RetryAttemptSeconds: Integer
  TargetId: String
  TargetType: String
  TargetVersion: String

```

## Properties

`Automatic`

The remediation is triggered automatically.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigRuleName`

The name of the AWS Config rule.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionControls`

An ExecutionControls object.

_Required_: No

_Type_: [ExecutionControls](aws-properties-config-remediationconfiguration-executioncontrols.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumAutomaticAttempts`

The maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.

For example, if you specify MaximumAutomaticAttempts as 5 with RetryAttemptSeconds as 50 seconds,

AWS Config will put a RemediationException on your behalf for the failing resource after the 5th failed attempt within 50 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

An object of the RemediationParameterValue. For more information, see [RemediationParameterValue](../../../../reference/config/latest/apireference/api-remediationparametervalue.md).

###### Note

The type is a map of strings to RemediationParameterValue.

_Required_: No

_Type_: Object of [RemediationParameterValue](aws-properties-config-remediationconfiguration-remediationparametervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of a resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryAttemptSeconds`

Time window to determine whether or not to add a remediation exception to prevent infinite remediation attempts.
If `MaximumAutomaticAttempts` remediation attempts have been made under `RetryAttemptSeconds`, a remediation exception will be added to the resource.
If you do not select a number, the default is 60 seconds.

For example, if you specify `RetryAttemptSeconds` as 50 seconds and `MaximumAutomaticAttempts` as 5,
AWS Config will run auto-remediations 5 times within 50 seconds before adding a remediation exception to the resource.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetId`

Target ID is the name of the SSM document.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetType`

The type of the target. Target executes remediation. For example, SSM document.

_Required_: Yes

_Type_: String

_Allowed values_: `SSM_DOCUMENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetVersion`

Version of the target. For example, version of the SSM document.

###### Note

If you make backward incompatible changes to the SSM document,
you must call PutRemediationConfiguration API again to ensure the remediations can run.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the remediation action with the associated SSM document.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Remeditation Configuration

The following example creates a remediation configuration using AWS Systems Manager document.

#### JSON

```json

{
    "BasicRemediationConfiguration": {
        "Type": "AWS::Config::RemediationConfiguration",
        "Properties": {
            "ConfigRuleName": "configRuleName",
            "Parameters": {
                "AutomationAssumeRole": {
                    "StaticValue": {
                        "Values": [
                            "automationAssumeRole"
                        ]
                    }
                },
                "InstanceId": {
                    "StaticValue": {
                        "Values": [
                            "instanceId"
                        ]
                    }
                }
            },
            "TargetId": "AWS-StartEC2Instance",
            "TargetType": "SSM_DOCUMENT",
            "TargetVersion": "1"
        }
    }
}
```

#### YAML

```yaml

BasicRemediationConfiguration:
    Type: "AWS::Config::RemediationConfiguration"
    Properties:
        ConfigRuleName: configRuleName
        Parameters:
            AutomationAssumeRole:
                StaticValue:
                    Values:
                    - automationAssumeRole
            InstanceId:
                StaticValue:
                    Values:
                        - instanceId
        TargetId: "AWS-StartEC2Instance"
        TargetType: "SSM_DOCUMENT"
        TargetVersion: "1"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConformancePackInputParameter

ExecutionControls

All content copied from https://docs.aws.amazon.com/.
