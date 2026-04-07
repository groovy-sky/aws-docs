This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBClusterParameterGroup

The `AWS::Neptune::DBClusterParameterGroup` resource creates a new Amazon Neptune DB cluster parameter group.

###### Note

Applying a parameter group to a DB cluster might require
instances to reboot, resulting in a database outage while the instances reboot.

###### Note

If you provide a custom `DBClusterParameterGroup` that you associate
with the `DBCluster`, it is best to specify an `EngineVersion`
property in the `DBCluster`. That `EngineVersion` needs to
be compatible with the value of the `Family` property in the
`DBClusterParameterGroup`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Neptune::DBClusterParameterGroup",
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

Type: AWS::Neptune::DBClusterParameterGroup
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

Provides the customer-specified description for this DB cluster parameter group.

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

Provides the name of the DB cluster parameter group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

The parameters to set for this DB cluster parameter group.

The parameters are expressed as a JSON object consisting of key-value pairs.

If you update the parameters, some interruption may occur depending on which
parameters you update.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that you want to attach to this parameter group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-neptune-dbclusterparametergroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
