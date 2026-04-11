# UpdateContinuousBackups

`UpdateContinuousBackups` enables or disables point in time recovery for
the specified table. A successful `UpdateContinuousBackups` call returns the
current `ContinuousBackupsDescription`. Continuous backups are
`ENABLED` on all tables at table creation. If point in time recovery is
enabled, `PointInTimeRecoveryStatus` will be set to ENABLED.

Once continuous backups and point in time recovery are enabled, you can restore to
any point in time within `EarliestRestorableDateTime` and
`LatestRestorableDateTime`.

`LatestRestorableDateTime` is typically 5 minutes before the current time.
You can restore your table to any point in time in the last 35 days. You can set the
`RecoveryPeriodInDays` to any value between 1 and 35 days.

## Request Syntax

```nohighlight

{
   "PointInTimeRecoverySpecification": {
      "PointInTimeRecoveryEnabled": boolean,
      "RecoveryPeriodInDays": number
   },
   "TableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[PointInTimeRecoverySpecification](#API_UpdateContinuousBackups_RequestSyntax)**

Represents the settings used to enable point in time recovery.

Type: [PointInTimeRecoverySpecification](api-pointintimerecoveryspecification.md) object

Required: Yes

**[TableName](#API_UpdateContinuousBackups_RequestSyntax)**

The name of the table. You can also provide the Amazon Resource Name (ARN) of the table in this
parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## Response Syntax

```nohighlight

{
   "ContinuousBackupsDescription": {
      "ContinuousBackupsStatus": "string",
      "PointInTimeRecoveryDescription": {
         "EarliestRestorableDateTime": number,
         "LatestRestorableDateTime": number,
         "PointInTimeRecoveryStatus": "string",
         "RecoveryPeriodInDays": number
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ContinuousBackupsDescription](#API_UpdateContinuousBackups_ResponseSyntax)**

Represents the continuous backups and point in time recovery settings on the
table.

Type: [ContinuousBackupsDescription](api-continuousbackupsdescription.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ContinuousBackupsUnavailableException**

Backups have not yet been enabled for this table.

HTTP Status Code: 400

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**TableNotFoundException**

A source table with the name `TableName` does not currently exist within
the subscriber's account or the subscriber is operating in the wrong AWS
Region.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/updatecontinuousbackups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/updatecontinuousbackups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

UpdateContributorInsights

All content copied from https://docs.aws.amazon.com/.
