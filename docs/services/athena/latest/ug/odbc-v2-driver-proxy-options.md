# Proxy options

## Proxy host

If you require users to go through a proxy, use this parameter to set the proxy host.
This parameter corresponds to the `ClientConfiguration.proxyHost` parameter
in the AWS SDK. For more information, see [AWS Client\
configuration](../../../../reference/sdk-for-cpp/v1/developer-guide/client-config.md) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**ProxyHostOptional`none``ProxyHost=127.0.0.1;`

## Proxy port

Use this parameter to set the proxy port. This parameter corresponds to the
`ClientConfiguration.proxyPort` parameter in the AWS SDK. For more
information, see [AWS Client\
configuration](../../../../reference/sdk-for-cpp/v1/developer-guide/client-config.md) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**ProxyPortOptional`none``ProxyPort=8888;`

## Proxy user name

Use this parameter to set the proxy user name. This parameter corresponds to the
`ClientConfiguration.proxyUserName` parameter in the AWS SDK. For more
information, see [AWS Client\
configuration](../../../../reference/sdk-for-cpp/v1/developer-guide/client-config.md) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**ProxyUIDOptional`none``ProxyUID=username;`

## Proxy password

Use this parameter to set the proxy password. This parameter corresponds to the
`ClientConfiguration.proxyPassword` parameter in the AWS SDK. For more
information, see [AWS Client\
configuration](../../../../reference/sdk-for-cpp/v1/developer-guide/client-config.md) in the _AWS SDK for C++ Developer Guide_.

**Connection string name****Parameter type****Default value****Connection string example**ProxyPWDOptional`none``ProxyPWD=password;`

## Non proxy host

Use this optional parameter to specify a host that the driver connects to without
using a proxy. This parameter corresponds to the
`ClientConfiguration.nonProxyHosts` parameter in the AWS SDK. For more
information, see [AWS Client\
configuration](../../../../reference/sdk-for-cpp/v1/developer-guide/client-config.md) in the _AWS SDK for C++ Developer Guide_.

The `NonProxyHost` connection parameter is passed to the
`CURLOPT_NOPROXY` curl option. For information about the
`CURLOPT_NOPROXY` format, see [CURLOPT\_NOPROXY](https://curl.se/libcurl/c/CURLOPT_NOPROXY.html) in the
curl documentation.

**Connection string name****Parameter type****Default value****Connection string example**NonProxyHostOptional`none``NonProxyHost=.amazonaws.com,localhost,.example.net,.example.com;`

## Use proxy

Enables user traffic through the specified proxy.

**Connection string name****Parameter type****Default value****Connection string example**UseProxyOptional`none``UseProxy=1;`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Advanced

Logging

All content copied from https://docs.aws.amazon.com/.
