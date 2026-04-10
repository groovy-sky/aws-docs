This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective DependencyConfig

Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DependencyKeyAttributes" : String,
  "DependencyOperationName" : String
}

```

### YAML

```yaml

  DependencyKeyAttributes: String
  DependencyOperationName: String

```

## Properties

`DependencyKeyAttributes`

If this SLO is related to a metric collected by Application Signals, you must use this field to specify which dependency the SLO metric is related to.

- `Type` designates the type of object this is.

- `ResourceType` specifies the type of the resource. This field is used only
when the value of the `Type` field is `Resource` or `AWS::Resource`.

- `Name` specifies the name of the object. This is used only if the value of the `Type` field
is `Service`, `RemoteService`, or `AWS::Service`.

- `Identifier` identifies the resource objects of this resource.
This is used only if the value of the `Type` field
is `Resource` or `AWS::Resource`.

- `Environment` specifies the location where this object is hosted, or what it belongs to.

_Required_: Yes

_Type_: String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DependencyOperationName`

When the SLO monitors a specific operation of the dependency, this field specifies the name of that operation in the dependency.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalendarInterval

Dimension

All content copied from https://docs.aws.amazon.com/.
