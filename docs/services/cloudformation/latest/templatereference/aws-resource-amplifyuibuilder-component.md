This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component

The AWS::AmplifyUIBuilder::Component resource specifies a component within an Amplify app.
A component is a user interface (UI) element that you can customize. Use
`ComponentChild` to configure an instance of a `Component`. A
`ComponentChild` instance inherits the configuration of the main
`Component`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AmplifyUIBuilder::Component",
  "Properties" : {
      "AppId" : String,
      "BindingProperties" : {Key: Value, ...},
      "Children" : [ ComponentChild, ... ],
      "CollectionProperties" : {Key: Value, ...},
      "ComponentType" : String,
      "EnvironmentName" : String,
      "Events" : {Key: Value, ...},
      "Name" : String,
      "Overrides" : {Key: Value, ...},
      "Properties" : {Key: Value, ...},
      "SchemaVersion" : String,
      "SourceId" : String,
      "Tags" : {Key: Value, ...},
      "Variants" : [ ComponentVariant, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AmplifyUIBuilder::Component
Properties:
  AppId: String
  BindingProperties:
    Key: Value
  Children:
    - ComponentChild
  CollectionProperties:
    Key: Value
  ComponentType: String
  EnvironmentName: String
  Events:
    Key: Value
  Name: String
  Overrides:
    Key: Value
  Properties:
    Key: Value
  SchemaVersion: String
  SourceId: String
  Tags:
    Key: Value
  Variants:
    - ComponentVariant

```

## Properties

`AppId`

The unique ID of the Amplify app associated with the component.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BindingProperties`

The information to connect a component's properties to data at runtime. You can't specify
`tags` as a valid property for `bindingProperties`.

_Required_: No

_Type_: Object of [ComponentBindingPropertiesValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Children`

A list of the component's `ComponentChild` instances.

_Required_: No

_Type_: Array of [ComponentChild](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentchild.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollectionProperties`

The data binding configuration for the component's properties. Use this for a collection
component. You can't specify `tags` as a valid property for
`collectionProperties`.

_Required_: No

_Type_: Object of [ComponentDataConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentdataconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentType`

The type of the component. This can be an Amplify custom UI component or
another custom component.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentName`

The name of the backend environment that is a part of the Amplify
app.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Events`

Describes the events that can be raised on the component. Use for the workflow feature in
Amplify Studio that allows you to bind events and actions to
components.

_Required_: No

_Type_: Object of [ComponentEvent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentevent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the component.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Overrides`

Describes the component's properties that can be overriden in a customized instance of the
component. You can't specify `tags` as a valid property for
`overrides`.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

Describes the component's properties. You can't specify `tags` as a valid
property for `properties`.

_Required_: No

_Type_: Object of [ComponentProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaVersion`

The schema version of the component when it was imported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceId`

The unique ID of the component in its original source system, such as Figma.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more key-value pairs to use when tagging the component.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variants`

A list of the component's variants. A variant is a unique style configuration of a main
component.

_Required_: No

_Type_: Array of [ComponentVariant](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentvariant.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time that the component was created.

`Id`

The unique ID of the component.

`ModifiedAt`

The time that the component was modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Amplify UI Builder

ActionParameters
