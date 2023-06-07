'''

docker run -d --name Etcd-server \     
--publish 2379:2379 \                
--publish 2380:2380 \                                                             
--env ALLOW_NONE_AUTHENTICATION=yes \
--env ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
--env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
--env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380 \
bitnami/etcd:latest

      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd1:2379
'''