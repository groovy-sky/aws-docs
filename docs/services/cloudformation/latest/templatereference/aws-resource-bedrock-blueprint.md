This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Blueprint

Details about a data automation blueprint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::Blueprint",
  "Properties" : {
      "BlueprintName" : String,
      "KmsEncryptionContext" : {Key: Value, ...},
      "KmsKeyId" : String,
      "Schema" : Json,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::Blueprint
Properties:
  BlueprintName: String
  KmsEncryptionContext:
    Key: Value
  KmsKeyId: String
  Schema: Json
  Tags:
    - Tag
  Type: String

```

## Properties

`BlueprintName`

The blueprint's name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsEncryptionContext`

Name-value pairs to include as an encryption context.

_Required_: No

_Type_: Object of String

_Pattern_: `^.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The AWS KMS key to use for encryption.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

The blueprint's schema.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-blueprint-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The blueprint's type.

_Required_: Yes

_Type_: String

_Allowed values_: `DOCUMENT | IMAGE | AUDIO | VIDEO`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`BlueprintArn`

The blueprint's ARN.

`BlueprintStage`

The blueprint's stage.

`CreationTime`

When the blueprint was created.

`LastModifiedTime`

When the blueprint was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
