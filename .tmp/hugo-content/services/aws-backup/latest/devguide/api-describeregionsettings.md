# DescribeRegionSettings

Returns the current service opt-in settings for the Region. If service opt-in is enabled
for a service, AWS Backup tries to protect that service's resources in this Region,
when the resource is included in an on-demand backup or scheduled backup plan. Otherwise,
AWS Backup does not try to protect that service's resources in this
Region.

## Request Syntax

```

GET /account-settings HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "ResourceTypeManagementPreference": {
      "string" : boolean
   },
   "ResourceTypeOptInPreference": {
      "string" : boolean
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ResourceTypeManagementPreference](#API_DescribeRegionSettings_ResponseSyntax)**

Returns whether AWS Backup fully manages the backups for a resource type.

For the benefits of full AWS Backup management, see [Full AWS Backup \
management](whatisbackup.md#full-management).

For a list of resource types and whether each supports full AWS Backup management,
see the [Feature availability by resource](backup-feature-availability.md#features-by-resource) table.

If `"DynamoDB":false`, you can enable full AWS Backup management for
DynamoDB backup by enabling [AWS Backup's advanced DynamoDB backup features](advanced-ddb-backup.md#advanced-ddb-backup-enable-cli).

Type: String to boolean map

Key Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

**[ResourceTypeOptInPreference](#API_DescribeRegionSettings_ResponseSyntax)**

The services along with the opt-in preferences in the Region.

Type: String to boolean map

Key Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ServiceUnavailableException**

The request failed due to a temporary failure of the server.

**Context**

**Type**

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/describeregionsettings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/describeregionsettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeRecoveryPoint

DescribeReportJob

All content copied from https://docs.aws.amazon.com/.
