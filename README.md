# external-dns-infoblox-webhook

Infoblox provider based on in-tree provider for ExternalDNS. Supported records:

| Record Type | Status     |
|-------------|------------|
| A           | supported  |
| CNAME       | supported  |
| TXT         | supported  |
| PTR         | not tested |


## Quick start

To run the provider, you must provide the following Environment Variables:

**Infoblox Environment Variables**:

| Environment Variable        | Default value | Required |
|-----------------------------|---------------|----------|
| INFOBLOX_HOST               | localhost     | true     |
| INFOBLOX_PORT               | 443           | true     |   
| INFOBLOX_WAPI_USER          |               | true     |
| INFOBLOX_WAPI_PASSWORD      |               | true     |
| INFOBLOX_VERSION            |               | true     |
| INFOBLOX_SSL_VERIFY         | true          | false    |
| INFOBLOX_DRY_RUN            | false         | false    |
| INFOBLOX_VIEW               | default       | false    |
| INFOBLOX_MAX_RESULTS        | 1500          | false    |
| INFOBLOX_CREATE_PTR         | false         | false    |
| INFOBLOX_DEFAULT_TTL        | 300           | false    |


**external-dns-infoblox-webhook Environment Variables**:

| Environment Variable           | Default value | Required |
|--------------------------------|---------------|----------|
| SERVER_HOST                    | 0.0.0.0       | true     |
| SERVER_PORT                    | 8888          | true     |   
| SERVER_READ_TIMEOUT            |               | false    |
| SERVER_WRITE_TIMEOUT           |               | false    |
| DOMAIN_FILTER                  |               | false    |
| EXCLUDE_DOMAIN_FILTER          |               | false    |
| REGEXP_DOMAIN_FILTER           |               | false    |
| REGEXP_DOMAIN_FILTER_EXCLUSION |               | false    |
| REGEXP_NAME_FILTER             |               | false    |


## Contribution
All PRs are welcome, but before you create a PR, make sure your changes pass the linters and the apache2 license is 
injected into the newly added files. The `make lint` command will do this for you. 

Another point is the tests. If you create/change functionality, make sure the tests are running, updated or necessary ones 
are added. The `make test` command is used to run the tests.

All commits MUST be SIGNED before merge into main branch.

## Running locally

To run provider in a local environment, you must provide all required settings through environment variables.
To run locally, set `SERVER_HOST` to `localhost`, otherwise leave it at `0.0.0.0`.
Infoblox Provider is a simple web server with several clearly defined routers:

| Route            | Method |
|------------------|--------|
| /healthz         | GET    |
| /records         | GET    |
| /records         | POST   |
| /adjustendpoints | POST   |

#### Reading Data
```shell
curl -H 'Accept: application/external.dns.webhook+json;version=1' localhost:8888/records
```

#### Writing Data

Here are the updating rules according to which the data in the DNS server will be updated:

- if updateNew is not part of Update Old , object should be created
- if updateOld is not part of Update New , object should be deleted
- if information is not present (TTL might change) , object should be updated
- if we rename the object, object should be deleted and created


Based on the rules I am providing some examples of `data.json` creating, changing and deleting records in DNS.

```shell
curl -X POST -H 'Accept: application/external.dns.webhook+json;version=1;' -H 'Content-Type: application/external.dns.webhook+json;version=1' -d @data.json localhost:8888/records
```

Create `test.cloud.example.com`
```json
{"Create":null,"UpdateOld":[{"dnsName":"test.cloud.example.com","targets":["1.3.2.1"],"recordType":"A","recordTTL":300}],"UpdateNew":null,"Delete":null}
```

Update `test.cloud.example.com`
```json
{"Create":null,"UpdateOld":[{"dnsName":"test.cloud.example.com","targets":["1.3.2.1"],"recordType":"A","recordTTL":300}],"UpdateNew":null,"Delete":[{"dnsName":"new-test.cloud.example.com","targets":["1.2.3.4","4.3.2.1"],"recordType":"A","recordTTL":300}]}
```

Delete `test-new.cloud.example.com`
```json
{"Create":null,"UpdateOld":null,"UpdateNew":null,"Delete":[{"dnsName":"new-test.cloud.example.","targets":["1.2.3.4","4.3.2.1"],"recordType":"A","recordTTL":300}]}
```
