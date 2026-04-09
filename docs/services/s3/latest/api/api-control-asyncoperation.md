# AsyncOperation

A container for the information about an asynchronous operation.

## Contents

**CreationTime**

The time that the request was sent to the service.

Type: Timestamp

Required: No

**Operation**

The specific operation for the asynchronous request.

Type: String

Valid Values: `CreateMultiRegionAccessPoint | DeleteMultiRegionAccessPoint | PutMultiRegionAccessPointPolicy`

Required: No

**RequestParameters**

The parameters associated with the request.

Type: [AsyncRequestParameters](api-control-asyncrequestparameters.md) data type

Required: No

**RequestStatus**

The current status of the request.

Type: String

Required: No

**RequestTokenARN**

The request token associated with the request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:.+`

Required: No

**ResponseDetails**

The details of the response.

Type: [AsyncResponseDetails](api-control-asyncresponsedetails.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/asyncoperation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/asyncoperation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/asyncoperation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AsyncErrorDetails

AsyncRequestParameters

All content copied from https://docs.aws.amazon.com/.
