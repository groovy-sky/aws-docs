This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener ListenerAttribute

Information about a listener attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The name of the attribute.

The following attribute is supported by Network Load Balancers, and Gateway Load Balancers.

- `tcp.idle_timeout.seconds` \- The tcp idle timeout value, in seconds. The
valid range is 60-6000 seconds. The default is 350 seconds.

The following attributes are only supported by Application Load Balancers.

- `routing.http.request.x_amzn_mtls_clientcert_serial_number.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Mtls-Clientcert-Serial-Number** HTTP request header.

- `routing.http.request.x_amzn_mtls_clientcert_issuer.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Mtls-Clientcert-Issuer** HTTP request header.

- `routing.http.request.x_amzn_mtls_clientcert_subject.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Mtls-Clientcert-Subject** HTTP request header.

- `routing.http.request.x_amzn_mtls_clientcert_validity.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Mtls-Clientcert-Validity** HTTP request header.

- `routing.http.request.x_amzn_mtls_clientcert_leaf.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Mtls-Clientcert-Leaf** HTTP request header.

- `routing.http.request.x_amzn_mtls_clientcert.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Mtls-Clientcert** HTTP request header.

- `routing.http.request.x_amzn_tls_version.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Tls-Version** HTTP request header.

- `routing.http.request.x_amzn_tls_cipher_suite.header_name` \-
Enables you to modify the header name of the
**X-Amzn-Tls-Cipher-Suite** HTTP request header.

- `routing.http.response.server.enabled` \-
Enables you to allow or remove the HTTP response server header.

- `routing.http.response.strict_transport_security.header_value` \-
Informs browsers that the site should only be accessed using HTTPS, and that
any future attempts to access it using HTTP should automatically be converted
to HTTPS.

- `routing.http.response.access_control_allow_origin.header_value` \-
Specifies which origins are allowed to access the server.

- `routing.http.response.access_control_allow_methods.header_value` \-
Returns which HTTP methods are allowed when accessing the server from a different
origin.

- `routing.http.response.access_control_allow_headers.header_value` \-
Specifies which headers can be used during the request.

- `routing.http.response.access_control_allow_credentials.header_value` \-
Indicates whether the browser should include credentials such as cookies or
authentication when making requests.

- `routing.http.response.access_control_expose_headers.header_value` \-
Returns which headers the browser can expose to the requesting client.

- `routing.http.response.access_control_max_age.header_value` \-
Specifies how long the results of a preflight request can be cached, in seconds.

- `routing.http.response.content_security_policy.header_value` \-
Specifies restrictions enforced by the browser to help minimize the risk of certain
types of security threats.

- `routing.http.response.x_content_type_options.header_value` \-
Indicates whether the MIME types advertised in the **Content-Type**
headers should be followed and not be changed.

- `routing.http.response.x_frame_options.header_value` \- Indicates
whether the browser is allowed to render a page in a **frame**,
**iframe**, **embed** or
**object**.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._]+$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the attribute.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example defines a TCP listener, specifying the `tcp.idle_timeout.seconds`
attribute.

#### YAML

```yaml

myTCPListener:
    Type: AWS::ElasticLoadBalancingV2::Listener
    Properties:
      LoadBalancerArn: !Ref myLoadBalancer
      Protocol: TCP
      Port: 80
      DefaultActions:
        - Type: forward
          TargetGroupArn: !Ref myTargetGroup
      ListenerAttributes:
        - Key: tcp.idle_timeout.seconds
          Value: 500
```

#### JSON

```json

{
    "myTCPListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "TCP",
            "Port": 80,
            "DefaultActions": [
                {
                    "Type": "forward",
                    "TargetGroupArn": {
                        "Ref": "myTargetGroup"
                    }
                }
            ],
            "ListenerAttributes": [
                {
                    "Key": "tcp.idle_timeout.seconds",
                    "Value": 500
                }
            ]
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JwtValidationConfig

MutualAuthentication
