# Install boot2docker (Mac or Windows)

```
https://github.com/boot2docker/boot2docker
```

# Open port for HAProxy dashboard

skip this step if you are not using boot2docker

```
VBoxManage modifyvm "boot2docker-vm" --natpf1 "tcp-port9000,tcp,,9000,,9000";
VBoxManage modifyvm "boot2docker-vm" --natpf1 "udp-port9000,udp,,9000,,9000";
```

# How to build:

```
docker build -rm --tag=multi-be-haproxy .
```

# How to run:

```
docker run -p :9000:9000 --privileged  -i -t multi-be-haproxy /bin/bash
```

# Start HAProxy and backend apps

```
./run.sh # pwd /app
```

The application runs inside tmux.  Use the following command to connect to it.

```
tmux attach-session -t forty9ten
```


