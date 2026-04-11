This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration AmiDistributionConfiguration

Define and configure the output AMIs of the pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmiTags" : {Key: Value, ...},
  "Description" : String,
  "KmsKeyId" : String,
  "LaunchPermissionConfiguration" : LaunchPermissionConfiguration,
  "Name" : String,
  "TargetAccountIds" : [ String, ... ]
}

```

### YAML

```yaml

  AmiTags:
    Key: Value
  Description: String
  KmsKeyId: String
  LaunchPermissionConfiguration:
    LaunchPermissionConfiguration
  Name: String
  TargetAccountIds:
    - String

```

## Properties

`AmiTags`

The tags to apply to AMIs distributed to this Region.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the AMI distribution configuration. Minimum and maximum length are
in characters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt the distributed image.
This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](../../../kms/latest/developerguide/concepts.md#key-id-key-ARN)
in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchPermissionConfiguration`

Launch permissions can be used to configure which AWS accounts can use the AMI to
launch instances.

_Required_: No

_Type_: [LaunchPermissionConfiguration](aws-properties-imagebuilder-distributionconfiguration-launchpermissionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the output AMI.

_Required_: No

_Type_: String

_Pattern_: `^[-_A-Za-z0-9{][-_A-Za-z0-9\s:{}\.]+[-_A-Za-z0-9}]$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetAccountIds`

The ID of an account to which you want to distribute an image.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ImageBuilder::DistributionConfiguration

ContainerDistributionConfiguration

All content copied from https://docs.aws.amazon.com/.
