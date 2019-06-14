



## kubeadm 搭建1.10.9高可用集群

| 主机名称   | IP   | 备注          |
| ------ | ---- | ----------- |
| node01 | ip1  | master/etcd |
| node02 | ip2  | master/etcd |
| node03 | ip3  | master/etcd |
| 负载均衡   | vip  |             |

#### 软件版本

| 软件名称           | 版本         |
| -------------- | ---------- |
| Docker         | 17.03.2-ce |
| Kubelet        | 1.10.9     |
| Kubeadm        | 1.10.9     |
| Kubectl        | 1.10.9     |
| Docker-compose | 1.17.1     |



#### 镜像版本

| 镜像名称                                     | 版本      |
| ---------------------------------------- | ------- |
| k8s.gcr.io/kube-proxy-amd64              | v1.10.9 |
| k8s.gcr.io/kube-scheduler-amd64          | v1.10.9 |
| k8s.gcr.io/kube-controller-manager-amd64 | v1.10.9 |
| k8s.gcr.io/kube-apiserver-amd64          | v1.10.9 |
| k8s.gcr.io/etcd-amd64                    | 3.1.12  |
| k8s.gcr.io/k8s-dns-sidecar-amd64         | 1.14.8  |
| k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64   | 1.14.8  |
| k8s.gcr.io/k8s-dns-kube-dns-amd64        | 1.14.8  |
| k8s.gcr.io/pause-amd64                   | 3.1     |

#### yum源

```
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
       http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
```



#### 安装docker

```
yum install https://mirrors.aliyun.com/docker-ce/linux/centos/7/x86_64/stable/Packages/docker-ce-selinux-17.03.2.ce-1.el7.centos.noarch.rpm  -y
yum install https://mirrors.aliyun.com/docker-ce/linux/centos/7/x86_64/stable/Packages/docker-ce-17.03.2.ce-1.el7.centos.x86_64.rpm  -y
```

修改配置文件 /usr/lib/systemd/system/docker.service

```
ExecStart=/usr/bin/dockerd   -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock  --registry-mirror=https://ms3cfraz.mirror.aliyuncs.com
```



#### 安装docker-compose

```
sudo curl -L https://github.com/docker/compose/releases/download/1.17.1/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

#### 安装kubeadm

```
yum install -y kubelet-1.10.9 kubeadm-1.10.9 kubectl-1.10.9 kubernetes-cni-0.6.0 --disableexcludes=kubernetes
systemctl enable docker && systemctl start docker
systemctl enable kubelet && systemctl start kubelet
```

#### 拉取镜像

```
images=(kube-proxy-amd64:v1.10.9 kube-scheduler-amd64:v1.10.9 kube-controller-manager-amd64:v1.10.9 kube-apiserver-amd64:v1.10.9
etcd-amd64:3.1.12 pause-amd64:3.1 kubernetes-dashboard-amd64:v1.8.3 k8s-dns-sidecar-amd64:1.14.8 k8s-dns-kube-dns-amd64:1.14.8
k8s-dns-dnsmasq-nanny-amd64:1.14.8)
for imageName in ${images[@]} ; do
  docker pull registry.aliyuncs.com/google_containers/$imageName
  docker tag registry.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
  docker rmi registry.aliyuncs.com/google_containers/$imageName
