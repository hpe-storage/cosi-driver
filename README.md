# HPE COSI Driver for Kubernetes

[![COSI Specification](https://img.shields.io/badge/COSI_Specification-v1alpha1-green)](https://github.com/kubernetes-sigs/container-object-storage-interface-spec/tree/v0.1.0)
[![Go Version](https://img.shields.io/badge/Go_version-v1.22-blue)](https://tip.golang.org/doc/go1.22)
[![Helm Version](https://img.shields.io/badge/Helm_version-v3-navy)](https://helm.sh/docs/intro/install/)

This repository contains the sources for the HPE Container Object Storage Interface (COSI) Driver.

## Releases

The COSI driver is released with a container images published on Quay. This is in turn referenced from deployment manifests, Helm charts. Releases are rolled up on the [HPE Storage Container Orchestrator Documentation](https://scod.hpedev.io/cosi_driver/index.html#compatibility_and_support) (SCOD) portal.

## Deploying and using the COSI driver on Kubernetes

All documentation for installing and using the COSI driver is available on [SCOD](https://scod.hpedev.io/cosi_driver/deployment.html).

Release vehicles include:

- [HPE COSI Driver for Kubernetes Helm chart](https://artifacthub.io/packages/helm/hpe-storage/hpe-cosi-driver)

Source deployment manifests are available in [hpe-storage/co-deployments](https://github.com/hpe-storage/co-deployments).

## Building

Instructions on how to build the HPE COSI Driver can be found in [BUILDING.md](BUILDING.md).

## Testing

Example Kubernetes object definitions used to build test cases for the COSI driver are available in [docs/examples](docs/examples) directory.

## Support

The HPE COSI Driver version 1.0.0 and onwards is fully supported by HPE and is Generally Available. Certain COSI features may be subject to alpha or beta status. Refer to the [official table](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/) of feature gates in the Kubernetes documentation to find availability of beta features.

Use [SCOD](https://scod.hpedev.io/legal/support) facility for formal support of your HPE storage products, including the COSI driver.

## Community

Please file any issues, questions or feature requests you may have [here](https://github.com/hpe-storage/cosi-driver/issues) (do not use this facility for support inquiries of your HPE storage product, see [SCOD](https://scod.hpedev.io/legal/support) for support)

## Contributing

We value all feedback and contributions. If you find any issues or want to contribute, please feel free to open an issue or file a PR. More details in [CONTRIBUTING.md](CONTRIBUTING.md)

## License

This is open source software licensed using the Apache License 2.0. Please see [LICENSE](LICENSE) for details.
