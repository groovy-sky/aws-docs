This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership ProtectedJobS3OutputConfigurationInput

Contains input information for protected jobs with an S3 output type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "KeyPrefix" : String
}

```

### YAML

```yaml

  Bucket: String
  KeyPrefix: String

```

## Properties

`Bucket`

The S3 bucket for job output.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPrefix`

The S3 prefix to unload the protected job results.

_Required_: No

_Type_: String

_Pattern_: `[\w!.=*/-]*`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MembershipSyntheticDataGenerationPaymentConfig

ProtectedQueryS3OutputConfiguration

All content copied from https://docs.aws.amazon.com/.
