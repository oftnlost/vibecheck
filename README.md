# vibecheck

a stateless web app that leverages sentiment analysis models to plot the _vibe_ of tweeters over time.


## what?

a full-stack project using the following technologies:

| role          | technologies |
| ------------- | ------------ |
| deployment    | [kubernetes](https://kubernetes.io/), [cri-o](https://cri-o.io/), [nginx ingress](https://kubernetes.github.io/ingress-nginx/) |
| backend logic | [golang](https://golang.org/) |
| web routing   | [gin](https://github.com/gin-gonic/gin) |
| frontend      | \<to do\> |
| external apis | \<to do\> |

_note:_ we intentially picked tools that we haven't used before to make this project as much of a learning experience as possible.


## available routes

`/api/test` : `(get)`

to test routing logic with go and gin. this should return `{"message": "test successful"}`.


## how to get started locally

_note:_ the following was tested on `Linux 5.10.0-1008-oem #9-Ubuntu SMP Tue Dec 15 14:22:38 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux`

this repo was designed to have the containers deployed across a cluster of nodes. the container orchestration tool of choice is kubernetes, with the cri-o container runtime. so, to test this locally, we used [minikube](https://minikube.sigs.k8s.io/docs/), which leverages [docker](https://www.docker.com/), to create an environment that simulates a cluster on a single machine.

after following the instructions to install kubernetes, cri-o, docker, and minikube, start minikube with the following command:

```bash
$ minikube start --container-runtime cri-o
```

to access the web server inside minikube, we used the nginx ingress controller, which can be enabled in minikube with the following command:

```bash
$ minikube addons enable ingress
```

to create the components of the web app, use the following command:

```bash
$ kubectl apply -f api-deployment.yml
```

at this stage, the web server should be running inside minikube behind the nginx ingress. with the following command, interogate the nginx ingress to find the IP where the web server can be reached locally.

```bash
$ kubectl describe ingress
Name:             nginx-ingress
Namespace:        default
Address:          192.168.49.2
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Rules:
  Host        Path  Backends
  ----        ----  --------
  *
              /   api-cluster-ip-service:5000 (10.244.0.3:5000,10.244.0.4:5000)
Annotations:  kubernetes.io/ingress.class: nginx
              nginx.ingress.kubernetes.io/rewrite-target: /
Events:
  Type    Reason  Age    From                      Message
  ----    ------  ----   ----                      -------
  Normal  CREATE  6m45s  nginx-ingress-controller  Ingress default/nginx-ingress
  Normal  UPDATE  6m44s  nginx-ingress-controller  Ingress default/nginx-ingress
```

with the above address, you can access the web server at 192.168.49.2


## how to deploy on a cluster

\<to do\>

