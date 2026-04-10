This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBClusterParameterGroup

The `AWS::RDS::DBClusterParameterGroup` resource creates a new Amazon RDS
DB cluster parameter group.

For information about configuring parameters for Amazon Aurora DB clusters, see
[Working with parameter \
groups](../../../amazonrds/latest/aurorauserguide/user-workingwithparamgroups.md) in the _Amazon Aurora User Guide_.

###### Note

If you apply a parameter group to a DB cluster, then its DB instances might need
to reboot. This can result in an outage while the DB instances are rebooting.

If you apply a change to parameter group associated with a stopped DB cluster,
then the updated stack waits until the DB cluster is started.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBClusterParameterGroup",
  "Properties" : {
      "DBClusterParameterGroupName" : String,
      "Description" : String,
      "Family" : String,
      "Parameters" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBClusterParameterGroup
Properties:
  DBClusterParameterGroupName: String
  Description: String
  Family: String
  Parameters: Json
  Tags:
    - Tag

```

## Properties

`DBClusterParameterGroupName`

The name of the DB cluster parameter group.

Constraints:

- Must not match the name of an existing DB cluster parameter group.

###### Note

This value is stored as a lowercase string.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9])*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description for the DB cluster parameter group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Family`

The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster
parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.

**Aurora MySQL**

Example: `aurora-mysql5.7`, `aurora-mysql8.0`

**Aurora PostgreSQL**

Example: `aurora-postgresql14`

**RDS for MySQL**

Example: `mysql8.0`

**RDS for PostgreSQL**

Example: `postgres13`

To list all of the available parameter group families for a DB engine, use the following command:

`aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`

For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:

`aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql`

###### Note

The output contains duplicates.

The following are the valid DB engine values:

- `aurora-mysql`

- `aurora-postgresql`

- `mysql`

- `postgres`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

Provides a list of parameters for the DB cluster parameter group.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to assign to the DB cluster parameter group.

_Required_: No

_Type_: Array of [Tag](aws-properties-rds-dbclusterparametergroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the DB cluster parameter group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a parameter group for an Aurora MySQL DB cluster

The following example creates a DB cluster parameter group for an Aurora MySQL DB
cluster and sets the `time_zone` and `character_set_database`
parameters:

#### JSON

```json

"RDSDBClusterParameterGroup": {
        "Type": "AWS::RDS::DBClusterParameterGroup",
        "Properties": {
            "Description": "CloudFormation Sample Aurora Cluster Parameter Group",
            "Family": "aurora-mysql8.0",
            "Parameters": {
                "time_zone": "US/Eastern",
                "character_set_database": "utf32"
            }
        }
    }
```

#### YAML

```yaml

RDSDBClusterParameterGroup:
  Type: 'AWS::RDS::DBClusterParameterGroup'
  Properties:
    Description: CloudFormation Sample Aurora Cluster Parameter Group
    Family: aurora5.6
    Parameters:
      time_zone: US/Eastern
      character_set_database: utf32
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
