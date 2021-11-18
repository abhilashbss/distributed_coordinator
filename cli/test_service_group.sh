
go run ../src/main.go --log_path ../multinode_testing_dir/node1/log.txt --node_conf_path ../multinode_testing_dir/node1/node_init.conf --cluster_conf_path ../multinode_testing_dir/node1/cluster_meta.conf & 
sleep 1
go run ../src/main.go --log_path ../multinode_testing_dir/node2/log.txt --node_conf_path ../multinode_testing_dir/node2/node_init.conf --cluster_conf_path ../multinode_testing_dir/node2/cluster_meta.conf & 
go run ../src/main.go --log_path ../multinode_testing_dir/node3/log.txt --node_conf_path ../multinode_testing_dir/node3/node_init.conf --cluster_conf_path ../multinode_testing_dir/node3/cluster_meta.conf & 
go run ../src/main.go --log_path ../multinode_testing_dir/node4/log.txt --node_conf_path ../multinode_testing_dir/node4/node_init.conf --cluster_conf_path ../multinode_testing_dir/node4/cluster_meta.conf & 






