# RestoreTestingSelectionForUpdate

This contains metadata about a restore testing selection.

## Contents

**IamRoleArn**

The Amazon Resource Name (ARN) of the IAM role that
AWS Backup uses to create the target resource; for example:
`arn:aws:iam::123456789012:role/S3Access`.

Type: String

Required: No

**ProtectedResourceArns**

You can include a list of specific ARNs, such as
`ProtectedResourceArns: ["arn:aws:...", "arn:aws:..."]`
or you can include a wildcard: `ProtectedResourceArns: ["*"]`,
but not both.

Type: Array of strings

Required: No

**ProtectedResourceConditions**

The conditions that you define for resources in
your restore testing plan using tags.

Type: [ProtectedResourceConditions](api-protectedresourceconditions.md) object

Required: No

**RestoreMetadataOverrides**

You can override certain restore metadata keys by including the parameter
`RestoreMetadataOverrides` in the body of
`RestoreTestingSelection`. Key values are not case sensitive.

See the complete list of [restore testing\
inferred metadata](restore-testing-inferred-metadata.md).

Type: String to string map

Required: No

**ValidationWindowHours**

This value represents the time, in hours, data is retained after
a restore test so that optional validation can be completed.

Accepted value is an integer between
0 and 168 (the hourly equivalent of seven days).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/restoretestingselectionforupdate.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/restoretestingselectionforupdate.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/restoretestingselectionforupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreTestingSelectionForList

ScanAction

All content copied from https://docs.aws.amazon.com/.
