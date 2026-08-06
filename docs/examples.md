# Creating records with ExternalDNS

Worked examples for creating Infoblox records through ExternalDNS, either from a
`DNSEndpoint` custom resource or from annotations on regular Kubernetes objects.

All examples assume ExternalDNS is already running with this webhook configured as its
provider, and that the parent zone (for example `cloud.example.com`) exists in the grid.

## Record types this provider supports

The provider maps ExternalDNS endpoints onto these Infoblox WAPI objects:

| Record type | Infoblox object | Notes                                                        |
|-------------|-----------------|--------------------------------------------------------------|
| A           | `record:a`      |                                                              |
| CNAME       | `record:cname`  |                                                              |
| TXT         | `record:txt`    | Also used by the TXT registry for ownership records          |
| NS          | `record:ns`     |                                                              |
| PTR         | `record:ptr`    | Created for A records when `INFOBLOX_CREATE_PTR=true`        |

Any other record type is silently ignored by the provider, so do not use `AAAA`, `MX`
or `SRV` in the examples below.

Two consequences worth calling out:

* ExternalDNS manages `A`, `AAAA` and `CNAME` by default (`--managed-record-types`).
  Because this provider has no `AAAA` mapping, IPv6 targets are dropped. ExternalDNS
  picks the record type from the target, so an IPv6 address on a Service or Ingress
  produces an `AAAA` endpoint that the provider will not create.
* To let ExternalDNS manage `NS` records from a `DNSEndpoint`, add `NS` to
  `--managed-record-types`, for example
  `--managed-record-types=A --managed-record-types=CNAME --managed-record-types=TXT --managed-record-types=NS`.
  `PTR` is not one of the record types ExternalDNS accepts there; PTR records come from
  `INFOBLOX_CREATE_PTR` instead (see below).

## DNSEndpoint (CRD source)

The `DNSEndpoint` CRD gives you direct control over the record name, type, targets and
TTL, without deriving anything from a Service or Ingress.

Run ExternalDNS with the CRD source:

```bash
external-dns \
  --source=crd \
  --provider=webhook \
  --webhook-provider-url=http://localhost:8888 \
  --domain-filter=cloud.example.com \
  --managed-record-types=A \
  --managed-record-types=CNAME \
  --managed-record-types=TXT
```

`--source=crd` defaults to `externaldns.k8s.io/v1alpha1` / kind `DNSEndpoint`
(`--crd-source-apiversion`, `--crd-source-kind`). The CRD itself is not installed by
ExternalDNS; apply the `DNSEndpoint` CRD manifest from the ExternalDNS repository first.

```yaml
apiVersion: externaldns.k8s.io/v1alpha1
kind: DNSEndpoint
metadata:
  name: example-records
  namespace: default
spec:
  endpoints:
    # A record with two targets
    - dnsName: app.cloud.example.com
      recordType: A
      recordTTL: 300
      targets:
        - 10.0.0.10
        - 10.0.0.11
    # CNAME pointing at the A record above
    - dnsName: www.cloud.example.com
      recordType: CNAME
      recordTTL: 300
      targets:
        - app.cloud.example.com
    # TXT record
    - dnsName: verify.cloud.example.com
      recordType: TXT
      recordTTL: 300
      targets:
        - "some-verification-token"
```

Field names are those of the ExternalDNS `Endpoint` type: `dnsName`, `recordType`,
`targets`, `recordTTL`, and optionally `setIdentifier`, `labels` and `providerSpecific`.
`recordTTL` is in seconds. Omitting it makes the provider fall back to
`INFOBLOX_DEFAULT_TTL` (and `INFOBLOX_USE_TTL` controls whether the TTL is sent at all).

### PTR records from a DNSEndpoint

With `INFOBLOX_CREATE_PTR=true` the provider creates the PTR companion for each `A`
record itself, so you do not need to declare it. The reverse zone must be inside the
domain filter:

```bash
DOMAIN_FILTER="cloud.example.com, 10.0.0.0/24"
INFOBLOX_CREATE_PTR=true
```

## Service annotations

Run ExternalDNS with the service source:

```bash
external-dns --source=service --provider=webhook --webhook-provider-url=http://localhost:8888
```

