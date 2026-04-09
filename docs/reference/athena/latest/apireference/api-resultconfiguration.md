# ResultConfiguration

The location in Amazon S3 where query and calculation results are stored and
the encryption option, if any, used for query and calculation results. These are known
as "client-side settings". If workgroup settings override client-side settings, then the
query uses the workgroup settings.

## Contents

**AclConfiguration**

Indicates that an Amazon S3 canned ACL should be set to control ownership of
stored query results. Currently the only supported canned ACL is
`BUCKET_OWNER_FULL_CONTROL`. This is a client-side setting. If workgroup
settings override client-side settings, then the query uses the ACL configuration that
is specified for the workgroup, and also uses the location for storing query results
specified in the workgroup. For more information, see [WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-EnforceWorkGroupConfiguration) and [Workgroup Settings Override Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: [AclConfiguration](api-aclconfiguration.md) object

Required: No

**EncryptionConfiguration**

If query and calculation results are encrypted in Amazon S3, indicates the
encryption option used (for example, `SSE_KMS` or `CSE_KMS`) and
key information. This is a client-side setting. If workgroup settings override
client-side settings, then the query uses the encryption configuration that is specified
for the workgroup, and also uses the location for storing query results specified in the
workgroup. See [WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-EnforceWorkGroupConfiguration)
and [Workgroup Settings Override Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) object

Required: No

**ExpectedBucketOwner**

The AWS account ID that you expect to be the owner of the Amazon S3 bucket specified by [ResultConfiguration:OutputLocation](#athena-Type-ResultConfiguration-OutputLocation).
If set, Athena uses the value for `ExpectedBucketOwner` when it
makes Amazon S3 calls to your specified output location. If the
`ExpectedBucketOwner`
AWS account ID does not match the actual owner of the Amazon S3
bucket, the call fails with a permissions error.

This is a client-side setting. If workgroup settings override client-side settings,
then the query uses the `ExpectedBucketOwner` setting that is specified for
the workgroup, and also uses the location for storing query results specified in the
workgroup. See [WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-EnforceWorkGroupConfiguration)
and [Workgroup Settings Override Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]+$`

Required: No

**OutputLocation**

The location in Amazon S3 where your query and calculation results are stored,
such as `s3://path/to/query/bucket/`. To run the query, you must specify the
query results location using one of the ways: either for individual queries using either
this setting (client-side), or in the workgroup, using [WorkGroupConfiguration](api-workgroupconfiguration.md). If none of them is set, Athena
issues an error that no output location is provided. If workgroup settings override
client-side settings, then the query uses the settings specified for the workgroup. See
[WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-EnforceWorkGroupConfiguration).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/resultconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/resultconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/resultconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryStagePlanNode

ResultConfigurationUpdates

All content copied from https://docs.aws.amazon.com/.
