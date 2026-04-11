This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipProtectedQueryOutputConfiguration

Contains configurations for protected query results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : ProtectedQueryS3OutputConfiguration
}

```

### YAML

```yaml

  S3:
    ProtectedQueryS3OutputConfiguration

```

## Properties

`S3`

Required configuration for a protected query with an `s3` output type.

_Required_: Yes

_Type_: [ProtectedQueryS3OutputConfiguration](aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MembershipProtectedJobResultConfiguration

MembershipProtectedQueryResultConfiguration

All content copied from https://docs.aws.amazon.com/.
