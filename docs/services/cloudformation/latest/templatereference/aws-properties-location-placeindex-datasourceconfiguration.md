This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::PlaceIndex DataSourceConfiguration

Specifies the data storage option requesting Places.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IntendedUse" : String
}

```

### YAML

```yaml

  IntendedUse: String

```

## Properties

`IntendedUse`

Specifies how the results of an operation will be stored by the caller.

Valid values include:

- `SingleUse` specifies that the results won't be stored.

- `Storage` specifies that the result can be cached or stored in a
database.

Default value: `SingleUse`

_Required_: No

_Type_: String

_Allowed values_: `SingleUse | Storage`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Location::PlaceIndex

Tag

All content copied from https://docs.aws.amazon.com/.
