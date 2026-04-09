# GetGroupConfiguration

Retrieves the service configuration associated with the specified resource group. For
details about the service configuration syntax, see [Service configurations for Resource Groups](about-slg.md).

**Minimum permissions**

To run this command, you must have the following permissions:

- `resource-groups:GetGroupConfiguration`

## Request Syntax

```nohighlight

POST /get-group-configuration HTTP/1.1
Content-type: application/json

{
   "Group": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[Group](#API_GetGroupConfiguration_RequestSyntax)**

The name or the Amazon resource name (ARN) of the resource group for which you want to retrive the service
configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[a-zA-Z0-9_\.-]{1,300}|[a-zA-Z0-9_\.-]{1,150}/[a-z0-9]{26}|arn:aws(-[a-z]+)*:resource-groups:[a-z]{2}(-[a-z]+)+-\d{1}:[0-9]{12}:group/([a-zA-Z0-9_\.-]{1,300}|[a-zA-Z0-9_\.-]{1,150}/[a-z0-9]{26})`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "GroupConfiguration": {
      "Configuration": [
         {
            "Parameters": [
               {
                  "Name": "string",
                  "Values": [ "string" ]
               }
            ],
            "Type": "string"
         }
      ],
      "FailureReason": "string",
      "ProposedConfiguration": [
         {
            "Parameters": [
               {
                  "Name": "string",
                  "Values": [ "string" ]
               }
            ],
            "Type": "string"
         }
      ],
      "Status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[GroupConfiguration](#API_GetGroupConfiguration_ResponseSyntax)**

A structure that describes the service configuration attached with the specified
group. For details about the service configuration syntax, see [Service configurations for\
Resource Groups](about-slg.md).

Type: [GroupConfiguration](api-groupconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The request includes one or more parameters that violate validation rules.

HTTP Status Code: 400

**ForbiddenException**

The caller isn't authorized to make the request. Check permissions.

HTTP Status Code: 403

**InternalServerErrorException**

An internal error occurred while processing the request. Try again later.

HTTP Status Code: 500

**MethodNotAllowedException**

The request uses an HTTP method that isn't allowed for the specified resource.

HTTP Status Code: 405

**NotFoundException**

One or more of the specified resources don't exist.

HTTP Status Code: 404

**TooManyRequestsException**

You've exceeded throttling limits by making too many requests in a period of
time.

HTTP Status Code: 429

## Examples

### Example

The following example shows the typical output of a resource group that
contains capacity reservations.

#### Sample Request

```

POST /get-group-configuration HTTP/1.1
Host: resource-groups.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.40 Python/3.8.8 Windows/10 exe/AMD64 prompt/off command/resource-groups.get-group-configuration
X-Amz-Date: 20220113T213346Z
X-Amz-Security-Token: <SECURITY-TOKEN>
Authorization: AWS4-HMAC-SHA256 Credential=<ACCESS-KEY>/20220113/us-west-2/resource-groups/aws4_request,SignedHeaders=host;x-amz-date;x-amz-security-token,Signature=<SIGV4-SIGNATURE>
Content-Length: 21

{
    "Group": "CRPGroup"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 13 Jan 2022 21:33:46 GMT
Content-Type: application/json
Content-Length: 461
x-amzn-RequestId: <VARIES>
x-amz-apigw-id: <VARIES>
X-Amzn-Trace-Id: Root=<VARIES>
Connection: keep-alive

{
    "GroupIdentifier":{
        "GroupName":"CRPGroup",
        "GroupArn":"arn:aws:resource-groups:us-west-2:123456789012:group/CRPGroup"
    },
    "GroupConfiguration":{
        "GroupParameters":[
            {
                "Name":"allowed-resource-types",
                "Values":[
                    "AWS::EC2::CapacityReservation"
                ]
            }
        ],
        "Configuration":[
            {
                "Type":"AWS::EC2::CapacityReservationPool"
            },
            {
                "Type":"AWS::ResourceGroups::Generic",
                "Parameters":[
                    {
                        "Name":"allowed-resource-types",
                        "Values":[
                            "AWS::EC2::CapacityReservation"
                        ]
                    }
                ]
            }
        ],
        "Status":"UPDATE_COMPLETE"
    }
}
```

### Example

The following example shows the typical output of a resource group that
contains a configuration for Amazon EC2 dedicated hosts.

#### Sample Request

```

POST /get-group-configuration HTTP/1.1
Host: resource-groups.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.40 Python/3.8.8 Windows/10 exe/AMD64 prompt/off command/resource-groups.get-group-configuration
X-Amz-Date: 20220113T213500Z
X-Amz-Security-Token: <SECURITY-TOKEN>
Authorization: AWS4-HMAC-SHA256 Credential=<ACCESS-KEY>/20220113/us-west-2/resource-groups/aws4_request,SignedHeaders=host;x-amz-date;x-amz-security-token,Signature=<SIGV4-SIGNATURE>
Content-Length: 32

{
    "Group": "HostManagementGroup"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 13 Jan 2022 21:35:00 GMT
Content-Type: application/json
Content-Length: 871
x-amzn-RequestId: <VARIES>
x-amz-apigw-id: <VARIES>
X-Amzn-Trace-Id: Root=<VARIES>
Connection: keep-alive

{
    "GroupIdentifier":{
        "GroupName":"HostManagementGroup",
        "GroupArn":"arn:aws:resource-groups:us-west-2:123456789012:group/HostManagementGroup"
    },
    "GroupConfiguration":{
        "GroupParameters":[
            {
                "Name":"allowed-resource-types",
                "Values":[
                    "AWS::EC2::Host"
                ]
            },
            {
                "Name":"deletion-protection",
                "Values":[
                    "UNLESS_EMPTY"
                ]
            }
        ],
        "ResourceTypeParameters":[
            {
                "ResourceType":"AWS::EC2::Host",
                "Name":"exclusive-membership",
                "Values":[
                    "ACROSS_SAME_SERVICE_LINK_TYPE"
                ]
            }
        ],
        "ServiceLinkParameters":[
            {
                "Name":"any-host-based-license-configuration",
                "Values":[
                    "true"
                ]
            }
        ],
        "Configuration":[
            {
                "Type":"AWS::EC2::HostManagement",
                "Parameters":[
                    {
                        "Name":"any-host-based-license-configuration",
                        "Values":[
                            "true"
                        ]
                    }
                ]
            },
            {
                "Type":"AWS::ResourceGroups::Generic",
                "Parameters":[
                    {
                        "Name":"allowed-resource-types",
                        "Values":[
                            "AWS::EC2::Host"
                        ]
                    },
                    {
                        "Name":"deletion-protection",
                        "Values":[
                            "UNLESS_EMPTY"
                        ]
                    }
                ]
            }
        ],
        "Status":"UPDATE_COMPLETE"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/resource-groups-2017-11-27/getgroupconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/getgroupconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetGroup

GetGroupQuery

All content copied from https://docs.aws.amazon.com/.
