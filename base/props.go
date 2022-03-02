package base

import (
	"context"
	"github.com/mszhangyi/infra"
	"github.com/mszhangyi/infra/utils"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	props   *systemConf
	EtcdKey string
)

type PropsStarter struct {
	infra.BaseStarter
}

type systemConf struct {
	Port        int    `json:"port"`
	Name        string `json:"name"`
	EmqAddr     string `json:"emq_addr"`
	EmqUser     string `json:"emq_user"`
	EmqPwd      string `json:"emq_pwd"`
	Environment string
	//mysql
	DataSource string `json:"data_source"`
	//log配置
	LogDir           string `json:"log_dir"`
	LogLevel         string `json:"log_level"`
	LogMaxAge        int    `json:"log_max_age"`
	LogRotationTime  int    `json:"log_rotation_time"`
	LogEnableLineLog bool   `json:"log_enableLineLog"`
	//redis
	RedisMaxIdle     int    `json:"redis_max_active"`
	RedisIdleTimeout int    `json:"redis_idle_timeout"`
	RedisMaxActive   int    `json:"redis_max_active"`
	RedisPwd         string `json:"redis_pwd"`
	RedisAddr        string `json:"redis_addr"`
	RedisSelectDb    int    `json:"redis_select_db"`
}

func (p *PropsStarter) Init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"81.68.243.67:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		log.Println("get failed, err:", err)
		return
	}
	props = &systemConf{}
	utils.ByteJsonByData(resp.Kvs[0].Value, props)
	InitLog()
}
