# Use `wscat` to connect to a WebSocket API and send messages to it

The `wscat`
utility is a convenient tool for testing a WebSocket API that you have created and
deployed in API Gateway. You can install and use `wscat` as follows:

1. Download `wscat` from [https://www.npmjs.com/package/wscat](https://www.npmjs.com/package/wscat).

2. Install `wscat` by running the following command:

```nohighlight

npm install -g wscat
```

3. To connect to your API, run the `wscat` command as shown in the
    following example. Note that this example assumes that the
    `Authorization` setting is `NONE`.

```nohighlight

wscat -c wss://aabbccddee.execute-api.us-east-1.amazonaws.com/test/
```

You need to replace `aabbccddee` with
    the actual API ID, which is displayed in the API Gateway console or returned by the
    AWS CLI [`create-api`](../../../cli/latest/reference/apigatewayv2/create-api.md) command.

In addition, if your API is in a Region other than `us-east-1`, you
    need to substitute the correct Region.

4. To test your API, enter a message such as the following while
    connected:

```nohighlight

{"{jsonpath-expression}":"{route-key}"}
```

where `{jsonpath-expression}` is a JSONPath
    expression and `{route-key}` is a route key for the
    API. For example:

```nohighlight

{"action":"action1"}
{"message":"test response body"}
```

For more information about JSONPath, see [JSONPath](https://goessner.net/articles/JsonPath) or [JSONPath for Java](https://github.com/json-path/JsonPath).

5. To disconnect from your API, enter `ctrl-C`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Invoke

Use @connections commands in your backend service

All content copied from https://docs.aws.amazon.com/.
