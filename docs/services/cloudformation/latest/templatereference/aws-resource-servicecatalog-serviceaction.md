This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::ServiceAction

Creates a self-service action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::ServiceAction",
  "Properties" : {
      "AcceptLanguage" : String,
      "Definition" : [ DefinitionParameter, ... ],
      "DefinitionType" : String,
      "Description" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::ServiceAction
Properties:
  AcceptLanguage: String
  Definition:
    - DefinitionParameter
  DefinitionType: String
  Description: String
  Name: String

```

## Properties

`AcceptLanguage`

The language code.

- `en` \- English (default)

- `jp` \- Japanese

- `zh` \- Chinese

_Required_: No

_Type_: String

_Allowed values_: `en | jp | zh`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Definition`

A map that defines the self-service action.

_Required_: Yes

_Type_: Array of [DefinitionParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicecatalog-serviceaction-definitionparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionType`

The self-service action definition type. For example, `SSM_AUTOMATION`.

_Required_: Yes

_Type_: String

_Allowed values_: `SSM_AUTOMATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The self-service action description.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The self-service action name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Id`

The self-service action identifier. For example, `act-fs7abcd89wxyz`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ServiceCatalog::ResourceUpdateConstraint

DefinitionParameter
