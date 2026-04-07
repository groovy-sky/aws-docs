This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBParameterGroup

`AWS::Neptune::DBParameterGroup` creates a new DB parameter group.
This type can be declared in a template and referenced in the `DBParameterGroupName`
parameter of `AWS::Neptune::DBInstance`.

###### Note

Applying a parameter group to a DB instance might require the
instance to reboot, resulting in a database outage for the duration of the reboot.

###### Note

If you provide a custom `DBParameterGroup` that you associate
with `DBInstance`, it is best to specify an `EngineVersion`
property in `DBCluster`. That `EngineVersion` needs to
be compatible with the value of the `Family` property in the
`DBParameterGroup`.

A DB parameter group is initially created with the default parameters for the database
engine used by the DB instance. To provide custom values for any of the parameters, you must
modify the group after creating it using _ModifyDBParameterGroup_. Once
you've created a DB parameter group, you need to associate it with your DB instance using
_ModifyDBInstance_. When you associate a new DB parameter group with a
running DB instance, you need to reboot the DB instance without failover for the new DB
parameter group and associated settings to take effect.

###### Important

After you create a DB parameter group, you should wait at least 5 minutes before
creating your first DB instance that uses that DB parameter group as the default parameter
group. This allows Amazon Neptune to fully complete the create action before the parameter
group is used as the default for a new DB instance. This is especially important for
parameters that are critical when creating the default database for a DB instance, such as
the character set for the default database defined by the
`character_set_database` parameter. You can use the _Parameter_
_Groups_ option of the Amazon Neptune console or the
_DescribeDBParameters_ command to verify that your DB parameter group has
been created or modified.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Neptune::DBParameterGroup",
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

Type: AWS::Neptune::DBParameterGroup
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

Provides the customer-specified description for this DB parameter group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Family`

Must be `neptune1` for engine versions prior to [1.2.0.0](https://docs.aws.amazon.com/neptune/latest/userguide/engine-releases-1.2.0.0.html), or
`neptune1.2` for engine version `1.2.0.0` and higher.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Provides the name of the DB parameter group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

The parameters to set for this DB parameter group.

The parameters are expressed as a JSON object consisting of key-value pairs.

Changes to dynamic parameters are applied immediately. During an update, if
you have static parameters (whether they were changed or not), it triggers CloudFormation
to reboot the associated DB instance without failover.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that you want to attach to this parameter group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-neptune-dbparametergroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
