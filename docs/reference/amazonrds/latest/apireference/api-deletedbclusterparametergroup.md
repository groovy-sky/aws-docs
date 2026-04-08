# DeleteDBClusterParameterGroup

Deletes a specified DB cluster parameter group. The DB cluster parameter group to be deleted can't be associated with any DB clusters.

For more information on Amazon Aurora, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterParameterGroupName**

The name of the DB cluster parameter group.

Constraints:

- Must be the name of an existing DB cluster parameter group.

- You can't delete a default DB cluster parameter group.

- Can't be associated with any DB clusters.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBParameterGroupNotFound**

`DBParameterGroupName` doesn't refer to an
existing DB parameter group.

HTTP Status Code: 404

**InvalidDBParameterGroupState**

The DB parameter group is in use or is in an invalid state. If you are attempting
to delete the parameter group, you can't delete it when the parameter group is in
this state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of DeleteDBClusterParameterGroup.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=DeleteDBClusterParameterGroup
    &DBClusterParameterGroupName=sample-cluster-pg
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20160913/us-west-2/rds/aws4_request
    &X-Amz-Date=20160913T172430Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=3f54b5ee720c2644296e98a1c0393a9abd91bc0847dfe7dd9be02ede8fd95ae5
```

#### Sample Response

```

<DeleteDBClusterParameterGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>ee0201e1-79d6-11e6-9b94-838991bd60c6</RequestId>
  </ResponseMetadata>
</DeleteDBClusterParameterGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deletedbclusterparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deletedbclusterparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBClusterEndpoint

DeleteDBClusterSnapshot

All content copied from https://docs.aws.amazon.com/.
