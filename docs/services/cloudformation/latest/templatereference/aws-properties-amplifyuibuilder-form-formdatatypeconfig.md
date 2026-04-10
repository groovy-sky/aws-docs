This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormDataTypeConfig

The `FormDataTypeConfig` property specifies the data type configuration for the data source associated with a form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceType" : String,
  "DataTypeName" : String
}

```

### YAML

```yaml

  DataSourceType: String
  DataTypeName: String

```

## Properties

`DataSourceType`

The data source type, either an Amplify DataStore model or a custom data type.

_Required_: Yes

_Type_: String

_Allowed values_: `DataStore | Custom`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTypeName`

The unique name of the data type you are using as the data source for the form.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormCTA

FormInputBindingPropertiesValue

All content copied from https://docs.aws.amazon.com/.
