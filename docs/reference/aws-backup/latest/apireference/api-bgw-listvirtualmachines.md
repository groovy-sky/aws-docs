---
title: "ListVirtualMachines"
---

# ListVirtualMachines

Lists your virtual machines.

## Request Syntax

```nohighlight

{
   "HypervisorArn": "string",
   "MaxResults": number,
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[HypervisorArn](#API_BGW_ListVirtualMachines_RequestSyntax)**

The Amazon Resource Name (ARN) of the hypervisor connected to your virtual machine.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 500.

Pattern: `arn:(aws|aws-cn|aws-us-gov):backup-gateway(:[a-zA-Z-0-9]+){3}\/[a-zA-Z-0-9]+`

Required: No

**[MaxResults](#API_BGW_ListVirtualMachines_RequestSyntax)**

The maximum number of virtual machines to list.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**[NextToken](#API_BGW_ListVirtualMachines_RequestSyntax)**

The next item following a partial list of returned resources. For example, if a request is
made to return `maxResults` number of resources, `NextToken` allows you
to return more items in your list starting at the location pointed to by the next
token.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `.+`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "VirtualMachines": [
      {
         "HostName": "string",
         "HypervisorId": "string",
         "LastBackupDate": number,
         "Name": "string",
         "Path": "string",
         "ResourceArn": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_BGW_ListVirtualMachines_ResponseSyntax)**

The next item following a partial list of returned resources. For example, if a request is
made to return `maxResults` number of resources, `NextToken` allows you
to return more items in your list starting at the location pointed to by the next
token.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `.+`

**[VirtualMachines](#API_BGW_ListVirtualMachines_ResponseSyntax)**

A list of your `VirtualMachine` objects, ordered by their Amazon Resource Names
(ARNs).

Type: Array of [VirtualMachine](api-bgw-virtualmachine.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

The operation did not succeed because an internal error occurred. Try again later.

**ErrorCode**

A description of which internal error occured.

HTTP Status Code: 500

**ThrottlingException**

TPS has been limited to protect against intentional or unintentional
high request volumes.

**ErrorCode**

Error: TPS has been limited to protect against intentional or unintentional
high request volumes.

HTTP Status Code: 400

**ValidationException**

The operation did not succeed because a validation error occurred.

**ErrorCode**

A description of what caused the validation error.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/backup-gateway-2021-01-01/ListVirtualMachines)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/backup-gateway-2021-01-01/ListVirtualMachines)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForResource

PutBandwidthRateLimitSchedule

All content copied from https://docs.aws.amazon.com/.
