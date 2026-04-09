# Ping

Ping is a SAML based plugin that works with the [PingFederate](https://www.pingidentity.com/en/platform/capabilities/authentication-authority/pingfederate.html) identity provider.

## Authentication type

**Connection string name****Parameter type****Default value****Connection string example**AuthenticationTypeRequired`IAM Credentials``AuthenticationType=Ping;`

## User ID

The user name for the PingFederate server.

**Connection string name****Parameter type****Default value****Connection string example**UIDRequired`none``UID=pingusername@domain.com;`

## Password

The password for the PingFederate server.

**Connection string name****Parameter type****Default value****Connection string example**PWDRequired`none``PWD=pingpassword;`

## Preferred role

The Amazon Resource Name (ARN) of the role to assume. If your SAML assertion has
multiple roles, you can specify this parameter to choose the role to be assumed. This
role should be present in the SAML assertion. For more information about ARN roles, see
[AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**preferred\_roleOptional`none``preferred_role=arn:aws:iam::123456789012:id/user1;`

## Session duration

The duration, in seconds, of the role session. For more information about session
duration, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the
_AWS Security Token Service API Reference_.

**Connection string name****Parameter type****Default value****Connection string example**durationOptional`900``duration=900;`

## IdP host

The address for your Ping server. To find your address, visit the following URL and
view the **SSO Application Endpoint** field.

```

https://your-pf-host-#:9999/pingfederate/your-pf-app#/spConnections
```

**Connection string name****Parameter type****Default value****Connection string example**idp\_hostRequired`none``idp_host=ec2-1-83-65-12.compute-1.amazonaws.com;`

## IdP port

The port number to use to connect to your IdP host.

**Connection string name****Parameter type****Default value****Connection string example**idp\_portRequired`None``idp_port=443;`

## Partner SPID

The service provider address. To find the service provider address, visit the
following URL and view the **SSO Application Endpoint** field.

```

https://your-pf-host-#:9999/pingfederate/your-pf-app#/spConnections
```

**Connection string name****Parameter type****Default value****Connection string example**partner\_spidRequired`None``partner_spid=https://us-east-1.signin.aws.amazon.com/platform/saml/<...>;`

## Ping URI param

Passes a URI argument for an authentication request to Ping. Use this parameter to
bypass the Lake Formation single role limitation. Configure Ping to recognize the passed parameter, and verify that the role passed exists in
the list of roles assigned to the user. Then, send a single role in the SAML
assertion.

**Connection string name****Parameter type****Default value****Connection string example**ping\_uri\_paramOptional`None``ping_uri_param=role=my_iam_role;`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Okta

Common
auth parameters

All content copied from https://docs.aws.amazon.com/.
