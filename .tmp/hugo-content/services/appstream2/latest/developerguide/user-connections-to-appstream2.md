# User Connections to Amazon WorkSpaces Applications

Users can connect to WorkSpaces Applications streaming instances through the default public
internet endpoint, or by using an interface VPC
endpoint (interface endpoint) that you create in your virtual private
cloud (VPC). For more information, see [Tutorial: Creating and Streaming from Interface VPC Endpoints](creating-streaming-from-interface-vpc-endpoints.md).

By default, WorkSpaces Applications is configured to route streaming connections over the
public internet. Internet connectivity is required to authenticate users and deliver
the web assets that WorkSpaces Applications requires to function. To allow this traffic, you
must allow the domains listed in [Allowed Domains](allowed-domains.md).

###### Note

For user authentication, WorkSpaces Applications supports user pools, Security Assertion Markup Language 2.0 (SAML 2.0), and the [CreateStreamingURL](../../../../reference/appstream2/latest/apireference/api-createstreamingurl.md) API action. For more information, see [User Authentication](authentication-authorization.md).

The following topics provide information about how to enable user connections to WorkSpaces Applications.

###### Contents

- [Bandwidth Recommendations](bandwidth-recommendations-user-connections.md)

- [IP Address and Port Requirements for WorkSpaces Applications User Devices](client-application-ports.md)

- [Allowed Domains](allowed-domains.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Customer Network Interface Ports

Bandwidth Recommendations

All content copied from https://docs.aws.amazon.com/.
