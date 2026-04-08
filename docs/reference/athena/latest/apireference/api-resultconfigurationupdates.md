# ResultConfigurationUpdates

The information about the updates in the query results, such as output location and
encryption configuration for the query results.

## Contents

**AclConfiguration**

The ACL configuration for the query results.

Type: [AclConfiguration](api-aclconfiguration.md) object

Required: No

**EncryptionConfiguration**

The encryption configuration for query and calculation results.

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) object

Required: No

**ExpectedBucketOwner**

The AWS account ID that you expect to be the owner of the Amazon S3 bucket specified by [ResultConfiguration:OutputLocation](api-resultconfiguration.md#athena-Type-ResultConfiguration-OutputLocation).
If set, Athena uses the value for `ExpectedBucketOwner` when it
makes Amazon S3 calls to your specified output location. If the
`ExpectedBucketOwner`
AWS account ID does not match the actual owner of the Amazon S3
bucket, the call fails with a permissions error.

If workgroup settings override client-side settings, then the query uses the
`ExpectedBucketOwner` setting that is specified for the workgroup, and
also uses the location for storing query results specified in the workgroup. See [WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-EnforceWorkGroupConfiguration) and [Workgroup Settings Override Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]+$`

Required: No

**OutputLocation**

The location in Amazon S3 where your query and calculation results are stored,
such as `s3://path/to/query/bucket/`. If workgroup settings override
client-side settings, then the query uses the location for the query results and the
encryption configuration that are specified for the workgroup. The "workgroup settings
override" is specified in `EnforceWorkGroupConfiguration` (true/false) in the
`WorkGroupConfiguration`. See [WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration.md#athena-Type-WorkGroupConfiguration-EnforceWorkGroupConfiguration).

Type: String

Required: No

**RemoveAclConfiguration**

If set to `true`, indicates that the previously-specified ACL configuration
for queries in this workgroup should be ignored and set to null. If set to
`false` or not set, and a value is present in the
`AclConfiguration` of `ResultConfigurationUpdates`, the
`AclConfiguration` in the workgroup's `ResultConfiguration` is
updated with the new value. For more information, see [Workgroup Settings Override\
Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: Boolean

Required: No

**RemoveEncryptionConfiguration**

If set to "true", indicates that the previously-specified encryption configuration
(also known as the client-side setting) for queries in this workgroup should be ignored
and set to null. If set to "false" or not set, and a value is present in the
`EncryptionConfiguration` in `ResultConfigurationUpdates` (the
client-side setting), the `EncryptionConfiguration` in the workgroup's
`ResultConfiguration` will be updated with the new value. For more
information, see [Workgroup Settings Override\
Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: Boolean

Required: No

**RemoveExpectedBucketOwner**

If set to "true", removes the AWS account ID previously specified for
[ResultConfiguration:ExpectedBucketOwner](api-resultconfiguration.md#athena-Type-ResultConfiguration-ExpectedBucketOwner). If set to "false" or not
set, and a value is present in the `ExpectedBucketOwner` in
`ResultConfigurationUpdates` (the client-side setting), the
`ExpectedBucketOwner` in the workgroup's `ResultConfiguration`
is updated with the new value. For more information, see [Workgroup Settings Override\
Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: Boolean

Required: No

**RemoveOutputLocation**

If set to "true", indicates that the previously-specified query results location (also
known as a client-side setting) for queries in this workgroup should be ignored and set
to null. If set to "false" or not set, and a value is present in the
`OutputLocation` in `ResultConfigurationUpdates` (the
client-side setting), the `OutputLocation` in the workgroup's
`ResultConfiguration` will be updated with the new value. For more
information, see [Workgroup Settings Override\
Client-Side Settings](../../../../services/athena/latest/ug/workgroups-settings-override.md).

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/resultconfigurationupdates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/resultconfigurationupdates.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/resultconfigurationupdates.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResultConfiguration

ResultReuseByAgeConfiguration
