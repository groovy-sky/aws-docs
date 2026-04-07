This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject

A data automation project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::DataAutomationProject",
  "Properties" : {
      "CustomOutputConfiguration" : CustomOutputConfiguration,
      "KmsEncryptionContext" : {Key: Value, ...},
      "KmsKeyId" : String,
      "OverrideConfiguration" : OverrideConfiguration,
      "ProjectDescription" : String,
      "ProjectName" : String,
      "ProjectType" : String,
      "StandardOutputConfiguration" : StandardOutputConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::DataAutomationProject
Properties:
  CustomOutputConfiguration:
    CustomOutputConfiguration
  KmsEncryptionContext:
    Key: Value
  KmsKeyId: String
  OverrideConfiguration:
    OverrideConfiguration
  ProjectDescription: String
  ProjectName: String
  ProjectType: String
  StandardOutputConfiguration:
    StandardOutputConfiguration
  Tags:
    - Tag

```

## Properties

`CustomOutputConfiguration`

Blueprints to apply to objects processed by the project.

_Required_: No

_Type_: [CustomOutputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-dataautomationproject-customoutputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsEncryptionContext`

The AWS KMS encryption context to use for encryption.

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

`OverrideConfiguration`

Additional settings for the project.

_Required_: No

_Type_: [OverrideConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectDescription`

The project's description.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectName`

The project's name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectType`

The type of bedrock data automation API that is compatible with this project.

_Required_: No

_Type_: String

_Allowed values_: `ASYNC | SYNC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StandardOutputConfiguration`

The project's standard output configuration.

_Required_: No

_Type_: [StandardOutputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-dataautomationproject-standardoutputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-dataautomationproject-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

When the project was created.

`LastModifiedTime`

When the project was last updated.

`ProjectArn`

The project's ARN.

`ProjectStage`

The project's stage.

`Status`

The project's status.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AudioExtractionCategory