done
```



#### 环境初始化

1. 设置主机名称

   ```
   hostnamectl set-hostname node01
   hostnamectl set-hostname node02
   hostnamectl set-hostname node03
   ```

   ​

2. 配置主机映射

   ```
   cat <<EOF > /etc/hosts
   127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
   ::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
   ip1 node01
   ip2 node02
   ip3 node03
   EOF
   ```

   ​

3. 执行ssh免密登陆

   ```
   ssh-keygen  #一路回车即可
   ssh-copy-id  node02
   ssh-copy-id  node03
   ```

   ​

4. 主机配置

   ```
   systemctl stop firewalld
   systemctl disable firewalld

   swapoff -a 
   sed -i 's/.*swap.*/#&/' /etc/fstab

   setenforce  0 

   cat <<EOF >  /etc/sysctl.d/k8s.conf
   net.bridge.bridge-nf-call-ip6tables = 1
   net.bridge.bridge-nf-call-iptables = 1
   EOF
   sysctl --system

   systemctl daemon-reload
   systemctl restart kubelet
   ```

5. 准备etcd证书(node01执行)

   ```
   wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64
   wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64
   wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64
   chmod +x cfssl_linux-amd64
   mv cfssl_linux-amd64 /usr/local/bin/cfssl
   chmod +x cfssljson_linux-amd64
   mv cfssljson_linux-amd64 /usr/local/bin/cfssljson
   chmod +x cfssl-certinfo_linux-amd64
   mv cfssl-certinfo_linux-amd64 /usr/local/bin/cfssl-certinfo
   export PATH=/usr/local/bin:$PATH
   ```

   创建ca配置文件

   ```
   mkdir /root/ssl
   cd /root/ssl
   cat >  ca-config.json <<EOF
   {
   "signing": {
   "default": {
     "expiry": "8760h"
   },
   "profiles": {
     "kubernetes-Soulmate": {
       "usages": [
           "signing",
           "key encipherment",
           "server auth",
           "client auth"
       ],
       "expiry": "8760h"
     }
   }
   }
   }
   EOF

   cat >  ca-csr.json <<EOF
   {
   "CN": "kubernetes-Soulmate",
   "key": {
   "algo": "rsa",
   "size": 2048
   },
   "names": [
   {
     "C": "CN",
     "ST": "shanghai",
     "L": "shanghai",
     "O": "k8s",
     "OU": "System"
   }
   ]
   }
   EOF

   cfssl gencert -initca ca-csr.json | cfssljson -bare ca

   cat > etcd-csr.json <<EOF
   {
     "CN": "etcd",
     "hosts": [
       "127.0.0.1",
       "ip1",
       "ip2",
       "ip3"
     ],
     "key": {
       "algo": "rsa",
       "size": 2048
     },
     "names": [
       {
         "C": "CN",
         "ST": "shanghai",
         "L": "shanghai",
         "O": "k8s",
         "OU": "System"
       }
     ]
   }
   EOF

   cfssl gencert -ca=ca.pem \
     -ca-key=ca-key.pem \
     -config=ca-config.json \
     -profile=kubernetes-Soulmate etcd-csr.json | cfssljson -bare etcd
   ```

   分发证书到node02 node03

   ```
   mkdir -p /etc/etcd/ssl
   cp etcd.pem etcd-key.pem ca.pem /etc/etcd/ssl/
   ssh -n node02 "mkdir -p /etc/etcd/ssl && exit"
   ssh -n node03 "mkdir -p /etc/etcd/ssl && exit"
   scp -r /etc/etcd/ssl/*.pem node02:/etc/etcd/ssl/
   scp -r /etc/etcd/ssl/*.pem node03:/etc/etcd/ssl/
   ```

   ​

   ​

#### 搭建etcd集群

Docker-compose:

以node01为例(其他node上的  ip1是当前node的ip)

```
version: "2"
services:
  etcd:
    image: k8s.gcr.io/etcd-amd64:3.1.12
    restart: always
    network_mode: host
    volumes:
      - "/data/etcd:/data/etcd"
      - "/etc/localtime:/etc/localtime"
      - "/etc/etcd/ssl:/etc/etcd/ssl"
    environment:
      ETCDCTL_API: "3"
    command: etcd  -name etcd1 -data-dir /data/etcd -advertise-client-urls https://ip1:2379 -listen-client-urls https://ip1:2379,http://127.0.0.1:2379 -initial-advertise-peer-urls https://ip1:2380 -listen-peer-urls https://ip1:2380 -initial-cluster-token etcd-cluster-0 -initial-cluster etcd1=https://ip1:2380,etcd2=https://ip2:2380,etcd3=https://ip3:2380 -initial-cluster-state new -cert-file /etc/etcd/ssl/etcd.pem -key-file /etc/etcd/ssl/etcd-key.pem -peer-cert-file /etc/etcd/ssl/etcd.pem -peer-key-file /etc/etcd/ssl/etcd-key.pem -trusted-ca-file /etc/etcd/ssl/ca.pem -peer-trusted-ca-file /etc/etcd/ssl/ca.pem
```

启动

```
docker-compose up -d
```

#### 配置kubeadm

修改kubelet配置文件

```
/etc/systemd/system/kubelet.service.d/10-kubeadm.conf
#修改这一行
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=cgroupfs"

systemctl daemon-reload
systemctl enable kubelet
```



Kubeadm master-config.yaml :

```
apiVersion: kubeadm.k8s.io/v1alpha1
kind: MasterConfiguration
etcd:
  endpoints:
  - https://ip1:2379
  - https://ip2:2379
  - https://ip3:2379
  caFile: /etc/etcd/ssl/ca.pem
  certFile: /etc/etcd/ssl/etcd.pem
  keyFile: /etc/etcd/ssl/etcd-key.pem
  dataDir: /var/lib/etcd
networking:
  podSubnet: 10.244.0.0/16
kubernetesVersion: 1.10.9
api:
  advertiseAddress: "vip"
token: "b99a00.a144ef80536d4344"
tokenTTL: "0s"
apiServerCertSANs:
- node01
- node02
- node03
- ip1
- ip2
- ip3
- vip
featureGates:
  CoreDNS: true
```

node01初始化集群

```
kubeadm init --config master-config.yaml
```

安装flannel网络

```
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
#版本信息：quay.io/coreos/flannel:v0.10.0-amd64

kubectl create -f  kube-flannel.yml
```



生成的证书分发到2号和3号机

```
scp -r /etc/kubernetes/pki  node02:/etc/kubernetes/
scp -r /etc/kubernetes/pki  node03:/etc/kubernetes/
```

在node02和03初始化

```
kubeadm init --config config.yaml
```

错误修改
 unable to check if the container runtime at “/var/run/dockershim.sock” is running: exit status 
```
rm -f /usr/bin/crictl
```
