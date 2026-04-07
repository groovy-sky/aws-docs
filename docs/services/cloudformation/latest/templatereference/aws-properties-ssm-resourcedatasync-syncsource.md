This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::ResourceDataSync SyncSource

Information about the source of the data included in the resource data sync.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsOrganizationsSource" : AwsOrganizationsSource,
  "IncludeFutureRegions" : Boolean,
  "SourceRegions" : [ String, ... ],
  "SourceType" : String
}

```

### YAML

```yaml

  AwsOrganizationsSource:
    AwsOrganizationsSource
  IncludeFutureRegions: Boolean
  SourceRegions:
    - String
  SourceType: String

```

## Properties

`AwsOrganizationsSource`

Information about the AwsOrganizationsSource resource data sync source. A sync source
of this type can synchronize data from AWS Organizations.

_Required_: No

_Type_: [AwsOrganizationsSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-resourcedatasync-awsorganizationssource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeFutureRegions`

Whether to automatically synchronize and aggregate data from new AWS Regions when those
Regions come online.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceRegions`

The `SyncSource` AWS Regions included in the resource data sync.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceType`

The type of data source for the resource data sync. `SourceType` is either
`AwsOrganizations` (if an organization is present in AWS Organizations) or
`SingleAccountMultiRegions`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Destination

AWS::SSM::ResourcePolicy
