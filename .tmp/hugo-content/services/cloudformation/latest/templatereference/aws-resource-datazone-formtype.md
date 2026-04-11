This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::FormType

The details of the metadata form type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::FormType",
  "Properties" : {
      "Description" : String,
      "DomainIdentifier" : String,
      "Model" : Model,
      "Name" : String,
      "OwningProjectIdentifier" : String,
      "Status" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::FormType
Properties:
  Description: String
  DomainIdentifier: String
  Model:
    Model
  Name: String
  OwningProjectIdentifier: String
  Status: String

```

## Properties

`Description`

The description of the metadata form type.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DomainIdentifier`

The identifier of the Amazon DataZone domain in which the form type exists.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Model`

The model of the form type.

_Required_: Yes

_Type_: [Model](aws-properties-datazone-formtype-model.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Name`

The name of the form type.

_Required_: Yes

_Type_: String

_Pattern_: `^(?![0-9_])\w+$|^_\w*[a-zA-Z0-9]\w*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OwningProjectIdentifier`

The identifier of the project that owns the form type.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Status`

The status of the form type.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainIdentifier` and
`FormTypeIdentifier`, which uniquely identifies a form type. For example: `{"Ref": "MyFormType"
    }` for the resource with the logical ID `MyFormType`, `Ref` returns
`DomainIdentifier|FormTypeIdentifier`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when the metadata form type was created.

`CreatedBy`

The Amazon DataZone user who created teh metadata form type.

`DomainId`

The identifier of the Amazon DataZone domain in which the form type exists.

`FormTypeIdentifier`

The ID of the metadata form type.

`OwningProjectId`

The identifier of the project that owns the form type.

`Revision`

The revision of the form type.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentParameter

Model

All content copied from https://docs.aws.amazon.com/.
