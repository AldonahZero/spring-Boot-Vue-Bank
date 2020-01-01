package com.cx.bank.mhg.bdemo.mapper;

import com.cx.bank.mhg.bdemo.domain.TUser;

import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Repository;

/**
 * TUserDAO继承基类
 */
@Repository
@Mapper
public interface TUserDAO extends MyBatisBaseDao<TUser, Integer> {
    //查询用户
    TUser selectByPrimaryKey(Integer id);
}