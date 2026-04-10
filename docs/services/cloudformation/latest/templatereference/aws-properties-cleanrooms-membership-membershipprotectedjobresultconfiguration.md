This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipProtectedJobResultConfiguration

Contains configurations for protected job results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputConfiguration" : MembershipProtectedJobOutputConfiguration,
  "RoleArn" : String
}

```

### YAML

```yaml

  OutputConfiguration:
    MembershipProtectedJobOutputConfiguration
  RoleArn: String

```

## Properties

`OutputConfiguration`

The output configuration for a protected job result.

_Required_: Yes

_Type_: [MembershipProtectedJobOutputConfiguration](aws-properties-cleanrooms-membership-membershipprotectedjoboutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The unique ARN for an IAM role that is used by AWS Clean Rooms to write protected
job results to the result location, given by the member who can receive results.

_Required_: Yes

_Type_: String

_Minimum_: `32`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MembershipProtectedJobOutputConfiguration

MembershipProtectedQueryOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
