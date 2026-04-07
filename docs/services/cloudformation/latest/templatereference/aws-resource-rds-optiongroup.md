This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::OptionGroup

The `AWS::RDS::OptionGroup` resource creates or updates an option group, to enable and
configure features that are specific to a particular DB engine.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::OptionGroup",
  "Properties" : {
      "EngineName" : String,
      "MajorEngineVersion" : String,
      "OptionConfigurations" : [ OptionConfiguration, ... ],
      "OptionGroupDescription" : String,
      "OptionGroupName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::OptionGroup
Properties:
  EngineName: String
  MajorEngineVersion: String
  OptionConfigurations:
    - OptionConfiguration
  OptionGroupDescription: String
  OptionGroupName: String
  Tags:
    - Tag

```

## Properties

`EngineName`

Specifies the name of the engine that this option group should be associated with.

Valid Values:

- `mariadb`

- `mysql`

- `oracle-ee`

- `oracle-ee-cdb`

- `oracle-se2`

- `oracle-se2-cdb`

- `postgres`

- `sqlserver-ee`

- `sqlserver-se`

- `sqlserver-ex`

- `sqlserver-web`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MajorEngineVersion`

Specifies the major version of the engine that this option group should be associated with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OptionConfigurations`

A list of all available options for an option group.

_Required_: Conditional

_Type_: Array of [OptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-optiongroup-optionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionGroupDescription`

The description of the option group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OptionGroupName`

The name of the option group to be created.

Constraints:

- Must be 1 to 255 letters, numbers, or hyphens

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

Example: `myoptiongroup`

If you don't specify a value for `OptionGroupName` property, a name is automatically created for the option group.

###### Note

This value is stored as a lowercase string.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to assign to the option group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-optiongroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the option group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create an option group with multiple option configurations](#aws-resource-rds-optiongroup--examples--Create_an_option_group_with_multiple_option_configurations)

- [Create an option group with multiple option settings](#aws-resource-rds-optiongroup--examples--Create_an_option_group_with_multiple_option_settings)

- [Microsoft SQL Server native backup and restore option](#aws-resource-rds-optiongroup--examples--Microsoft_SQL_Server_native_backup_and_restore_option)

### Create an option group with multiple option configurations

The following example creates an option group with two option configurations
( `OEM` and `APEX`). For more information about these
options, see [Adding options to Oracle DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.Options.html)
in the _Amazon RDS User Guide_.

#### JSON

```json

{
    "OracleOptionGroup": {
        "Type": "AWS::RDS::OptionGroup",
        "Properties": {
            "EngineName": "oracle-ee",
            "MajorEngineVersion": "12.1",
            "OptionGroupDescription": "A test option group",
            "OptionConfigurations": [
                {
                    "OptionName": "OEM",
                    "DBSecurityGroupMemberships": [
                        "default"
                    ],
                    "Port": "5500"
                },
                {
                    "OptionName": "APEX"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

---
OracleOptionGroup:
  Type: AWS::RDS::OptionGroup
  Properties:
    EngineName: oracle-ee
    MajorEngineVersion: "12.1"
    OptionConfigurations:
      -
        DBSecurityGroupMemberships:
          - default
        OptionName: OEM
        Port: "5500"
      -
        OptionName: APEX
    OptionGroupDescription: A test option group
```

### Create an option group with multiple option settings

The following example creates an option group that specifies two option settings
for the `MEMCACHED` option. For more information about this option, see
[MySQL\
memcached support](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.memcached.html) in the _Amazon RDS User_
_Guide_.

#### JSON

```json

{
    "SQLOptionGroup": {
        "Type": "AWS::RDS::OptionGroup",
        "Properties": {
            "EngineName": "mysql",
            "MajorEngineVersion": "8.0",
            "OptionGroupDescription": "A test option group",
            "OptionConfigurations": [
                {
                    "OptionName": "MEMCACHED",
                    "VpcSecurityGroupMemberships": [
                        "sg-a1238db7"
                    ],
                    "Port": "1234",
                    "OptionSettings": [
                        {
                            "Name": "CHUNK_SIZE",
                            "Value": "32"
                        },
                        {
                            "Name": "BINDING_PROTOCOL",
                            "Value": "ascii"
                        }
                    ]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

---
SQLOptionGroup:
  Properties:
    EngineName: mysql
    MajorEngineVersion: "8.0"
    OptionConfigurations:
      -
        OptionName: MEMCACHED
        OptionSettings:
          -
            Name: CHUNK_SIZE
            Value: "32"
          -
            Name: BINDING_PROTOCOL
            Value: ascii
        Port: "1234"
        VpcSecurityGroupMemberships:
          - sg-a1238db7
    OptionGroupDescription: "A test option group"
  Type: AWS::RDS::OptionGroup

```

### Microsoft SQL Server native backup and restore option

The following example creates an option group that specifies the Microsoft SQL
Server native backup and restore option. For more information about this option, see
[Support for Native Backup and Restore in SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.BackupRestore.html) in the
_Amazon RDS User Guide_.

#### JSON

```json

{
    "myOptionGroup": {
        "Type": "AWS::RDS::OptionGroup",
        "Properties": {
            "EngineName": "sqlserver-se",
            "MajorEngineVersion": "12.00",
            "OptionGroupDescription": "SQL Server Native Backup and Restore",
            "OptionConfigurations": [
                {
                    "OptionName": "SQLSERVER_BACKUP_RESTORE",
                    "OptionSettings": [
                        {
                            "Name": "IAM_ROLE_ARN",
                            "Value": "arn:aws:iam::333333333333333:role/service-role/sqlserverrestore"
                        }
                    ]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

---
myOptionGroup:
  Type: 'AWS::RDS::OptionGroup'
  Properties:
    EngineName: sqlserver-se
    MajorEngineVersion: '12.00'
    OptionGroupDescription: SQL Server Native Backup and Restore
    OptionConfigurations:
      - OptionName: SQLSERVER_BACKUP_RESTORE
        OptionSettings:
          - Name: IAM_ROLE_ARN
            Value: 'arn:aws:iam::333333333333333:role/service-role/sqlserverrestore'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

OptionConfiguration
