# DescribeValidDBInstanceModifications

You can call `DescribeValidDBInstanceModifications` to learn what modifications you can make to
your DB instance. You can use this information when you call `ModifyDBInstance`.

This command doesn't apply to RDS Custom.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The customer identifier or the ARN of your DB instance.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**ValidDBInstanceModificationsMessage**

Information about valid modifications that you can make to your DB instance.
Contains the result of a successful call to the
`DescribeValidDBInstanceModifications` action.
You can use this information when you call
`ModifyDBInstance`.

Type: [ValidDBInstanceModificationsMessage](api-validdbinstancemodificationsmessage.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describevaliddbinstancemodifications.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describevaliddbinstancemodifications.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeTenantDatabases

DisableHttpEndpoint

All content copied from https://docs.aws.amazon.com/.
