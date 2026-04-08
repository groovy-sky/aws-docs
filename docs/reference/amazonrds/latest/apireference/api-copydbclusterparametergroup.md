# CopyDBClusterParameterGroup

Copies the specified DB cluster parameter group.

###### Note

You can't copy a default DB cluster parameter group. Instead, create a new custom DB cluster parameter group, which copies
the default parameters and values for the specified DB cluster parameter group family.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceDBClusterParameterGroupIdentifier**

The identifier or Amazon Resource Name (ARN) for the source DB cluster parameter group.
For information about
creating an ARN,
see [Constructing an ARN for Amazon RDS](../../../../services/amazonrds/latest/aurorauserguide/user-tagging-arn.md#USER_Tagging.ARN.Constructing) in the _Amazon Aurora User Guide_.

Constraints:

- Must specify a valid DB cluster parameter group.

Type: String

Required: Yes

**TargetDBClusterParameterGroupDescription**

A description for the copied DB cluster parameter group.

Type: String

Required: Yes

**TargetDBClusterParameterGroupIdentifier**

The identifier for the copied DB cluster parameter group.

Constraints:

- Can't be null, empty, or blank

- Must contain from 1 to 255 letters, numbers, or hyphens

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

Example: `my-cluster-param-group1`

Type: String

Required: Yes

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**DBClusterParameterGroup**

Contains the details of an Amazon RDS DB cluster parameter group.

This data type is used as a response element in the `DescribeDBClusterParameterGroups` action.

Type: [DBClusterParameterGroup](api-dbclusterparametergroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBParameterGroupAlreadyExists**

A DB parameter group with the same name exists.

HTTP Status Code: 400

**DBParameterGroupNotFound**

`DBParameterGroupName` doesn't refer to an
existing DB parameter group.

HTTP Status Code: 404

**DBParameterGroupQuotaExceeded**

The request would result in the user exceeding the allowed number of DB parameter
groups.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CopyDBClusterParameterGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CopyDBClusterParameterGroup
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SourceDBClusterParameterGroupIdentifier=arn%3Aaws%3Ards%3Aus-east-1%3A815981987263%3cluster-pg%3Amy-cluster-pg
   &TargetDBParameterGroupIdentifier=new-cluster-pg
   &TargetDBParameterGroupDescription=New%20cluster%20group
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20160705/us-east-1/rds/aws4_request
   &X-Amz-Date=20160705T143101Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=9164337efa99caf850e874a1cb7ef62f3cea29d0b448b9e0e7c53b288ddffed2

```

#### Sample Response

```

<CopyDBClusterParameterGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateDBClusterParameterGroupResult>
    <DBClusterParameterGroup>
      <DBParameterGroupFamily>aurora5.6</DBParameterGroupFamily>
      <Description>New cluster group</Description>
      <DBClusterParameterGroupName>new-cluster-pg</DBClusterParameterGroupName>
    </DBClusterParameterGroup>
  </CreateDBClusterParameterGroupResult>
  <ResponseMetadata>
    <RequestId>ae81a963-cd9d-11e4-8b88-8351746a4c92</RequestId>
  </ResponseMetadata>
</CopyDBClusterParameterGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/copydbclusterparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/copydbclusterparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CancelExportTask

CopyDBClusterSnapshot

All content copied from https://docs.aws.amazon.com/.
