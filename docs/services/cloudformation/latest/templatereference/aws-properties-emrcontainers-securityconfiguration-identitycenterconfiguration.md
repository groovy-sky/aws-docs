This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration IdentityCenterConfiguration

The `IdentityCenterConfiguration` property type specifies Property description not available. for an [AWS::EMRContainers::SecurityConfiguration](aws-resource-emrcontainers-securityconfiguration.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableIdentityCenter" : Boolean,
  "IdentityCenterApplicationAssignmentRequired" : Boolean,
  "IdentityCenterInstanceARN" : String
}

```

### YAML

```yaml

  EnableIdentityCenter: Boolean
  IdentityCenterApplicationAssignmentRequired: Boolean
  IdentityCenterInstanceARN: String

```

## Properties

`EnableIdentityCenter`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityCenterApplicationAssignmentRequired`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityCenterInstanceARN`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IAMConfiguration

InTransitEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
