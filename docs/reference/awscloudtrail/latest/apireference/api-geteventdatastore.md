# GetEventDataStore

###### Important

CloudTrail Lake will no longer be open to new customers starting May 31, 2026. If you would like to use CloudTrail Lake, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see [CloudTrail Lake availability change](../../../../services/awscloudtrail/latest/userguide/cloudtrail-lake-service-availability-change.md).

Returns information about an event data store specified as either an ARN or the ID
portion of the ARN.

## Request Syntax

```nohighlight

{
   "EventDataStore": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[EventDataStore](#API_GetEventDataStore_RequestSyntax)**

The ARN (or ID suffix of the ARN) of the event data store about which you want
information.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

Required: Yes

## Response Syntax

```nohighlight

{
   "AdvancedEventSelectors": [
      {
         "FieldSelectors": [
            {
               "EndsWith": [ "string" ],
               "Equals": [ "string" ],
               "Field": "string",
               "NotEndsWith": [ "string" ],
               "NotEquals": [ "string" ],
               "NotStartsWith": [ "string" ],
               "StartsWith": [ "string" ]
            }
         ],
         "Name": "string"
      }
   ],
   "BillingMode": "string",
   "CreatedTimestamp": number,
   "EventDataStoreArn": "string",
   "FederationRoleArn": "string",
   "FederationStatus": "string",
   "KmsKeyId": "string",
   "MultiRegionEnabled": boolean,
   "Name": "string",
   "OrganizationEnabled": boolean,
   "PartitionKeys": [
      {
         "Name": "string",
         "Type": "string"
      }
   ],
   "RetentionPeriod": number,
   "Status": "string",
   "TerminationProtectionEnabled": boolean,
   "UpdatedTimestamp": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AdvancedEventSelectors](#API_GetEventDataStore_ResponseSyntax)**

The advanced event selectors used to select events for the data store.

Type: Array of [AdvancedEventSelector](api-advancedeventselector.md) objects

**[BillingMode](#API_GetEventDataStore_ResponseSyntax)**

The billing mode for the event data store.

Type: String

Valid Values: `EXTENDABLE_RETENTION_PRICING | FIXED_RETENTION_PRICING`

**[CreatedTimestamp](#API_GetEventDataStore_ResponseSyntax)**

The timestamp of the event data store's creation.

Type: Timestamp

**[EventDataStoreArn](#API_GetEventDataStore_ResponseSyntax)**

The event data store Amazon Resource Number (ARN).

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

**[FederationRoleArn](#API_GetEventDataStore_ResponseSyntax)**

If Lake query federation is enabled, provides the ARN of the federation role used to access the resources for the federated event data store.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 125.

Pattern: `^[a-zA-Z0-9._/\-:@=\+,\.]+$`

**[FederationStatus](#API_GetEventDataStore_ResponseSyntax)**

Indicates the [Lake query federation](../../../../services/awscloudtrail/latest/userguide/query-federation.md) status. The status is
`ENABLED` if Lake query federation is enabled, or `DISABLED` if Lake query federation is disabled. You cannot delete an event data store if the `FederationStatus` is `ENABLED`.

Type: String

Valid Values: `ENABLING | ENABLED | DISABLING | DISABLED`

**[KmsKeyId](#API_GetEventDataStore_ResponseSyntax)**

Specifies the AWS KMS key ID that encrypts the events delivered by CloudTrail. The value is a fully specified ARN to a AWS KMS key in the
following format.

`arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 350.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

**[MultiRegionEnabled](#API_GetEventDataStore_ResponseSyntax)**

Indicates whether the event data store includes events from all Regions, or only from
the Region in which it was created.

Type: Boolean

**[Name](#API_GetEventDataStore_ResponseSyntax)**

The name of the event data store.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `^[a-zA-Z0-9._\-]+$`

**[OrganizationEnabled](#API_GetEventDataStore_ResponseSyntax)**

Indicates whether an event data store is collecting logged events for an organization in
AWS Organizations.

Type: Boolean

**[PartitionKeys](#API_GetEventDataStore_ResponseSyntax)**

The partition keys for the event data store. To improve query performance and efficiency, CloudTrail Lake organizes
event data into partitions based on values derived from partition keys.

Type: Array of [PartitionKey](api-partitionkey.md) objects

Array Members: Maximum number of 2 items.

**[RetentionPeriod](#API_GetEventDataStore_ResponseSyntax)**

The retention period of the event data store, in days.

Type: Integer

Valid Range: Minimum value of 7. Maximum value of 3653.

**[Status](#API_GetEventDataStore_ResponseSyntax)**

The status of an event data store.

Type: String

Valid Values: `CREATED | ENABLED | PENDING_DELETION | STARTING_INGESTION | STOPPING_INGESTION | STOPPED_INGESTION`

**[TerminationProtectionEnabled](#API_GetEventDataStore_ResponseSyntax)**

Indicates that termination protection is enabled.

Type: Boolean

**[UpdatedTimestamp](#API_GetEventDataStore_ResponseSyntax)**

Shows the time that an event data store was updated, if applicable.
`UpdatedTimestamp` is always either the same or newer than the time shown in
`CreatedTimestamp`.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**EventDataStoreARNInvalidException**

The specified event data store ARN is not valid or does not map to an event data store
in your account.

HTTP Status Code: 400

**EventDataStoreNotFoundException**

The specified event data store was not found.

HTTP Status Code: 400

**InvalidParameterException**

The request includes a parameter that is not valid.

HTTP Status Code: 400

**NoManagementAccountSLRExistsException**

This exception is thrown when the management account does not have a service-linked
role.

HTTP Status Code: 400

**OperationNotPermittedException**

This exception is thrown when the requested operation is not permitted.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/geteventdatastore.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/geteventdatastore.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetEventConfiguration

GetEventSelectors

All content copied from https://docs.aws.amazon.com/.
