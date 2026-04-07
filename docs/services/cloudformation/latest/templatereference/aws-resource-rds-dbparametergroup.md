This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBParameterGroup

The `AWS::RDS::DBParameterGroup` resource creates a custom parameter group
for an RDS database family.

This type can be declared in a template and referenced in the
`DBParameterGroupName` property of an `
                    AWS::RDS::DBInstance
                ` resource.

For information about configuring parameters for Amazon RDS DB instances, see
[Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html)
in the _Amazon RDS User Guide_.

For information about configuring parameters for Amazon Aurora DB instances, see
[Working with parameter \
groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the _Amazon Aurora User Guide_.

###### Note

Applying a parameter group to a DB instance may require the DB instance to reboot,
resulting in a database outage for the duration of the reboot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBParameterGroup",
  "Properties" : {
      "DBParameterGroupName" : String,
      "Description" : String,
      "Family" : String,
      "Parameters" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBParameterGroup
Properties:
  DBParameterGroupName: String
  Description: String
  Family: String
  Parameters: Json
  Tags:
    - Tag

```

## Properties

`DBParameterGroupName`

The name of the DB parameter group.

Constraints:

- Must be 1 to 255 letters, numbers, or hyphens.

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

If you don't specify a value for `DBParameterGroupName` property, a name is automatically created for the DB parameter group.

###### Note

This value is stored as a lowercase string.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9])*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

Provides the customer-specified description for this DB parameter group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Family`

The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family.

To list all of the available parameter group families for a DB engine, use the following command:

`aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`

For example, to list all of the available parameter group families for the MySQL DB engine, use the following command:

`aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine mysql`

###### Note

The output contains duplicates.

The following are the valid DB engine values:

- `aurora-mysql`

- `aurora-postgresql`

- `db2-ae`

- `db2-se`

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

`Parameters`

A mapping of parameter names and values for the parameter update. You must specify at
least one parameter name and value.

For more information about parameter groups, see [Working with\
parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the _Amazon RDS User Guide_, or [Working\
with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the
_Amazon Aurora User Guide_.

###### Note

AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default
apply method for each parameter is used.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to assign to the DB parameter group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-dbparametergroup-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the DB parameter group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DBParameterGroupName`

The name of the DB parameter group.

## Examples

### Create a parameter group for a MySQL DB instance

The following example creates a parameter group for a MySQL DB instance and
sets the `sql_mode`, `max_allowed_packet`, and `innodb_buffer_pool_size` parameters.

#### JSON

```json

"RDSDBParameterGroup": {
        "Type": "AWS::RDS::DBParameterGroup",
        "Properties": {
            "Description": "CloudFormation Sample MySQL Parameter Group",
            "Family": "mysql8.0",
            "Parameters": {
                "sql_mode": "IGNORE_SPACE",
                "max_allowed_packet": 1024,
                "innodb_buffer_pool_size": "{DBInstanceClassMemory*3/4}"
            }
        }
    }
```

#### YAML

```yaml

RDSDBParameterGroup:
  Type: 'AWS::RDS::DBParameterGroup'
  Properties:
    Description: CloudFormation Sample MySQL Parameter Group
    Family: mysql8.0
    Parameters:
      sql_mode: IGNORE_SPACE
      max_allowed_packet: 1024
      innodb_buffer_pool_size: '{DBInstanceClassMemory*3/4}'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
