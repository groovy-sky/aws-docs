# DeleteDBSubnetGroup

Deletes a DB subnet group.

###### Note

The specified database subnet group must not be associated with any DB instances.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSubnetGroupName**

The name of the database subnet group to delete.

###### Note

You can't delete the default subnet group.

Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.

Example: `mydbsubnetgroup`

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBSubnetGroupNotFoundFault**

`DBSubnetGroupName` doesn't refer to an existing DB subnet group.

HTTP Status Code: 404

**InvalidDBSubnetGroupStateFault**

The DB subnet group cannot be deleted because it's in use.

HTTP Status Code: 400

**InvalidDBSubnetStateFault**

The DB subnet isn't in the _available_ state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of DeleteDBSubnetGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DeleteDBSubnetGroup
   &DBSubnetGroupName=myawsuser-dbsubnetgroup
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140425/us-east-1/rds/aws4_request
   &X-Amz-Date=20140425T180721Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=0f461da21ec03527fdc98acba8a11c36863a399065f9b4ff891ab7cb5e70de74

```

#### Sample Response

```

<DeleteDBSubnetGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>6295e5ab-bbf3-11d3-f4c6-37db295f7674</RequestId>
  </ResponseMetadata>
</DeleteDBSubnetGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deletedbsubnetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deletedbsubnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBSnapshot

DeleteEventSubscription

All content copied from https://docs.aws.amazon.com/.
