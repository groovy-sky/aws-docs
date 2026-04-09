# DescribeLimits

Returns the current provisioned-capacity quotas for your AWS account in
a Region, both for the Region as a whole and for any one DynamoDB table that you create
there.

When you establish an AWS account, the account has initial quotas on
the maximum read capacity units and write capacity units that you can provision across
all of your DynamoDB tables in a given Region. Also, there are per-table
quotas that apply when you create a table there. For more information, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) page in the _Amazon DynamoDB_
_Developer Guide_.

Although you can increase these quotas by filing a case at [AWS Support Center](https://console.aws.amazon.com/support/home), obtaining the
increase is not instantaneous. The `DescribeLimits` action lets you write
code to compare the capacity you are currently using to those quotas imposed by your
account so that you have enough time to apply for an increase before you hit a
quota.

For example, you could use one of the AWS SDKs to do the
following:

1. Call `DescribeLimits` for a particular Region to obtain your
    current account quotas on provisioned capacity there.

2. Create a variable to hold the aggregate read capacity units provisioned for
    all your tables in that Region, and one to hold the aggregate write capacity
    units. Zero them both.

3. Call `ListTables` to obtain a list of all your DynamoDB
    tables.

4. For each table name listed by `ListTables`, do the
    following:

- Call `DescribeTable` with the table name.

- Use the data returned by `DescribeTable` to add the read
capacity units and write capacity units provisioned for the table itself
to your variables.

- If the table has one or more global secondary indexes (GSIs), loop
over these GSIs and add their provisioned capacity values to your
variables as well.

5. Report the account quotas for that Region returned by
    `DescribeLimits`, along with the total current provisioned
    capacity levels you have calculated.

This will let you see whether you are getting close to your account-level
quotas.

The per-table quotas apply only when you are creating a new table. They restrict the
sum of the provisioned capacity of the new table itself and all its global secondary
indexes.

For existing tables and their GSIs, DynamoDB doesn't let you increase provisioned
capacity extremely rapidly, but the only quota that applies is that the aggregate
provisioned capacity over all your tables and GSIs cannot exceed either of the
per-account quotas.

###### Note

`DescribeLimits` should only be called periodically. You can expect
throttling errors if you call it more than once in a minute.

The `DescribeLimits` Request element has no content.

## Response Syntax

```nohighlight

{
   "AccountMaxReadCapacityUnits": number,
   "AccountMaxWriteCapacityUnits": number,
   "TableMaxReadCapacityUnits": number,
   "TableMaxWriteCapacityUnits": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AccountMaxReadCapacityUnits](#API_DescribeLimits_ResponseSyntax)**

The maximum total read capacity units that your account allows you to provision across
all of your tables in this Region.

Type: Long

Valid Range: Minimum value of 1.

**[AccountMaxWriteCapacityUnits](#API_DescribeLimits_ResponseSyntax)**

The maximum total write capacity units that your account allows you to provision
across all of your tables in this Region.

Type: Long

Valid Range: Minimum value of 1.

**[TableMaxReadCapacityUnits](#API_DescribeLimits_ResponseSyntax)**

The maximum read capacity units that your account allows you to provision for a new
table that you are creating in this Region, including the read capacity units
provisioned for its global secondary indexes (GSIs).

Type: Long

Valid Range: Minimum value of 1.

**[TableMaxWriteCapacityUnits](#API_DescribeLimits_ResponseSyntax)**

The maximum write capacity units that your account allows you to provision for a new
table that you are creating in this Region, including the write capacity units
provisioned for its global secondary indexes (GSIs).

Type: Long

Valid Range: Minimum value of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

## Examples

### DescribeLimits

This example illustrates one usage of DescribeLimits.

#### Sample Request

```

POST / HTTP/1.1
Host: dynamodb.<region>.<domain>;
Accept-Encoding: identity
Content-Length: <PayloadSizeBytes>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.0
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=<Headers>, Signature=<Signature>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDB_20120810.DescribeLimits
{ }
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
x-amz-crc32: <Checksum>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
    "AccountMaxReadCapacityUnits": 20000,
    "AccountMaxWriteCapacityUnits": 20000,
    "TableMaxReadCapacityUnits": 10000,
    "TableMaxWriteCapacityUnits": 10000
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/describelimits.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/describelimits.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeKinesisStreamingDestination

DescribeTable

All content copied from https://docs.aws.amazon.com/.
