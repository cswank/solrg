# solrg
A leader aware solr client


solr zk upconfig -d /tmp/solrg -n solrg -z zk:2181

{"things":{
    "replicationFactor":"2",
    "router":{"name":"compositeId"},
    "maxShardsPerNode":"3",
    "autoAddReplicas":"false",
    "shards":{
      "shard1":{
        "range":"80000000-d554ffff",
        "state":"active",
        "replicas":{
    "core_node2":{
            "core":"things_shard1_replica1",
            "base_url":"http://172.19.0.2:8983/solr",
            "node_name":"172.19.0.2:8983_solr",
            "state":"active",
            "leader":"true"},
          "core_node6":{
            "core":"things_shard1_replica2",
            "base_url":"http://172.19.0.4:8983/solr",
            "node_name":"172.19.0.4:8983_solr",
            "state":"active"}}},
      "shard2":{
        "range":"d5550000-2aa9ffff",
        "state":"active",
        "replicas":{
    "core_node1":{
            "core":"things_shard2_replica2",
            "base_url":"http://172.19.0.6:8983/solr",
            "node_name":"172.19.0.6:8983_solr",
            "state":"active",
            "leader":"true"},
          "core_node4":{
            "core":"things_shard2_replica1",
            "base_url":"http://172.19.0.3:8983/solr",
            "node_name":"172.19.0.3:8983_solr",
            "state":"active"}}},
      "shard3":{
        "range":"2aaa0000-7fffffff",
        "state":"active",
        "replicas":{
          "core_node3":{
            "core":"things_shard3_replica1",
            "base_url":"http://172.19.0.2:8983/solr",
            "node_name":"172.19.0.2:8983_solr",
            "state":"active",
            "leader":"true"},
          "core_node5":{
            "core":"things_shard3_replica2",
            "base_url":"http://172.19.0.4:8983/solr",
            "node_name":"172.19.0.4:8983_solr",
            "state":"active"}}}}}

