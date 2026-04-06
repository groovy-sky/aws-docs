# VPCLinks

Represents your VPC links as a collection. A collection offers a paginated view of your VPC links.

## URI

`/v2/vpclinks`

## HTTP methods

### GET

**Operation ID:** `GetVpcLinks`

Gets a collection of VPC links.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``VpcLinks`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateVpcLink`

Creates a VPC link.

ResponsesStatus codeResponse modelDescription`201``VpcLink`

The request has succeeded and has resulted in the creation of a resource.

`400``BadRequestException`

One of the parameters in the request is invalid.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

## Schemas

### Request bodies

```json

{
  "name": "string",
  "subnetIds": [
    "string"
  ],
  "securityGroupIds": [
    "string"
  ],
  "tags": {
  }
}
```

### Response bodies

```json

{
  "items": [
    {
      "vpcLinkId": "string",
      "name": "string",
      "subnetIds": [
        "string"
      ],
      "securityGroupIds": [
        "string"
      ],
      "tags": {
      },
      "createdDate": "string",
      "vpcLinkStatus": enum,
      "vpcLinkStatusMessage": "string",
      "vpcLinkVersion": enum
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "vpcLinkId": "string",
  "name": "string",
  "subnetIds": [
    "string"
  ],
  "securityGroupIds": [
    "string"
  ],
  "tags": {
  },
  "createdDate": "string",
  "vpcLinkStatus": enum,
  "vpcLinkStatusMessage": "string",
  "vpcLinkVersion": enum
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "limitType": "string"
}
```

## Properties

### BadRequestException

The request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### CreateVpcLinkInput

Represents the input parameters for a `CreateVpcLink` request.

PropertyTypeRequiredDescription`name`

string

True

The name of the VPC link.

`securityGroupIds`

Array of type string

False

A list of security group IDs for the VPC link.

`subnetIds`

Array of type string

True

A list of subnet IDs to include in the VPC link.

`tags`

[Tags](#vpclinks-model-tags)

False

A list of tags.

### LimitExceededException

A limit has been exceeded. See the accompanying error message for details.

PropertyTypeRequiredDescription`limitType`

string

False

The limit type.

`message`

string

False

Describes the error encountered.

### Tags

Represents a collection of tags associated with the resource.

PropertyTypeRequiredDescription

`*`

string

False

### VpcLink

Represents a VPC link.

PropertyTypeRequiredDescription`createdDate`

string

Format: date-time

False

The timestamp when the VPC link was created.

`name`

string

True

The name of the VPC link.

`securityGroupIds`

Array of type string

True

A list of security group IDs for the VPC link.

`subnetIds`

Array of type string

True

A list of subnet IDs to include in the VPC link.

`tags`

[Tags](#vpclinks-model-tags)

False

Tags for the VPC link.

`vpcLinkId`

string

True

The ID of the VPC link.

`vpcLinkStatus`

[VpcLinkStatus](#vpclinks-model-vpclinkstatus)

False

The status of the VPC link.

`vpcLinkStatusMessage`

string

False

A message summarizing the cause of the status of the VPC link.

`vpcLinkVersion`

[VpcLinkVersion](#vpclinks-model-vpclinkversion)

False

The version of the VPC link.

### VpcLinkStatus

The status of the VPC link.

- `PENDING`

- `AVAILABLE`

- `DELETING`

- `FAILED`

- `INACTIVE`

### VpcLinkVersion

The version of the VPC link.

- `V2`

### VpcLinks

Represents a collection of VPCLinks.

PropertyTypeRequiredDescription`items`

Array of type VpcLink

False

A collection of VPC links.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetVpcLinks

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetVpcLinks)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetVpcLinks)

### CreateVpcLink

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateVpcLink)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateVpcLink)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VPCLink

Operations
