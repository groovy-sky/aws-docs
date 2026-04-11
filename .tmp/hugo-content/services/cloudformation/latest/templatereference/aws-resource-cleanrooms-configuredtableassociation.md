This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTableAssociation

Creates a configured table association. A configured table association links a
configured table with a collaboration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::ConfiguredTableAssociation",
  "Properties" : {
      "ConfiguredTableAssociationAnalysisRules" : [ ConfiguredTableAssociationAnalysisRule, ... ],
      "ConfiguredTableIdentifier" : String,
      "Description" : String,
      "MembershipIdentifier" : String,
      "Name" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::ConfiguredTableAssociation
Properties:
  ConfiguredTableAssociationAnalysisRules:
    - ConfiguredTableAssociationAnalysisRule
  ConfiguredTableIdentifier: String
  Description: String
  MembershipIdentifier: String
  Name: String
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`ConfiguredTableAssociationAnalysisRules`

An analysis rule for a configured table association. This analysis rule specifies how
data from the table can be used within its associated collaboration. In the console, the
`ConfiguredTableAssociationAnalysisRule` is referred to as the
_collaboration analysis rule_.

_Required_: No

_Type_: Array of [ConfiguredTableAssociationAnalysisRule](aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrule.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfiguredTableIdentifier`

A unique identifier for the configured table to be associated to. Currently accepts a
configured table ID.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the configured table association.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MembershipIdentifier`

The unique ID for the membership this configured table association belongs to.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the configured table association, in lowercase. The table is identified by
this name when running protected queries against the underlying data.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The service will assume this role to access catalog metadata and query the table.

_Required_: Yes

_Type_: String

_Minimum_: `32`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag
consists of a key and an optional value, both of which you define. When you use tagging,
you can also use tag-based access control in IAM policies to control access
to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-cleanrooms-configuredtableassociation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ConfiguredTableAssociation` and the ID of the
Membership. For example:
`c1baf760-935e-4b2d-b36e-af8daaeb6e48|81a97460-2c40-46ce-a2fd-4ccda7398b2c`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the specified configured table
association.

Example:
`arn:aws:cleanrooms:us-east-1:111122223333:configuredtable/a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`

`ConfiguredTableAssociationIdentifier`

Returns the unique identifier of the specified configured table association.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`

## Examples

### A configured table association

The following is an example of a configured table association which associates a
configured table to a collaboration.

#### JSON

```json

"ExampleAssociation": {
  "Type" : "AWS::CleanRooms::ConfiguredTableAssociation",
  "Properties" : {
    "Name" : "Example Table",
    "Description" : "Example associated configured table",
    "ConfiguredTableIdentifier" : "a1b2c3d4-5678-90ab-cdef-EXAMPLE22222",
    "MembershipIdentifier" : "a1b2c3d4-5678-90ab-cdef-EXAMPLE33333",
    "RoleArn" : "arn:aws:iam::123456789012:role/service-role/cleanrooms-association-example"
  }
}
```

#### YAML

```yaml

ExampleAssociation:
  Type: AWS::CleanRooms::ConfiguredTableAssociation
  Properties:
    Name: Example Table
    Description: Example associated configured table
    ConfiguredTableIdentifier: a1b2c3d4-5678-90ab-cdef-EXAMPLE22222"
    MembershipIdentifier: a1b2c3d4-5678-90ab-cdef-EXAMPLE33333
    RoleArn: arn:aws:iam::123456789012:role/service-role/cleanrooms-association-example
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ConfiguredTableAssociationAnalysisRule

All content copied from https://docs.aws.amazon.com/.
