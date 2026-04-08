# CreateVPCAssociationAuthorization

Authorizes the AWS account that created a specified VPC to submit an
`AssociateVPCWithHostedZone` request to associate the VPC with a
specified hosted zone that was created by a different account. To submit a
`CreateVPCAssociationAuthorization` request, you must use the account
that created the hosted zone. After you authorize the association, use the account that
created the VPC to submit an `AssociateVPCWithHostedZone` request.

###### Note

If you want to associate multiple VPCs that you created by using one account with
a hosted zone that you created by using a different account, you must submit one
authorization request for each VPC.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone/Id/authorizevpcassociation HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<CreateVPCAssociationAuthorizationRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <VPC>
      <VPCId>string</VPCId>
      <VPCRegion>string</VPCRegion>
   </VPC>
</CreateVPCAssociationAuthorizationRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_CreateVPCAssociationAuthorization_RequestSyntax)**

The ID of the private hosted zone that you want to authorize associating a VPC
with.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[CreateVPCAssociationAuthorizationRequest](#API_CreateVPCAssociationAuthorization_RequestSyntax)**

Root level tag for the CreateVPCAssociationAuthorizationRequest parameters.

Required: Yes

**[VPC](#API_CreateVPCAssociationAuthorization_RequestSyntax)**

A complex type that contains the VPC ID and region for the VPC that you want to
authorize associating with your hosted zone.

Type: [VPC](api-vpc.md) object

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<CreateVPCAssociationAuthorizationResponse>
   <HostedZoneId>string</HostedZoneId>
   <VPC>
      <VPCId>string</VPCId>
      <VPCRegion>string</VPCRegion>
   </VPC>
</CreateVPCAssociationAuthorizationResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[CreateVPCAssociationAuthorizationResponse](#API_CreateVPCAssociationAuthorization_ResponseSyntax)**

Root level tag for the CreateVPCAssociationAuthorizationResponse parameters.

Required: Yes

**[HostedZoneId](#API_CreateVPCAssociationAuthorization_ResponseSyntax)**

The ID of the hosted zone that you authorized associating a VPC with.

Type: String

Length Constraints: Maximum length of 32.

**[VPC](#API_CreateVPCAssociationAuthorization_ResponseSyntax)**

The VPC that you authorized associating with a hosted zone.

Type: [VPC](api-vpc.md) object

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

**TooManyVPCAssociationAuthorizations**

You've created the maximum number of authorizations that can be created for the
specified hosted zone. To authorize another VPC to be associated with the hosted zone,
submit a `DeleteVPCAssociationAuthorization` request to remove an existing
authorization. To get a list of existing authorizations, submit a
`ListVPCAssociationAuthorizations` request.

**message**

HTTP Status Code: 400

## Examples

### Example request

This example illustrates one usage of CreateVPCAssociationAuthorization.

```

POST /2013-04-01/hostedzone/Z1PA6795UKMFR9/authorizevpcassociation HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<CreateVPCAssociationAuthorizationRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <VPC>
      <VPCId>vpc-a1b2c3d4e5</VPCId>
      <VPCRegion>us-east-2</VPCRegion>
   </VPC>
</CreateVPCAssociationAuthorizationRequest>
```

### Example Response

This example illustrates one usage of CreateVPCAssociationAuthorization.

```

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<CreateVPCAssociationAuthorizationResponse>
   <HostedZoneId>Z1PA6795UKMFR9</HostedZoneId>
   <VPC>
      <VPCId>vpc-a1b2c3d4e5</VPCId>
      <VPCRegion>us-east-2</VPCRegion>
   </VPC>
</CreateVPCAssociationAuthorizationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/createvpcassociationauthorization.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/createvpcassociationauthorization.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTrafficPolicyVersion

DeactivateKeySigningKey
