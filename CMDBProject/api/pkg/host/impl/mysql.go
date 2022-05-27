package impl

//一些大公司禁止使用ORM，使用ORM不好维护，而且DBA也不好管理。
const (
	insertResourceSQL = `INSERT INTO resource (
		id,
		vendor,
		region,
		zone,
		create_at,
		expire_at,
		category,
		type,
		instance_id,
		name,description,
		status,update_at,
		sync_at,
		sync_accout,
		public_ip,
		private_ip,
		pay_type,
		describe_hash,
		resource_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,
		cpu,memory,
		gpu_amount,
		gpu_spec,
		os_type,
		os_name,
		serial_number,
		image_id,
		internet_max_bandwidth_out,
		internet_max_bandwidth_in,
		key_pair_name,
		security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`

	// INSERT INTO `host` ( resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number )
	// VALUES
	// ( "111", 1, 2048, 1, 'n', 'linux', 'centos8', '00000' );
	insertDescribeSQL = `INSERT INTO host ( 
		resource_id, 
		cpu, 
		memory, 
		gpu_amount, 
		gpu_spec, 
		os_type, 
		os_name, 
		serial_number 
	)
	VALUES
		( ?,?,?,?,?,?,?,? );
	`
	updateResourceSQL = `UPDATE resource SET 
		expire_at=?,
		category=?,
		type=?,
		name=?,
		description=?,
		status=?,
		update_at=?,
		sync_at=?,
		sync_accout=?,
		public_ip=?,
		private_ip=?,
		pay_type=?,
		describe_hash=?,
		resource_hash=?
	WHERE id = ?`
	updateHostSQL = `UPDATE host SET 
		cpu=?,memory=?,
		gpu_amount=?,
		gpu_spec=?,
		os_type=?,
		os_name=?,
		image_id=?,
		internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,
		key_pair_name=?,
		security_groups=?
	WHERE resource_id = ?`

	queryHostSQL      = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
	deleteHostSQL     = `DELETE FROM host WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)
