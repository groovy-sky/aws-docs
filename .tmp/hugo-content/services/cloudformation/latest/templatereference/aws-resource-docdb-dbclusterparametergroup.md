This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDB::DBClusterParameterGroup

The `AWS::DocDB::DBClusterParameterGroup` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBClusterParameterGroup.
For more information, see [DBClusterParameterGroup](../../../documentdb/latest/developerguide/api-dbclusterparametergroup.md)
in the _Amazon DocumentDB Developer Guide_.

Parameters in a cluster parameter group apply to all of the instances in a cluster.

A cluster parameter group is initially created with the default parameters for the
database engine used by instances in the cluster. To provide custom values for any of
the parameters, you must modify the group after you create it. After you create a DB
cluster parameter group, you must associate it with your cluster. For the new cluster
parameter group and associated settings to take effect, you must then reboot the DB
instances in the cluster without failover.

###### Important

After you create a cluster parameter group, you should wait at least 5 minutes
before creating your first cluster that uses that cluster parameter group as the
default parameter group. This allows Amazon DocumentDB to fully complete the create action
before the cluster parameter group is used as the default for a new cluster. This
step is especially important for parameters that are critical when creating the default
database for a cluster, such as the character set for the default database defined by
the `character_set_database` parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DocDB::DBClusterParameterGroup",
  "Properties" : {
      "Description" : String,
      "Family" : String,
      "Name" : String,
      "Parameters" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DocDB::DBClusterParameterGroup
Properties:
  Description: String
  Family: String
  Name: String
  Parameters: Json
  Tags:
    - Tag

```

## Properties

`Description`

The description for the cluster parameter group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Family`

The cluster parameter group family name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the DB cluster parameter group.

Constraints:

- Must not match the name of an existing
`DBClusterParameterGroup`.

###### Note

This value is stored as a lowercase string.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

Provides a list of parameters for the cluster parameter group.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be assigned to the cluster parameter group.

_Required_: No

_Type_: Array of [Tag](aws-properties-docdb-dbclusterparametergroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DBClusterParameterGroup's name, such as `sample-db-cluster-param-group`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

#### JSON

```json

{
   "Type" : "AWS::DocDB::DBClusterParameterGroup",
   "Properties" : {
      "Description" : "description",
      "Family" : "docdb3.6",
      "Name" : "sampleParameterGroup",
      "Parameters" :  {
        "audit_logs": "disabled",
        "tls": "enabled",
        "ttl_monitor": "enabled"
     },
      "Tags" : [{ "Key": "String","Value": "String" }]
   }
}
```

#### YAML

```yaml

Type: "AWS::DocDB::DBClusterParameterGroup"
Properties:
   Description: "description"
   Family: "docdb3.6"
   Name: "sampleParameterGroup"
   Parameters:
        audit_logs: "disabled"
        tls: "enabled"
        ttl_monitor: "enabled"
   Tags:
      -
         Key: "String"
         Value: "String"
```

## See also

- [DBClusterParameterGroup](../../../documentdb/latest/developerguide/api-dbclusterparametergroup.md)

- [CreateDBClusterParameterGroup](../../../documentdb/latest/developerguide/api-createdbclusterparametergroup.md)

- [DeleteDBClusterParameterGroup](../../../documentdb/latest/developerguide/api-deletedbclusterparametergroup.md)

- [DescribeDBClusterParameterGroups](../../../documentdb/latest/developerguide/api-describedbclusterparametergroups.md)

- [ModifyDBClusterParameterGroup](../../../documentdb/latest/developerguide/api-modifydbclusterparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
