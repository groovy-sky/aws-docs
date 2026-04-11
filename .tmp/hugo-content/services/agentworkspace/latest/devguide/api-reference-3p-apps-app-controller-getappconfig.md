# Get the application configuration in Amazon Connect Agent Workspace

Returns the application configuration for the given application ARN in the Amazon Connect
agent workspace.

**Signature**

```typescript

getAppConfig(appArn: string): Promise<AppConfig>
```

**Usage**

```typescript

const applicationConfig: AppConfig = await appControllerClient.getAppConfig(arn);
```

**Input**

**Parameter****Type****Description**arn _Required_stringThe ARN of the application

**Output - AppConfig**

**Parameter****Type****Description**arnstringThe AmazonResourceName(ARN) of the applicationnamespacestringThe immutable application namespace used at the time of
registrationidstringThe unique identifier of the applicationnamestringName of the applicationdescriptionstringDescription of the applicationaccessUrlstringURL to access the applicationinitializationTimeoutnumberThe maximum time allowed to establish a connection with the
workspacecontactHandling.scopePER\_CONTACT \| CROSS\_CONTACTSIndicates whether the application refreshes for each
contact

**Permissions required:**

```typescript

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getAppCatalog()

getApps()

All content copied from https://docs.aws.amazon.com/.
