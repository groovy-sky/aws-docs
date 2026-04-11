# DeleteVPCAssociationAuthorization

Removes authorization to submit an `AssociateVPCWithHostedZone` request to
associate a specified VPC with a hosted zone that was created by a different account.
You must use the account that created the hosted zone to submit a
`DeleteVPCAssociationAuthorization` request.

###### Important

Sending this request only prevents the AWS account that created the
VPC from associating the VPC with the Amazon Route 53 hosted zone in the future. If
the VPC is already associated with the hosted zone,
`DeleteVPCAssociationAuthorization` won't disassociate the VPC from
the hosted zone. If you want to delete an existing association, use
`DisassociateVPCFromHostedZone`.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone/Id/deauthorizevpcassociation HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<DeleteVPCAssociationAuthorizationRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <VPC>
      <VPCId>string</VPCId>
      <VPCRegion>string</VPCRegion>
   </VPC>
</DeleteVPCAssociationAuthorizationRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DeleteVPCAssociationAuthorization_RequestSyntax)**

When removing authorization to associate a VPC that was created by one AWS account with a hosted zone that was created with a different AWS account, the ID of the hosted zone.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[DeleteVPCAssociationAuthorizationRequest](#API_DeleteVPCAssociationAuthorization_RequestSyntax)**

Root level tag for the DeleteVPCAssociationAuthorizationRequest parameters.

Required: Yes

**[VPC](#API_DeleteVPCAssociationAuthorization_RequestSyntax)**

When removing authorization to associate a VPC that was created by one AWS account with a hosted zone that was created with a different AWS account, a complex type that includes the ID and region of the
VPC.

Type: [VPC](api-vpc.md) object

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModification**

Another user submitted a request to create, update, or delete the object at the same
time that you did. Retry the request.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**InvalidVPCId**

The VPC ID that you specified either isn't a valid ID or the current account is not
authorized to access this VPC.

**message**

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

**VPCAssociationAuthorizationNotFound**

The VPC that you specified is not authorized to be associated with the hosted
zone.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of DeleteVPCAssociationAuthorization.

```

POST /2013-04-01/hostedzone/Z1PA6795UKMFR9/deauthorizevpcassociation HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<DeleteVPCAssociationAuthorizationRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <VPC>
      <VPCId>vpc-a1b2c3d4e5</VPCId>
      <VPCRegion>us-east-2</VPCRegion>
   </VPC>
</DeleteVPCAssociationAuthorizationRequest>
```

### Example Response

This example illustrates one usage of DeleteVPCAssociationAuthorization.

```

HTTP/1.1 200
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletevpcassociationauthorization.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletevpcassociationauthorization.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteTrafficPolicyInstance

DisableHostedZoneDNSSEC
