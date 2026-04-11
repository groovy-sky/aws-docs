# IndexDocument

Container for the `Suffix` element.

## Contents

**Suffix**

A suffix that is appended to a request that is for a directory on the website endpoint. (For
example, if the suffix is `index.html` and you make a request to
`samplebucket/images/`, the data that is returned will be for the object with the key name
`images/index.html`.) The suffix must not be empty and must not include a slash
character.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../userguide/object-keys.md#object-key-xml-related-constraints).

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/indexdocument.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/indexdocument.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/indexdocument.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Grantee

Initiator

All content copied from https://docs.aws.amazon.com/.
