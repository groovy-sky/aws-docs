This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup ResultConfiguration

The location in Amazon S3 where query and calculation results are stored and
the encryption option, if any, used for query and calculation results. These are known
as "client-side settings". If workgroup settings override client-side settings, then the
query uses the workgroup settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AclConfiguration" : AclConfiguration,
  "EncryptionConfiguration" : EncryptionConfiguration,
  "ExpectedBucketOwner" : String,
  "OutputLocation" : String
}

```

### YAML

```yaml

  AclConfiguration:
    AclConfiguration
  EncryptionConfiguration:
    EncryptionConfiguration
  ExpectedBucketOwner: String
  OutputLocation: String

```

## Properties

`AclConfiguration`

Indicates that an Amazon S3 canned ACL should be set to control ownership of stored
query results. Currently the only supported canned ACL is
`BUCKET_OWNER_FULL_CONTROL`. This is a client-side setting. If workgroup
settings override client-side settings, then the query uses the ACL configuration that
is specified for the workgroup, and also uses the location for storing query results
specified in the workgroup. See [EnforceWorkGroupConfiguration](../userguide/aws-properties-athena-workgroup-workgroupconfiguration.md#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration).

_Required_: No

_Type_: [AclConfiguration](aws-properties-athena-workgroup-aclconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

If query results are encrypted in Amazon S3, indicates the encryption option used (for
example, `SSE_KMS` or `CSE_KMS`) and key information. This is a
client-side setting. If workgroup settings override client-side settings, then the query
uses the encryption configuration that is specified for the workgroup, and also uses the
location for storing query results specified in the workgroup. See [EnforceWorkGroupConfiguration](../userguide/aws-properties-athena-workgroup-workgroupconfiguration.md#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration) and [Override client-side\
settings](../../../athena/latest/ug/workgroups-settings-override.md).

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-athena-workgroup-encryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpectedBucketOwner`

The account ID that you expect to be the owner of the Amazon S3 bucket specified by
`ResultConfiguration:OutputLocation`. If set, Athena uses the value for
`ExpectedBucketOwner` when it makes Amazon S3 calls to your specified
output location. If the `ExpectedBucketOwner` account ID does not match the
actual owner of the Amazon S3 bucket, the call fails with a permissions error.

This is a client-side setting. If workgroup settings override client-side settings,
then the query uses the `ExpectedBucketOwner` setting that is specified for
the workgroup, and also uses the location for storing query results specified in the
workgroup. See [EnforceWorkGroupConfiguration](../userguide/aws-properties-athena-workgroup-workgroupconfiguration.md#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration).

_Required_: No

_Type_: String

_Pattern_: `^[0-9]+$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputLocation`

The location in Amazon S3 where your query results are stored, such as
`s3://path/to/query/bucket/`. To run a query, you must specify the query
results location using either a client-side setting for individual queries or a location
specified by the workgroup. If workgroup settings override client-side settings, then
the query uses the location specified for the workgroup. If no query location is set,
Athena issues an error. For more information, see [Work with query results and recent\
queries](../../../athena/latest/ug/querying.md) and [EnforceWorkGroupConfiguration](../userguide/aws-properties-athena-workgroup-workgroupconfiguration.md#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringConfiguration

ResultConfigurationUpdates

All content copied from https://docs.aws.amazon.com/.
