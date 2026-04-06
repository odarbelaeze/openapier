# OpenAPI Tags Documentation

This document lists all the supported `@` tags for the `openapier` tool.

## Spec-level Tags

These tags are used to define general information about the API (Info, Servers, etc.).

| Tag | Usage |
| :--- | :--- |
| `@externalDocs.description` | `@externalDocs.description <description>` |
| `@externalDocs.url` | `@externalDocs.url <url>` |
| `@info.contact.email` | `@info.contact.email <email>` |
| `@info.contact.name` | `@info.contact.name <name>` |
| `@info.contact.url` | `@info.contact.url <url>` |
| `@info.description` | `@info.description <description>` |
| `@info.license.name` | `@info.license.name <name>` |
| `@info.license.url` | `@info.license.url <url>` |
| `@info.summary` | `@info.summary <summary>` |
| `@info.termsOfService` | `@info.termsOfService <url>` |
| `@info.title` | `@info.title <title>` |
| `@info.version` | `@info.version <version>` |
| `@security` | `@security <name> [scope1] [scope2] ...` |
| `@server.description` | `@server.description <description>` |
| `@server.url` | `@server.url <url>` |
| `@server.variable.default` | `@server.variable.default <variable> <default>` |
| `@server.variable.description` | `@server.variable.description <variable> <description>` |
| `@server.variable.enum` | `@server.variable.enum <variable> [value1] [value2] ...` |
| `@tag.description` | `@tag.description <description>` |
| `@tag.externalDocs.description` | `@tag.externalDocs.description <description>` |
| `@tag.externalDocs.url` | `@tag.externalDocs.url <url>` |
| `@tag.name` | `@tag.name <name>` |

## Operation-level Tags

These tags are used to define individual API operations (Paths, Parameters, Responses, etc.).

| Tag | Usage |
| :--- | :--- |
| `@deprecated` | `@deprecated` |
| `@description` | `@description <description>` |
| `@externaldocs.description` | `@externalDocs.description <description>` |
| `@externaldocs.url` | `@externalDocs.url <url>` |
| `@id` | `@id <operationId>` |
| `@param` | `@param <name> <type> <in> [description...]` |
| `@param.allowEmptyValue` | `@param.allowEmptyValue <param> [param]...` |
| `@param.allowReserved` | `@param.allowReserved <param> [param]...` |
| `@param.deprecated` | `@param.deprecated <param> [param]...` |
| `@param.description` | `@param.description <name> <description>` |
| `@param.required` | `@param.required <param> [param]...` |
| `@requestBody` | `@requestBody <content_type> <type> [description]` |
| `@response` | `@response <status_code> <content_type> <type> [description]` |
| `@router` | `// @router <path> [method]` |
| `@security` | `@security <name> [scope1] [scope2] ...` |
| `@server.description` | `@server.description <description>` |
| `@server.url` | `@server.url <url>` |
| `@server.variable.default` | `@server.variable.default <variable> <default>` |
| `@server.variable.description` | `@server.variable.description <variable> <description>` |
| `@server.variable.enum` | `@server.variable.enum <variable> [value1] [value2] ...` |
| `@summary` | `@summary <summary>` |
| `@tags` | `@tags <tag1> [tag2]...` |
