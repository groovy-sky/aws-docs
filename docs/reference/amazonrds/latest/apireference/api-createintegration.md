# CreateIntegration

Creates a zero-ETL integration with Amazon Redshift.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**IntegrationName**

The name of the integration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**SourceArn**

The Amazon Resource Name (ARN) of the database to use as the source for
replication.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `arn:aws[a-z\-]*:rds(-[a-z]*)?:[a-z0-9\-]*:[0-9]*:(cluster|db):[a-z][a-z0-9]*(-[a-z0-9]+)*`

Required: Yes

**TargetArn**

The ARN of the Redshift data warehouse to use as the target for replication.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Required: Yes

**AdditionalEncryptionContext**AdditionalEncryptionContext.entry.N.key (key)AdditionalEncryptionContext.entry.N.value (value)

An optional set of non-secret key–value pairs that contains additional contextual
information about the data. For more information, see [Encryption\
context](../../../../services/kms/latest/developerguide/concepts.md#encrypt_context) in the _AWS Key Management Service Developer_
_Guide_.

You can only include this parameter if you specify the `KMSKeyId` parameter.

Type: String to string map

Required: No

**DataFilter**

Data filtering options for the integration. For more information, see
[Data filtering for Aurora zero-ETL integrations with Amazon Redshift](../../../../services/amazonrds/latest/aurorauserguide/zero-etl-filtering.md)
or
[Data filtering for Amazon RDS zero-ETL integrations with Amazon Redshift](../../../../services/amazonrds/latest/userguide/zero-etl-filtering.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 25600.

Pattern: `[a-zA-Z0-9_ "\\\-$,*.:?+\/]*`

Required: No

**Description**

A description of the integration.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `.*`

Required: No

**KMSKeyId**

The AWS Key Management System (AWS KMS) key identifier for the key to use to
encrypt the integration. If you don't specify an encryption key, RDS uses a default
AWS owned key.

Type: String

Required: No

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**AdditionalEncryptionContext**AdditionalEncryptionContext.entry.N.key (key)AdditionalEncryptionContext.entry.N.value (value)

The encryption context for the integration. For more information, see [Encryption context](../../../../services/kms/latest/developerguide/concepts.md#encrypt_context) in the _AWS Key Management Service Developer_
_Guide_.

Type: String to string map

**CreateTime**

The time when the integration was created, in Universal Coordinated Time
(UTC).

Type: Timestamp

**DataFilter**

Data filters for the integration. These filters determine which tables
from the source database are sent to the target Amazon Redshift data warehouse.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 25600.

Pattern: `[a-zA-Z0-9_ "\\\-$,*.:?+\/]*`

**Description**

A description of the integration.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `.*`

**Errors.IntegrationError.N**

Any errors associated with the integration.

Type: Array of [IntegrationError](api-integrationerror.md) objects

**IntegrationArn**

The ARN of the integration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `arn:aws[a-z\-]*:rds(-[a-z]*)?:[a-z0-9\-]*:[0-9]*:integration:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

**IntegrationName**

The name of the integration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

**KMSKeyId**

The AWS Key Management System (AWS KMS) key identifier for the key used to to
encrypt the integration.

Type: String

**SourceArn**

The Amazon Resource Name (ARN) of the database used as the source for
replication.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `arn:aws[a-z\-]*:rds(-[a-z]*)?:[a-z0-9\-]*:[0-9]*:(cluster|db):[a-z][a-z0-9]*(-[a-z0-9]+)*`

**Status**

The current status of the integration.

Type: String

Valid Values: `creating | active | modifying | failed | deleting | syncing | needs_attention`

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

**TargetArn**

The ARN of the Redshift data warehouse used as the target for replication.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**IntegrationAlreadyExistsFault**

The integration you are trying to create already exists.

HTTP Status Code: 400

**IntegrationConflictOperationFault**

A conflicting conditional operation is currently in progress against this resource.
Typically occurs when there are multiple requests being made to the same resource at the same time,
and these requests conflict with each other.

HTTP Status Code: 400

**IntegrationQuotaExceededFault**

You can't crate any more zero-ETL integrations because the quota has been reached.

HTTP Status Code: 400

**KMSKeyNotAccessibleFault**

An error occurred accessing an AWS KMS key.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CreateIntegration.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CreateIntegration
   &IntegrationName=my-integration
   &SourceArn=arn%3Aaws%3Ards%3Aus-east-1%3A123456789012%3Adb%3Asource-db
   &TargetArn=arn%3Aaws%3Aredshift-serverless%3Aus-east-1%3A123456789012%3Anamespace%3A0844171c-1e01-4d9f-be52-89e6c44083e5
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20141031/us-east-1/rds/aws4_request
   &X-Amz-Date=20230110T005253Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=8a684aebe6d5219bb3572316a341963324d6ef339bd0dcfa5854f1a01d401214
```

#### Sample Response

```

<CreateIntegrationResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
    <CreateIntegrationResult>
        <SourceArn>arn:aws:rds:us-east-1:123456789012:cluster:source-db</SourceArn>
        <IntegrationName>my-integration</IntegrationName>
        <IntegrationArn>arn:aws:rds:us-east-1:123456789012:integration:f30acbd8-aaab-4c3c-afb5-09d51d041037</IntegrationArn>
        <TargetArn>arn:aws:redshift-serverless:us-east-1:123456789012:namespace/0844171c-1e01-4d9f-be52-89e6c44083e5</TargetArn>
        <Tags/>
        <CreateTime>2023-12-14T00:15:21.358Z</CreateTime>
        <KMSKeyId>arn:aws:kms:us-east-1:211223847500:key/eda7134d-cd39-4af1-b62b-ad2415b6bccc</KMSKeyId>
        <Status>creating</Status>
    </CreateIntegrationResult>
    <ResponseMetadata>
        <RequestId>f5a16865-4415-4054-890c-2f5b2c3c67a8</RequestId>
    </ResponseMetadata>
</CreateIntegrationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createintegration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createintegration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createintegration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createintegration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createintegration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createintegration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createintegration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createintegration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createintegration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createintegration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateGlobalCluster

CreateOptionGroup

All content copied from https://docs.aws.amazon.com/.
