# DescribeDBSnapshotAttributes

Returns a list of DB snapshot attribute names and values for a manual DB snapshot.

When sharing snapshots with other AWS accounts, `DescribeDBSnapshotAttributes`
returns the `restore` attribute and a list of IDs for the AWS accounts that are
authorized to copy or restore the manual DB snapshot. If `all` is included in the list of
values for the `restore` attribute, then the manual DB snapshot is public and
can be copied or restored by all AWS accounts.

To add or remove access for an AWS account to copy or restore a manual DB snapshot, or to make the
manual DB snapshot public or private, use the `ModifyDBSnapshotAttribute` API action.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSnapshotIdentifier**

The identifier for the DB snapshot to describe the attributes for.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**DBSnapshotAttributesResult**

Contains the results of a successful call to the `DescribeDBSnapshotAttributes`
API action.

Manual DB snapshot attributes are used to authorize other AWS accounts
to copy or restore a manual DB snapshot. For more information, see the `ModifyDBSnapshotAttribute`
API action.

Type: [DBSnapshotAttributesResult](api-dbsnapshotattributesresult.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBSnapshotNotFound**

`DBSnapshotIdentifier` doesn't refer to an existing DB snapshot.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeDBSnapshotAttributes.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=DescribeDBSnapshotAttributes
    &DBSnapshotIdentifier=manual-snapshot1
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20151027/us-east-1/rds/aws4_request
    &X-Amz-Date=20151027T210706Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=27413f450dfac3d68b2197453e52109bacd3863f9df1a02d6e40022165bb2e09

```

#### Sample Response

```

<DescribeDBSnapshotAttributesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBSnapshotAttributesResult>
    <DBSnapshotAttributesResult>
      <DBSnapshotAttributes>
        <DBSnapshotAttribute>
          <AttributeName>restore</AttributeName>
          <AttributeValues>
            <AttributeValue>012345678901</AttributeValue>
          </AttributeValues>
        </DBSnapshotAttribute>
      </DBSnapshotAttributes>
      <DBSnapshotIdentifier>manual-snapshot1</DBSnapshotIdentifier>
    </DBSnapshotAttributesResult>
  </DescribeDBSnapshotAttributesResult>
  <ResponseMetadata>
    <RequestId>ae5be4a2-7cee-11e5-a056-f1c189649a47</RequestId>
  </ResponseMetadata>
</DescribeDBSnapshotAttributesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbsnapshotattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbsnapshotattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBShardGroups

DescribeDBSnapshots

All content copied from https://docs.aws.amazon.com/.
