This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipProtectedJobOutputConfiguration

Contains configurations for protected job results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : ProtectedJobS3OutputConfigurationInput
}

```

### YAML

```yaml

  S3:
    ProtectedJobS3OutputConfigurationInput

```

## Properties

`S3`

Contains the configuration to write the job results to S3.

_Required_: Yes

_Type_: [ProtectedJobS3OutputConfigurationInput](aws-properties-cleanrooms-membership-protectedjobs3outputconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MembershipPaymentConfiguration

MembershipProtectedJobResultConfiguration

All content copied from https://docs.aws.amazon.com/.
