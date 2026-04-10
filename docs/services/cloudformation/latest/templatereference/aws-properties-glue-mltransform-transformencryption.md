This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::MLTransform TransformEncryption

The encryption-at-rest settings of the transform that apply to accessing user data. Machine learning
transforms can access user data encrypted in Amazon S3 using KMS.

Additionally, imported labels and trained transforms can now be encrypted using a customer provided
KMS key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MLUserDataEncryption" : MLUserDataEncryption,
  "TaskRunSecurityConfigurationName" : String
}

```

### YAML

```yaml

  MLUserDataEncryption:
    MLUserDataEncryption
  TaskRunSecurityConfigurationName: String

```

## Properties

`MLUserDataEncryption`

The encryption-at-rest settings of the transform that apply to accessing user data.

_Required_: No

_Type_: [MLUserDataEncryption](aws-properties-glue-mltransform-mluserdataencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskRunSecurityConfigurationName`

The name of the security configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MLUserDataEncryption

TransformParameters

All content copied from https://docs.aws.amazon.com/.
