This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::Hub

The `AWS::SecurityHub::Hub` resource specifies the enablement of the
AWS Security Hub CSPM service in your AWS account. The service is enabled in the current AWS Region
or the specified Region. You create a separate `Hub` resource in
each Region in which you want to enable Security Hub CSPM.

When you use this resource to enable Security Hub CSPM, default security standards are enabled.
To disable default standards, set the `EnableDefaultStandards` property to `false`.
You can use the [`AWS::SecurityHub::Standard`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-standard.html) resource to
enable additional standards.

When you use this resource to enable Security Hub CSPM, new controls are automatically enabled for your enabled
standards. To disable automatic enablement of new controls, set the `AutoEnableControls` property to `false`.

You must create an `AWS::SecurityHub::Hub` resource for an account before
you can create other types of Security Hub CSPM resources for the account through CloudFormation. Use a
[DependsOn\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html), such as `"DependsOn": "Hub"`, to ensure that you've
created an `AWS::SecurityHub::Hub` resource before creating other Security Hub CSPM resources for an account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::Hub",
  "Properties" : {
      "AutoEnableControls" : Boolean,
      "ControlFindingGenerator" : String,
      "EnableDefaultStandards" : Boolean,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::Hub
Properties:
  AutoEnableControls: Boolean
  ControlFindingGenerator: String
  EnableDefaultStandards: Boolean
  Tags:
    Key: Value

```

## Properties

`AutoEnableControls`

Whether to automatically enable new controls when they are added to standards that are
enabled.

By default, this is set to `true`, and new controls are enabled
automatically. To not automatically enable new controls, set this to `false`.

When you automatically enable new controls, you can interact with the controls in
the console and programmatically immediately after release. However, automatically enabled controls have a temporary default status of
`DISABLED`. It can take up to several days for Security Hub CSPM to process the control release and designate the
control as `ENABLED` in your account. During the processing period, you can manually enable or disable a
control, and Security Hub CSPM will maintain that designation regardless of whether you have `AutoEnableControls` set to
`true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlFindingGenerator`

Specifies whether an account has consolidated control findings turned on or off.
If the value for this field is set to
`SECURITY_CONTROL`, Security Hub CSPM generates a single finding for a control check even when the check
applies to multiple enabled standards.

If the value for this field is set to `STANDARD_CONTROL`, Security Hub CSPM generates separate findings
for a control check when the check applies to multiple enabled standards.

The value for this field in a member account matches the value in the administrator
account. For accounts that aren't part of an organization, the default value of this field
is `SECURITY_CONTROL` if you enabled Security Hub CSPM on or after February 23,
2023.

_Required_: No

_Type_: String

_Pattern_: `^(SECURITY_CONTROL|STANDARD_CONTROL)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDefaultStandards`

Whether to enable the security standards that Security Hub CSPM has designated as
automatically enabled. If you don't provide a value for
`EnableDefaultStandards`, it is set to `true`, and the
designated standards are automatically enabled in each AWS Region where
you enable Security Hub CSPM. If you don't want to enable the designated standards,
set `EnableDefaultStandards` to `false`.

Currently, the automatically enabled standards are the Center for Internet Security
(CIS) AWS Foundations Benchmark v1.2.0 and AWS
Foundational Security Best Practices (FSBP).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `HubArn` for the hub resource created, such as `arn:aws:securityhub:us-east-1:123456789012:hub/default`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ARN`

The Amazon Resource Name (ARN) of the `Hub` resource that was retrieved.

`SubscribedAt`

The date and time when Security Hub CSPM was enabled in your account.

## Examples

The following examples show how to declare an `AWS::SecurityHub::Hub` resource.

- [Creating a Hub resource that enables default standards and turns on consolidated control findings](#aws-resource-securityhub-hub--examples--Creating_a_Hub_resource_that_enables_default_standards_and_turns_on_consolidated_control_findings)

- [Creating a Hub resource that disables default standards and turns off consolidated control findings](#aws-resource-securityhub-hub--examples--Creating_a_Hub_resource_that_disables_default_standards_and_turns_off_consolidated_control_findings)

### Creating a `Hub` resource that enables default standards and turns on consolidated control findings

In this example, the default standards are automatically enabled, and
consolidated control findings is turned on.

#### JSON

```json

{
    "Description": "Example template to create a Hub resource",
    "Resources": {
        "ExampleHubWithTags": {
            "Type": "AWS::SecurityHub::Hub",
            "Properties": {
                "Tags": {
                    "key1": "value1",
                    "key2": "value2"
                },
                "EnableDefaultStandards": true,
                "ControlFindingGenerator": "SECURITY_CONTROL"
            }
        }
    },
    "Outputs": {
        "HubArn": {
            "Value": {
                "Ref": "ExampleHubWithTags"
            }
        }
    }
}
```

#### YAML

```yaml

Description: Example template to create a Hub resource
Resources:
  ExampleHubWithTags:
    Type: 'AWS::SecurityHub::Hub'
    Properties:
      Tags:
        key1: value1
        key2: value2
      EnableDefaultStandards: true
      ControlFindingGenerator: 'SECURITY_CONTROL'

Outputs:
  HubArn:
    Value: !Ref ExampleHubWithTags
```

### Creating a `Hub` resource that disables default standards and turns off consolidated control findings

In this example, the default standards are disabled, and consolidated control
findings is turned off.

#### JSON

```json

{
    "Description": "Example template to create a Hub resource",
    "Resources": {
        "ExampleHubWithTags": {
            "Type": "AWS::SecurityHub::Hub",
            "Properties": {
                "Tags": {
                    "key1": "value1",
                    "key2": "value2"
                },
                "EnableDefaultStandards": false,
                "ControlFindingGenerator": "STANDARD_CONTROL"
            }
        }
    },
    "Outputs": {
        "HubArn": {
            "Value": {
                "Ref": "ExampleHubWithTags"
            }
        }
    }
}
```

#### YAML

```yaml

Description: Example template to create a Hub resource
Resources:
  ExampleHubWithTags:
    Type: 'AWS::SecurityHub::Hub'
    Properties:
      Tags:
        key1: value1
        key2: value2
      EnableDefaultStandards: false
      ControlFindingGenerator: 'STANDARD_CONTROL'

Outputs:
  HubArn:
    Value: !Ref ExampleHubWithTags
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityHub::FindingAggregator

AWS::SecurityHub::HubV2
