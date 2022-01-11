## Update on 10-04-2021

CSI is moved to and supported in [Alluxio main repo](https://github.com/Alluxio/alluxio).
**You shall stop using this repo which is out of date**.

The source code is moved to https://github.com/Alluxio/alluxio/tree/master/integration/docker/csi.
The documentation is moved to https://github.com/Alluio/alluxio/tree/master/integration/kubernetes/CSI_README.md.
This feature is also available in `alluxio/alluxio` docker image on dockerhub.


## Update on 07-03-2021
This feature has been moved back to Alluxio main repo https://github.com/Alluxio/alluxio/tree/master/integration/csi,
and no longer maintained in this separate repo.

## What is Alluxio
[Alluxio](https://www.alluxio.io) (formerly known as Tachyon)
is a virtual distributed storage system. It bridges the gap between
computation frameworks and storage systems, enabling computation applications to connect to
numerous storage systems through a common interface. Read more about
[Alluxio Overview](https://docs.alluxio.io/os/user/stable/en/Overview.html).

The Alluxio project originated from a research project called Tachyon at AMPLab, UC Berkeley,
which was the data layer of the Berkeley Data Analytics Stack ([BDAS](https://amplab.cs.berkeley.edu/bdas/)).
For more details, please refer to Haoyuan Li's PhD dissertation
[Alluxio: A Virtual Distributed File System](https://www2.eecs.berkeley.edu/Pubs/TechRpts/2018/EECS-2018-29.html).

## Container Storage Interface for Alluxio

This repository contains the CSI implementation to provide POSIX access to Alluxio in
containerized environments such as Kubernetes.

## How to build

Build alluxio-csi docker image

`docker build . -t alluxio/alluxio-csi:<version_tag>`

## Example

### Deploy

Go to `deploy` folder, Run following commands:
```bash
kubectl apply -f csi-alluxio-driver.yml
kubectl apply -f csi-alluxio-daemon.yml
``` 

### Example Nginx application
The `/examples` folder contains `PersistentVolume`, `PersistentVolumeClaim` and an nginx `Pod` mounting the alluxio volume under `/data`.

You will need to update the alluxio `MASTER_HOST_NAME` and the share information under `volumeAttributes` in `alluxio-pv.yaml` file to match your alluxio master configuration.

Run following commands to create nginx pod mounted with alluxio volume:
```bash
kubectl apply -f alluxio-pv.yml
kubectl apply -f alluxio-pvc.yml
kubectl apply -f nginx.yml
```

## Useful Links

- [Alluxio Github](https://github.com/Alluxio/alluxio)
- [Alluxio Website](https://www.alluxio.io/)
- [Downloads](https://www.alluxio.io/download)
- [Releases and Notes](https://www.alluxio.io/download/releases/)
- [Documentation](https://www.alluxio.io/docs/)
