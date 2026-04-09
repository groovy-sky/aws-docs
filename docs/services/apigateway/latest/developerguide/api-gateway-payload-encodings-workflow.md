# Content type conversions in API Gateway

The combination of your API's `binaryMediaTypes`, the headers in client
requests, and the integration `contentHandling` property determine how API Gateway
encodes payloads.

The following table shows how API Gateway converts the request payload for specific
configurations of a request's `Content-Type` header, the
`binaryMediaTypes` list of a [RestApi](../api/api-restapi.md) resource, and the `contentHandling` property value of the [Integration](../api/api-integration.md) resource.

Method request payloadRequest `Content-Type` header`binaryMediaTypes``contentHandling`Integration request payloadText dataAny data typeUndefinedUndefinedUTF8-encoded stringText dataAny data typeUndefined`CONVERT_TO_BINARY`Base64-decoded binary blobText dataAny data typeUndefined`CONVERT_TO_TEXT`UTF8-encoded stringText dataA text data typeSet with matching media typesUndefinedText dataText dataA text data typeSet with matching media types`CONVERT_TO_BINARY`Base64-decoded binary blobText dataA text data typeSet with matching media types`CONVERT_TO_TEXT`Text dataBinary dataA binary data typeSet with matching media typesUndefinedBinary dataBinary dataA binary data typeSet with matching media types`CONVERT_TO_BINARY`Binary dataBinary dataA binary data typeSet with matching media types`CONVERT_TO_TEXT`Base64-encoded string

The following table shows how API Gateway converts the response payload for specific
configurations of a request's `Accept` header, the `binaryMediaTypes` list of a [RestApi](../api/api-restapi.md) resource, and the `contentHandling` property value of the [IntegrationResponse](../api/api-integrationresponse.md)
resource.

###### Important

When a request contains multiple media types in its `Accept` header, API Gateway honors only the first `Accept` media type. If you can't control the order of the `Accept` media types and the media type of your binary
content isn't the first in the list, add the first `Accept` media type in the `binaryMediaTypes` list of your API. API Gateway handles all content types in
this list as binary.

For example, to send a JPEG file using an `<img>` element in a browser, the browser might send `Accept:image/webp,image/*,*/*;q=0.8` in a request. By
adding `image/webp` to the `binaryMediaTypes` list, the endpoint receives the JPEG file as binary.

Integration response payloadRequest `Accept` header`binaryMediaTypes``contentHandling`Method response payloadText or binary dataA text typeUndefinedUndefinedUTF8-encoded stringText or binary dataA text typeUndefined`CONVERT_TO_BINARY`Base64-decoded blobText or binary dataA text typeUndefined`CONVERT_TO_TEXT`UTF8-encoded stringText dataA text typeSet with matching media typesUndefinedText dataText dataA text typeSet with matching media types`CONVERT_TO_BINARY`Base64-decoded blobText dataA text typeSet with matching media types`CONVERT_TO_TEXT`UTF8-encoded stringText dataA binary typeSet with matching media typesUndefinedBase64-decoded blobText dataA binary typeSet with matching media types`CONVERT_TO_BINARY`Base64-decoded blobText dataA binary typeSet with matching media types`CONVERT_TO_TEXT`UTF8-encoded stringBinary dataA text typeSet with matching media typesUndefinedBase64-encoded stringBinary dataA text typeSet with matching media types`CONVERT_TO_BINARY`Binary dataBinary dataA text typeSet with matching media types`CONVERT_TO_TEXT`Base64-encoded stringBinary dataA binary typeSet with matching media typesUndefinedBinary dataBinary dataA binary typeSet with matching media types`CONVERT_TO_BINARY`Binary dataBinary dataA binary typeSet with matching media types`CONVERT_TO_TEXT`Base64-encoded string

When converting a text payload to a binary blob, API Gateway assumes that the text data is a
base64-encoded string and outputs the binary data as a base64-decoded blob. If the
conversion fails, it returns a `500` response, which
indicates an API configuration error. You don't provide a mapping template for such a
conversion, although you must enable the [passthrough behaviors](integration-passthrough-behaviors.md) on the
API.

When converting a binary payload to a text string, API Gateway always applies a base64
encoding on the binary data. You can define a mapping template for such a payload, but
can only access the base64-encoded string in the mapping template through `$input.body`, as shown in the following excerpt of an
example mapping template.

```nohighlight

{
    "data": "$input.body"
}
```

To have the binary payload passed through without modification, you must enable the
[passthrough behaviors](integration-passthrough-behaviors.md) on
the API.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Binary media types

Enabling binary support using the API Gateway console

All content copied from https://docs.aws.amazon.com/.
