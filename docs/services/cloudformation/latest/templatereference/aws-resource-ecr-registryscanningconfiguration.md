This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::RegistryScanningConfiguration

The scanning configuration for a private registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::RegistryScanningConfiguration",
  "Properties" : {
      "Rules" : [ ScanningRule, ... ],
      "ScanType" : String
    }
}

```

### YAML

```yaml

Type: AWS::ECR::RegistryScanningConfiguration
Properties:
  Rules:
    - ScanningRule
  ScanType: String

```

## Properties

`Rules`

The scanning rules associated with the registry.

_Required_: Yes

_Type_: Array of [ScanningRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecr-registryscanningconfiguration-scanningrule.html)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanType`

The type of scanning configured for the registry.

_Required_: Yes

_Type_: String

_Allowed values_: `BASIC | ENHANCED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RegistryId`

The account ID of the destination registry.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ECR::RegistryPolicy

RepositoryFilter
