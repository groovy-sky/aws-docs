# ApplyPendingMaintenanceAction

Applies a pending maintenance action to a resource (for example, to a DB instance).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ApplyAction**

The pending maintenance action to apply to this resource.

Valid Values:

- `ca-certificate-rotation`

- `db-upgrade`

- `hardware-maintenance`

- `os-upgrade`

- `system-update`

For more information about these actions, see
[Maintenance actions for Amazon Aurora](../../../../services/amazonrds/latest/aurorauserguide/user-upgradedbinstance-maintenance.md#maintenance-actions-aurora) or
[Maintenance actions for Amazon RDS](../../../../services/amazonrds/latest/userguide/user-upgradedbinstance-maintenance.md#maintenance-actions-rds).

Type: String

Required: Yes

**OptInType**

A value that specifies the type of opt-in request, or undoes an opt-in request. An opt-in
request of type `immediate` can't be undone.

Valid Values:

- `immediate` \- Apply the maintenance action immediately.

- `next-maintenance` \- Apply the maintenance action during
the next maintenance window for the resource.

- `undo-opt-in` \- Cancel any existing `next-maintenance`
opt-in requests.

Type: String

Required: Yes

**ResourceIdentifier**

The RDS Amazon Resource Name (ARN) of the resource that the
pending maintenance action applies to. For information about
creating an ARN,
see [Constructing an RDS Amazon Resource Name (ARN)](../../../../services/amazonrds/latest/userguide/user-tagging-arn.md#USER_Tagging.ARN.Constructing).

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**ResourcePendingMaintenanceActions**

Describes the pending maintenance actions for a resource.

Type: [ResourcePendingMaintenanceActions](api-resourcependingmaintenanceactions.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**ResourceNotFoundFault**

The specified resource ID was not found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of ApplyPendingMaintenanceAction.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=ApplyPendingMaintenanceAction
   &ResourceIdentifier=arn:aws:rds:us-east-1:123456781234:db:my-instance
   &ApplyAction=system-update
   &OptInType=immediate
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20141216/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194732Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=6e25c542bf96fe24b28c12976ec92d2f856ab1d2a158e21c35441a736e4fde2b

```

#### Sample Response

```

<ApplyPendingMaintenanceActionResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ApplyPendingMaintenanceActionResult>
     <ResourcePendingMaintenanceActions>
      <ResourceIdentifier>arn:aws:rds:us-east-1:123456781234:db:my-instance</ResourceIdentifier>
      <PendingMaintenanceActionDetails>
        <PendingMaintenanceAction>
          <Action>system-update</Action>
          <OptInStatus>immediate</OptInStatus>
        </PendingMaintenanceAction>
      </PendingMaintenanceActionDetails>
     /ResourcePendingMaintenanceActions>
  </ApplyPendingMaintenanceActionResult>
  <ResponseMetadata>
    <RequestId>dcfe0682-870c-11e4-9833-b3ad657ea9da</RequestId>
  </ResponseMetadata>
</ApplyPendingMaintenanceActionResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/applypendingmaintenanceaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/applypendingmaintenanceaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddTagsToResource

AuthorizeDBSecurityGroupIngress

All content copied from https://docs.aws.amazon.com/.
