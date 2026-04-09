# Checking Confluence (Cloud) connectivity

Before you sync your Confluence (Cloud) data source connector after [configuring it](confluence-cloud-console.md), we recommend you check the
connection between Amazon Q Business and Confluence (Cloud). The following are the cURL
commands you need to check Confluence (Cloud) connectivity.

###### Topics

- [Checking basic authentication connectivity](#confluence-cloud-connection-check-basic)

## Checking basic authentication connectivity

To check connectivity for a Confluence (Cloud) data source connector using
basic authentication, use the following cURL command:

```nohighlight

curl --location 'https://<confluence_host-url>/wiki/rest/api/user/current'
--header 'Authorization: Basic <base64(email:api_token)>'
```

If your data source is connected as expected, the JSON response should
resemble the following:

```json

{
    "type": "known",
    "accountId": "accountId",
    "accountType": "atlassian",
    "email": "email",
    "publicName": "Administrator",
    "profilePicture": {
        "path": "/wiki/aa-avatar/<accountId>",
        "width": 48,
        "height": 48,
        "isDefault": false
    },
    "displayName": "Administrator",
    "isExternalCollaborator": false,
    "_expandable": {
        "operations": "",
        "personalSpace": ""
    },
    "_links": {
        "self": "https://<host_url>/wiki/rest/api/user?accountId=<accountId>",
        "base": "https://<host_url>/wiki",
        "context": "/wiki"
    }
}
```

If your Confluence (Cloud) connector is not connected correctly, you will see
the following error:

- CNF-5123: The profile value is invalid. Try again after
sometime.

To troubleshoot the issue, check your Confluence (Cloud) URL and make sure
it's correct.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

How Amazon Q works with Confluence (Cloud) access and refresh tokens

Using the console

All content copied from https://docs.aws.amazon.com/.
