/*
   @Time : 2019-06-28 13:49:25
   @Author : 
   @File : service
   @Software: blog
*/
package service

import (
	"blog/conf"
	"blog/dao"
	"blog/proto"
)

type Service struct {
	dao *dao.Dao
	HelloService proto.HelloService
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
	}
	return s
}


