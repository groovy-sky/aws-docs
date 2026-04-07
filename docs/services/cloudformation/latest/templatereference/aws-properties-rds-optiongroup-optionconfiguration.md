This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::OptionGroup OptionConfiguration

The `OptionConfiguration` property type specifies an individual option, and
its settings, within an `AWS::RDS::OptionGroup` resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DBSecurityGroupMemberships" : [ String, ... ],
  "OptionName" : String,
  "OptionSettings" : [ OptionSetting, ... ],
  "OptionVersion" : String,
  "Port" : Integer,
  "VpcSecurityGroupMemberships" : [ String, ... ]
}

```

### YAML

```yaml

  DBSecurityGroupMemberships:
    - String
  OptionName: String
  OptionSettings:
    - OptionSetting
  OptionVersion: String
  Port: Integer
  VpcSecurityGroupMemberships:
    - String

```

## Properties

`DBSecurityGroupMemberships`

A list of DB security groups used for this option.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionName`

The configuration of options to include in a group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionSettings`

The option settings to include in an option group.

_Required_: No

_Type_: Array of [OptionSetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-optiongroup-optionsetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionVersion`

The version for the option.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The optional port for the option.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupMemberships`

A list of VPC security group names used for this option.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Specify an option configuration

The following example template uses `OptionName` and
`OptionVersion` parameters when creating an
`AWS::RDS::OptionGroup` resource.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "APEX has a dependency on XMLDB, so, there must be at least one XMLDB when there is an APEX",
    "Parameters": {
        "OptionName": {
            "Type": "String"
        },
        "OptionVersion": {
            "Type": "String"
        }
    },
    "Resources": {
        "myOptionGroup": {
            "Type": "AWS::RDS::OptionGroup",
            "Properties": {
                "EngineName": "oracle-ee",
                "MajorEngineVersion": "11.2",
                "OptionGroupDescription": "testing creating optionGroup with APEX version",
                "OptionConfigurations": [
                    {
                        "OptionName": "XMLDB"
                    },
                    {
                        "OptionName": {
                            "Ref": "OptionName"
                        },
                        "OptionVersion": {
                            "Ref": "OptionVersion"
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

---
AWSTemplateFormatVersion: 2010-09-09
Description: "APEX has a dependency on XMLDB, so, there must be at least one XMLDB when there is an APEX"
Parameters:
  OptionName:
    Type: String
  OptionVersion:
    Type: String
Resources:
  myOptionGroup:
    Properties:
      EngineName: oracle-ee
      MajorEngineVersion: "11.2"
      OptionConfigurations:
        -
          OptionName: XMLDB
        -
          OptionName: OptionName
          OptionVersion: OptionVersion
      OptionGroupDescription: "testing creating optionGroup with APEX version"
    Type: AWS::RDS::OptionGroup
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RDS::OptionGroup

OptionSetting
