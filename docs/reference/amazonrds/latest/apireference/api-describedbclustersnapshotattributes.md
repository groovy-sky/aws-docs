# DescribeDBClusterSnapshotAttributes

Returns a list of DB cluster snapshot attribute names and values for a manual DB cluster snapshot.

When sharing snapshots with other AWS accounts, `DescribeDBClusterSnapshotAttributes`
returns the `restore` attribute and a list of IDs for the AWS accounts that are
authorized to copy or restore the manual DB cluster snapshot. If `all` is included in the list of
values for the `restore` attribute, then the manual DB cluster snapshot is public and
can be copied or restored by all AWS accounts.

To add or remove access for an AWS account to copy or restore a manual DB cluster snapshot, or to make the
manual DB cluster snapshot public or private, use the `ModifyDBClusterSnapshotAttribute` API action.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterSnapshotIdentifier**

The identifier for the DB cluster snapshot to describe the attributes for.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**DBClusterSnapshotAttributesResult**

Contains the results of a successful call to the `DescribeDBClusterSnapshotAttributes`
API action.

Manual DB cluster snapshot attributes are used to authorize other AWS accounts
to copy or restore a manual DB cluster snapshot. For more information, see the `ModifyDBClusterSnapshotAttribute`
API action.

Type: [DBClusterSnapshotAttributesResult](api-dbclustersnapshotattributesresult.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterSnapshotNotFoundFault**

`DBClusterSnapshotIdentifier` doesn't refer to an existing DB cluster snapshot.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeDBClusterSnapshotAttributes.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=DescribeDBClusterSnapshotAttributes
    &DBClusterSnapshotIdentifier=manual-cluster-snapshot1
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20230227/us-east-1/rds/aws4_request
    &X-Amz-Date=20230227T210706Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=27413f450dfac3d68b2197453e52109bacd3863f9df1a02d6e40022165bb2e09

```

#### Sample Response

```

<DescribeDBClusterSnapshotAttributesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBClusterSnapshotAttributesResult>
    <DBClusterSnapshotAttributesResult>
      <DBClusterSnapshotAttributes>
        <DBClusterSnapshotAttribute>
          <AttributeName>restore</AttributeName>
          <AttributeValues>
            <AttributeValue>012345678901</AttributeValue>
          </AttributeValues>
        </DBClusterSnapshotAttribute>
      </DBClusterSnapshotAttributes>
      <DBSnapshotIdentifier>manual-cluster-snapshot1</DBSnapshotIdentifier>
    </DBClusterSnapshotAttributesResult>
  </DescribeDBClusterSnapshotAttributesResult>
  <ResponseMetadata>
    <RequestId>ae5be4a2-7cee-11e5-a056-f1c189649a47</RequestId>
  </ResponseMetadata>
</DescribeDBClusterSnapshotAttributesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbclustersnapshotattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbclustersnapshotattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBClusters

DescribeDBClusterSnapshots

All content copied from https://docs.aws.amazon.com/.