A `LoadBalancer` Service gets an `A` record per load balancer IPv4 address:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: app
  namespace: default
  annotations:
    external-dns.alpha.kubernetes.io/hostname: app.cloud.example.com
    external-dns.alpha.kubernetes.io/ttl: "300"
spec:
  type: LoadBalancer
  selector:
    app: app
  ports:
    - port: 80
      targetPort: 8080
```

To publish a name that does not follow the Service's own address, set the target
explicitly. A non-IP target becomes a `CNAME`, an IPv4 target becomes an `A` record:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: app-alias
  namespace: default
  annotations:
    external-dns.alpha.kubernetes.io/hostname: alias.cloud.example.com
    # non-IP target -> CNAME
    external-dns.alpha.kubernetes.io/target: app.cloud.example.com
    external-dns.alpha.kubernetes.io/ttl: "300"
spec:
  type: ClusterIP
  selector:
    app: app
  ports:
    - port: 80
```

## Ingress annotations

Run ExternalDNS with the ingress source:

```bash
external-dns --source=ingress --provider=webhook --webhook-provider-url=http://localhost:8888
```

By default the hostnames come from `spec.rules[].host` and the targets from the Ingress
load balancer status:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: app
  namespace: default
  annotations:
    external-dns.alpha.kubernetes.io/ttl: "300"
spec:
  ingressClassName: nginx
  rules:
    - host: app.cloud.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: app
                port:
                  number: 80
```

Additional names can be added with the `hostname` annotation, and
`ingress-hostname-source` controls which of the two sources is used. Its accepted values
are `defined-hosts-only` (only `spec.rules[].host`) and `annotation-only` (only the
`hostname` annotation); if the annotation is absent, both sources are used:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: app-extra-names
  namespace: default
  annotations:
    external-dns.alpha.kubernetes.io/hostname: www.cloud.example.com,web.cloud.example.com
    external-dns.alpha.kubernetes.io/ingress-hostname-source: annotation-only
    # override the target instead of using the ingress load balancer address
    external-dns.alpha.kubernetes.io/target: 10.0.0.20
    external-dns.alpha.kubernetes.io/ttl: 5m
spec:
  ingressClassName: nginx
  rules:
    - host: app.cloud.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: app
                port:
                  number: 80
```

## Annotations used above

Every annotation below is defined in the ExternalDNS `source` package and applies to any
provider, this one included.

| Annotation                                                | Purpose                                                                            |
|-----------------------------------------------------------|------------------------------------------------------------------------------------|
| `external-dns.alpha.kubernetes.io/hostname`                | Comma separated list of names to create                                            |
| `external-dns.alpha.kubernetes.io/target`                  | Overrides the target; IPv4 gives an `A` record, anything non-IP gives a `CNAME`     |
| `external-dns.alpha.kubernetes.io/ttl`                     | Record TTL, as seconds (`"300"`) or a Go duration (`5m`)                            |
| `external-dns.alpha.kubernetes.io/controller`              | Set to `dns-controller` to opt an object in when ExternalDNS filters on it          |
| `external-dns.alpha.kubernetes.io/ingress-hostname-source` | `defined-hosts-only` or `annotation-only`                                           |
| `external-dns.alpha.kubernetes.io/internal-hostname`       | Names published from the internal (ClusterIP / node internal) address               |
| `external-dns.alpha.kubernetes.io/access`                  | For headless/node sources: `public` or `private` node address                       |
| `external-dns.alpha.kubernetes.io/endpoints-type`          | For headless Services: `NodeExternalIP` or `HostIP`                                |
| `external-dns.alpha.kubernetes.io/set-identifier`          | Endpoint set identifier; only meaningful for providers with weighted routing        |

Annotations that are specific to other providers, such as
`external-dns.alpha.kubernetes.io/alias`, `external-dns.alpha.kubernetes.io/cloudflare-proxied`
and the `aws-`, `scw-` and `ibmcloud-` prefixed ones, have no effect on this provider.

## Verifying the result

The webhook exposes the records it sees on `/records`:

```shell
curl -H 'Accept: application/external.dns.webhook+json;version=1' localhost:8888/records
```
