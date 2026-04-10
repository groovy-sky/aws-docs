This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::OptionGroup OptionSetting

The `OptionSetting` property type specifies the value for an option within
an `OptionSetting` property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the option that has settings that you can set.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The current value of the option setting.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Enable ingress to a DB security group

The following example creates an option group for the Oracle Enterprise Edition
engine, version 19, and configures it with the Oracle Enterprise Manager option,
specifying necessary option settings and a port.

#### JSON

```json

{
  "Resources": {
    "MyOptionGroup": {
      "Type": "AWS::RDS::OptionGroup",
      "Properties": {
        "EngineName": "oracle-ee",
        "MajorEngineVersion": "19",
        "OptionGroupDescription": "My Oracle OEM Option Group",
        "Options": [
          {
            "OptionName": "OEM",
            "Port": 5500,
            "OptionSettings": [
              {
                "Name": "OEM_HOST",
                "Value": "oem.example.com"
              },
              {
                "Name": "OEM_PORT",
                "Value": "5500"
              },
              {
                "Name": "SECURITY_GROUP_IDS",
                "Value": "sg-0123456789abcdef0"
              }
            ]
          }
        ],
        "Tags": [
          {
            "Key": "Name",
            "Value": "MyOptionGroup"
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
  MyOptionGroup:
    Type: AWS::RDS::OptionGroup
    Properties:
      EngineName: oracle-ee
      MajorEngineVersion: 19
      OptionGroupDescription: My Oracle OEM Option Group
      Options:
        - OptionName: OEM
          Port: 5500
          OptionSettings:
            - Name: OEM_HOST
              Value: oem.example.com
            - Name: OEM_PORT
              Value: 5500
            - Name: SECURITY_GROUP_IDS
              Value: sg-0123456789abcdef0
      Tags:
        - Key: Name
          Value: MyOptionGroup
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OptionConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
