# Making requests to the Amazon EC2 API

###### Important

As of **October 14, 2022**, HTTP responses from
the Amazon EC2 APIs no longer include a reason-phrase element. As recommended by
[RFC7230](https://datatracker.ietf.org/doc/html/rfc7230),
you should ensure that your applications do not make use of the reason-phrase content.
Ensure that your applications use the 3-digit status-code element included in the
HTTP response instead.

We provide the Query API for Amazon EC2, as well as software development kits (SDK) for AWS
that enable you to access Amazon EC2 from your preferred programming language. For more information,
see the [Amazon EC2 Developer Guide](../../../../services/ec2/latest/devguide.md).

###### Contents

- [Required knowledge](#required-knowledge)

- [Available APIs for Amazon EC2](#using-libraries)

- [Query requests for Amazon EC2](query-requests.md)

- [Troubleshooting API request errors](query-api-troubleshooting.md)

- [Cross-origin resource sharing support and Amazon EC2](cors-support.md)

- [VM Import Manifest](manifest.md)

## Required knowledge

If you plan to access Amazon EC2 through an API, you should be familiar with the
following:

- XML

- Web services

- HTTP requests

- One or more programming languages, such as Java, PHP, Perl, Python, Ruby, C#,
or C++.

## Available APIs for Amazon EC2

The Amazon EC2 Query API provides HTTP or HTTPS requests that use the HTTP verb GET or POST
and a Query parameter named `Action`.

AWS provides libraries, sample code, tutorials, and other resources for software
developers who prefer to build applications using language-specific APIs instead of
submitting a request over HTTP or HTTPS. These libraries provide basic functions that
automatically take care of tasks such as cryptographically signing your requests,
retrying requests, and handling error responses, so that it is easier for you to get
started.

For more information, see [Create Amazon EC2 resources using an AWS SDK](../../../ec2/latest/devguide/sdk-general-information-section.md) in the _Amazon EC2 Developer Guide_.

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpnTunnelOptionsSpecification

Query requests
