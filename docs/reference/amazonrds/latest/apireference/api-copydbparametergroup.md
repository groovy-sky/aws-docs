# CopyDBParameterGroup

Copies the specified DB parameter group.

###### Note

You can't copy a default DB parameter group. Instead, create a new custom DB parameter group, which copies the default
parameters and values for the specified DB parameter group family.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceDBParameterGroupIdentifier**

The identifier or ARN for the source DB parameter group.
For information about
creating an ARN,
see [Constructing an ARN for Amazon RDS](../../../../services/amazonrds/latest/userguide/user-tagging-arn.md#USER_Tagging.ARN.Constructing) in the _Amazon RDS User Guide_.

Constraints:

- Must specify a valid DB parameter group.

Type: String

Required: Yes

**TargetDBParameterGroupDescription**

A description for the copied DB parameter group.

Type: String

Required: Yes

**TargetDBParameterGroupIdentifier**

The identifier for the copied DB parameter group.

Constraints:

- Can't be null, empty, or blank

- Must contain from 1 to 255 letters, numbers, or hyphens

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

Example: `my-db-parameter-group`

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

**DBParameterGroup**

Contains the details of an Amazon RDS DB parameter group.

This data type is used as a response element in the `DescribeDBParameterGroups` action.

Type: [DBParameterGroup](api-dbparametergroup.md) object

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

This example illustrates one usage of CopyDBParameterGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CopyDBParameterGroup
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SourceDBParameterGroupIdentifier=arn%3Aaws%3Ards%3Aus-west-2%3A815981987263%3pg%3Amy-remote-param-group
   &TargetDBParameterGroupIdentifier=new-local-param-group
   &TargetDBParameterGroupDescription=description
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140429/us-east-1/rds/aws4_request
   &X-Amz-Date=20140429T175351Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=9164337efa99caf850e874a1cb7ef62f3cea29d0b448b9e0e7c53b288ddffed2

```

#### Sample Response

```

<CopyDBParameterGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CopyDBParameterGroupResult>
    <DBParameterGroup>
      <DBParameterGroupFamily>mysql5.6</DBParameterGroupFamily>
      <Description>description</Description>
      <DBParameterGroupName>new-local-param-group</DBParameterGroupName>
    </DBParameterGroup>
  </CopyDBParameterGroupResult>
  <ResponseMetadata>
    <RequestId>2928d60e-beb6-11d3-8e5c-3ccda5460c46</RequestId>
  </ResponseMetadata>
</CopyDBParameterGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/copydbparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/copydbparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyDBClusterSnapshot

CopyDBSnapshot

All content copied from https://docs.aws.amazon.com/.
