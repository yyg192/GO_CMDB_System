package impl

import (
	"context"
	"database/sql"
	"fmt"

	"CMDBProject/api/conf"
	"CMDBProject/api/pkg/host"

	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

/**
这里就是实现你的Service的具体功能，对数据库进行增删改查，
但是我这里好像没有和数据库进行解偶后面再解偶
***/

var (
	HostService = &HostServiceImpl{} //这里要继承interface的Service，然后实现
)

type HostServiceImpl struct {
	db *sql.DB //这里没有给数据库解偶，暂时没用到其他数据库，先这样吧
	//这里怎么定义日志？？
	log *logrus.Logger
}

func (s *HostServiceImpl) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db
	s.log = conf.C().LogInfo.Log
	//注意如果你为了保护LogInfo或Log把它搞成小写，那他可能不会被toml给反射到
	// go比较鸡肋的点出现了。
	return nil
}

func (s *HostServiceImpl) SaveHost(ctx context.Context, h *host.Host) (*host.Host, error) {
	h.Id = xid.New().String() //生成唯一的全局id
	h.ResourceId = h.Id
	h.SyncAt = ftime.Now().Timestamp() //这里有点问题，我不太想要他的包

	tx, err := s.db.BeginTx(ctx, nil) //开启一个事务
	//http://cngoLib.com/database-sql.html#db-begintx
	//https://zhuanlan.zhihu.com/p/117476959
	if err != nil {
		return nil, fmt.Errorf("HostServiceImpl.SaveHost function error, start mysql contxt error, %s", err)
	}
	defer func() { //这个defer是最后才执行的！！！！
		if err != nil {
			//开启context失败，准备回滚事务
			if err := tx.Rollback(); err != nil {
				//如果连回滚事务都失败
				s.log.Error("HostServiceImpl.SaveHost function error, mysql rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				s.log.Error("HostServiceImpl.SaveHost function error, mysql commit error, %s", err)
			}
		}
	}()

	resource_stmt, err := tx.Prepare(insertResourceSQL)
	//准备插入Resource的语句
	if err != nil {
		return nil, err
	}
	defer resource_stmt.Close() //在获得资源的地方就准备好defer

	_, err = resource_stmt.Exec(
		h.Id, h.Vendor, h.Region, h.CreateAt, h.ExpireAt, h.Type,
		h.Name, h.Description, h.Status, h.SyncAt, h.SyncAccount, h.PublicIP,
		h.PrivateIP,
	)
	//执行sql语句

	if err != nil {

		return nil, err
	}

	describe_stmt, err := tx.Prepare(insertDescribeSQL)
	if err != nil {
		return nil, err
	}
	defer describe_stmt.Close()

	_, err = describe_stmt.Exec(
		h.Id, h.CPU, h.Memory, h.GPUAmount, h.GPUSpec,
		h.OSType, h.OSName, h.SerialNumber,
	)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (s *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	query := sqlbuilder.NewQuery(queryHostSQL).Order("create_at").Desc().Limit(int64(req.Offset()), uint(req.PageSize))
	//限制query从req.Offset()查询到req.PageSize() 回头改一下，
	//改成start和end
	querySqlStr, args := query.Build()
	s.log.Debug("QueryHost sql sentence: %s, args: %v", querySqlStr, args)

	query_stmt, err := s.db.Prepare(querySqlStr)
	if err != nil {
		return nil, fmt.Errorf("query sql statment prepare failed")
	}
	defer query_stmt.Close()

	rows, err := query_stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("stmt query error, %s", err)
	}

	hostSet := host.NewDefaultHostSet() //我要返回的是一整个页面的主机
	for rows.Next() {
		oneHost := host.NewDefaultHost()
		if err := rows.Scan(
			&oneHost.Id, &oneHost.Vendor, &oneHost.Region, &oneHost.Zone, &oneHost.CreateAt, &oneHost.ExpireAt,
			&oneHost.Category, &oneHost.Type, &oneHost.InstanceId, &oneHost.Name,
			&oneHost.Description, &oneHost.Status, &oneHost.UpdateAt, &oneHost.SyncAt, &oneHost.SyncAccount,
			&oneHost.PublicIP, &oneHost.PrivateIP, &oneHost.PayType, &oneHost.ResourceHash, &oneHost.DescribeHash,
			&oneHost.Id, &oneHost.CPU,
			&oneHost.Memory, &oneHost.GPUAmount, &oneHost.GPUSpec, &oneHost.OSType, &oneHost.OSName,
			&oneHost.SerialNumber, &oneHost.ImageID, &oneHost.InternetMaxBandwidthOut, &oneHost.InternetMaxBandwidthIn,
			&oneHost.KeyPairName, &oneHost.SecurityGroups,
		); err != nil {
			return nil, err
		}
		hostSet.Add(oneHost)
	}

	countStr, countArgs := query.BuildCount()
	//Count获取总数量, build一个count语句
	countStmt, err := s.db.Prepare(countStr)
	if err != nil {
		return nil, fmt.Errorf("prepare count stmt error, %s", err)
	}
	defer countStmt.Close()

	//QueryRow只查询一行，查询主机总量
	if err := countStmt.QueryRow(countArgs...).Scan(&hostSet.Total); err != nil {
		return nil, fmt.Errorf("count stmt query error, %s", err)
	}
	return hostSet, nil
}

func (s *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}
	switch req.UpdateMode {
	case host.PUT:
		ins.Update(req.UpdateHostData.Resource, req.UpdateHostData.Describe)
	case host.PATCH:
		ins.Patch(req.UpdateHostData.Resource, req.UpdateHostData.Describe)
	}
	if err := ins.Validate(); err != nil {
		/// 检查更新后的参数是否合法
		return nil, err
	}
	updateRes_stmt, err := s.db.Prepare(updateResourceSQL)
	if err != nil {
		return nil, fmt.Errorf("HostServiceImpl.UpdateHost function error, updateResourceSQL error: %v", err)

	}
	defer updateRes_stmt.Close()
	_, err = updateRes_stmt.Exec(ins.Vendor, ins.Region, ins.Zone, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return nil, err
	}
	return ins, nil
}
func (s *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)
	querySQL, args := query.Where("id = ?", req.Id).BuildQuery()
	s.log.Debug("DescribeHost sql sentence: %s, args: %v", querySQL, args)
	describe_stmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, fmt.Errorf("prepare query host sql error, %s", err)
	}
	defer describe_stmt.Close()
	//这就是查一个主机的详细信息
	ins := host.NewDefaultHost()
	err = describe_stmt.QueryRow(args...).Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
		&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name,
		&ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
		&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.ResourceHash, &ins.DescribeHash,
		&ins.Id, &ins.CPU,
		&ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
		&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
		&ins.KeyPairName, &ins.SecurityGroups,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("critical error! host %s not found", req.Id)
		}
		return nil, fmt.Errorf("query host describe info error, %s", err)
	}
	return ins, nil
}
func (s *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	var ()
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil) //开启一个事务
	if err != nil {
		return nil, fmt.Errorf("HostServiceIml.DeleteHost function error, bgein context error, %v", err)
	}
	defer func() {
		if err != nil {
			if err = tx.Rollback(); err != nil {
				s.log.Error("HostServiceImpl.DeleteHost function error, mysql rollback error, %v", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				s.log.Error("HostServiceImpl.DeleteHost function error, mysql commit error, %s", err)
			}
		}
	}()

	deleteRes_stmt, err := tx.Prepare(deleteResourceSQL)
	if err != nil {
		return nil, fmt.Errorf("HostServiceIml.DeleteHost function error, the describe sentence prepare error, %v", err)
	}
	defer deleteRes_stmt.Close()

	_, err = deleteRes_stmt.Exec(req.Id)
	if err != nil {
		return nil, err
	}

	deleteHost_stmt, err := tx.Prepare(deleteHostSQL)
	if err != nil {
		return nil, err
	}
	defer deleteHost_stmt.Close()

	_, err = deleteHost_stmt.Exec(req.Id)
	if err != nil {
		return nil, err
	}

	return ins, nil

}
